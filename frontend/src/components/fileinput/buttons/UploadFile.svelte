<script>
  import { Upload as UploadCatbox } from "../../../../wailsjs/go/types/Catbox.js";
  import { Upload as UploadLobfile } from "../../../../wailsjs/go/types/Lobfile.js";
  import { Upload as UploadPomf } from "../../../../wailsjs/go/types/Pomf.js";
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";
  import { currentHost } from "../../../stores";
  import Icon from "@iconify/svelte";

  export let width;
  export let file;
  /**
   *
   * @param {Object} fileToUpload the file that will be uploaded
   * */
  const upload = (fileToUpload) => {
    if (!fileToUpload.fileName && $currentHost === "Select a host") {
      LogInfo(`file name or current host not set`);
      return;
    }

    LogInfo(`file name: ${fileToUpload.fileName}`);
    switch ($currentHost) {
      case "Catbox":
        UploadCatbox(fileToUpload.fileName);
        break;
      case "Lobfile":
        UploadLobfile(fileToUpload.fileName);
        break;
      case "Pomf":
        UploadPomf(fileToUpload.fileName);
        break;
      default:
        LogInfo("????");
        break;
    }
  };
</script>

<button
  id="upload"
  class="button-4"
  style=" width: {width}px;"
  on:click={() => upload(file)}
>
  <div>
    <Icon
      icon="ci:file-upload"
      width="26"
      height="26"
      style="color:black; display: inline-block; vertical-align: middle;"
    ></Icon>
    <!-- <svg -->
    <!--   class="icon" -->
    <!--   style="display: inline-block; vertical-align: middle;" -->
    <!--   xmlns="http://www.w3.org/2000/svg" -->
    <!--   height="26px" -->
    <!--   viewBox="0 -960 960 960" -->
    <!--   width="26px" -->
    <!--   fill="#000000" -->
    <!--   ><path -->
    <!--     d="M320-200v-560l440 280-440 280Zm80-280Zm0 134 210-134-210-134v268Z" -->
    <!--   /></svg -->
  </div>
</button>
