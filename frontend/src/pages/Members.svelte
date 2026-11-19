<script>
  import { onMount } from "svelte";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
  import { membersAPI } from "../lib/api.js";

  const pageSize = 20;
  let members = $state([]);
  let offset = $state(0);
  let hasNextPage = $state(false);
  let isLoading = $state(true);
  let errorMessage = $state("");

  onMount(loadMembers);

  async function loadMembers(nextOffset = offset) {
    isLoading = true;
    errorMessage = "";
    try {
      const response = await membersAPI.list(pageSize, nextOffset);
      members = response.data || [];
      offset = nextOffset;
      hasNextPage = members.length === pageSize;
    } catch (error) {
      errorMessage = error.data?.error || error.message || "Gagal memuat daftar member.";
    } finally {
      isLoading = false;
    }
  }
</script>

<main class="min-h-screen bg-slate-50 text-slate-900">
  <div class="mx-auto max-w-7xl px-5 py-5 sm:px-8">
    <Navbar />
    <header class="mb-7 mt-10">
      <div>
        <p class="text-sm font-semibold text-blue-600">PerpusFun</p>
        <h1 class="mt-1 text-3xl font-bold tracking-tight text-slate-900">Daftar member</h1>
        <p class="mt-2 text-sm text-slate-500">Anggota terdaftar di PerpusFun.</p>
      </div>
    </header>

    {#if isLoading}
      <p class="rounded-2xl border border-slate-200 bg-white p-8 text-center text-sm text-slate-500">Memuat daftar member...</p>
    {:else if errorMessage}
      <div role="alert" class="rounded-2xl border border-red-100 bg-red-50 p-6 text-center text-sm text-red-700">
        <p>{errorMessage}</p>
        <button onclick={() => loadMembers()} class="mt-3 font-semibold underline">Coba lagi</button>
      </div>
    {:else if members.length}
      <div class="overflow-x-auto rounded-2xl border border-slate-200 bg-white shadow-sm">
        <table class="w-full min-w-[520px] text-left">
          <thead class="border-b border-slate-200 bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th class="px-6 py-4 font-semibold">Nama member</th>
              <th class="px-6 py-4 font-semibold">Tanggal bergabung</th>
            </tr>
          </thead>
          <tbody>
            {#each members as member (member.id)}
              <tr class="border-b border-slate-100 last:border-0 hover:bg-slate-50">
                <td class="px-6 py-4 text-sm font-medium text-slate-800">{member.full_name}</td>
                <td class="px-6 py-4 text-sm text-slate-600">
                  {member.created_at
                    ? new Date(member.created_at).toLocaleDateString("id-ID", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                      })
                    : "-"}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>

      <nav aria-label="Navigasi daftar member" class="mt-6 flex items-center justify-between">
        <button
          onclick={() => loadMembers(Math.max(0, offset - pageSize))}
          disabled={offset === 0 || isLoading}
          class="rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-semibold text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50"
        >
          Sebelumnya
        </button>
        <span class="text-sm text-slate-500">Halaman {Math.floor(offset / pageSize) + 1}</span>
        <button
          onclick={() => loadMembers(offset + pageSize)}
          disabled={!hasNextPage || isLoading}
          class="rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-semibold text-white transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-50"
        >
          Berikutnya
        </button>
      </nav>
    {:else}
      <p class="rounded-2xl border border-slate-200 bg-white p-8 text-center text-sm text-slate-500">Belum ada member terdaftar.</p>
    {/if}
  </div>
  <Footer />
</main>
