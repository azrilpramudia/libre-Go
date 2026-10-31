<script>
  import { Link } from "svelte-routing";
  import { booksAPI } from "../lib/api.js";

  let title = $state("");
  let results = $state([]);
  let importedBooks = $state([]);
  let importedKeys = $state([]);
  let isSearching = $state(false);
  let importingKey = $state("");
  let errorMessage = $state("");
  let notification = $state("");

  async function searchBooks(event) {
    event.preventDefault();
    const query = title.trim();
    if (!query || isSearching) return;

    isSearching = true;
    errorMessage = "";
    notification = "";
    try {
      const data = await booksAPI.search(query);
      console.log(data);
      results = Array.isArray(data.data?.docs) ? data.data.docs : [];
    } catch (error) {
      errorMessage = error.data?.error || error.message || "Gagal mencari buku di OpenLibrary.";
      results = [];
    } finally {
      isSearching = false;
    }
  }

  async function importBook(result) {
    if (!result.key || importingKey) return;

    importingKey = result.key;
    errorMessage = "";
    notification = "";

    try {
      const data = await booksAPI.import(
        result.title,
        result.author_name?.join(", ") || "",
        result.isbn?.[0] || "",
        result.cover_i || 0,
        result.key,
      );
      console.log(data);
      importedBooks = data.data ? [data.data, ...importedBooks] : importedBooks;
      importedKeys = [...importedKeys, result.key];
      notification = `"${result.title}" berhasil diimpor ke database.`;
    } catch (error) {
      errorMessage = error.data?.error || error.message || `Gagal mengimpor "${result.title}".`;
    } finally {
      importingKey = "";
    }
  }

  function coverUrl(coverId) {
    return coverId
      ? `https://covers.openlibrary.org/b/id/${coverId}-M.jpg`
      : "";
  }
</script>

<main class="min-h-screen bg-[#f4f9f6] p-6 text-gray-800 sm:p-8">
  <div class="mx-auto max-w-6xl">
    <header class="mb-8 flex flex-wrap items-center justify-between gap-4">
      <div>
        <h1 class="text-3xl font-extrabold text-gray-900">Import Buku</h1>
        <p class="mt-2 text-sm text-gray-500">Cari buku di OpenLibrary lalu tambahkan ke koleksi.</p>
      </div>
      <Link to="/koleksi" class="rounded-full bg-[#9B9AFF] px-5 py-2.5 font-bold text-white transition hover:bg-[#8382ff]">
        ← Kembali ke Koleksi
      </Link>
    </header>

    <form onsubmit={searchBooks} class="mb-6 flex flex-col gap-3 rounded-3xl bg-white p-5 shadow-sm sm:flex-row">
      <label for="book-search" class="sr-only">Judul buku</label>
      <input
        id="book-search"
        bind:value={title}
        required
        placeholder="Masukkan judul buku..."
        class="min-w-0 flex-1 rounded-full border border-gray-200 px-5 py-3 outline-none transition focus:border-[#9B9AFF]"
      />
      <button
        type="submit"
        disabled={isSearching || !title.trim()}
        class="rounded-full bg-[#9B9AFF] px-7 py-3 font-bold text-white transition hover:bg-[#8382ff] disabled:cursor-wait disabled:opacity-60"
      >
        {isSearching ? "Mencari..." : "Cari Buku"}
      </button>
    </form>

    {#if notification}
      <p role="status" class="mb-5 rounded-2xl bg-green-100 px-5 py-4 font-medium text-green-800">{notification}</p>
    {/if}
    {#if errorMessage}
      <p role="alert" class="mb-5 rounded-2xl bg-red-100 px-5 py-4 font-medium text-red-700">{errorMessage}</p>
    {/if}

    {#if results.length}
      <div class="grid gap-5 md:grid-cols-2">
        {#each results as result (result.key)}
          <article class="flex gap-4 rounded-3xl border border-gray-100 bg-white p-5 shadow-sm">
            <div class="h-36 w-24 shrink-0 overflow-hidden rounded-xl bg-[#AADEFF]/30">
              {#if result.cover_i}
                <img src={coverUrl(result.cover_i)} alt={`Sampul ${result.title}`} class="h-full w-full object-cover" />
              {:else}
                <div class="flex h-full items-center justify-center text-4xl">📖</div>
              {/if}
            </div>
            <div class="flex min-w-0 flex-1 flex-col items-start">
              <h2 class="text-lg font-bold text-gray-900">{result.title}</h2>
              <p class="mt-1 text-sm text-gray-600">{result.author_name?.join(", ") || "Penulis tidak diketahui"}</p>
              {#if result.isbn?.length}
                <p class="mt-2 text-xs text-gray-500">ISBN: {result.isbn[0]}</p>
              {/if}
              <button
                type="button"
                onclick={() => importBook(result)}
                disabled={!result.key || importingKey !== "" || importedKeys.includes(result.key)}
                class="mt-auto rounded-full bg-[#9B9AFF] px-5 py-2.5 text-sm font-bold text-white transition hover:bg-[#8382ff] disabled:cursor-not-allowed disabled:opacity-60"
              >
                {#if importingKey === result.key}
                  Mengimpor...
                {:else if importedKeys.includes(result.key)}
                  Sudah Diimpor
                {:else}
                  Import ke Database
                {/if}
              </button>
            </div>
          </article>
        {/each}
      </div>
    {:else if !isSearching && title.trim() && !errorMessage}
      <p class="rounded-3xl bg-white p-8 text-center text-gray-500">Belum ada hasil. Silakan cari judul buku.</p>
    {/if}

    {#if importedBooks.length}
      <section class="mt-12">
        <h2 class="mb-4 text-2xl font-extrabold text-gray-900">Buku yang baru diimpor</h2>
        <ul class="space-y-3">
          {#each importedBooks as book (book.id)}
            <li class="rounded-2xl bg-white p-4 shadow-sm">
              <span class="font-bold">{book.title}</span>
              <span class="ml-2 text-sm text-gray-500">Stok: {book.stock}</span>
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  </div>
</main>
