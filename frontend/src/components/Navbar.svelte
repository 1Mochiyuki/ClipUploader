<script>
  import { push } from "svelte-spa-router";

  export let addFileSection = () => {};
</script>

<div
  class="button-div"
  style="text-align: center;  justify-content: space-between;"
>
  <button
    class="navbar-button"
    on:click={addFileSection}
    style="background-color: rgba(0,0,0,0); border: 0px solid;  box-shadow: none;"
    ><svg
      xmlns="http://www.w3.org/2000/svg"
      height="30px"
      viewBox="0 -960 960 960"
      width="30px"
      fill="#FFFFFF"
      ><path
        d="M360-240q-33 0-56.5-23.5T280-320v-480q0-33 23.5-56.5T360-880h360q33 0 56.5 23.5T800-800v480q0 33-23.5 56.5T720-240H360Zm0-80h360v-480H360v480ZM200-80q-33 0-56.5-23.5T120-160v-560h80v560h440v80H200Zm160-240v-480 480Z"
      /></svg
    ></button
  >
  <div>
    <button class="navbar-button" on:click={addFileSection}
      >Add File Section
    </button>
    <button class="navbar-button"> Upload All </button>
  </div>

  <button
    class="navbar-button"
    on:click={() => push("/settings")}
    style="background-color: rgba(0,0,0,0); border: 0px solid;  box-shadow: none;"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="30px"
      viewBox="0 -960 960 960"
      width="30px"
      fill="#FFFFFF"
      ><path
        d="m370-80-16-128q-13-5-24.5-12T307-235l-119 50L78-375l103-78q-1-7-1-13.5v-27q0-6.5 1-13.5L78-585l110-190 119 50q11-8 23-15t24-12l16-128h220l16 128q13 5 24.5 12t22.5 15l119-50 110 190-103 78q1 7 1 13.5v27q0 6.5-2 13.5l103 78-110 190-118-50q-11 8-23 15t-24 12L590-80H370Zm70-80h79l14-106q31-8 57.5-23.5T639-327l99 41 39-68-86-65q5-14 7-29.5t2-31.5q0-16-2-31.5t-7-29.5l86-65-39-68-99 42q-22-23-48.5-38.5T533-694l-13-106h-79l-14 106q-31 8-57.5 23.5T321-633l-99-41-39 68 86 64q-5 15-7 30t-2 32q0 16 2 31t7 30l-86 65 39 68 99-42q22 23 48.5 38.5T427-266l13 106Zm42-180q58 0 99-41t41-99q0-58-41-99t-99-41q-59 0-99.5 41T342-480q0 58 40.5 99t99.5 41Zm-2-140Z"
      /></svg
    >
  </button>
</div>

<style>
  .navbar-button {
    appearance: none;
    background-color: #fafbfc;
    border: 1px solid rgba(27, 31, 35, 0.15);
    border-radius: 6px;
    box-shadow:
      rgba(27, 31, 35, 0.04) 0 1px 0,
      rgba(255, 255, 255, 0.25) 0 1px 0 inset;
    box-sizing: border-box;
    color: #24292e;
    cursor: pointer;
    display: inline-block;
    font-family: -apple-system, system-ui, "navbar-button", Helvetica, Arial,
      sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
    font-size: 16px;
    font-weight: 500;
    line-height: 20px;
    list-style: none;
    padding: 6px 16px;
    position: relative;
    transition: background-color 0.2s cubic-bezier(0.3, 0, 0.5, 1);
    user-select: none;
    -webkit-user-select: none;
    touch-action: manipulation;
    vertical-align: middle;
    white-space: nowrap;
    word-wrap: break-word;
  }

  .navbar-button:hover {
    background-color: #f3f4f6;
    text-decoration: none;
    transition-duration: 0.1s;
  }

  .navbar-button:disabled {
    background-color: #fafbfc;
    border-color: rgba(27, 31, 35, 0.15);
    color: #959da5;
    cursor: default;
  }

  .navbar-button:active {
    background-color: #dee2e7;
    box-shadow: rgba(225, 228, 232, 0.2) 0 0.0625em 0 inset;
    transition: none 0s;
  }

  .navbar-button:focus {
    outline: 0.0625em transparent;
  }

  .navbar-button:before {
    display: none;
  }

  .navbar-button:-webkit-details-marker {
    display: none;
  }

  .button-div {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 3.125em;
  }
</style>
