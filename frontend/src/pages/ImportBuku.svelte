<script>
  import { Link } from "svelte-routing";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
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

<main class="min-h-screen bg-slate-50 text-slate-900">
  <div class="mx-auto max-w-7xl px-5 py-5 sm:px-8">
    <Navbar />
    <header class="mb-7 mt-10">
      <div>
        <p class="text-sm font-semibold text-blue-600">Perluas koleksi</p>
        <h1 class="mt-1 text-3xl font-bold tracking-tight text-slate-900">Import buku</h1>
        <p class="mt-2 text-sm text-slate-500">Cari buku di OpenLibrary lalu tambahkan ke koleksi.</p>
      </div>
      <Link to="/koleksi" class="inline-flex w-fit rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-semibold text-slate-700 transition hover:border-blue-200 hover:bg-blue-50 hover:text-blue-700">
        Kembali ke koleksi
      </Link>
    </header>

    <form onsubmit={searchBooks} class="mb-7 flex flex-col gap-3 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:flex-row sm:p-5">
      <label for="book-search" class="sr-only">Judul buku</label>
      <input
        id="book-search"
        bind:value={title}
        required
        placeholder="Masukkan judul buku..."
        class="min-w-0 flex-1 rounded-xl border border-slate-200 px-4 py-3 text-sm outline-none transition placeholder:text-slate-400 focus:border-blue-400 focus:ring-4 focus:ring-blue-100"
      />
      <button
        type="submit"
        disabled={isSearching || !title.trim()}
        class="rounded-xl bg-blue-600 px-6 py-3 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-wait disabled:opacity-60"
      >
        {isSearching ? "Mencari..." : "Cari Buku"}
      </button>
    </form>

    {#if notification}
      <p role="status" class="mb-5 rounded-xl border border-blue-100 bg-blue-50 px-5 py-4 text-sm font-medium text-blue-800">{notification}</p>
    {/if}
    {#if errorMessage}
      <p role="alert" class="mb-5 rounded-xl border border-red-100 bg-red-50 px-5 py-4 text-sm font-medium text-red-700">{errorMessage}</p>
    {/if}

    {#if results.length}
      <div class="grid gap-4 md:grid-cols-2">
        {#each results as result (result.key)}
          <article class="flex gap-4 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm">
            <div class="h-36 w-24 shrink-0 overflow-hidden rounded-xl bg-blue-50">
              {#if result.cover_i}
                <img src={coverUrl(result.cover_i)} alt={`Sampul ${result.title}`} class="h-full w-full object-cover" />
              {:else}
                <div class="flex h-full items-center justify-center text-4xl text-blue-300" aria-hidden="true">▤</div>
              {/if}
            </div>
            <div class="flex min-w-0 flex-1 flex-col items-start">
              <h2 class="text-base font-semibold leading-6 text-slate-900">{result.title}</h2>
              <p class="mt-1 text-sm text-slate-600">{result.author_name?.join(", ") || "Penulis tidak diketahui"}</p>
              {#if result.isbn?.length}
                <p class="mt-2 text-xs text-slate-500">ISBN: {result.isbn[0]}</p>
              {/if}
              <button
                type="button"
                onclick={() => importBook(result)}
                disabled={!result.key || importingKey !== "" || importedKeys.includes(result.key)}
                class="mt-auto rounded-lg bg-blue-600 px-4 py-2.5 text-xs font-semibold text-white transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-60"
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
      <p class="rounded-2xl border border-slate-200 bg-white p-8 text-center text-sm text-slate-500">Belum ada hasil. Silakan cari judul buku.</p>
    {/if}

    {#if importedBooks.length}
      <section class="mt-12">
        <h2 class="mb-4 text-xl font-bold text-slate-900">Buku yang baru diimpor</h2>
        <ul class="space-y-3">
          {#each importedBooks as book (book.id)}
            <li class="rounded-xl border border-slate-200 bg-white p-4">
              <span class="font-bold">{book.title}</span>
              <span class="ml-2 text-sm text-slate-500">Stok: {book.stock}</span>
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  </div>
  <Footer />
</main>
