<script>
  import { Router, Route } from "svelte-routing";
  import Home from "./pages/Home.svelte";
  import Koleksi from "./pages/Koleksi.svelte";
  import Login from "./pages/Login.svelte";

  export let url = "";
</script>

<Router {url}>
  
  <!-- Cara Svelte 5: Masukkan komponen ke dalam tag Route, BUKAN sebagai atribut -->
  <Route path="/">
    <Home />
  </Route>
  
  <Route path="/koleksi">
    <Koleksi />
  </Route>
  
  <Route path="/login">
    <Login />
  </Route>

</Router>