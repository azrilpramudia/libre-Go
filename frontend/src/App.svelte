<script>
  import { Router, Route } from "svelte-routing";
  
  // Import semua halaman yang ada di folder pages
  import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import Daftar from "./pages/Daftar.svelte";
  import Dashboard from "./pages/Dashboard.svelte";
  import Koleksi from "./pages/Koleksi.svelte";

  export let url = "";
</script>

<Router {url}>
  <!-- Rute Halaman Utama -->
  <Route path="/">
    <Home />
  </Route>

  <!-- Rute Halaman Login -->
  <Route path="/login">
    <Login />
  </Route>

  <!-- Rute Halaman Pendaftaran -->
  <Route path="/daftar">
    <Daftar />
  </Route>

  <Route path="/koleksi">
    <Koleksi />
  </Route>

  <!-- Rute Halaman Dashboard (Protected) -->
  <Route path="/dashboard">
    <Dashboard />
  </Route>
</Router>