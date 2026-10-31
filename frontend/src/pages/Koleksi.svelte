<script>
  import { onMount } from "svelte";
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from "../components/BookModal.svelte";
  import { booksAPI } from "../lib/api.js";

  let bukuTersedia = $state([]);
  let selectedBuku = $state(null);
  let isLoading = $state(true);
  let errorMessage = $state("");

  onMount(loadBooks);

  async function loadBooks() {
    isLoading = true;
    errorMessage = "";
    try {
      const response = await booksAPI.list(50);
      bukuTersedia = response.data || [];
    } catch (error) {
      errorMessage = error.data?.error || error.message || "Gagal memuat koleksi buku.";
    } finally {
      isLoading = false;
    }
  }
</script>

<main class="min-h-screen bg-[#f4f9f6] p-8 text-gray-800">
  <div class="max-w-7xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-extrabold text-gray-900">📚 Koleksi Buku</h1>
      <Link to="/" class="bg-[#9B9AFF] text-white px-4 py-2 rounded-full font-bold hover:bg-[#8382ff] transition">
        ← Kembali ke Home
      </Link>
    </div>

    {#if isLoading}
      <p class="py-12 text-center text-gray-500">Memuat koleksi buku...</p>
    {:else if errorMessage}
      <div role="alert" class="rounded-2xl bg-red-100 p-5 text-center text-red-700">
        <p>{errorMessage}</p>
        <button onclick={loadBooks} class="mt-3 font-bold underline">Coba lagi</button>
      </div>
    {:else if bukuTersedia.length}
      <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {#each bukuTersedia as buku (buku.id)}
          <BookCard {buku} onbukaDetail={(selected) => (selectedBuku = selected)} />
        {/each}
      </div>
    {:else}
      <p class="rounded-2xl bg-white p-8 text-center text-gray-500">Belum ada buku di koleksi.</p>
    {/if}
  </div>

  {#if selectedBuku}
    <BookModal buku={selectedBuku} onclose={() => (selectedBuku = null)} onborrowed={loadBooks} />
  {/if}
</main>
