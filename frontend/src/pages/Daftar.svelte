<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, authIsLoading } from "../lib/auth.js";

  let showPassword = $state(false);
  let nama = $state("");
  let email = $state("");
  let phone = $state("");
  let password = $state("");
  let errorMessage = $state("");

  async function handleFormSubmit(event) {
    event.preventDefault();
    errorMessage = "";
    try {
      await auth.register(nama, email, phone, password);
      await auth.login(email, password);
      navigate("/dashboard");
    } catch (error) {
      errorMessage = error.data?.error || error.message || "Pendaftaran gagal.";
    }
  }
</script>

<main class="flex min-h-screen items-center justify-center bg-slate-50 px-5 py-10 text-slate-900">
  <section class="w-full max-w-lg rounded-3xl border border-slate-200 bg-white p-7 shadow-sm sm:p-9">
    <Link to="/" class="inline-flex items-center gap-3" aria-label="Kembali ke halaman utama">
      <img src="/logo-perpusfun.svg" alt="" class="h-10 w-10 shrink-0" />
      <span class="text-lg font-bold tracking-tight">Perpus<span class="text-blue-600">Fun</span></span>
    </Link>

    <div class="mt-8">
      <p class="text-sm font-semibold text-blue-600">Selamat datang</p>
      <h1 class="mt-2 text-3xl font-bold tracking-tight">Buat akun baru</h1>
      <p class="mt-2 text-sm text-slate-500">Isi data berikut untuk bergabung dengan PerpusFun.</p>
    </div>

    <form class="mt-7 space-y-4" onsubmit={handleFormSubmit}>
      {#if errorMessage}
        <div role="alert" class="rounded-xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700">{errorMessage}</div>
      {/if}

      <div>
        <label for="nama" class="mb-1.5 block text-sm font-medium text-slate-700">Nama lengkap</label>
        <input
          type="text"
          id="nama"
          bind:value={nama}
          required
          autocomplete="name"
          placeholder="Nama kamu"
          class="w-full rounded-xl border border-slate-200 px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
        />
      </div>

      <div>
        <label for="email" class="mb-1.5 block text-sm font-medium text-slate-700">Email</label>
        <input
          type="email"
          id="email"
          bind:value={email}
          required
          autocomplete="email"
          placeholder="nama@email.com"
          class="w-full rounded-xl border border-slate-200 px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
        />
      </div>

      <div>
        <label for="phone" class="mb-1.5 block text-sm font-medium text-slate-700">Nomor telepon <span class="font-normal text-slate-400">(opsional)</span></label>
        <input
          type="tel"
          id="phone"
          bind:value={phone}
          autocomplete="tel"
          placeholder="08xxxxxxxxxx"
          class="w-full rounded-xl border border-slate-200 px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
        />
      </div>

      <div>
        <label for="password" class="mb-1.5 block text-sm font-medium text-slate-700">Password</label>
        <div class="relative">
          <input
            type={showPassword ? "text" : "password"}
            id="password"
            bind:value={password}
            required
            minlength="6"
            autocomplete="new-password"
            placeholder="Minimal 6 karakter"
            class="w-full rounded-xl border border-slate-200 px-4 py-3 pr-20 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
          />
          <button
            type="button"
            onclick={() => (showPassword = !showPassword)}
            class="absolute right-3 top-1/2 -translate-y-1/2 rounded-lg px-2 py-1 text-xs font-semibold text-blue-700 hover:bg-blue-50"
          >
            {showPassword ? "Sembunyikan" : "Lihat"}
          </button>
        </div>
      </div>

      <button
        type="submit"
        disabled={$authIsLoading}
        class="w-full rounded-xl bg-blue-600 px-4 py-3 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-wait disabled:opacity-60"
      >
        {$authIsLoading ? "Mendaftarkan..." : "Daftar"}
      </button>
    </form>

    <p class="mt-6 text-center text-sm text-slate-500">
      Sudah punya akun?
      <Link to="/login" class="ml-1 font-semibold text-blue-700 hover:text-blue-800">Masuk</Link>
    </p>
    <Link to="/" class="mt-4 block text-center text-sm font-medium text-slate-500 transition hover:text-blue-700">Kembali ke beranda</Link>
  </section>
</main>
