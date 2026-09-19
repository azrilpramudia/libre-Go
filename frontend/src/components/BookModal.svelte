<script>
  import { createEventDispatcher } from 'svelte';
  export let buku;

  const dispatch = createEventDispatcher();

  // Fungsi ini mengirim sinyal 'close' ke file utama untuk menutup modal
  function tutupModal() {
    dispatch('close');
  }
</script>

<!-- Latar belakang gelap (Overlay) yang bisa diklik untuk menutup -->
<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
<div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-gray-900/40 backdrop-blur-sm transition-all" on:click={tutupModal}>
  
  <!-- Kotak Modal (menggunakan stopPropagation agar tidak ikut tertutup saat kotaknya diklik) -->
  <div class="bg-white/90 backdrop-blur-xl rounded-3xl p-8 max-w-lg w-full shadow-2xl border border-white relative animate-fade-in-up" on:click|stopPropagation>
    
    <!-- Tombol Silang (X) di pojok kanan atas -->
    <button class="absolute top-5 right-5 text-gray-400 hover:text-gray-700 transition" on:click={tutupModal}>
      ✕
    </button>
    
    <!-- Header Modal: Cover & Info Singkat -->
    <div class="flex gap-6 items-center mb-6">
       <div class="{buku.cover} w-24 h-32 rounded-2xl flex items-center justify-center text-4xl shadow-inner shrink-0">
         📖
       </div>
       <div>
         <h2 class="text-2xl font-bold text-gray-800 leading-tight mb-1">{buku.judul}</h2>
         <p class="text-gray-600 font-medium mb-3">Oleh: {buku.penulis}</p>
         <span class="text-xs font-bold px-3 py-1.5 rounded-full {buku.status === 'Tersedia' ? 'bg-[#DFFCE2] text-green-700' : 'bg-red-100 text-red-600'}">
           {buku.status}
         </span>
       </div>
    </div>
    
    <!-- Body Modal: Sinopsis (Simulasi) -->
    <div class="mb-8">
      <h3 class="text-lg font-bold text-gray-800 mb-2">Sinopsis</h3>
      <p class="text-sm text-gray-600 leading-relaxed bg-gray-50 p-4 rounded-2xl border border-gray-100">
        Buku "{buku.judul}" ini menceritakan tentang petualangan luar biasa yang tak terlupakan. 
        (Ini adalah sinopsis sementara sambil menunggu integrasi API dari backend Golang).
      </p>
    </div>

    <!-- Footer Modal: Tombol Aksi -->
    <div class="flex justify-end gap-3">
       <button class="px-5 py-2.5 rounded-full border-2 border-gray-200 text-gray-600 hover:bg-gray-100 hover:border-gray-300 transition font-bold text-sm" on:click={tutupModal}>
         Tutup
       </button>
       <!-- Tombol Pinjam hanya aktif jika buku Tersedia -->
       {#if buku.status === 'Tersedia'}
         <button class="px-5 py-2.5 rounded-full bg-[#9B9AFF] hover:bg-[#8382ff] text-white transition font-bold text-sm shadow-md">
           Pinjam Buku
         </button>
       {:else}
         <button class="px-5 py-2.5 rounded-full bg-gray-300 text-gray-500 cursor-not-allowed font-bold text-sm" disabled>
           Sedang Dipinjam
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