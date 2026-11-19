<script>
  import { Link } from "svelte-routing";
</script>

<footer class="mt-20 border-t border-slate-200 bg-white">
  <div class="mx-auto flex max-w-7xl flex-col gap-6 px-6 py-8 sm:flex-row sm:items-center sm:justify-between">
    <Link to="/" class="flex items-center gap-3">
      <img src="/logo-perpusfun.svg" alt="" class="h-9 w-9 shrink-0" />
      <span class="font-bold text-slate-900">Perpus<span class="text-blue-600">Fun</span></span>
    </Link>
    <p class="text-sm text-slate-500">Temukan bacaan berikutnya, kapan saja.</p>
    <nav aria-label="Navigasi footer" class="flex gap-5 text-sm font-medium text-slate-600">
      <Link to="/koleksi" class="transition hover:text-blue-600">Koleksi</Link>
      <Link to="/login" class="transition hover:text-blue-600">Akun</Link>
    </nav>
  </div>
</footer>
