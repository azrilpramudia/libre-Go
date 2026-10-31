<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, currentMember, isAuthenticated } from "../lib/auth.js";

  let { searchQuery = $bindable("") } = $props();

  function scrollToSection(sectionId) {
    document.getElementById(sectionId)?.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }

  function logout() {
    auth.logout();
    navigate("/");
  }
</script>

<nav class="flex flex-wrap justify-between items-center mb-16 bg-white/70 backdrop-blur-md px-6 py-4 rounded-full border border-gray-100 shadow-sm">
  <div class="flex items-center gap-2">
    <div class="text-3xl">📚</div>
    <span class="text-xl font-extrabold text-[#9B9AFF] tracking-wide">PerpusFun</span>
  </div>

  <div class="hidden lg:flex gap-8 text-sm font-semibold text-gray-600">
    {#if $isAuthenticated}
      <Link to="/koleksi">
        <span class="hover:text-[#9B9AFF] transition cursor-pointer">Koleksi Buku</span>
      </Link>
    {:else}
      <Link to="/" onclick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
        <span class="hover:text-[#9B9AFF] transition cursor-pointer">Home</span>
      </Link>
      <button
        type="button"
        onclick={() => scrollToSection("tentang")}
        class="hover:text-[#9B9AFF] transition cursor-pointer"
      >
        Tentang
      </button>
      <button
        type="button"
        onclick={() => scrollToSection("keunggulan")}
        class="hover:text-[#9B9AFF] transition cursor-pointer"
      >
        Keunggulan
      </button>
    {/if}
  </div>

  <div class="flex items-center gap-4">
    {#if $isAuthenticated}
      <div class="flex items-center gap-4">
        <div class="relative hidden md:block w-48 lg:w-64">
          <input
            type="text"
            bind:value={searchQuery}
            placeholder="Cari buku..."
            class="w-full pl-5 pr-10 py-2.5 rounded-full border border-gray-200 bg-gray-50 focus:outline-none focus:border-[#9B9AFF] text-sm text-gray-700 transition"
          />
          <span class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 text-sm">🔍</span>
        </div>
        <Link to="/dashboard">
          <span class="text-sm font-bold text-gray-700 hover:text-[#9B9AFF] transition cursor-pointer">
            Hi, {$currentMember?.full_name?.split(" ")[0] || "Member"}
          </span>
        </Link>
        <button
          type="button"
          onclick={logout}
          class="rounded-full bg-[#9B9AFF] px-5 py-2.5 text-sm font-bold text-white transition hover:bg-[#8382ff]"
        >
          Logout
        </button>
      </div>
    {:else}
      <Link to="/login">
        <span class="rounded-full bg-[#9B9AFF] px-5 py-2.5 text-sm font-bold text-white transition hover:bg-[#8382ff]">
          Login / Register
        </span>
      </Link>
    {/if}
  </div>
</nav>
