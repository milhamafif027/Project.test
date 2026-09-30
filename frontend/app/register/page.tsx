"use client";

import React from "react";
import Link from "next/link";
import RegisterForm from "../components/RegisterForm";

export default function RegisterPageRoute() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-slate-50 via-white to-emerald-50 text-slate-800 flex flex-col justify-between p-6">
      {/* Top Navigation */}
      <div className="w-full max-w-md mx-auto pt-4 flex justify-between items-center">
        <Link
          href="/"
          className="text-xs text-slate-600 hover:text-slate-900 bg-white px-3.5 py-1.5 rounded-full border border-slate-200 shadow-sm transition-all flex items-center space-x-1.5"
        >
          <span>← Kembali ke Beranda</span>
        </Link>
        <span className="text-xs text-emerald-600 font-medium">
          Halaman Registrasi
        </span>
      </div>

      {/* Form Container */}
      <div className="w-full max-w-md mx-auto my-auto py-6">
        <RegisterForm />
      </div>

      {/* Footer */}
      <footer className="w-full text-center py-4 text-xs text-slate-400">
        &copy; 2026 Fullstack Assessment &bull; Go & Next.js
      </footer>
    </main>
  );
}
