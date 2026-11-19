<script>
  import { Link, navigate } from "svelte-routing";
  import { currentMember, isAuthenticated } from "../lib/auth.js";
  import { loansAPI } from "../lib/api.js";

  let { buku, onclose, onborrowed } = $props();

  let isBorrowing = $state(false);
  let borrowError = $state("");
  let borrowSuccess = $state(false);
  let available = $state(0);

  $effect(() => {
    available = buku.available;
  });

  function tutupModal() {
    onclose?.();
  }

  async function pinjamBuku() {
    borrowError = "";
    borrowSuccess = false;

    if (!$isAuthenticated || !$currentMember?.id) {
      tutupModal();
      navigate("/login");
      return;
    }

    isBorrowing = true;
    try {
      await loansAPI.borrow(buku.id, $currentMember.id);
      borrowSuccess = true;
      available -= 1;
      onborrowed?.();
    } catch (error) {
      borrowError = error.data?.error || error.message || "Gagal meminjam buku.";
    } finally {
      isBorrowing = false;
    }
  }
</script>

<div class="fixed inset-0 z-50 flex items-center justify-center overflow-y-auto p-4">
  <button
    type="button"
    aria-label="Tutup detail buku"
    class="absolute inset-0 bg-slate-950/45 backdrop-blur-sm"
    onclick={tutupModal}
  ></button>

  <div
    role="dialog"
    aria-modal="true"
    aria-labelledby="book-modal-title"
    tabindex="-1"
    class="relative z-10 my-auto w-full max-w-xl rounded-3xl border border-slate-200 bg-white p-6 shadow-2xl sm:p-8"
  >
    <button
      type="button"
      aria-label="Tutup"
      class="absolute right-5 top-5 flex h-9 w-9 items-center justify-center rounded-full bg-slate-100 text-slate-500 transition hover:bg-slate-200 hover:text-slate-900"
      onclick={tutupModal}
    >
      ✕
    </button>

    <div class="flex flex-col gap-5 pr-8 sm:flex-row">
      <div class="h-40 w-32 shrink-0 overflow-hidden rounded-xl bg-blue-50">
        {#if buku.cover_url}
          <img src={buku.cover_url} alt={`Sampul ${buku.title}`} class="h-full w-full object-cover" />
        {:else}
          <div class="flex h-full items-center justify-center text-5xl text-blue-300" aria-hidden="true">▤</div>
        {/if}
      </div>
      <div class="pt-1">
        <p class="text-xs font-semibold uppercase tracking-wider text-blue-600">Informasi buku</p>
        <h2 id="book-modal-title" class="mt-2 text-2xl font-bold leading-snug text-slate-900">{buku.title}</h2>
        <p class="mt-2 text-sm text-slate-500">ISBN: {buku.isbn || "Tidak tersedia"}</p>
        <span class="mt-4 inline-flex rounded-full px-3 py-1 text-xs font-semibold {available > 0 ? 'bg-blue-50 text-blue-700' : 'bg-slate-100 text-slate-600'}">
          {available > 0 ? `${available} salinan tersedia` : "Stok habis"}
        </span>
      </div>
    </div>

    <dl class="mt-6 grid grid-cols-2 gap-3 rounded-2xl bg-slate-50 p-4 text-sm">
      <div>
        <dt class="text-slate-500">Total stok</dt>
        <dd class="mt-1 font-semibold text-slate-900">{buku.stock}</dd>
      </div>
      <div>
        <dt class="text-slate-500">Tersedia</dt>
        <dd class="mt-1 font-semibold text-slate-900">{available}</dd>
      </div>
    </dl>

    {#if borrowError}
      <p role="alert" class="mt-4 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">{borrowError}</p>
    {/if}
    {#if borrowSuccess}
      <p role="status" class="mt-4 rounded-xl bg-blue-50 px-4 py-3 text-sm text-blue-700">Buku berhasil dipinjam.</p>
    {/if}

    <div class="mt-7 flex flex-wrap justify-end gap-3">
      <button
        type="button"
        onclick={tutupModal}
        class="rounded-xl border border-slate-200 px-4 py-2.5 text-sm font-semibold text-slate-700 transition hover:bg-slate-50"
      >
        Tutup
      </button>
      <Link
        to={`/books/${buku.id}`}
        onclick={tutupModal}
        class="rounded-xl border border-blue-200 px-4 py-2.5 text-sm font-semibold text-blue-700 transition hover:bg-blue-50"
      >
        Detail buku
      </Link>
      <button
        type="button"
        onclick={pinjamBuku}
        disabled={available < 1 || isBorrowing || borrowSuccess}
        class="rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:bg-slate-300"
      >
        {isBorrowing ? "Memproses..." : borrowSuccess ? "Berhasil dipinjam" : "Pinjam buku"}
      </button>
    </div>
  </div>
</div>
