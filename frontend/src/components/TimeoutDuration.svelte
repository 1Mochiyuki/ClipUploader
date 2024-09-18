<script>
  import { LogInfo } from "../../wailsjs/runtime/runtime.js";
  export let timeoutDuration;
  $: timeoutDuration = `${timeoutDuration}`.toString();
  let labelText = "Timeout Duration (Seconds)";
  let error = "";

  function handleInput() {
    if (!timeoutDuration) {
      error = "Please enter a number";
      return;
    } else if (isNaN(Number(timeoutDuration))) {
      error = "Please enter a valid number";
      return;
    }
    error = "";
    LogInfo(`Valid number entered: ${timeoutDuration}`);

    // Here you can perform any action with the valid number
  }
</script>

<div class="input-container">
  <label for="numberInput">{labelText}</label>
  <input
    type="number"
    id="numberInput"
    bind:value={timeoutDuration}
    on:click={() => LogInfo(`duration: ${timeoutDuration}`)}
    on:input={handleInput}
    placeholder="Timeout Duration"
  />
  {#if error}
    <p class="error">{error}</p>
  {/if}
</div>

<style>
  .input-container {
    max-width: 300px;
    margin: 0 auto;
    padding: 20px;
    border-radius: 5px;
  }

  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }

  input {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 16px;
    transition: border-color 0.3s ease;
  }

  input:focus {
    outline: none;
    border-color: #4caf50;
  }

  .error {
    color: red;
    margin-top: 5px;
    font-size: 14px;
  }
</style>
