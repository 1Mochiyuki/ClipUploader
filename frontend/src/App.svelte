<script>
  import Router from "svelte-spa-router";
  import routes from "./routes";
</script>

<main>
  <Router {routes} />
</main>

<style>
  :global(.visually-hidden) {
    position: absolute !important;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px 1px 1px 1px);
    clip: rect(1px, 1px, 1px, 1px);
    white-space: nowrap;
  }
</style>
