<script>
  import { createEventDispatcher } from "svelte";
  import { ChooseFile } from "../../wailsjs/go/main/App.js";
  export let filePath;
  export function HandleChooseFile() {
    ChooseFile().then((x) => (filePath.fileName = x));
  }
  const dispatch = createEventDispatcher();
  function handleChooseFile() {
    console.log(`id: ${filePath.id}`);
    dispatch(filePath.id);
  }
</script>

<div class="btn-div">
  <button
    on:click={handleChooseFile}
    class="button-4"
    style="padding-bottom: 3px;"
    ><svg
      xmlns="http://www.w3.org/2000/svg"
      height="24px"
      viewBox="0 -960 960 960"
      width="24px"
      fill="#000000"
      ><path
        d="m336-280 144-144 144 144 56-56-144-144 144-144-56-56-144 144-144-144-56 56 144 144-144 144 56 56ZM480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"
      /></svg
    ></button
  >

  <button
    class="button-4"
    style="padding-bottom: 3px;"
    on:click={HandleChooseFile}
    ><svg
      xmlns="http://www.w3.org/2000/svg"
      height="24px"
      viewBox="0 -960 960 960"
      width="24px"
      fill="#000000"
      ><path
        d="M440-200h80v-167l64 64 56-57-160-160-160 160 57 56 63-63v167ZM240-80q-33 0-56.5-23.5T160-160v-640q0-33 23.5-56.5T240-880h320l240 240v480q0 33-23.5 56.5T720-80H240Zm280-520v-200H240v640h480v-440H520ZM240-800v200-200 640-640Z"
      /></svg
    ></button
  >
  <button class="button-4"> Im a FileUpload svelte component :3 </button>

  {#if filePath.fileName != undefined && filePath.fileName !== ""}
    <label>Selected File: {filePath.fileName}</label>
  {/if}
</div>
