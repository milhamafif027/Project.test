import type { Metadata } from "next";
import { Plus_Jakarta_Sans } from "next/font/google";
import "./globals.css";

const jakarta = Plus_Jakarta_Sans({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700", "800"],
  variable: "--font-jakarta",
});

export const metadata: Metadata = {
  title: "Fullstack Assessment - Go & Next.js",
  description: "Sistem registrasi dengan notifikasi email otomatis",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="id" className={jakarta.variable}>
      <body
        className={`${jakarta.className} antialiased selection:bg-blue-500 selection:text-white`}
      >
        {children}
      </body>
    </html>
  );
}
