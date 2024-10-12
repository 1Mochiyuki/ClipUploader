<script>
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";
  import { currentFiles as storedFiles, updateFileName } from "../../../stores";
  import {
    ChooseFile,
    RemovePathFromFile,
  } from "../../../../wailsjs/go/main/App.js";

  export let id;
  export let width;
  let file = $storedFiles.find(
    (f) => f.id === id || { fileName: "Select a file" },
  );

  /**
   * @param {Object} file the currently selected file.
   **/
  const handleChooseFile = async (file) => {
    await ChooseFile().then((absPath) => {
      file.absPath = absPath;
    });
    await RemovePathFromFile(file.absPath).then((name) => {
      LogInfo(`id from SelectFile.svelte: ${file.id}`);
      updateFileName(file.id, name);
    });
    LogInfo(`name outside promise: ${file.fileName}`);
  };
</script>

<button
  class="button-4"
  style="width: {width}px; "
  on:click={() => handleChooseFile(file)}
>
  <div>
    <svg
      class="icon"
      style="  display: inline-block; vertical-align: middle;"
      xmlns="http://www.w3.org/2000/svg"
      height="26px"
      viewBox="0 -960 960 960"
      width="26px"
      fill="#000000"
      ><path
        d="M440-200h80v-167l64 64 56-57-160-160-160 160 57 56 63-63v167ZM240-80q-33 0-56.5-23.5T160-160v-640q0-33 23.5-56.5T240-880h320l240 240v480q0 33-23.5 56.5T720-80H240Zm280-520v-200H240v640h480v-440H520ZM240-800v200-200 640-640Z"
      /></svg
    >
  </div>
</button>
