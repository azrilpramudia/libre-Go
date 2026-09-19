<footer class="mt-32 bg-[#9B9AFF] text-white rounded-t-[4rem] relative z-20 overflow-hidden shadow-[0_-10px_40px_rgba(155,154,255,0.3)]">
  
  <!-- Ornamen Dekoratif Glowing di dalam Footer -->
  <div class="absolute -top-24 -right-10 w-72 h-72 bg-[#B8B8FF] rounded-full mix-blend-screen blur-3xl opacity-80 pointer-events-none"></div>
  <div class="absolute -bottom-20 -left-10 w-80 h-80 bg-[#AADEFF] rounded-full mix-blend-screen blur-3xl opacity-60 pointer-events-none"></div>

  <div class="relative max-w-7xl mx-auto px-8 pt-20 pb-8 z-10">
    <!-- Grid Konten Footer -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-12 mb-16">
      
      <!-- Kolom 1: Branding & Info Utama -->
      <div class="col-span-1 md:col-span-5 space-y-6">
        <div class="flex items-center gap-3">
          <!-- Ikon Buku dengan style miring sedikit biar "fun" -->
          <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center text-2xl shadow-lg transform -rotate-6">
            📚
          </div>
          <span class="text-3xl font-extrabold tracking-wide">PerpusFun</span>
        </div>
        <p class="text-white/90 leading-relaxed font-medium max-w-sm">
          Membaca jadi lebih seru dan asyik. Kelola koleksi buku digitalmu dan kumpulkan poin reward tanpa batas!
        </p>
      </div>

      <!-- Kolom 2: Menu Navigasi Cepat -->
      <div class="col-span-1 md:col-span-3 space-y-6">
        <h4 class="text-xl font-bold tracking-wide text-white drop-shadow-sm">Jelajahi</h4>
        <ul class="space-y-4 text-white/80 font-medium">
          <li><a href="#" class="hover:text-white hover:translate-x-2 transition-all duration-300 inline-block">Koleksi Terbaru</a></li>
          <li><a href="#" class="hover:text-white hover:translate-x-2 transition-all duration-300 inline-block">Cara Pinjam</a></li>
          <li><a href="#" class="hover:text-white hover:translate-x-2 transition-all duration-300 inline-block">Leaderboard</a></li>
        </ul>
      </div>

      <!-- Kolom 3: Call to Action (CTA) -->
      <div class="col-span-1 md:col-span-4 space-y-6">
        <h4 class="text-xl font-bold tracking-wide text-white drop-shadow-sm">Join Komunitas</h4>
        <p class="text-white/80 font-medium text-sm leading-relaxed">
          Punya pertanyaan atau sekadar ingin diskusi buku bareng developer?
        </p>
        <button class="w-full bg-white text-[#9B9AFF] hover:bg-[#f4f9f6] font-bold py-3.5 px-6 rounded-full shadow-lg transition transform hover:-translate-y-1 flex items-center justify-center gap-2">
          <span>💬</span> Gabung Discord Sekarang
        </button>
      </div>
    </div>

    <!-- Garis Pemisah & Copyright -->
    <div class="border-t border-white/30 pt-8 flex flex-col md:flex-row justify-between items-center gap-4 text-sm font-medium text-white/70">
      <p>© 2026 PerpusFun. Dibuat menggunakan Svelte.</p>
      <div class="flex gap-6">
        <a href="#" class="hover:text-white transition">Syarat & Ketentuan</a>
        <a href="#" class="hover:text-white transition">Kebijakan Privasi</a>
      </div>
    </div>
  </div>
</footer>