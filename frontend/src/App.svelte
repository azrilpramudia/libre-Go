<script>
  import { Router, Route } from "svelte-routing";
  
  // Import semua halaman yang ada di folder pages
  import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import Daftar from "./pages/Daftar.svelte"; // Kita buat ini setelah ini

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
</Router>