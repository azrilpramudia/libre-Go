<script>
  import { Link } from "svelte-routing";
  import BookCard from "../components/BookCard.svelte";
  import BookModal from '../components/BookModal.svelte';
  import Footer from '../components/Footer.svelte';
  import Navbar from "../components/Navbar.svelte";
  import { booksAPI } from "../lib/api.js";
  import { currentMember, isAuthenticated } from "../lib/auth.js";

  let bukuPopuler = $state([]);
  let searchQuery = $state("");
  let isLoadingBooks = $state(false);
  let booksError = $state("");
  let isAuth = $derived($isAuthenticated);
  let filteredBuku = $derived(
    bukuPopuler.filter((buku) =>
      buku.title?.toLowerCase().includes(searchQuery.toLowerCase())
    )
  );
  let showModal = $state(false);
  let selectedBuku = $state(null);

  $effect(() => {
    if (isAuth) {
      loadBooks();
    } else {
      bukuPopuler = [];
      searchQuery = "";
      booksError = "";
      isLoadingBooks = false;
    }
  });

  async function loadBooks() {
    isLoadingBooks = true;
    booksError = "";
    try {
      const response = await booksAPI.list(8);
      bukuPopuler = response.data || [];
    } catch (error) {
      booksError = error.data?.error || error.message || "Gagal memuat koleksi buku.";
    } finally {
      isLoadingBooks = false;
    }
  }

  function handleBukaDetail(buku) {
    selectedBuku = buku;
    showModal = true;
  }
</script>

<main class="relative min-h-screen bg-[#fdfdfd] text-gray-900 font-sans overflow-x-hidden flex flex-col">
  <!-- Efek Background Latar -->
  <div class="absolute -top-40 -left-40 w-[600px] h-[600px] bg-[#DFFCE2]/40 rounded-full blur-[120px] pointer-events-none"></div>
  <div class="absolute top-[20%] -right-20 w-[700px] h-[700px] bg-[#B8B8FF]/20 rounded-full blur-[150px] pointer-events-none"></div>
  <div class="absolute bottom-0 left-[30%] w-[500px] h-[500px] bg-[#AADEFF]/30 rounded-full blur-[100px] pointer-events-none"></div>

  <div class="relative z-10 max-w-7xl mx-auto px-6 py-6 w-full flex-grow flex flex-col">
    
    <Navbar bind:searchQuery />

    {#if isAuth}
      <section class="mb-12 rounded-3xl border border-white/50 bg-white/60 p-8 shadow-sm backdrop-blur-lg">
        <p class="text-sm font-bold uppercase tracking-wide text-[#9B9AFF]">Dashboard Pengguna</p>
        <h1 class="mt-2 text-3xl font-extrabold text-gray-800">
          Selamat datang, {$currentMember?.full_name || "Member"}! 👋
        </h1>
        <p class="mt-2 text-gray-600">Temukan buku berikutnya untuk dibaca.</p>
      </section>

      <section class="w-full pb-16">
        <div class="mb-8 flex flex-col gap-4 border-b border-gray-200 pb-4 md:flex-row md:items-end md:justify-between">
          <div>
            <h2 class="text-3xl font-extrabold text-gray-800">
              Jelajahi <span class="bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff] bg-clip-text text-transparent">Koleksi</span>
            </h2>
            <p class="mt-2 text-sm font-medium text-gray-500">Temukan buku-buku menarik untuk dibaca hari ini.</p>
          </div>
          <div class="flex flex-wrap items-center gap-3">
            {#if searchQuery.trim()}
              <span class="rounded-full bg-[#DFFCE2] px-4 py-2 text-sm font-bold text-green-700">
                Hasil untuk: "{searchQuery}"
              </span>
            {/if}
            <Link to="/koleksi" class="text-sm font-bold text-[#8382ff] hover:underline">
              Lihat semua buku →
            </Link>
          </div>
        </div>

        <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-4">
          {#if isLoadingBooks}
            <p class="col-span-full py-12 text-center text-gray-500">Memuat koleksi buku...</p>
          {:else if booksError}
            <div role="alert" class="col-span-full rounded-2xl bg-red-100 p-5 text-center text-red-700">{booksError}</div>
          {:else if filteredBuku.length > 0}
            {#each filteredBuku as buku (buku.id)}
              <BookCard {buku} onbukaDetail={handleBukaDetail} />
            {/each}
          {:else}
            <div class="col-span-full rounded-3xl border border-gray-200 bg-white/50 py-16 text-center">
              <div class="mb-4 text-4xl">📚</div>
              <h3 class="text-xl font-bold text-gray-700">
                {searchQuery.trim() ? "Buku tidak ditemukan" : "Koleksi buku belum tersedia"}
              </h3>
              <p class="mt-2 text-gray-500">
                {searchQuery.trim() ? "Coba kata pencarian lain." : "Silakan kembali lagi nanti."}
              </p>
            </div>
          {/if}
        </div>
      </section>
    {:else}
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
          <Link to="/login" class="rounded-full bg-[#9B9AFF] px-8 py-3.5 text-sm font-bold text-white shadow-[0_8px_20px_rgba(155,154,255,0.4)] transition hover:-translate-y-1 hover:bg-[#8382ff]">
            Mulai Sekarang →
          </Link>
          <Link to="/koleksi" class="rounded-full border border-[#B8B8FF] bg-white px-8 py-3.5 text-sm font-bold text-[#9B9AFF] shadow-sm transition hover:bg-gray-50">
            Lihat Koleksi
          </Link>
        </div>

        <div class="flex gap-8 pt-6">
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
    <section id="tentang" class="w-full mt-32 relative z-20 scroll-mt-8">
      <div class="bg-white/60 backdrop-blur-lg rounded-[3rem] p-10 md:p-16 shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-white/50 relative overflow-hidden">
        
        <!-- Ornamen dekoratif -->
        <div class="absolute -top-10 -right-10 w-40 h-40 bg-[#DFFCE2] rounded-full blur-2xl opacity-80 pointer-events-none"></div>
        <div class="absolute -bottom-10 -left-10 w-40 h-40 bg-[#AADEFF] rounded-full blur-2xl opacity-60 pointer-events-none"></div>

        <div class="relative z-10">
          
          <!-- Teks Penjelasan -->
          <div class="max-w-3xl space-y-6">
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

        </div>
      </div>
    </section>
    <section id="keunggulan" class="mt-24 w-full scroll-mt-8">
      <div class="mb-10 text-center">
        <p class="text-sm font-bold uppercase tracking-wide text-[#9B9AFF]">Keunggulan</p>
        <h2 class="mt-2 text-3xl font-extrabold text-gray-800">Kenapa memilih PerpusFun?</h2>
        <p class="mx-auto mt-3 max-w-2xl text-gray-600">
          Temukan dan pinjam buku dengan mudah melalui pengalaman perpustakaan digital yang praktis.
        </p>
      </div>
      <div class="grid gap-6 md:grid-cols-3">
        <article class="rounded-3xl border border-gray-100 bg-white p-7 shadow-sm">
          <div class="mb-4 text-4xl">🔎</div>
          <h3 class="mb-2 text-lg font-extrabold text-gray-800">Pencarian mudah</h3>
          <p class="text-sm leading-relaxed text-gray-600">Cari buku dari koleksi perpustakaan dan temukan bacaan yang sesuai minatmu.</p>
        </article>
        <article class="rounded-3xl border border-gray-100 bg-white p-7 shadow-sm">
          <div class="mb-4 text-4xl">📚</div>
          <h3 class="mb-2 text-lg font-extrabold text-gray-800">Koleksi terpusat</h3>
          <p class="text-sm leading-relaxed text-gray-600">Lihat ketersediaan buku dan riwayat peminjaman dalam satu platform.</p>
        </article>
        <article class="rounded-3xl border border-gray-100 bg-white p-7 shadow-sm">
          <div class="mb-4 text-4xl">💻</div>
          <h3 class="mb-2 text-lg font-extrabold text-gray-800">Praktis di web</h3>
          <p class="text-sm leading-relaxed text-gray-600">Jelajahi katalog dan kelola peminjaman langsung dari perangkatmu.</p>
        </article>
      </div>
    </section>
    {/if}
  </div>

  {#if showModal}
    <BookModal buku={selectedBuku} onclose={() => (showModal = false)} onborrowed={loadBooks} />
  {/if}

  <Footer />
</main>