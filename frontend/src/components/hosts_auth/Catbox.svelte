<script>
  import { LogInfo } from "wailsjs/runtime/runtime";

  let error = "";
  let userHash = "";
  function handleInput() {
    if (!userHash) {
      error = "Please enter a string";
      return;
    } else if (userHash.length !== 25) {
      error = "Not a valid Catbox User Hash";
      return;
    }

    LogInfo("valid catbox userhash");

    error = "";
  }
</script>

<div class="input-container">
  <label for="userHash">Catbox User Hash</label>
  <input
    id="userHash"
    type="text"
    bind:value={userHash}
    placeholder="Catbox User Hash"
    on:input={handleInput}
  />
  {#if error}
    <p class="error">{error}</p>
  {/if}
</div>

<style>
  .input-container {
    max-width: 22%;
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
    text-align: center;
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
