<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, authIsLoading, authError } from "../lib/auth.js";

  let isLogin = $state(true);
  let nama = $state("");
  let email = $state("");
  let phone = $state("");
  let password = $state("");
  let localError = $state("");

  async function handleSubmit(e) {
    e.preventDefault();
    localError = "";

    try {
      if (isLogin) {
        await auth.login(email, password);
        navigate("/dashboard");
      } else {
        await auth.register(nama, email, phone, password);
        await auth.login(email, password);
        navigate("/dashboard");
      }
    } catch (err) {
      localError = err.data?.error || err.message || "Terjadi kesalahan.";
    }
  }

  function clearError() {
    localError = "";
    auth.clearError();
  }
</script>

<main class="flex min-h-screen items-center justify-center bg-slate-50 px-5 py-10 text-slate-900">
  <section class="w-full max-w-md rounded-3xl border border-slate-200 bg-white p-7 shadow-sm sm:p-9">
    <Link to="/" class="inline-flex items-center gap-3" aria-label="Kembali ke halaman utama">
      <img src="/logo-perpusfun.svg" alt="" class="h-10 w-10 shrink-0" />
      <span class="text-lg font-bold tracking-tight">Perpus<span class="text-blue-600">Fun</span></span>
    </Link>

    <div class="mt-9">
      <p class="text-sm font-semibold text-blue-600">{isLogin ? "Selamat datang kembali" : "Mulai membaca bersama kami"}</p>
      <h1 class="mt-2 text-3xl font-bold tracking-tight text-slate-900">{isLogin ? "Masuk ke akun" : "Buat akun baru"}</h1>
      <p class="mt-2 text-sm leading-6 text-slate-500">
        {isLogin ? "Masuk untuk mengelola buku dan peminjamanmu." : "Daftar untuk mulai menjelajahi koleksi buku."}
      </p>
    </div>

    <form onsubmit={handleSubmit} class="mt-7 space-y-4">
      {#if localError || $authError}
        <div role="alert" class="rounded-xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700">
          {localError || $authError}
          <button type="button" onclick={clearError} class="ml-2 font-semibold underline">Tutup</button>
        </div>
      {/if}

      {#if !isLogin}
        <div>
          <label for="inputNama" class="mb-1.5 block text-sm font-medium text-slate-700">Nama lengkap</label>
          <input
            id="inputNama"
            type="text"
            bind:value={nama}
            required
            autocomplete="name"
            placeholder="Nama kamu"
            class="w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
          />
        </div>
        <div>
          <label for="inputPhone" class="mb-1.5 block text-sm font-medium text-slate-700">Nomor telepon <span class="font-normal text-slate-400">(opsional)</span></label>
          <input
            id="inputPhone"
            type="tel"
            bind:value={phone}
            autocomplete="tel"
            placeholder="08xxxxxxxxxx"
            class="w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
          />
        </div>
      {/if}

      <div>
        <label for="inputEmail" class="mb-1.5 block text-sm font-medium text-slate-700">Email</label>
        <input
          id="inputEmail"
          type="email"
          bind:value={email}
          required
          autocomplete="email"
          placeholder="nama@email.com"
          class="w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
        />
      </div>

      <div>
        <label for="inputPassword" class="mb-1.5 block text-sm font-medium text-slate-700">Password</label>
        <input
          id="inputPassword"
          type="password"
          bind:value={password}
          required
          autocomplete={isLogin ? "current-password" : "new-password"}
          minlength={isLogin ? undefined : 6}
          placeholder="Masukkan password"
          class="w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
        />
      </div>

      <button
        type="submit"
        disabled={$authIsLoading}
        class="mt-2 w-full rounded-xl bg-blue-600 px-4 py-3 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-wait disabled:opacity-60"
      >
        {$authIsLoading ? "Memproses..." : isLogin ? "Masuk" : "Daftar"}
      </button>
    </form>

    <p class="mt-6 text-center text-sm text-slate-500">
      {isLogin ? "Belum punya akun?" : "Sudah punya akun?"}
      <button
        type="button"
        onclick={() => (isLogin = !isLogin)}
        class="ml-1 font-semibold text-blue-700 hover:text-blue-800"
      >
        {isLogin ? "Daftar" : "Masuk"}
      </button>
    </p>
    <Link to="/" class="mt-5 block text-center text-sm font-medium text-slate-500 transition hover:text-blue-700">Kembali ke beranda</Link>
  </section>
</main>
