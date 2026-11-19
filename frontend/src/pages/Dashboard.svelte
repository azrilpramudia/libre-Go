<script>
  import { Link, navigate } from "svelte-routing";
  import { currentMember, isAuthenticated } from "../lib/auth.js";
  import { loansAPI } from "../lib/api.js";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";

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

<main class="min-h-screen bg-slate-50 text-slate-900">
  <div class="mx-auto max-w-7xl px-5 py-5 sm:px-8">
    <Navbar />
    <section class="mt-10 rounded-3xl border border-slate-200 bg-white p-7 shadow-sm sm:p-9">
      <div class="flex flex-col gap-6 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p class="text-sm font-semibold text-blue-600">Ringkasan akun</p>
          <h1 class="mt-2 text-3xl font-bold tracking-tight text-slate-900">
            Halo, {$currentMember?.full_name || "Member"}
          </h1>
          <p class="mt-2 text-sm text-slate-600">{$currentMember?.email}</p>
          {#if $currentMember?.phone}
            <p class="mt-1 text-sm text-slate-500">{$currentMember.phone}</p>
          {/if}
        </div>
        <div class="min-w-40 rounded-2xl bg-blue-50 p-5">
          <p class="text-sm font-medium text-blue-700">Total peminjaman</p>
          <p class="mt-2 text-3xl font-bold text-blue-900">{loans.length}</p>
        </div>
      </div>
    </section>

    <section class="mt-10">
      <div class="mb-5 flex flex-wrap items-end justify-between gap-4">
        <div>
          <p class="text-sm font-semibold text-blue-600">Aktivitas akun</p>
          <h2 class="mt-1 text-2xl font-bold tracking-tight text-slate-900">Riwayat peminjaman</h2>
        </div>
        <button
          type="button"
          onclick={fetchLoans}
          disabled={isLoadingLoans}
          class="rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-semibold text-slate-700 transition hover:border-blue-200 hover:bg-blue-50 hover:text-blue-700 disabled:opacity-50"
        >
          {isLoadingLoans ? "Memuat..." : "Perbarui"}
        </button>
      </div>

        <!-- Loading State -->
        {#if isLoadingLoans}
          <div class="rounded-2xl border border-slate-200 bg-white p-12 text-center">
            <div class="mx-auto h-9 w-9 animate-spin rounded-full border-4 border-blue-100 border-t-blue-600"></div>
            <p class="mt-4 text-sm text-slate-500">Memuat data peminjaman...</p>
          </div>
        {/if}

        <!-- Error State -->
        {#if loansError && !isLoadingLoans}
          <div role="alert" class="rounded-2xl border border-red-100 bg-red-50 px-5 py-4 text-sm text-red-700">
            {loansError}
            <button
              type="button"
              onclick={() => (loansError = "")}
              class="ml-2 font-semibold underline"
            >
              Tutup
            </button>
          </div>
        {/if}

        <!-- Loans Table -->
        {#if !isLoadingLoans && loans.length > 0}
          <div class="overflow-x-auto rounded-2xl border border-slate-200 bg-white shadow-sm">
            <table class="w-full min-w-245 text-left">
              <thead class="bg-slate-50">
                <tr class="border-b border-slate-200">
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">ID peminjaman</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">ID buku</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">Tanggal pinjam</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">Jatuh tempo</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">Dikembalikan</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">Status</th>
                  <th class="px-5 py-4 text-xs font-semibold uppercase tracking-wide text-slate-500">Denda</th>
                </tr>
              </thead>
              <tbody>
                {#each loans as loan (loan.id)}
                  <tr class="border-b border-slate-100 last:border-0 hover:bg-slate-50">
                    <td class="px-5 py-4 font-mono text-xs text-slate-600">{loan.id.slice(0, 8)}...</td>
                    <td class="px-5 py-4 font-mono text-xs text-slate-600">{loan.book_id.slice(0, 8)}...</td>
                    <td class="px-5 py-4 text-sm text-slate-600">{formatDate(loan.borrowed_at)}</td>
                    <td class="px-5 py-4 text-sm text-slate-600">{formatDate(loan.due_at)}</td>
                    <td class="px-5 py-4 text-sm text-slate-600">{formatDate(loan.returned_at)}</td>
                    <td class="px-5 py-4">
                      <span class="rounded-full px-3 py-1 text-xs font-semibold {getStatusBadgeColor(loan.status)}">
                        {loan.status || "Unknown"}
                      </span>
                    </td>
                    <td class="px-5 py-4 text-sm font-semibold text-slate-700">
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
          <div class="rounded-2xl border border-slate-200 bg-white p-10 text-center">
            <h3 class="text-lg font-semibold text-slate-900">Belum ada peminjaman</h3>
            <p class="mt-2 text-sm text-slate-500">Mulai jelajahi koleksi buku untuk menemukan bacaanmu.</p>
            <Link to="/koleksi" class="mt-5 inline-flex rounded-xl bg-blue-600 px-5 py-3 text-sm font-semibold text-white transition hover:bg-blue-700">
              Jelajahi koleksi
            </Link>
          </div>
        {/if}
      </section>
  </div>
  <Footer />
</main>
