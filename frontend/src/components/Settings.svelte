<script>
  import { push } from "svelte-spa-router";

  const hosts = ["catbox", "pomf", "other"];
  export let selectedHost;
  export let timeoutDuration = 0;
</script>

<div>
  <button
    on:click={() => push("/")}
    class=""
    style="background-color: rgba(0,0,0,0); border: 2px solid #FFFFFF;  box-shadow: none; float: left"
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
  <div>
    <div class="center" style=" border: 2px solid gray;">
      <label for="selected-host">Select a host</label>
      <select name="selected-host" bind:value={selectedHost}>
        {#each hosts as host}
          <option value={host}>{host}</option>
        {/each}
      </select>
      <span>{!selectedHost ? "Nothing selected" : selectedHost}</span>
    </div>

    <div class="center" style=" border: 2px solid gray;">
      <label for="timeout-duration">Timeout Duration</label>
      <input
        name="timeout-duration"
        type="number"
        placeholder="5000"
        bind:value={timeoutDuration}
      />

      <label for="timeoutDuration">Timeout Duration: {timeoutDuration}</label>
    </div>
  </div>
</div>

<style>
  .center {
    margin: auto;
    width: 50%;
    border: 3px solid green;
    padding: 10px;
  }
</style>
