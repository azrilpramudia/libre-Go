<script>
  import { Link } from "svelte-routing";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import { booksAPI } from "../lib/api.js";

  let { id } = $props();
  let book = $state(null);
  let isLoading = $state(true);
  let errorMessage = $state("");

  $effect(() => {
    const bookId = id;
    let active = true;

    async function loadBook() {
      isLoading = true;
      errorMessage = "";
      book = null;
      try {
        const data = await booksAPI.getById(bookId);
        if (active) book = data.data || null;
      } catch (error) {
        if (active) {
          errorMessage = error.data?.error || error.message || "Gagal memuat detail buku.";
        }
      } finally {
        if (active) isLoading = false;
      }
    }

    if (bookId) loadBook();
    return () => {
      active = false;
    };
  });

  let readerUrl = $derived(
    book?.openlibrary_id
      ? `https://openlibrary.org${book.openlibrary_id}?embed=true`
      : "",
  );
</script>

<main class="min-h-screen bg-slate-50 text-slate-900">
  <div class="mx-auto max-w-7xl px-5 py-5 sm:px-8">
    <Navbar />
    <header class="mb-7 mt-10 flex flex-wrap items-end justify-between gap-4">
      <div>
        <p class="text-sm font-semibold text-blue-600">PerpusFun</p>
        <h1 class="mt-1 text-3xl font-bold tracking-tight text-slate-900">Detail buku</h1>
      </div>
      <Link to="/koleksi" class="rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-semibold text-slate-700 transition hover:border-blue-200 hover:bg-blue-50 hover:text-blue-700">
        Kembali ke koleksi
      </Link>
    </header>

    {#if isLoading}
      <p class="rounded-2xl border border-slate-200 bg-white p-8 text-center text-sm text-slate-500">Memuat detail buku...</p>
    {:else if errorMessage}
      <p role="alert" class="rounded-2xl border border-red-100 bg-red-50 p-6 text-center text-sm text-red-700">{errorMessage}</p>
    {:else if book}
      <section class="flex flex-col gap-7 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm sm:flex-row sm:p-8">
        <div class="h-72 w-full shrink-0 overflow-hidden rounded-xl bg-blue-50 sm:w-52">
          {#if book.cover_url}
            <img src={book.cover_url} alt={`Sampul ${book.title}`} class="h-full w-full object-cover" />
          {:else}
            <div class="flex h-full items-center justify-center text-7xl text-blue-300" aria-hidden="true">▤</div>
          {/if}
        </div>
        <div class="flex-1">
          <p class="text-sm font-semibold text-blue-600">Informasi buku</p>
          <h2 class="mt-2 text-2xl font-bold leading-snug text-slate-900 sm:text-3xl">{book.title}</h2>
          <dl class="mt-6 grid gap-4 border-t border-slate-100 pt-5 text-sm sm:grid-cols-2">
            <div><dt class="text-slate-500">ISBN</dt><dd class="mt-1 font-medium text-slate-800">{book.isbn || "Tidak tersedia"}</dd></div>
            <div><dt class="text-slate-500">Total stok</dt><dd class="mt-1 font-medium text-slate-800">{book.stock}</dd></div>
            <div><dt class="text-slate-500">Tersedia</dt><dd class="mt-1 font-medium text-slate-800">{book.available}</dd></div>
          </dl>
        </div>
      </section>

      <section class="mt-10">
        <h2 class="mb-4 text-xl font-bold text-slate-900">Baca buku</h2>
        {#if readerUrl}
          <iframe
            title={`Baca ${book.title} di OpenLibrary`}
            src={readerUrl}
            class="h-[75vh] min-h-[480px] w-full rounded-2xl border border-slate-200 bg-white shadow-sm"
            allow="fullscreen"
          ></iframe>
          <p class="mt-3 text-sm leading-6 text-slate-500">
            Pratinjau dan akses baca bergantung pada ketersediaan edisi di OpenLibrary. Jika tidak tampil, kunjungi
            <a href={`https://openlibrary.org${book.openlibrary_id}`} target="_blank" rel="noreferrer" class="font-semibold text-blue-700 underline">
              OpenLibrary
            </a>.
          </p>
        {:else}
          <p class="rounded-2xl border border-slate-200 bg-white p-6 text-sm text-slate-500">
            Pratinjau baca OpenLibrary tidak tersedia untuk buku ini.
          </p>
        {/if}
      </section>
    {/if}
  </div>
  <Footer />
</main>
