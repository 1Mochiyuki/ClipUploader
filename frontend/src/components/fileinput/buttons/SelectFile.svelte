<script>
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";
  import { currentFiles as storedFiles, updateFileName } from "../../../stores";
  import {
    ChooseFile,
    RemovePathFromFile,
  } from "../../../../wailsjs/go/main/App.js";
  import Icon from "@iconify/svelte";

  export let fileId;
  export let width;
  let file = $storedFiles.find((f) => f.id === fileId);

  /**
   * @param {Object} file the currently selected file.
   **/
  const handleChooseFile = async (file) => {
    LogInfo(`[CHOOSE FILE] ${file.id}`);
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
    <Icon
      icon="material-symbols:folder-open-outline"
      width="26"
      height="26"
      style="color: black; display: inline-block; vertical-align: middle;"
    ></Icon>
  </div>
</button>
