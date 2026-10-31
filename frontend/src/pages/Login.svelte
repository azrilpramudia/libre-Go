<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, authIsLoading, authError } from "../lib/auth.js";

  let isLogin = $state(true);
  let nama = $state("");
  let email = $state("");
  let phone = $state("");
  let password = $state("");
  let localError = $state("");

  async function handleSubmit(e) {
    e.preventDefault();
    localError = "";

    try {
      if (isLogin) {
        // Login
        await auth.login(email, password);
        // Navigate to dashboard on success
        navigate("/dashboard");
      } else {
        // Register
        await auth.register(nama, email, phone, password);
        // After successful register, automatically login
        await auth.login(email, password);
        navigate("/dashboard");
      }
    } catch (err) {
      localError = err.data?.error || err.message || "An error occurred";
    }
  }

  function clearError() {
    localError = "";
    auth.clearError();
  }
</script>

<main class="relative min-h-screen bg-[#f4f9f6] flex justify-center items-center overflow-hidden p-6 font-sans">
  
  <!-- Efek Background Orbs -->
  <div class="absolute top-0 left-10 w-96 h-96 bg-[#AADEFF]/60 rounded-full blur-[100px] pointer-events-none animate-pulse"></div>
  <div class="absolute bottom-0 right-10 w-96 h-96 bg-[#9B9AFF]/40 rounded-full blur-[100px] pointer-events-none"></div>

  <!-- Container Form (Glassmorphism) -->
  <div class="relative z-10 w-full max-w-md bg-white/40 backdrop-blur-2xl border border-white/60 p-10 rounded-[2.5rem] shadow-[0_20px_50px_rgba(0,0,0,0.05)]">
    
    <!-- Tombol Kembali ke Home -->
    <div class="mb-8">
      <Link to="/">
        <span class="text-sm font-bold text-gray-500 hover:text-[#9B9AFF] transition cursor-pointer">
          ← Kembali
        </span>
      </Link>
    </div>

    <!-- Header Form -->
    <h2 class="text-4xl font-extrabold text-gray-900 mb-2">
      {isLogin ? 'Selamat Datang! 👋' : 'Gabung Member ✨'}
    </h2>
    <p class="text-sm text-gray-600 mb-8 font-medium">
      {isLogin ? 'Silakan login untuk mulai meminjam buku.' : 'Daftar sekarang dan nikmati ribuan koleksi digital.'}
    </p>

    <!-- Form -->
    <form onsubmit={handleSubmit} class="space-y-5">
      
      <!-- Error Message -->
      {#if localError || $authError}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-2xl text-sm font-medium">
          {localError || $authError}
          <button type="button" onclick={clearError} class="ml-2 font-bold hover:underline">✕</button>
        </div>
      {/if}

      <!-- Input Nama (Hanya muncul jika di mode Register) -->
      {#if !isLogin}
        <div>
          <label for="inputNama" class="block text-xs font-bold text-gray-700 uppercase tracking-wider mb-2">Nama Lengkap</label>
          <input 
            id="inputNama"
            type="text" 
            bind:value={nama} 
            required 
            placeholder="Ketik nama kamu..."
            class="w-full bg-white/60 border border-white/80 rounded-2xl px-5 py-3 focus:outline-none focus:ring-2 focus:ring-[#9B9AFF] transition shadow-sm"
          />
        </div>

        <div>
          <label for="inputPhone" class="block text-xs font-bold text-gray-700 uppercase tracking-wider mb-2">No. Telepon</label>
          <input 
            id="inputPhone"
            type="tel" 
            bind:value={phone} 
            placeholder="08xxxxxxxxxx"
            class="w-full bg-white/60 border border-white/80 rounded-2xl px-5 py-3 focus:outline-none focus:ring-2 focus:ring-[#9B9AFF] transition shadow-sm"
          />
        </div>
      {/if}

      <div>
        <label for="inputEmail" class="block text-xs font-bold text-gray-700 uppercase tracking-wider mb-2">Email</label>
        <input 
          id="inputEmail"
          type="email" 
          bind:value={email} 
          required 
          placeholder="contoh@email.com"
          class="w-full bg-white/60 border border-white/80 rounded-2xl px-5 py-3 focus:outline-none focus:ring-2 focus:ring-[#9B9AFF] transition shadow-sm"
        />
      </div>

      <div>
        <label for="inputPassword" class="block text-xs font-bold text-gray-700 uppercase tracking-wider mb-2">Password</label>
        <input 
          id="inputPassword"
          type="password" 
          bind:value={password} 
          required 
          placeholder="••••••••"
          class="w-full bg-white/60 border border-white/80 rounded-2xl px-5 py-3 focus:outline-none focus:ring-2 focus:ring-[#9B9AFF] transition shadow-sm"
        />
      </div>

      <!-- Tombol Submit -->
      <button 
        type="submit" 
        disabled={$authIsLoading}
        class="w-full mt-4 bg-[#9B9AFF] hover:bg-[#8382ff] text-white font-bold py-4 rounded-2xl shadow-lg transition transform hover:-translate-y-1 disabled:opacity-70 disabled:transform-none"
      >
        {$authIsLoading ? 'Memproses...' : (isLogin ? 'Masuk Sekarang' : 'Daftar Akun')}
      </button>
    </form>

    <!-- Toggle Login/Register -->
    <div class="mt-8 text-center text-sm font-medium text-gray-600">
      {isLogin ? 'Belum punya akun?' : 'Sudah jadi member?'} 
      <!-- Mengubah nilai isLogin saat di-klik -->
      <button 
        type="button" 
        onclick={() => isLogin = !isLogin} 
        class="text-[#9B9AFF] font-bold hover:underline ml-1 focus:outline-none"
      >
        {isLogin ? 'Daftar di sini' : 'Login dong'}
      </button>
    </div>

  </div>
</main>