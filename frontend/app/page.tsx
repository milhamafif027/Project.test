import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-slate-50 via-white to-emerald-50 text-slate-800 flex flex-col justify-between">
      {/* Navbar Minimalis */}
      <header className="w-full max-w-6xl mx-auto px-6 py-6 flex items-center justify-between border-b border-slate-200">
        <div className="flex items-center space-x-2">
          <div className="h-3 w-3 rounded-full bg-emerald-600 animate-pulse"></div>
          <span className="font-bold tracking-tight text-slate-900">
            Fullstack Assessment
          </span>
        </div>
        <span className="text-xs text-emerald-600 bg-emerald-50 px-3 py-1 rounded-full border border-emerald-100 font-medium">
          Next.js & Golang Fiber
        </span>
      </header>

      {/* Main Hero Section / Content */}
      <div className="max-w-4xl mx-auto px-6 py-16 text-center my-auto w-full">
        <div className="inline-flex items-center space-x-2 rounded-full bg-emerald-50 px-3 py-1 text-xs font-medium text-emerald-600 border border-emerald-200 mb-6 shadow-sm">
          <span>Sistem Registrasi & Notifikasi Email Otomatis</span>
        </div>

        {/* Judul Utama */}
        <h1 className="text-3xl sm:text-5xl font-extrabold tracking-tight text-slate-900 leading-tight mb-6">
          Platform Pendaftaran Pengguna dengan{" "}
          <span className="text-emerald-600">Go & Next.js</span>
        </h1>

        <p className="text-slate-600 max-w-2xl mx-auto text-sm sm:text-base mb-8 leading-relaxed">
          Mengintegrasikan performa tinggi backend Go (Fiber & MySQL) dengan
          antarmuka modern Next.js, dilengkapi pengiriman notifikasi via SMTP
          Ethereal.
        </p>

        {/* Tombol Menuju Pendaftaran */}
        <div>
          <Link
            href="/register"
            className="inline-flex items-center space-x-2 rounded-xl bg-emerald-600 px-8 py-3.5 text-sm font-semibold text-white shadow-lg shadow-emerald-600/20 hover:bg-emerald-500 focus:outline-none focus:ring-2 focus:ring-emerald-500/40 active:scale-[0.98] transition-all"
          >
            <span>Mulai Pendaftaran</span>
            <svg
              className="w-4 h-4 ml-1"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M9 5l7 7-7 7"
              />
            </svg>
          </Link>
        </div>
      </div>

      {/* Footer */}
      <footer className="w-full text-center py-6 text-xs text-slate-400 border-t border-slate-200">
        Project Test Fullstack Engineer &bull; Dibangun dengan Next.js dan
        Golang
      </footer>
    </main>
  );
}
