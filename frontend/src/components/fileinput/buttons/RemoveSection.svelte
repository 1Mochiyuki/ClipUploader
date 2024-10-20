<script>
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";
  import { currentFiles as storedFiles } from "../../../stores";
  import Icon from "@iconify/svelte";

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
    <Icon
      icon="mdi:remove"
      width="26"
      height="26"
      style="color: black; vertical-align: middle; display: inline-block;"
    ></Icon>
  </div>
</button>
