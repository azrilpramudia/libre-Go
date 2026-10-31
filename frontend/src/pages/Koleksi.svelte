<script>
  import { onMount } from "svelte";
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from "../components/BookModal.svelte";
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

<main class="min-h-screen bg-[#f4f9f6] p-8 text-gray-800">
  <div class="max-w-7xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <div>
        <h1 class="text-3xl font-extrabold text-gray-900">📚 Koleksi Buku</h1>
        {#if bukuTersedia.length}
          <p class="mt-2 text-sm text-gray-500">Menampilkan {bukuTersedia.length} buku dari koleksi.</p>
        {/if}
      </div>
      <Link to="/" class="bg-[#9B9AFF] text-white px-4 py-2 rounded-full font-bold hover:bg-[#8382ff] transition">
        ← Kembali ke Home
      </Link>
    </div>
    <div class="mb-6 flex justify-end">
      <Link to="/books/import" class="rounded-full border border-[#9B9AFF] px-5 py-2.5 text-sm font-bold text-[#8382ff] transition hover:bg-[#9B9AFF]/10">
        Import Buku dari OpenLibrary
      </Link>
    </div>

    {#if isLoading && bukuTersedia.length === 0}
      <p class="py-12 text-center text-gray-500">Memuat koleksi buku...</p>
    {:else if errorMessage && bukuTersedia.length === 0}
      <div role="alert" class="rounded-2xl bg-red-100 p-5 text-center text-red-700">
        <p>{errorMessage}</p>
        <button onclick={() => loadBooks({ reset: true })} class="mt-3 font-bold underline">Coba lagi</button>
      </div>
    {:else if bukuTersedia.length}
      <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {#each bukuTersedia as buku (buku.id)}
          <BookCard {buku} onbukaDetail={(selected) => (selectedBuku = selected)} />
        {/each}
      </div>
      {#if errorMessage}
        <div role="alert" class="mt-6 rounded-2xl bg-red-100 p-4 text-center text-red-700">
          <p>{errorMessage}</p>
          <button onclick={() => loadBooks()} class="mt-2 font-bold underline">Coba lagi</button>
        </div>
      {/if}
      {#if hasMore}
        <div class="mt-8 flex justify-center">
          <button
            onclick={() => loadBooks()}
            disabled={isLoading}
            class="rounded-full bg-[#9B9AFF] px-8 py-3 font-bold text-white shadow transition hover:bg-[#8382ff] disabled:cursor-wait disabled:opacity-60"
          >
            {isLoading ? "Memuat buku..." : "Muat lebih banyak"}
          </button>
        </div>
      {:else}
        <p class="mt-8 text-center text-sm text-gray-500">Semua buku sudah ditampilkan.</p>
      {/if}
    {:else}
      <p class="rounded-2xl bg-white p-8 text-center text-gray-500">Belum ada buku di koleksi.</p>
    {/if}
  </div>

  {#if selectedBuku}
    <BookModal buku={selectedBuku} onclose={() => (selectedBuku = null)} onborrowed={refreshBooks} />
  {/if}
</main>
