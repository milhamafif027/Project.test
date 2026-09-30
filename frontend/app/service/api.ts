const API_BASE_URL = "http://localhost:8080/api";

export interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}

export async function registerUser(payload: RegisterPayload) {
  try {
    const response = await fetch(`${API_BASE_URL}/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Gagal melakukan pendaftaran ke server.");
    }

    return data;
  } catch (error: unknown) {
    const errMessage =
      error instanceof Error ? error.message : "Terjadi kesalahan jaringan.";
    throw new Error(errMessage);
  }
}
