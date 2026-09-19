<script>
  let bukuTersedia = [
    { id: 1, judul: "Petualangan Si Kancil", penulis: "Anonim", status: "Tersedia" },
    { id: 2, judul: "Belajar Svelte itu Asyik!", penulis: "Indra", status: "Dipinjam" },
  ];
</script>

<main class="min-h-screen bg-[#f4f9f6] p-8 text-gray-800">
  <div class="max-w-7xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-extrabold text-gray-900">📚 Koleksi Buku</h1>
      <!-- Tombol kembali ke Beranda -->
      <a href="/" class="bg-[#9B9AFF] text-white px-4 py-2 rounded-full font-bold hover:bg-[#8382ff] transition">
        ← Kembali ke Home
      </a>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
      {#each bukuTersedia as buku}
        <div class="bg-white p-6 rounded-3xl shadow-sm border border-gray-100 hover:shadow-md transition">
          <div class="bg-[#AADEFF]/30 h-32 rounded-2xl mb-4 flex items-center justify-center text-4xl">📖</div>
          <h3 class="font-bold text-lg">{buku.judul}</h3>
          <p class="text-sm text-gray-500 mb-4">Oleh: {buku.penulis}</p>
          <span class="text-xs font-bold px-3 py-1 rounded-full {buku.status === 'Tersedia' ? 'bg-[#DFFCE2] text-green-700' : 'bg-red-100 text-red-600'}">
            {buku.status}
          </span>
        </div>
      {/each}
    </div>
  </div>
</main>