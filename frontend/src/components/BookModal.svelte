<script>
  import { navigate } from 'svelte-routing';
  import { currentMember, isAuthenticated } from '../lib/auth.js';
  import { loansAPI } from '../lib/api.js';

  let { buku, onclose, onborrowed } = $props();

  let isBorrowing = $state(false);
  let borrowError = $state('');
  let borrowSuccess = $state(false);
  let available = $state(0);

  $effect(() => {
    available = buku.available;
  });

  function tutupModal() {
    onclose?.();
  }

  async function pinjamBuku() {
    borrowError = '';
    borrowSuccess = false;

    if (!$isAuthenticated || !$currentMember?.id) {
      tutupModal();
      navigate('/login');
      return;
    }

    isBorrowing = true;
    try {
      await loansAPI.borrow(buku.id, $currentMember.id);
      borrowSuccess = true;
      available -= 1;
      onborrowed?.();
    } catch (error) {
      borrowError = error.data?.error || error.message || 'Gagal meminjam buku.';
    } finally {
      isBorrowing = false;
    }
  }
</script>

<div class="fixed inset-0 z-50 flex items-center justify-center p-4">
  <button
    type="button"
    aria-label="Tutup detail buku"
    class="absolute inset-0 bg-gray-900/40 backdrop-blur-sm"
    onclick={tutupModal}
  ></button>

  <div
    role="dialog"
    aria-modal="true"
    aria-labelledby="book-modal-title"
    tabindex="-1"
    class="relative z-10 bg-white/90 backdrop-blur-xl rounded-3xl p-8 max-w-lg w-full shadow-2xl border border-white animate-fade-in-up"
  >
    
    <!-- Tombol Silang (X) di pojok kanan atas -->
    <button class="absolute top-5 right-5 text-gray-400 hover:text-gray-700 transition" onclick={tutupModal}>
      ✕
    </button>
    
    <!-- Header Modal: Cover & Info Singkat -->
    <div class="flex gap-6 items-center mb-6">
       <div class="bg-[#AADEFF]/30 w-24 h-32 rounded-2xl flex items-center justify-center text-4xl shadow-inner shrink-0 overflow-hidden">
         {#if buku.cover_url}
           <img src={buku.cover_url} alt={`Sampul ${buku.title}`} class="w-full h-full object-cover" />
         {:else}
           📖
         {/if}
       </div>
       <div>
         <h2 id="book-modal-title" class="text-2xl font-bold text-gray-800 leading-tight mb-1">{buku.title}</h2>
         <p class="text-gray-600 font-medium mb-3">ISBN: {buku.isbn || "Tidak tersedia"}</p>
         <span class="text-xs font-bold px-3 py-1.5 rounded-full {available > 0 ? 'bg-[#DFFCE2] text-green-700' : 'bg-red-100 text-red-600'}">
           {available > 0 ? `${available} tersedia` : "Habis"}
         </span>
       </div>
    </div>
    
    <div class="mb-8 rounded-2xl bg-gray-50 p-4 text-sm text-gray-600">
      <p>Total stok: <strong>{buku.stock}</strong></p>
      <p>Salinan tersedia: <strong>{available}</strong></p>
    </div>

    {#if borrowError}
      <p role="alert" class="mb-4 rounded-xl bg-red-100 px-4 py-3 text-sm text-red-700">{borrowError}</p>
    {/if}
    {#if borrowSuccess}
      <p role="status" class="mb-4 rounded-xl bg-green-100 px-4 py-3 text-sm text-green-700">Buku berhasil dipinjam.</p>
    {/if}

    <!-- Footer Modal: Tombol Aksi -->
    <div class="flex justify-end gap-3">
       <button class="px-5 py-2.5 rounded-full border-2 border-gray-200 text-gray-600 hover:bg-gray-100 hover:border-gray-300 transition font-bold text-sm" onclick={tutupModal}>
         Tutup
       </button>
       <!-- Tombol Pinjam hanya aktif jika buku Tersedia -->
       {#if available > 0}
         <button onclick={pinjamBuku} disabled={isBorrowing || borrowSuccess} class="px-5 py-2.5 rounded-full bg-[#9B9AFF] hover:bg-[#8382ff] text-white transition font-bold text-sm shadow-md disabled:opacity-60">
           {isBorrowing ? "Memproses..." : borrowSuccess ? "Berhasil Dipinjam" : "Pinjam Buku"}
         </button>
       {:else}
         <button class="px-5 py-2.5 rounded-full bg-gray-300 text-gray-500 cursor-not-allowed font-bold text-sm" disabled>
           Stok Habis
         </button>
       {/if}
    </div>

  </div>
</div>

<style>
  /* Sedikit animasi CSS agar modal muncul dengan mulus */
  .animate-fade-in-up {
    animation: fadeInUp 0.3s ease-out forwards;
  }
  @keyframes fadeInUp {
    from { opacity: 0; transform: translateY(20px) scale(0.95); }
    to { opacity: 1; transform: translateY(0) scale(1); }
  }
</style>