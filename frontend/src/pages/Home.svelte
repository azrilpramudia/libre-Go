<script>
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from "../components/BookModal.svelte";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import { booksAPI } from "../lib/api.js";
  import { currentMember, isAuthenticated } from "../lib/auth.js";

  let bukuPopuler = $state([]);
  let searchQuery = $state("");
  let isLoadingBooks = $state(false);
  let booksError = $state("");
  let isAuth = $derived($isAuthenticated);
  let filteredBuku = $derived(
    bukuPopuler.filter((buku) =>
      buku.title?.toLowerCase().includes(searchQuery.toLowerCase())
    )
  );
  let showModal = $state(false);
  let selectedBuku = $state(null);

  $effect(() => {
    if (isAuth) {
      loadBooks();
    } else {
      bukuPopuler = [];
      searchQuery = "";
      booksError = "";
      isLoadingBooks = false;
    }
  });

  async function loadBooks() {
    isLoadingBooks = true;
    booksError = "";
    try {
      const response = await booksAPI.list(8);
      bukuPopuler = response.data || [];
    } catch (error) {
      booksError = error.data?.error || error.message || "Gagal memuat koleksi buku.";
    } finally {
      isLoadingBooks = false;
    }
  }

  function handleBukaDetail(buku) {
    selectedBuku = buku;
    showModal = true;
  }
</script>

<main class="min-h-screen overflow-hidden bg-slate-50 text-slate-900">
  <div class="mx-auto w-full max-w-7xl px-5 py-5 sm:px-8">
    <Navbar bind:searchQuery />

    {#if isAuth}
      <section class="mt-10 rounded-3xl border border-blue-100 bg-white p-7 shadow-sm sm:p-10">
        <p class="text-sm font-semibold text-blue-600">Perpustakaan digital</p>
        <div class="mt-2 flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <h1 class="text-3xl font-bold tracking-tight text-slate-900 sm:text-4xl">
              Halo, {$currentMember?.full_name || "Member"}
            </h1>
            <p class="mt-2 text-slate-600">Temukan bacaan baru dari koleksi perpustakaan.</p>
          </div>
          <Link to="/koleksi" class="inline-flex w-fit items-center gap-2 rounded-xl bg-blue-600 px-5 py-3 text-sm font-semibold text-white transition hover:bg-blue-700">
            Jelajahi semua buku <span aria-hidden="true">→</span>
          </Link>
        </div>
      </section>

      <section class="py-12">
        <div class="mb-6 flex flex-wrap items-end justify-between gap-4">
          <div>
            <p class="text-sm font-semibold text-blue-600">Pilihan untukmu</p>
            <h2 class="mt-1 text-2xl font-bold tracking-tight text-slate-900">Koleksi buku</h2>
          </div>
          <Link to="/koleksi" class="text-sm font-semibold text-blue-700 transition hover:text-blue-800">Lihat semua</Link>
        </div>

        <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
          {#if isLoadingBooks}
            <p class="col-span-full rounded-2xl border border-slate-200 bg-white p-8 text-center text-sm text-slate-500">Memuat koleksi buku...</p>
          {:else if booksError}
            <div role="alert" class="col-span-full rounded-2xl border border-red-100 bg-red-50 p-5 text-sm text-red-700">{booksError}</div>
          {:else if filteredBuku.length > 0}
            {#each filteredBuku as buku (buku.id)}
              <BookCard {buku} onbukaDetail={handleBukaDetail} />
            {/each}
          {:else}
            <div class="col-span-full rounded-2xl border border-slate-200 bg-white px-6 py-14 text-center">
              <h3 class="font-semibold text-slate-800">{searchQuery.trim() ? "Buku tidak ditemukan" : "Koleksi belum tersedia"}</h3>
              <p class="mt-2 text-sm text-slate-500">{searchQuery.trim() ? "Coba kata pencarian lain." : "Silakan kembali lagi nanti."}</p>
            </div>
          {/if}
        </div>
      </section>
    {:else}
      <section class="grid items-center gap-10 py-14 md:grid-cols-2 md:py-20">
        <div class="max-w-xl">
          <span class="inline-flex rounded-full border border-blue-100 bg-blue-50 px-3 py-1.5 text-xs font-semibold text-blue-700">
            Ruang baru untuk membaca
          </span>
          <h1 class="mt-6 text-4xl font-bold leading-tight tracking-tight text-slate-900 sm:text-5xl lg:text-6xl">
            Temukan cerita <span class="text-blue-600">berikutnya.</span>
          </h1>
          <p class="mt-5 max-w-lg text-base leading-7 text-slate-600">
            Jelajahi koleksi buku, lihat ketersediaan, dan kelola peminjaman dengan mudah dalam satu perpustakaan digital.
          </p>
          <div class="mt-8 flex flex-wrap gap-3">
            <Link to="/koleksi" class="rounded-xl bg-blue-600 px-5 py-3 text-sm font-semibold text-white shadow-sm transition hover:bg-blue-700">
              Jelajahi koleksi <span aria-hidden="true">→</span>
            </Link>
            <Link to="/daftar" class="rounded-xl border border-slate-300 bg-white px-5 py-3 text-sm font-semibold text-slate-700 transition hover:border-blue-200 hover:bg-blue-50">
              Buat akun
            </Link>
          </div>
          <div class="mt-10 flex gap-8 border-t border-slate-200 pt-6">
            <div>
              <p class="text-xl font-bold text-slate-900">24/7</p>
              <p class="mt-1 text-xs text-slate-500">Akses katalog</p>
            </div>
            <div class="h-10 w-px bg-slate-200"></div>
            <div>
              <p class="text-xl font-bold text-slate-900">Mudah</p>
              <p class="mt-1 text-xs text-slate-500">Kelola peminjaman</p>
            </div>
          </div>
        </div>

        <div class="relative mx-auto w-full max-w-lg">
          <div class="absolute -inset-4 rounded-[2rem] bg-blue-100/70"></div>
          <div class="relative overflow-hidden rounded-[1.75rem] bg-blue-600 p-7 text-white sm:p-10">
            <div class="flex items-start justify-between">
              <div>
                <p class="text-sm font-medium text-blue-100">PerpusFun</p>
                <p class="mt-2 text-2xl font-bold leading-snug sm:text-3xl">Banyak ide bermula dari satu halaman.</p>
              </div>
              <span class="ml-4 flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-white/15 text-2xl" aria-hidden="true">▤</span>
            </div>
            <div class="mt-10 grid grid-cols-3 gap-3" aria-hidden="true">
              <div class="h-32 rounded-xl bg-white/15"></div>
              <div class="mt-5 h-32 rounded-xl bg-white/25"></div>
              <div class="h-32 rounded-xl bg-white/10"></div>
            </div>
            <div class="mt-6 flex items-center justify-between border-t border-white/20 pt-5 text-sm text-blue-100">
              <span>Koleksi yang terus bertambah</span>
              <span aria-hidden="true">✦</span>
            </div>
          </div>
        </div>
      </section>

      <section id="tentang" class="scroll-mt-8 border-y border-slate-200 py-14">
        <div class="max-w-3xl">
          <p class="text-sm font-semibold text-blue-600">Tentang PerpusFun</p>
          <h2 class="mt-2 text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">Membaca jadi lebih mudah diakses.</h2>
          <p class="mt-4 leading-7 text-slate-600">
            PerpusFun membantu pembaca menemukan buku, memeriksa ketersediaan, dan mengelola peminjaman tanpa proses yang rumit.
          </p>
        </div>
      </section>

      <section id="keunggulan" class="scroll-mt-8 py-14">
        <div class="mb-7 max-w-2xl">
          <p class="text-sm font-semibold text-blue-600">Dirancang untuk pembaca</p>
          <h2 class="mt-2 text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">Semua yang dibutuhkan, tanpa kerumitan.</h2>
        </div>
        <div class="grid gap-4 md:grid-cols-3">
          <article class="rounded-2xl border border-slate-200 bg-white p-6">
            <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-50 font-semibold text-blue-700" aria-hidden="true">01</span>
            <h3 class="mt-5 font-semibold text-slate-900">Cari dengan mudah</h3>
            <p class="mt-2 text-sm leading-6 text-slate-600">Temukan judul buku dan informasi ketersediaan dalam satu tempat.</p>
          </article>
          <article class="rounded-2xl border border-slate-200 bg-white p-6">
            <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-50 font-semibold text-blue-700" aria-hidden="true">02</span>
            <h3 class="mt-5 font-semibold text-slate-900">Kelola peminjaman</h3>
            <p class="mt-2 text-sm leading-6 text-slate-600">Pantau buku yang dipinjam dan tanggal pengembaliannya.</p>
          </article>
          <article class="rounded-2xl border border-slate-200 bg-white p-6">
            <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-50 font-semibold text-blue-700" aria-hidden="true">03</span>
            <h3 class="mt-5 font-semibold text-slate-900">Akses kapan saja</h3>
            <p class="mt-2 text-sm leading-6 text-slate-600">Jelajahi katalog perpustakaan dari perangkat yang kamu gunakan.</p>
          </article>
        </div>
      </section>
    {/if}
  </div>

  {#if showModal}
    <BookModal buku={selectedBuku} onclose={() => (showModal = false)} onborrowed={loadBooks} />
  {/if}

  <Footer />
</main>
