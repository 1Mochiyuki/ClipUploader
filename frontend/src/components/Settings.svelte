<script>
  import { push } from "svelte-spa-router";
  import HostInput from "./HostInput.svelte";
  import TimeoutDuration from "./TimeoutDuration.svelte";

  export let selectedHost;
  export let timeoutDuration = 0;

  const handleSaveSettings = () => {
    alert(`host: ${selectedHost}\ntimeout duration: ${timeoutDuration}`);
    console.log(`host: ${selectedHost}\ntimeout duration: ${timeoutDuration}`);
  };
</script>

<div>
  <button
    on:click={() => push("/")}
    class="back"
    style="background-color: rgba(0,0,0,0); border: 0;  box-shadow: none; float: left; "
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="30px"
      viewBox="0 -960 960 960"
      width="30px"
      fill="#FFFFFF"
      ><path
        d="M400-240 160-480l240-240 56 58-142 142h486v80H314l142 142-56 58Z"
      /></svg
    >
  </button>
  <div style="justify-content: space-between; padding-top: 20px;">
    <div>
      <HostInput bind:value={selectedHost}></HostInput>

      <TimeoutDuration bind:timeoutDuration></TimeoutDuration>
    </div>
  </div>
  <div>
    <button class="button-4" on:click={handleSaveSettings}>Save</button>
  </div>
</div>

<style>
  .back:hover {
    cursor: pointer;
  }
</style>
