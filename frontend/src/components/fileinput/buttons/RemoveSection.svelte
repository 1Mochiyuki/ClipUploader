<script>
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";
  import { currentFiles as storedFiles } from "../../../stores";
  export let file;
  export let currentFiles;
  export let totalFiles;
  export let width;
  /**
   * @param {Object} file the currently selected file.
   **/

  const removeSection = (file) => {
    LogInfo(`length: ${totalFiles} id: ${file.id}`);

    storedFiles.set(currentFiles.filter((x) => x.id !== file.id));
    storedFiles.update(() => (currentFiles = $storedFiles));
    LogInfo(`${currentFiles} `);
  };
</script>

<button
  on:click={() => removeSection(file)}
  class="button-4"
  style="width: {width}px;"
>
  <div style="vertical-align: middle;">
    <svg
      class="icon"
      style="  display: inline-block; vertical-align: middle;"
      xmlns="http://www.w3.org/2000/svg"
      height="26px"
      viewBox="0 -960 960 960"
      width="26px"
      fill="#000000"
      ><path
        d="m336-280 144-144 144 144 56-56-144-144 144-144-56-56-144 144-144-144-56 56 144 144-144 144 56 56ZM480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"
      /></svg
    >
  </div>
</button>
