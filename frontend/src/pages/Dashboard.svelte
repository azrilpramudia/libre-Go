<script>
  import { Link, navigate } from "svelte-routing";
  import { auth, currentMember, isAuthenticated } from "../lib/auth.js";
  import { loansAPI } from "../lib/api.js";

  let loans = $state([]);
  let isLoadingLoans = $state(false);
  let loansError = $state("");

  // Redirect to login if not authenticated
  $effect(() => {
    if (!$isAuthenticated) {
      navigate("/login");
    }
  });

  // Fetch loans when component mounts
  $effect(() => {
    if ($currentMember?.id) {
      fetchLoans();
    }
  });

  async function fetchLoans() {
    if (!$currentMember?.id) return;

    isLoadingLoans = true;
    loansError = "";

    try {
      const response = await loansAPI.listByMember($currentMember.id);
      if (response.success) {
        loans = response.data || [];
      } else {
        loansError = response.error || "Failed to load loans";
      }
    } catch (err) {
      loansError = err.data?.error || err.message || "Error loading loans";
    } finally {
      isLoadingLoans = false;
    }
  }

  function handleLogout() {
    auth.logout();
    navigate("/");
  }

  function getStatusBadgeColor(status) {
    const statusLower = status?.toLowerCase() || "";
    if (statusLower === "active" || statusLower === "borrowed") {
      return "bg-blue-100 text-blue-700";
    } else if (statusLower === "returned") {
      return "bg-green-100 text-green-700";
    } else {
      return "bg-gray-100 text-gray-700";
    }
  }

  function formatDate(dateString) {
    if (!dateString) return "-";
    const date = new Date(dateString);
    return date.toLocaleDateString("id-ID", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  }
</script>

<main class="relative min-h-screen bg-[#fdfdfd] text-gray-900 font-sans overflow-x-hidden flex flex-col">
  <!-- Background Orbs -->
  <div class="absolute -top-40 -left-40 w-[600px] h-[600px] bg-[#DFFCE2]/40 rounded-full blur-[120px] pointer-events-none"></div>
  <div class="absolute top-[20%] -right-20 w-[700px] h-[700px] bg-[#B8B8FF]/20 rounded-full blur-[150px] pointer-events-none"></div>

  <div class="relative z-10 max-w-6xl mx-auto px-6 py-6 w-full flex-grow flex flex-col">
    <!-- Navbar -->
    <nav class="flex flex-wrap justify-between items-center mb-10 bg-white/70 backdrop-blur-md px-6 py-4 rounded-full border border-gray-100 shadow-sm">
      <div class="flex items-center gap-2">
        <div class="text-3xl">📚</div>
        <span class="text-xl font-extrabold text-[#9B9AFF] tracking-wide">PerpusFun</span>
      </div>

      <div class="flex items-center gap-4">
        <Link to="/koleksi">
          <span class="text-sm font-semibold text-gray-600 hover:text-[#9B9AFF] transition cursor-pointer">
            Koleksi
          </span>
        </Link>
        <button
          onclick={handleLogout}
          class="text-sm font-bold text-white bg-[#9B9AFF] hover:bg-[#8382ff] py-2.5 px-6 rounded-full shadow transition transform hover:-translate-y-1"
        >
          Logout
        </button>
      </div>
    </nav>

    <!-- Main Content -->
    <div class="space-y-8">
      <!-- Profile Section -->
      <section class="bg-white/60 backdrop-blur-lg rounded-3xl p-8 shadow-sm border border-white/50">
        <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
          <div>
            <h1 class="text-4xl font-extrabold text-gray-800 mb-2">
              Selamat Datang, {$currentMember?.full_name || "Member"}! 👋
            </h1>
            <p class="text-gray-600 font-medium">
              Email: <span class="text-[#9B9AFF] font-bold">{$currentMember?.email}</span>
            </p>
            {#if $currentMember?.phone}
              <p class="text-gray-600 font-medium">
                Telepon: <span class="text-gray-700 font-bold">{$currentMember.phone}</span>
              </p>
            {/if}
          </div>

          <div class="bg-gradient-to-br from-[#9B9AFF] to-[#B8B8FF] rounded-2xl p-6 text-white shadow-lg">
            <div class="text-4xl font-extrabold">{loans.length}</div>
            <p class="text-sm font-semibold mt-1 opacity-90">Buku Dipinjam</p>
          </div>
        </div>
      </section>

      <!-- Active Loans Section -->
      <section>
        <div class="flex justify-between items-center mb-6 border-b border-gray-200 pb-4">
          <h2 class="text-2xl font-extrabold text-gray-800">
            Riwayat <span class="text-transparent bg-clip-text bg-gradient-to-r from-[#9B9AFF] to-[#60c4ff]">Peminjaman</span>
          </h2>
          <button
            onclick={fetchLoans}
            disabled={isLoadingLoans}
            class="text-sm font-bold text-[#9B9AFF] hover:text-[#8382ff] transition disabled:opacity-50"
          >
            🔄 Refresh
          </button>
        </div>

        <!-- Loading State -->
        {#if isLoadingLoans}
          <div class="bg-white/60 backdrop-blur-lg rounded-3xl p-12 shadow-sm border border-white/50 text-center">
            <div class="inline-block">
              <div class="w-12 h-12 border-4 border-[#9B9AFF]/20 border-t-[#9B9AFF] rounded-full animate-spin"></div>
            </div>
            <p class="text-gray-600 font-medium mt-4">Memuat data peminjaman...</p>
          </div>
        {/if}

        <!-- Error State -->
        {#if loansError && !isLoadingLoans}
          <div class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-2xl text-sm font-medium">
            {loansError}
            <button
              type="button"
              onclick={() => (loansError = "")}
              class="ml-2 font-bold hover:underline"
            >
              ✕
            </button>
          </div>
        {/if}

        <!-- Loans Table -->
        {#if !isLoadingLoans && loans.length > 0}
          <div class="overflow-x-auto bg-white/60 backdrop-blur-lg rounded-3xl shadow-sm border border-white/50">
            <table class="w-full">
              <thead>
                <tr class="border-b border-gray-200">
                  <th class="text-left px-6 py-4 font-bold text-gray-700">ID Peminjaman</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">ID Buku</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">Tanggal Pinjam</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">Jatuh Tempo</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">Dikembalikan</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">Status</th>
                  <th class="text-left px-6 py-4 font-bold text-gray-700">Denda</th>
                </tr>
              </thead>
              <tbody>
                {#each loans as loan (loan.id)}
                  <tr class="border-b border-gray-100 hover:bg-white/40 transition">
                    <td class="px-6 py-4 text-sm text-gray-700 font-mono">{loan.id.slice(0, 8)}...</td>
                    <td class="px-6 py-4 text-sm text-gray-700 font-mono">{loan.book_id.slice(0, 8)}...</td>
                    <td class="px-6 py-4 text-sm text-gray-600">{formatDate(loan.borrowed_at)}</td>
                    <td class="px-6 py-4 text-sm text-gray-600">{formatDate(loan.due_at)}</td>
                    <td class="px-6 py-4 text-sm text-gray-600">{formatDate(loan.returned_at)}</td>
                    <td class="px-6 py-4">
                      <span class="text-xs font-bold px-3 py-1 rounded-full {getStatusBadgeColor(loan.status)}">
                        {loan.status || "Unknown"}
                      </span>
                    </td>
                    <td class="px-6 py-4 text-sm font-bold text-gray-700">
                      Rp {loan.fine_amount?.toLocaleString("id-ID") || "0"}
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
          </div>
        {/if}

        <!-- Empty State -->
        {#if !isLoadingLoans && loans.length === 0 && !loansError}
          <div class="bg-white/60 backdrop-blur-lg rounded-3xl p-12 shadow-sm border border-white/50 text-center">
            <div class="text-5xl mb-4">📚</div>
            <h3 class="text-xl font-bold text-gray-700 mb-2">Belum Ada Peminjaman</h3>
            <p class="text-gray-500 font-medium mb-6">Mulai pinjam buku dari koleksi kami sekarang!</p>
            <Link to="/koleksi">
              <button class="bg-[#9B9AFF] hover:bg-[#8382ff] text-white font-bold py-3 px-8 rounded-full shadow-lg transition transform hover:-translate-y-1">
                Jelajahi Koleksi →
              </button>
            </Link>
          </div>
        {/if}
      </section>
    </div>
  </div>
</main>
