<script>
  import { onMount } from "svelte";
  import { Link } from "svelte-routing";
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

<main class="min-h-screen bg-[#f4f9f6] p-6 text-gray-800 sm:p-8">
  <div class="mx-auto max-w-5xl">
    <header class="mb-8 flex flex-wrap items-center justify-between gap-4">
      <div>
        <h1 class="text-3xl font-extrabold text-gray-900">Daftar Member</h1>
        <p class="mt-2 text-sm text-gray-500">Anggota terdaftar di PerpusFun.</p>
      </div>
      <Link to="/" class="rounded-full bg-[#9B9AFF] px-5 py-2.5 font-bold text-white transition hover:bg-[#8382ff]">
        ← Kembali ke Home
      </Link>
    </header>

    {#if isLoading}
      <p class="rounded-3xl bg-white p-8 text-center text-gray-500">Memuat daftar member...</p>
    {:else if errorMessage}
      <div role="alert" class="rounded-3xl bg-red-100 p-6 text-center text-red-700">
        <p>{errorMessage}</p>
        <button onclick={() => loadMembers()} class="mt-3 font-bold underline">Coba lagi</button>
      </div>
    {:else if members.length}
      <div class="overflow-x-auto rounded-3xl border border-gray-100 bg-white shadow-sm">
        <table class="w-full text-left">
          <thead class="border-b border-gray-100 bg-gray-50 text-sm text-gray-600">
            <tr>
              <th class="px-6 py-4 font-bold">Nama Member</th>
              <th class="px-6 py-4 font-bold">Tanggal Bergabung</th>
            </tr>
          </thead>
          <tbody>
            {#each members as member (member.id)}
              <tr class="border-b border-gray-100 last:border-0">
                <td class="px-6 py-4 font-semibold text-gray-800">{member.full_name}</td>
                <td class="px-6 py-4 text-sm text-gray-600">
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
          class="rounded-full border border-gray-200 bg-white px-5 py-2.5 text-sm font-bold text-gray-700 disabled:cursor-not-allowed disabled:opacity-50"
        >
          Sebelumnya
        </button>
        <span class="text-sm text-gray-500">Halaman {Math.floor(offset / pageSize) + 1}</span>
        <button
          onclick={() => loadMembers(offset + pageSize)}
          disabled={!hasNextPage || isLoading}
          class="rounded-full bg-[#9B9AFF] px-5 py-2.5 text-sm font-bold text-white disabled:cursor-not-allowed disabled:opacity-50"
        >
          Berikutnya
        </button>
      </nav>
    {:else}
      <p class="rounded-3xl bg-white p-8 text-center text-gray-500">Belum ada member terdaftar.</p>
    {/if}
  </div>
</main>
