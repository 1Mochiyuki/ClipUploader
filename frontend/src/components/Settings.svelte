<script>
  import { push } from "svelte-spa-router";
</script>

<div class="div-button">
  <button
    on:click={() => push("/")}
    class="button-4"
    style="background-color: rgba(0,0,0,0); border: 0px solid;  box-shadow: none; float: left"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="30px"
      viewBox="0 -960 960 960"
      width="24px"
      fill="#FFFFFF"
      ><path
        d="M400-240 160-480l240-240 56 58-142 142h486v80H314l142 142-56 58Z"
      /></svg
    >
  </button>
</div>

<style>
</style>
