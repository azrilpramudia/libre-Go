<script>
  import { Link } from "svelte-routing";
  
  // Data sementara (dummy) sebelum disambungkan dengan backend Golang
  let bukuPopuler = [
    { id: 1, judul: "Petualangan Si Kancil", penulis: "Anonim", status: "Tersedia", cover: "bg-[#AADEFF]" },
    { id: 2, judul: "Belajar Svelte 5", penulis: "Frontend Ninja", status: "Dipinjam", cover: "bg-[#B8B8FF]" },
    { id: 3, judul: "Misteri Hutan Berbisik", penulis: "Budi Santoso", status: "Tersedia", cover: "bg-[#9B9AFF]" },
    { id: 4, judul: "Resep Kue Kering", penulis: "Chef Juna", status: "Tersedia", cover: "bg-[#DFFCE2]" },
  ];
</script>

<main class="relative min-h-screen bg-[#f4f9f6] text-gray-900 font-sans overflow-x-hidden flex flex-col">
  
  <!-- Efek Background Glowing Orbs -->
  <div class="absolute -top-20 -left-20 w-[500px] h-[500px] bg-[#B8B8FF]/60 rounded-full blur-[100px] pointer-events-none"></div>
  <div class="absolute top-[10%] -right-20 w-[600px] h-[600px] bg-[#AADEFF]/50 rounded-full blur-[120px] pointer-events-none"></div>
  <div class="absolute top-[60%] left-[20%] w-[700px] h-[700px] bg-[#9B9AFF]/30 rounded-full blur-[150px] pointer-events-none"></div>
  
  <!-- Konten Utama -->
  <div class="relative z-10 max-w-7xl mx-auto px-6 py-8 w-full flex-grow flex flex-col">
    
    <!-- Bagian Navbar -->
    <nav class="flex justify-between items-center mb-10 lg:mb-16">
      <div class="flex items-center gap-2">
        <div class="text-3xl drop-shadow-md">📚</div>
        <span class="text-xl font-bold tracking-wide text-gray-800">PerpusFun</span>
      </div>
      
      <div class="hidden md:flex gap-8 text-sm font-semibold text-gray-600">
        <Link to="/koleksi"><span class="hover:text-[#9B9AFF] transition cursor-pointer">01 / Koleksi</span></Link>
        <a href="#" class="hover:text-[#9B9AFF] transition">02 / Cara Pinjam</a>
        <a href="#" class="hover:text-[#9B9AFF] transition">03 / Kategori</a>
        <a href="#" class="hover:text-[#9B9AFF] transition">04 / FAQ</a>
      </div>
      
      <div class="flex gap-4 items-center">
        <!-- Link menuju Login -->
        <Link to="/login">
          <span class="text-sm font-bold text-gray-700 hover:text-[#9B9AFF] transition cursor-pointer">
            Login
          </span>
        </Link>
        <Link to="/login">
          <span class="bg-[#9B9AFF] hover:bg-[#8382ff] text-white text-sm font-bold py-2 px-6 rounded-full shadow-lg transition transform hover:-translate-y-1 cursor-pointer inline-block">
            Daftar Member →
          </span>
        </Link>
      </div>
    </nav>

    <!-- Bagian Hero / Tengah -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
      
      <!-- Kolom Kiri (Teks Utama & Statistik) -->
      <div class="lg:col-span-5 space-y-8 relative z-20">
        <div class="inline-flex items-center gap-3 bg-white/60 backdrop-blur-md border border-white/80 px-4 py-2 rounded-full text-sm font-bold text-[#9B9AFF] shadow-sm">
          <span class="bg-[#B8B8FF] text-white px-3 py-1 rounded-full text-xs">Program Membaca</span>
          Pelajari caranya →
        </div>
        
        <h1 class="text-6xl lg:text-7xl font-extrabold leading-tight text-gray-900 drop-shadow-sm">
          Kelola <br />
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff]">
            koleksi digital
          </span> <br />
          hingga 600K
        </h1>
        
        <div class="flex gap-10 pt-8 border-t border-gray-300/50">
          <div>
            <h3 class="text-3xl font-extrabold text-gray-900">100%</h3>
            <p class="text-sm text-gray-600 mt-1 font-medium">Gratis Akses</p>
          </div>
          <div>
            <h3 class="text-3xl font-extrabold text-gray-900">1</h3>
            <p class="text-sm text-gray-600 mt-1 font-medium">Menit Daftar</p>
          </div>
          <div>
            <h3 class="text-3xl font-extrabold text-gray-900">24/7</h3>
            <p class="text-sm text-gray-600 mt-1 font-medium">Support Live</p>
          </div>
        </div>
      </div>

      <!-- Kolom Tengah (Ilustrasi HP Glassmorphism & Tangan) -->
      <div class="lg:col-span-4 flex justify-center relative h-[600px]">
         <div class="absolute top-[10%] right-[10%] w-[240px] h-[480px] z-0">
            <div class="absolute bottom-0 right-0 w-[180px] h-[300px] bg-[#1a1a1a] rounded-[40px] transform rotate-12 translate-x-10 translate-y-10"></div>
            <div class="absolute top-10 -left-10 w-[70px] h-[140px] bg-[#1a1a1a] rounded-full transform -rotate-[30deg]"></div>
            <div class="absolute top-32 -left-16 w-[70px] h-[130px] bg-[#1a1a1a] rounded-full transform -rotate-[20deg]"></div>
            <div class="absolute top-52 -left-12 w-[65px] h-[120px] bg-[#1a1a1a] rounded-full transform -rotate-[15deg]"></div>
         </div>
         
         <div class="relative z-10 w-[270px] h-[550px] bg-white/20 backdrop-blur-xl border border-white/60 rounded-[3rem] shadow-[0_20px_50px_rgba(0,0,0,0.1)] p-4 flex flex-col items-center justify-center overflow-hidden self-center">
            <div class="absolute top-0 left-0 w-full h-1/3 bg-gradient-to-b from-white/60 to-transparent"></div>
            <div class="w-40 h-40 rounded-full border border-white/50 bg-white/20 flex items-center justify-center relative z-20 shadow-inner">
               <div class="w-28 h-28 rounded-full bg-gradient-to-tr from-[#AADEFF] to-[#DFFCE2] opacity-90 animate-pulse"></div>
               <div class="absolute flex flex-col gap-3 z-30 transform -translate-x-2">
                  <div class="w-10 h-3 bg-white/90 rounded-full ml-10 shadow-sm"></div>
                  <div class="w-16 h-3 bg-white/90 rounded-full shadow-sm"></div>
                  <div class="w-12 h-3 bg-white/90 rounded-full ml-14 shadow-sm"></div>
               </div>
            </div>
         </div>

         <div class="absolute bottom-[25%] right-[5%] w-[80px] h-[180px] bg-[#1a1a1a] rounded-full transform -rotate-12 z-20"></div>
      </div>

      <!-- Kolom Kanan (Deskripsi & Tombol Aksi) -->
      <div class="lg:col-span-3 space-y-6 relative z-10 lg:pl-6">
        <div class="absolute -top-16 right-10 w-24 h-24 bg-[#9B9AFF]/40 backdrop-blur-xl rounded-full shadow-lg border border-white/30"></div>
        <div class="absolute -bottom-10 left-0 w-16 h-16 bg-[#B8B8FF]/50 backdrop-blur-xl rounded-full shadow-lg border border-white/50 flex items-center justify-center text-sm">✨</div>
        
        <p class="text-sm text-gray-700 leading-relaxed bg-white/50 backdrop-blur-xl p-6 rounded-3xl border border-white/70 shadow-sm relative z-20 font-medium">
          Tunjukkan kegemaran membacamu, selesaikan tantangan literasi, dan dapatkan poin untuk berpartisipasi dalam event perpustakaan eksklusif tanpa biaya sama sekali.
        </p>
        
        <div class="flex flex-col sm:flex-row lg:flex-col xl:flex-row gap-3 pt-2 relative z-20">
          <Link to="/koleksi">
            <span class="inline-block bg-[#9B9AFF] hover:bg-[#8382ff] text-white text-sm font-bold py-3 px-6 rounded-full shadow-lg transition cursor-pointer text-center">
              Mulai Baca →
            </span>
          </Link>
          <button class="bg-white/60 hover:bg-white text-gray-800 border border-gray-300 text-sm font-bold py-3 px-6 rounded-full shadow-sm transition">
            Gabung Discord
          </button>
        </div>
      </div>

    </div>

    <!-- ========================================== -->
    <!-- SEKSI BARU: Search Bar & Koleksi Buku       -->
    <!-- ========================================== -->
    <section class="w-full mt-32 relative z-20 pb-16">
      
      <!-- Header Section & Search Bar -->
      <div class="mb-12 flex flex-col md:flex-row justify-between items-center gap-6 bg-white/40 backdrop-blur-md p-6 rounded-[2rem] border border-white/60 shadow-sm">
        <h2 class="text-3xl font-extrabold text-gray-800">
          Jelajahi <span class="text-transparent bg-clip-text bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff]">Koleksi</span>
        </h2>
        
        <!-- Search Bar -->
        <div class="relative w-full md:w-96">
          <input 
            type="text" 
            placeholder="Cari judul buku, penulis..." 
            class="w-full pl-6 pr-14 py-3 rounded-full border-2 border-white/70 bg-white/50 backdrop-blur-md shadow-inner focus:outline-none focus:border-[#9B9AFF] text-gray-700 placeholder-gray-500 transition-all font-medium" 
          />
          <button class="absolute right-2 top-1/2 transform -translate-y-1/2 bg-[#9B9AFF] hover:bg-[#8382ff] text-white p-2 rounded-full transition shadow-md flex items-center justify-center">
            <span class="text-sm">🔍</span>
          </button>
        </div>
      </div>

      <!-- Grid Kartu Buku -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        {#each bukuPopuler as buku}
          <!-- Kartu Buku -->
          <div class="bg-white/60 backdrop-blur-lg rounded-3xl p-5 shadow-lg hover:-translate-y-3 transition-transform duration-300 border border-white/80 group flex flex-col">
            
            <!-- Cover Buku Abstrak -->
            <div class="{buku.cover} h-48 rounded-2xl mb-5 flex items-center justify-center text-6xl shadow-inner group-hover:scale-105 transition-transform duration-300">
              📖
            </div>
            
            <!-- Info Buku -->
            <h3 class="text-xl font-bold text-gray-800 leading-tight mb-2">{buku.judul}</h3>
            <p class="text-sm text-gray-600 font-medium mb-5">Oleh: {buku.penulis}</p>
            
            <!-- Status & Tombol Pinjam -->
            <div class="flex justify-between items-center mt-auto border-t border-gray-200/60 pt-4">
              <span class="text-xs font-bold px-3 py-1.5 rounded-full {buku.status === 'Tersedia' ? 'bg-[#DFFCE2] text-green-700' : 'bg-red-100 text-red-600'}">
                {buku.status}
              </span>
              <button class="bg-[#B8B8FF] hover:bg-[#9B9AFF] text-white text-xs font-bold py-2 px-4 rounded-full transition shadow">
                Detail
              </button>
            </div>
          </div>
        {/each}
      </div>
      
    </section>

  </div>
</main>