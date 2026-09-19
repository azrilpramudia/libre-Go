<script>
  import { Link } from "svelte-routing";
  // Mengimpor fungsi animasi dari Svelte
  import { fly, fade } from "svelte/transition";
  
  let showPassword = false;
  function togglePassword() {
    showPassword = !showPassword;
  }
</script>

<main class="min-h-screen bg-[#F8FAFC] flex items-center justify-center p-4 sm:p-8 relative overflow-hidden font-sans">
  <!-- Latar Belakang dengan Animasi Denyut (Blob) -->
  <div class="animate-blob absolute top-[-15%] left-[-10%] w-[500px] h-[500px] bg-[#DFFCE2] rounded-full blur-[100px] opacity-70 pointer-events-none"></div>
  <div class="animate-blob animation-delay-2000 absolute bottom-[-10%] right-[-5%] w-[600px] h-[600px] bg-[#B8B8FF] rounded-full blur-[120px] opacity-60 pointer-events-none"></div>

  <!-- Main Card Container dengan efek Fade-in -->
  <div in:fade={{ duration: 800 }} class="relative z-10 bg-white/80 backdrop-blur-2xl w-full max-w-5xl rounded-[2.5rem] shadow-[0_20px_60px_-15px_rgba(0,0,0,0.05)] border border-white flex flex-col md:flex-row overflow-hidden">
    
    <!-- Bagian Kiri -->
    <div class="w-full md:w-5/12 relative overflow-hidden bg-gradient-to-br from-[#9B9AFF] to-[#8382ff] p-10 lg:p-12 text-white flex flex-col justify-between">
      
      <!-- Ornamen statis -->
      <div class="absolute -top-20 -right-20 w-64 h-64 bg-white opacity-10 rounded-full blur-3xl pointer-events-none"></div>
      <div class="absolute bottom-10 -left-10 w-40 h-40 bg-[#B8B8FF] opacity-40 rounded-full blur-2xl pointer-events-none"></div>
      
      <!-- Teks dengan efek terbang dari kiri -->
      <div class="relative z-10" in:fly={{ x: -30, duration: 800, delay: 200 }}>
        <Link to="/">
          <div class="flex items-center gap-3 cursor-pointer mb-12 w-fit group">
            <div class="w-10 h-10 bg-white text-[#9B9AFF] rounded-xl flex items-center justify-center text-xl font-bold shadow-sm group-hover:-rotate-12 group-hover:scale-110 transition transform duration-300">
              📚
            </div>
            <span class="text-xl font-extrabold tracking-wide">PerpusFun</span>
          </div>
        </Link>
        
        <h2 class="text-4xl font-extrabold leading-tight mb-6 drop-shadow-sm">
          Mulai <br/> Petualangan <br/> Literasimu.
        </h2>
        <p class="text-white/90 text-sm font-medium leading-relaxed max-w-xs">
          Gabung sekarang dan nikmati ribuan koleksi buku digital secara gratis. Baca kapan saja, di mana saja.
        </p>
      </div>

      <!-- Lencana Mengambang dengan animasi kustom CSS -->
      <div in:fly={{ y: 30, duration: 800, delay: 400 }} class="animate-float relative z-10 mt-12 bg-white/20 backdrop-blur-md border border-white/30 rounded-2xl p-4 shadow-lg max-w-[220px] hidden md:block">
        <div class="flex items-center gap-4">
          <div class="text-3xl animate-pulse">✨</div>
          <div>
            <p class="text-xs text-white/80 font-semibold mb-0.5">Akses Spesial</p>
            <p class="text-sm font-bold">1000+ Buku Gratis!</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Bagian Kanan: Form -->
    <div class="w-full md:w-7/12 p-8 sm:p-10 lg:p-14 flex flex-col justify-center bg-white/40">
      <div class="max-w-md mx-auto w-full">
        
        <div in:fly={{ y: 20, duration: 600, delay: 300 }}>
          <h3 class="text-3xl font-extrabold text-gray-800 mb-2">Daftar Akun Baru</h3>
          <p class="text-sm text-gray-500 mb-8 font-medium">
            Sudah punya akun? 
            <Link to="/login"><span class="text-[#9B9AFF] font-bold hover:text-[#8382ff] hover:underline transition cursor-pointer">Login di sini</span></Link>
          </p>
        </div>

        <form class="space-y-5" on:submit|preventDefault={() => console.log('Form disubmit!')}>
          
          <!-- Input Nama -->
          <div class="relative group" in:fly={{ y: 20, duration: 600, delay: 400 }}>
            <label class="block text-sm font-bold text-gray-700 mb-1.5 ml-1 transition-colors group-focus-within:text-[#9B9AFF]" for="nama">Nama Lengkap</label>
            <div class="relative">
              <span class="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400 group-focus-within:text-[#9B9AFF] transition-colors group-hover:scale-110">👤</span>
              <input type="text" id="nama" placeholder="Indra Suliwa" class="w-full pl-11 pr-5 py-3.5 rounded-2xl border-2 border-gray-100 bg-white focus:outline-none focus:ring-4 focus:ring-[#9B9AFF]/20 focus:border-[#9B9AFF] transition-all font-medium text-gray-700 hover:border-gray-200">
            </div>
          </div>

          <!-- Input Email -->
          <div class="relative group" in:fly={{ y: 20, duration: 600, delay: 500 }}>
            <label class="block text-sm font-bold text-gray-700 mb-1.5 ml-1 transition-colors group-focus-within:text-[#9B9AFF]" for="email">Alamat Email</label>
            <div class="relative">
              <span class="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400 group-focus-within:text-[#9B9AFF] transition-colors group-hover:scale-110">✉️</span>
              <input type="email" id="email" placeholder="contoh@email.com" class="w-full pl-11 pr-5 py-3.5 rounded-2xl border-2 border-gray-100 bg-white focus:outline-none focus:ring-4 focus:ring-[#9B9AFF]/20 focus:border-[#9B9AFF] transition-all font-medium text-gray-700 hover:border-gray-200">
            </div>
          </div>

          <!-- Input Password -->
          <div class="relative group" in:fly={{ y: 20, duration: 600, delay: 600 }}>
            <label class="block text-sm font-bold text-gray-700 mb-1.5 ml-1 transition-colors group-focus-within:text-[#9B9AFF]" for="password">Password</label>
            <div class="relative">
              <span class="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400 group-focus-within:text-[#9B9AFF] transition-colors group-hover:scale-110">🔒</span>
              <input type="{showPassword ? 'text' : 'password'}" id="password" placeholder="••••••••" class="w-full pl-11 pr-16 py-3.5 rounded-2xl border-2 border-gray-100 bg-white focus:outline-none focus:ring-4 focus:ring-[#9B9AFF]/20 focus:border-[#9B9AFF] transition-all font-medium text-gray-700 hover:border-gray-200">
              <button type="button" on:click={togglePassword} class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-[#9B9AFF] focus:outline-none transition-transform hover:scale-105">
                {#if showPassword}
                  <span class="text-xs font-bold bg-[#9B9AFF]/10 text-[#9B9AFF] px-2 py-1 rounded-md">Tutup</span>
                {:else}
                  <span class="text-xs font-bold bg-gray-100 px-2 py-1 rounded-md">Lihat</span>
                {/if}
              </button>
            </div>
          </div>

          <!-- Checkbox S&K -->
          <div class="flex items-center gap-3 pt-2" in:fly={{ y: 20, duration: 600, delay: 700 }}>
            <input type="checkbox" id="terms" class="w-5 h-5 rounded border-2 border-gray-300 text-[#9B9AFF] focus:ring-[#9B9AFF] transition cursor-pointer">
            <label for="terms" class="text-sm text-gray-500 font-medium cursor-pointer hover:text-gray-700 transition">
              Saya setuju dengan <a href="#" class="text-[#9B9AFF] hover:underline">Syarat & Ketentuan</a>
            </label>
          </div>

          <!-- Tombol Submit -->
          <div in:fly={{ y: 20, duration: 600, delay: 800 }}>
            <button type="submit" class="w-full mt-6 bg-[#9B9AFF] hover:bg-[#8382ff] text-white font-extrabold text-lg py-4 rounded-2xl shadow-[0_8px_20px_rgba(155,154,255,0.3)] hover:shadow-[0_12px_25px_rgba(155,154,255,0.5)] transition-all transform hover:-translate-y-1 active:translate-y-0 active:scale-95 flex justify-center items-center gap-2">
              Daftar Sekarang 🚀
            </button>
          </div>

          <!-- Divider -->
          <div class="flex items-center gap-4 my-6 opacity-60" in:fade={{ duration: 600, delay: 900 }}>
            <div class="h-px bg-gray-300 flex-1"></div>
            <span class="text-xs font-bold text-gray-400 uppercase tracking-wider">Atau</span>
            <div class="h-px bg-gray-300 flex-1"></div>
          </div>

          <!-- Tombol Google -->
          <div in:fly={{ y: 20, duration: 600, delay: 1000 }}>
            <button type="button" class="w-full bg-white hover:bg-gray-50 text-gray-700 font-bold py-3.5 rounded-2xl border-2 border-gray-100 shadow-sm hover:shadow-md transition-all active:scale-95 flex items-center justify-center gap-3">
              <svg width="24" height="24" class="flex-shrink-0 group-hover:scale-110 transition-transform" viewBox="0 0 24 24">
                <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
              </svg>
              Daftar dengan Google
            </button>
          </div>

        </form>
      </div>
    </div>
  </div>
</main>

<style>
  /* Animasi CSS kustom untuk memberikan nyawa pada elemen statis */
  @keyframes float {
    0%, 100% { transform: translateY(0px) rotate(-3deg); }
    50% { transform: translateY(-12px) rotate(0deg); }
  }
  .animate-float {
    animation: float 5s ease-in-out infinite;
  }
  
  @keyframes blob {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.05); }
  }
  .animate-blob {
    animation: blob 8s ease-in-out infinite;
  }
  .animation-delay-2000 {
    animation-delay: 2s;
  }
</style>