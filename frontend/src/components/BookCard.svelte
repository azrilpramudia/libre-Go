<script>
  import { createEventDispatcher } from 'svelte'; 
  export let buku;
  
  const dispatch = createEventDispatcher();
</script>

<div class="bg-white/60 backdrop-blur-lg rounded-3xl p-5 shadow-lg hover:-translate-y-3 transition-transform duration-300 border border-white/80 group flex flex-col">
  <!-- Cover Buku Abstrak -->
  <div class="{buku.cover} h-48 rounded-2xl mb-5 flex items-center justify-center text-6xl shadow-inner group-hover:scale-105 transition-transform duration-300">
    📖
  </div>

  <!-- Info Buku -->
  <h3 class="text-xl font-bold text-gray-800 leading-tight mb-2">
    {buku.judul}
  </h3>
  <p class="text-sm text-gray-600 font-medium mb-5">
    Oleh: {buku.penulis}
  </p>

  <!-- Status & Tombol Detail -->
  <div class="flex justify-between items-center mt-auto border-t border-gray-200/60 pt-4">
    <span class="text-xs font-bold px-3 py-1.5 rounded-full {buku.status === 'Tersedia' ? 'bg-[#DFFCE2] text-green-700' : 'bg-red-100 text-red-600'}">
      {buku.status}
    </span>
    
    <!-- Tombol HTML diletakkan di sini beserta event on:click nya -->
    <button 
      on:click={() => dispatch('bukaDetail', buku)} 
      class="bg-[#B8B8FF] hover:bg-[#9B9AFF] text-white text-xs font-bold py-2 px-4 rounded-full transition shadow"
    >
      Detail
    </button>
  </div>
</div>