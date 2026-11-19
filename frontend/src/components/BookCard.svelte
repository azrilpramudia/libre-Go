<script>
  let { buku, onbukaDetail } = $props();
</script>

<article class="group flex h-full flex-col overflow-hidden rounded-2xl border border-slate-200 bg-white p-3 shadow-sm transition hover:-translate-y-1 hover:shadow-lg hover:shadow-slate-200/70">
  <div class="relative aspect-[3/4] overflow-hidden rounded-xl bg-blue-50">
    {#if buku.cover_url}
      <img
        src={buku.cover_url}
        alt={`Sampul ${buku.title}`}
        loading="lazy"
        class="h-full w-full object-cover transition duration-300 group-hover:scale-[1.03]"
      />
    {:else}
      <div class="flex h-full items-center justify-center text-5xl text-blue-300" aria-hidden="true">▤</div>
    {/if}
    <span class="absolute left-3 top-3 rounded-full bg-white/95 px-2.5 py-1 text-xs font-semibold {buku.available > 0 ? 'text-blue-700' : 'text-slate-500'}">
      {buku.available > 0 ? `${buku.available} tersedia` : "Habis"}
    </span>
  </div>

  <div class="flex flex-1 flex-col px-2 pb-2 pt-4">
    <h3 class="line-clamp-2 min-h-12 text-base font-semibold leading-6 text-slate-900">{buku.title}</h3>
    <p class="mt-1 truncate text-xs text-slate-500">ISBN: {buku.isbn || "Tidak tersedia"}</p>
    <div class="mt-4 flex items-center justify-between border-t border-slate-100 pt-3">
      <span class="text-xs text-slate-500">Stok {buku.stock ?? 0}</span>
      <button
        type="button"
        onclick={() => onbukaDetail?.(buku)}
        class="rounded-lg bg-blue-50 px-3 py-2 text-xs font-semibold text-blue-700 transition hover:bg-blue-600 hover:text-white"
      >
        Lihat detail
      </button>
    </div>
  </div>
</article>
