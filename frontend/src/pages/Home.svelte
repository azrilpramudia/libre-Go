<script>
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from '../components/BookModal.svelte';
  import Footer from '../components/Footer.svelte';

  let bukuPopuler = [
    { id: 1, judul: "Petualangan Si Kancil", penulis: "Anonim", status: "Tersedia", cover: "bg-[#AADEFF]" },
    { id: 2, judul: "Belajar Svelte 5", penulis: "Frontend Ninja", status: "Dipinjam", cover: "bg-[#B8B8FF]" },
    { id: 3, judul: "Misteri Hutan Berbisik", penulis: "Budi Santoso", status: "Tersedia", cover: "bg-[#9B9AFF]" },
    { id: 4, judul: "Resep Kue Kering", penulis: "Chef Juna", status: "Tersedia", cover: "bg-[#DFFCE2]" },
  ];

  let searchQuery = "";
  
  $: filteredBuku = bukuPopuler.filter((buku) =>
    buku.judul.toLowerCase().includes(searchQuery.toLowerCase()) ||
    buku.penulis.toLowerCase().includes(searchQuery.toLowerCase())
  );

  let showModal = false;
  /** @type {{id: number, judul: string, penulis: string, status: string, cover: string} | null} */
  let selectedBuku = null;

  /** @param {CustomEvent} event */
  function handleBukaDetail(event) {
    selectedBuku = event.detail; 
    showModal = true;            
  }
</script>

<main class="relative min-h-screen bg-[#fdfdfd] text-gray-900 font-sans overflow-x-hidden flex flex-col">
  <!-- Efek Background Latar -->
  <div class="absolute -top-40 -left-40 w-[600px] h-[600px] bg-[#DFFCE2]/40 rounded-full blur-[120px] pointer-events-none"></div>
  <div class="absolute top-[20%] -right-20 w-[700px] h-[700px] bg-[#B8B8FF]/20 rounded-full blur-[150px] pointer-events-none"></div>
  <div class="absolute bottom-0 left-[30%] w-[500px] h-[500px] bg-[#AADEFF]/30 rounded-full blur-[100px] pointer-events-none"></div>

  <div class="relative z-10 max-w-7xl mx-auto px-6 py-6 w-full flex-grow flex flex-col">
    
    <!-- Navbar -->
    <nav class="flex flex-wrap justify-between items-center mb-16 bg-white/70 backdrop-blur-md px-6 py-4 rounded-full border border-gray-100 shadow-sm">
      <div class="flex items-center gap-2">
        <div class="text-3xl">📚</div>
        <span class="text-xl font-extrabold text-[#9B9AFF] tracking-wide">PerpusFun</span>
      </div>

      <div class="hidden lg:flex gap-8 text-sm font-semibold text-gray-600">
        <Link to="/"><span class="hover:text-[#9B9AFF] transition cursor-pointer">Home</span></Link>
        <Link to="/tentang"><span class="hover:text-[#9B9AFF] transition cursor-pointer">Tentang</span></Link>
        <Link to="/keunggulan"><span class="hover:text-[#9B9AFF] transition cursor-pointer">Keunggulan</span></Link>
        <Link to="/koleksi"><span class="hover:text-[#9B9AFF] transition cursor-pointer">Koleksi Buku</span></Link>
      </div>

      <div class="flex items-center gap-4">
        <!-- Search Bar di dalam Navbar -->
        <div class="relative hidden md:block w-48 lg:w-64">
          <input
            type="text"
            bind:value={searchQuery}
            placeholder="Cari buku..."
            class="w-full pl-5 pr-10 py-2.5 rounded-full border border-gray-200 bg-gray-50 focus:outline-none focus:border-[#9B9AFF] text-sm text-gray-700 transition"
          />
          <span class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 text-sm">🔍</span>
        </div>

        <!-- TOMBOL LOGIN & DAFTAR BARU -->
        <Link to="/login">
          <span class="text-sm font-bold text-gray-700 hover:text-[#9B9AFF] transition cursor-pointer hidden sm:block">
            Login
          </span>
        </Link>
        <Link to="/daftar">
          <span class="bg-[#9B9AFF] hover:bg-[#8382ff] text-white text-sm font-bold py-2.5 px-6 rounded-full shadow-[0_4px_15px_rgba(155,154,255,0.4)] transition transform hover:-translate-y-1 cursor-pointer whitespace-nowrap">
            Daftar
          </span>
        </Link>
        
      </div>
    </nav>
    

    <!-- Seksyen Hero -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center min-h-[70vh]">
      <!-- Bahagian Kiri -->
      <div class="space-y-8 z-20">
        <h1 class="text-5xl lg:text-6xl font-extrabold leading-tight text-gray-800">
          Perpustakaan <br />
          <span class="text-[#9B9AFF]">Digital Platform</span> <br />
          PERPUSFUN
        </h1>
        
        <p class="text-gray-600 text-lg max-w-md font-medium leading-relaxed">
          Akses Literasi Tanpa Batas, Hanya dalam Genggamanmu. Membawa pengetahuan ke era digital.
        </p>

        <div class="flex flex-wrap gap-4 pt-2">
          <button class="bg-[#9B9AFF] hover:bg-[#8382ff] text-white text-sm font-bold py-3.5 px-8 rounded-full shadow-[0_8px_20px_rgba(155,154,255,0.4)] transition transform hover:-translate-y-1">
            Mulai Sekarang →
          </button>
          <button class="bg-white hover:bg-gray-50 text-[#9B9AFF] border border-[#B8B8FF] text-sm font-bold py-3.5 px-8 rounded-full shadow-sm transition">
            Lihat Koleksi
          </button>
        </div>

        <div class="flex gap-8 pt-6">
          <div>
            <h3 class="text-2xl font-extrabold text-[#9B9AFF]">1000+</h3>
            <p class="text-xs text-gray-500 font-semibold mt-1">Koleksi Buku</p>
          </div>
          <div class="w-px bg-gray-200"></div>
          <div>
            <h3 class="text-2xl font-extrabold text-[#9B9AFF]">24/7</h3>
            <p class="text-xs text-gray-500 font-semibold mt-1">Akses</p>
          </div>
          <div class="w-px bg-gray-200"></div>
          <div>
            <h3 class="text-2xl font-extrabold text-[#9B9AFF]">100%</h3>
            <p class="text-xs text-gray-500 font-semibold mt-1">Digital</p>
          </div>
        </div>
      </div>

      <!-- Bahagian Kanan (Ilustrasi HP) -->
      <div class="relative h-[600px] flex justify-center items-center z-10 hidden lg:flex">
        <div class="absolute top-16 left-4 bg-white px-5 py-2.5 rounded-full shadow-lg text-sm font-bold text-[#9B9AFF] flex items-center gap-2 z-30 animate-bounce" style="animation-duration: 3s;">
          <span class="w-2.5 h-2.5 rounded-full bg-[#9B9AFF]"></span> Akses 24/7
        </div>
        <div class="absolute bottom-20 right-0 bg-white px-5 py-2.5 rounded-full shadow-lg text-sm font-bold text-orange-400 flex items-center gap-2 z-30">
          <span class="w-2.5 h-2.5 rounded-full bg-orange-400"></span> Multi-Platform
        </div>

        <div class="w-[240px] h-[480px] bg-white rounded-[2.5rem] shadow-[0_20px_50px_rgba(0,0,0,0.15)] border-[10px] border-gray-900 absolute transform rotate-12 translate-x-16 translate-y-12 z-20 overflow-hidden flex flex-col">
          <div class="bg-gradient-to-br from-[#9B9AFF] to-[#B8B8FF] h-40 w-full p-4 flex flex-col justify-between">
             <div class="w-1/2 h-2 bg-white/40 rounded-full mx-auto"></div>
             <div class="text-white font-bold text-xl leading-tight">Perpus<br>Digital.</div>
          </div>
          <div class="p-4 flex-grow bg-gray-50 flex flex-col gap-3">
             <div class="w-2/3 h-3 bg-gray-200 rounded-full mb-2"></div>
             <div class="w-full h-24 bg-[#DFFCE2] rounded-xl shadow-inner flex items-center justify-center text-3xl">🌿</div>
             <div class="w-full h-24 bg-[#AADEFF] rounded-xl shadow-inner flex items-center justify-center text-3xl">🌊</div>
          </div>
        </div>

        <div class="w-[240px] h-[480px] bg-gray-50 rounded-[2.5rem] shadow-xl border-[10px] border-gray-800 absolute transform -rotate-[15deg] -translate-x-12 -translate-y-8 z-10 overflow-hidden">
          <div class="w-1/2 h-2 bg-gray-300 rounded-full mx-auto mt-4 mb-6"></div>
          <div class="px-4 space-y-4">
             <div class="w-full h-32 bg-white rounded-xl shadow-sm border border-gray-100 flex items-center justify-center text-4xl">📚</div>
             <div class="w-full h-32 bg-[#B8B8FF]/30 rounded-xl shadow-sm flex items-center justify-center text-4xl">✨</div>
          </div>
        </div>
      </div>
    </div>
<!-- ========================================== -->
    <!-- SEKSYEN TENTANG (TAMBAHAN BARU)            -->
    <!-- ========================================== -->
    <section class="w-full mt-32 relative z-20">
      <div class="bg-white/60 backdrop-blur-lg rounded-[3rem] p-10 md:p-16 shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-white/50 relative overflow-hidden">
        
        <!-- Ornamen dekoratif -->
        <div class="absolute -top-10 -right-10 w-40 h-40 bg-[#DFFCE2] rounded-full blur-2xl opacity-80 pointer-events-none"></div>
        <div class="absolute -bottom-10 -left-10 w-40 h-40 bg-[#AADEFF] rounded-full blur-2xl opacity-60 pointer-events-none"></div>

        <div class="relative z-10 flex flex-col lg:flex-row gap-12 items-center">
          
          <!-- Teks Penjelasan -->
          <div class="lg:w-1/2 space-y-6">
            <div class="inline-block bg-[#B8B8FF]/20 text-[#9B9AFF] font-bold px-4 py-1.5 rounded-full text-sm mb-2 shadow-sm border border-[#B8B8FF]/30">
              Tentang Kami
            </div>
            <h2 class="text-3xl md:text-4xl font-extrabold text-gray-800 leading-tight">
              Revolusi Membaca di <br/> Era <span class="text-transparent bg-clip-text bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff]">Digital</span>
            </h2>
            <p class="text-gray-600 font-medium leading-relaxed">
              PerpusFun hadir untuk mematahkan stigma bahwa perpustakaan itu kaku dan membosankan. Dibangun dengan antarmuka Svelte yang super ringan dan ditenagai ketangguhan Golang di balik layar, kami menjamin pengalaman membacamu lancar tanpa hambatan!
            </p>
            
            <ul class="space-y-4 mt-6">
              <li class="flex items-center gap-4 text-sm font-bold text-gray-700">
                <span class="w-8 h-8 rounded-full bg-[#DFFCE2] text-green-600 flex items-center justify-center shadow-sm">✓</span> 
                Ribuan Koleksi Gratis
              </li>
              <li class="flex items-center gap-4 text-sm font-bold text-gray-700">
                <span class="w-8 h-8 rounded-full bg-[#DFFCE2] text-green-600 flex items-center justify-center shadow-sm">✓</span> 
                Tanpa Iklan yang Mengganggu
              </li>
            </ul>
          </div>

          <!-- Kartu Fitur (Grid) -->
          <div class="lg:w-1/2 grid grid-cols-1 sm:grid-cols-2 gap-6 w-full">
            
            <!-- Kartu 1 -->
            <div class="bg-white rounded-3xl p-6 shadow-sm border border-gray-100 transform sm:translate-y-6 hover:-translate-y-2 transition duration-300">
              <div class="w-14 h-14 bg-[#9B9AFF]/10 rounded-2xl flex items-center justify-center text-3xl mb-5 shadow-inner">
                🚀
              </div>
              <h4 class="font-extrabold text-gray-800 mb-2 text-lg">Akses Ngebut</h4>
              <p class="text-sm text-gray-500 font-medium leading-relaxed">Sistem pencarian instan yang membuatmu bisa menemukan buku incaran dalam hitungan detik.</p>
            </div>
            
            <!-- Kartu 2 -->
            <div class="bg-white rounded-3xl p-6 shadow-sm border border-gray-100 hover:-translate-y-2 transition duration-300">
              <div class="w-14 h-14 bg-[#AADEFF]/20 rounded-2xl flex items-center justify-center text-3xl mb-5 shadow-inner">
                📱
              </div>
              <h4 class="font-extrabold text-gray-800 mb-2 text-lg">Multi Perangkat</h4>
              <p class="text-sm text-gray-500 font-medium leading-relaxed">Mulai baca dari laptop, lanjutkan dari HP. Sinkronisasi otomatis di mana saja kamu berada.</p>
            </div>

          </div>
        </div>
      </div>
    </section>
    <!-- ========================================== -->
    <!-- SEKSYEN KOLEKSI (DIKEMBALIKAN & SELALU TAMPIL) -->
    <section class="w-full mt-24 relative z-20 pb-16">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-end mb-10 border-b border-gray-200 pb-4 gap-4">
        <div>
          <h2 class="text-3xl font-extrabold text-gray-800">
            Jelajahi <span class="text-transparent bg-clip-text bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff]">Koleksi</span>
          </h2>
          <p class="text-gray-500 mt-2 text-sm font-medium">Temukan buku-buku menarik untuk dibaca hari ini.</p>
        </div>
        
        <!-- Indikator jika sedang mencari -->
        {#if searchQuery.length > 0}
          <div class="bg-[#DFFCE2] text-green-700 px-4 py-2 rounded-full text-sm font-bold shadow-sm">
            Hasil untuk: "{searchQuery}"
          </div>
        {/if}
      </div>

      <!-- Grid Kartu Buku -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        {#each filteredBuku as buku (buku.id)}
          <BookCard {buku} on:bukaDetail={handleBukaDetail} />
        {:else}
          <!-- Pesan jika buku yang dicari tidak ada -->
          <div class="col-span-full text-center py-16 bg-white/50 rounded-3xl border border-gray-200">
            <div class="text-4xl mb-4">🕵️‍♂️</div>
            <h3 class="text-xl font-bold text-gray-700">Waduh, bukunya tidak ditemukan!</h3>
            <p class="text-gray-500 mt-2">Coba cari dengan kata kunci judul atau nama penulis yang lain.</p>
          </div>
        {/each}
      </div>
    </section>
  </div>

  {#if showModal}
    <BookModal buku={selectedBuku} on:close={() => showModal = false} />
  {/if}

  <Footer />
</main>