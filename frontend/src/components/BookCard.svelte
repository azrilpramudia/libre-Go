<script>
  // 'export let' adalah cara Svelte menerima Props dari komponen induk
  export let buku;
</script>

<div class="bg-white rounded-3xl p-6 shadow-xl hover:-translate-y-2 transition-transform duration-300 border-b-4 border-fun-lavender">
  <div class="bg-fun-blue h-40 rounded-2xl mb-4 flex items-center justify-center text-5xl shadow-inner">
    📖
  </div>
  <h3 class="text-xl font-bold text-gray-800 leading-tight mb-1">{buku.judul}</h3>
  <p class="text-sm text-gray-500 mb-4">Oleh: {buku.penulis}</p>
  
  <div class="flex justify-between items-center mt-auto">
    <span class="text-xs font-bold px-3 py-1 rounded-full {buku.status === 'Tersedia' ? 'bg-fun-mint text-green-700' : 'bg-red-100 text-red-600'}">
      {buku.status}
    </span>
    <button class="bg-fun-purple hover:bg-purple-500 text-white text-sm font-bold py-2 px-4 rounded-full transition shadow">
      Detail
    </button>
  </div>
</div>