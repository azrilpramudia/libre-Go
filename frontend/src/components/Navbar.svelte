<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, currentMember, isAuthenticated } from "../lib/auth.js";

  let { searchQuery = $bindable("") } = $props();

  function scrollToSection(sectionId) {
    const scroll = () =>
      document.getElementById(sectionId)?.scrollIntoView({
        behavior: "smooth",
        block: "start",
      });

    if (window.location.pathname !== "/") {
      navigate("/");
      window.setTimeout(scroll, 100);
      return;
    }

    scroll();
  }

  function logout() {
    auth.logout();
    navigate("/");
  }
</script>

<nav class="flex flex-wrap items-center justify-between gap-4 rounded-2xl border border-slate-200 bg-white px-5 py-4 shadow-sm sm:px-6">
  <Link to="/" class="flex shrink-0 items-center gap-3" aria-label="PerpusFun, halaman utama">
    <img src="/logo-perpusfun.svg" alt="" class="h-10 w-10 shrink-0" />
    <span class="text-lg font-bold tracking-tight text-slate-900">Perpus<span class="text-blue-600">Fun</span></span>
  </Link>

  <div class="order-3 flex w-full items-center gap-6 text-sm font-medium text-slate-600 sm:order-none sm:w-auto">
    {#if $isAuthenticated}
      <Link to="/koleksi" class="transition hover:text-blue-600">Koleksi Buku</Link>
      <Link to="/dashboard" class="transition hover:text-blue-600">Peminjaman</Link>
    {:else}
      <Link to="/" class="transition hover:text-blue-600">Beranda</Link>
      <button type="button" onclick={() => scrollToSection("tentang")} class="transition hover:text-blue-600">Tentang</button>
      <button type="button" onclick={() => scrollToSection("keunggulan")} class="transition hover:text-blue-600">Keunggulan</button>
    {/if}
  </div>

  <div class="ml-auto flex items-center gap-3">
    {#if $isAuthenticated}
      <label class="relative hidden md:block">
        <span class="sr-only">Cari buku</span>
        <input
          type="search"
          bind:value={searchQuery}
          placeholder="Cari buku..."
          class="w-44 rounded-xl border border-slate-200 bg-slate-50 py-2 pl-3 pr-9 text-sm text-slate-700 outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:bg-white focus:ring-4 focus:ring-blue-100 lg:w-56"
        />
        <span aria-hidden="true" class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400">⌕</span>
      </label>
      <Link to="/dashboard" class="hidden text-sm font-medium text-slate-600 transition hover:text-blue-600 sm:block">
        {$currentMember?.full_name?.split(" ")[0] || "Akun"}
      </Link>
      <button
        type="button"
        onclick={logout}
        class="rounded-xl border border-slate-200 px-4 py-2 text-sm font-semibold text-slate-700 transition hover:border-blue-200 hover:bg-blue-50 hover:text-blue-700"
      >
        Keluar
      </button>
    {:else}
      <Link to="/login" class="rounded-xl bg-blue-600 px-4 py-2 text-sm font-semibold text-white transition hover:bg-blue-700">
        Masuk
      </Link>
    {/if}
  </div>
</nav>
