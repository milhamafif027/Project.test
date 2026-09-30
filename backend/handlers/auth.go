package handlers

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format data tidak valid",
		})
	}

	// Cek email terdaftar
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email sudah terdaftar",
		})
	}

	// Simpan user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyimpan user ke database",
		})
	}

	// Ambil konfigurasi SMTP secara dinamis dari file .env
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SMTP_USER")
	senderPass := os.Getenv("SMTP_PASS")

	subject := "Subject: Pendaftaran Berhasil!\n"
	body := fmt.Sprintf("Halo %s,\n\nSelamat! Akun Anda dengan email %s sudah berhasil terdaftar di sistem kami.", user.Name, user.Email)
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", senderEmail, senderPass, smtpHost)
	err := sendMailTLS(smtpHost+":"+smtpPort, auth, senderEmail, []string{user.Email}, message)
	if err != nil {
		fmt.Println("Gagal mengirim email:", err)
	} else {
		fmt.Println("Email notifikasi berhasil dikirim ke:", user.Email)
	}

	return c.JSON(fiber.Map{
		"message": "User berhasil mendaftar dan email notifikasi telah diproses!",
		"user":    user,
	})
}

// Helper TLS Mailer
func sendMailTLS(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Quit()

	smtpHost := os.Getenv("SMTP_HOST")
	if err = client.StartTLS(&tls.Config{InsecureSkipVerify: true, ServerName: smtpHost}); err != nil {
		return err
	}

	if err = client.Auth(a); err != nil {
		return err
	}

	if err = client.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	return w.Close()
}