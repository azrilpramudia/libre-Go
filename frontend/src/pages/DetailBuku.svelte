<script>
  import { Link } from "svelte-routing";
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
        console.log(data);
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

<main class="min-h-screen bg-[#f4f9f6] p-6 text-gray-800 sm:p-8">
  <div class="mx-auto max-w-5xl">
    <header class="mb-8 flex flex-wrap items-center justify-between gap-4">
      <h1 class="text-3xl font-extrabold text-gray-900">Detail Buku</h1>
      <Link to="/koleksi" class="rounded-full bg-[#9B9AFF] px-5 py-2.5 font-bold text-white transition hover:bg-[#8382ff]">
        ← Kembali ke Koleksi
      </Link>
    </header>

    {#if isLoading}
      <p class="rounded-3xl bg-white p-8 text-center text-gray-500">Memuat detail buku...</p>
    {:else if errorMessage}
      <p role="alert" class="rounded-3xl bg-red-100 p-6 text-center text-red-700">{errorMessage}</p>
    {:else if book}
      <section class="flex flex-col gap-8 rounded-3xl bg-white p-6 shadow-sm sm:flex-row sm:p-8">
        <div class="h-72 w-full shrink-0 overflow-hidden rounded-2xl bg-[#AADEFF]/30 sm:w-52">
          {#if book.cover_url}
            <img src={book.cover_url} alt={`Sampul ${book.title}`} class="h-full w-full object-cover" />
          {:else}
            <div class="flex h-full items-center justify-center text-7xl">📖</div>
          {/if}
        </div>
        <div class="flex-1">
          <h2 class="text-3xl font-extrabold text-gray-900">{book.title}</h2>
          <dl class="mt-6 space-y-3 text-gray-600">
            <div><dt class="inline font-bold text-gray-800">ISBN: </dt><dd class="inline">{book.isbn || "Tidak tersedia"}</dd></div>
            <div><dt class="inline font-bold text-gray-800">Stok: </dt><dd class="inline">{book.stock}</dd></div>
            <div><dt class="inline font-bold text-gray-800">Tersedia: </dt><dd class="inline">{book.available}</dd></div>
          </dl>
        </div>
      </section>

      <section class="mt-10">
        <h2 class="mb-4 text-2xl font-extrabold text-gray-900">Baca Buku</h2>
        {#if readerUrl}
          <iframe
            title={`Baca ${book.title} di OpenLibrary`}
            src={readerUrl}
            class="h-[75vh] min-h-[480px] w-full rounded-3xl border border-gray-200 bg-white shadow-sm"
            allow="fullscreen"
          ></iframe>
          <p class="mt-3 text-sm text-gray-500">
            Pratinjau dan akses baca bergantung pada ketersediaan edisi di OpenLibrary. Jika tidak tampil, kunjungi
            <a href={`https://openlibrary.org${book.openlibrary_id}`} target="_blank" rel="noreferrer" class="font-semibold text-[#8382ff] underline">
              OpenLibrary
            </a>.
          </p>
        {:else}
          <p class="rounded-2xl bg-white p-6 text-gray-500">
            Pratinjau baca OpenLibrary tidak tersedia untuk buku ini.
          </p>
        {/if}
      </section>
    {/if}
  </div>
</main>
