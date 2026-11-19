<script>
  import { onMount } from "svelte";
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from "../components/BookModal.svelte";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import { booksAPI } from "../lib/api.js";

  const pageSize = 20;
  let bukuTersedia = $state([]);
  let selectedBuku = $state(null);
  let isLoading = $state(false);
  let hasMore = $state(true);
  let errorMessage = $state("");

  onMount(() => loadBooks({ reset: true }));

  async function loadBooks({ reset = false } = {}) {
    if (isLoading) return;

    isLoading = true;
    errorMessage = "";
    try {
      const offset = reset ? 0 : bukuTersedia.length;
      const response = await booksAPI.list(pageSize, offset);
      const page = response.data || [];

      bukuTersedia = reset ? page : [...bukuTersedia, ...page];
      hasMore = page.length === pageSize;
    } catch (error) {
      errorMessage = error.data?.error || error.message || "Gagal memuat koleksi buku.";
    } finally {
      isLoading = false;
    }
  }

  function refreshBooks() {
    bukuTersedia = [];
    hasMore = true;
    loadBooks({ reset: true });
  }
</script>

<main class="min-h-screen bg-slate-50 text-slate-900">
  <div class="mx-auto max-w-7xl px-5 py-5 sm:px-8">
    <Navbar />
    <header class="mb-8 mt-10 flex flex-wrap items-end justify-between gap-5">
      <div>
        <p class="text-sm font-semibold text-blue-600">PerpusFun</p>
        <h1 class="mt-1 text-3xl font-bold tracking-tight text-slate-900">Koleksi buku</h1>
        {#if bukuTersedia.length}
          <p class="mt-2 text-sm text-slate-500">Menampilkan {bukuTersedia.length} buku dari koleksi.</p>
        {/if}
      </div>
      <Link to="/books/import" class="rounded-xl border border-blue-200 bg-white px-4 py-2.5 text-sm font-semibold text-blue-700 transition hover:bg-blue-50">
        Import dari OpenLibrary
      </Link>
    </header>

    {#if isLoading && bukuTersedia.length === 0}
      <p class="rounded-2xl border border-slate-200 bg-white py-12 text-center text-sm text-slate-500">Memuat koleksi buku...</p>
    {:else if errorMessage && bukuTersedia.length === 0}
      <div role="alert" class="rounded-2xl border border-red-100 bg-red-50 p-5 text-center text-sm text-red-700">
        <p>{errorMessage}</p>
        <button onclick={() => loadBooks({ reset: true })} class="mt-3 font-semibold underline">Coba lagi</button>
      </div>
    {:else if bukuTersedia.length}
      <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        {#each bukuTersedia as buku (buku.id)}
          <BookCard {buku} onbukaDetail={(selected) => (selectedBuku = selected)} />
        {/each}
      </div>
      {#if errorMessage}
        <div role="alert" class="mt-6 rounded-2xl border border-red-100 bg-red-50 p-4 text-center text-sm text-red-700">
          <p>{errorMessage}</p>
          <button onclick={() => loadBooks()} class="mt-2 font-semibold underline">Coba lagi</button>
        </div>
      {/if}
      {#if hasMore}
        <div class="mt-8 flex justify-center">
          <button
            onclick={() => loadBooks()}
            disabled={isLoading}
            class="rounded-xl bg-blue-600 px-5 py-3 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-wait disabled:opacity-60"
          >
            {isLoading ? "Memuat buku..." : "Muat lebih banyak"}
          </button>
        </div>
      {:else}
        <p class="mt-8 text-center text-sm text-slate-500">Semua buku sudah ditampilkan.</p>
      {/if}
    {:else}
      <p class="rounded-2xl border border-slate-200 bg-white p-10 text-center text-sm text-slate-500">Belum ada buku di koleksi.</p>
    {/if}
  </div>

  {#if selectedBuku}
    <BookModal buku={selectedBuku} onclose={() => (selectedBuku = null)} onborrowed={refreshBooks} />
  {/if}
  <Footer />
</main>
