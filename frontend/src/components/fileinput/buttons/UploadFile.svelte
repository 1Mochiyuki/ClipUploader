<script>
  import { Upload as UploadPomf } from "../../../../wailsjs/go/types/Pomf.js";
  import { Upload as UploadLobfile } from "../../../../wailsjs/go/types/Lobfile.js";
  import { Upload as UploadCatbox } from "../../../../wailsjs/go/types/Catbox.js";
  import { currentHost, currentFiles as storedFiles } from "../../../stores";
  import { LogInfo } from "../../../../wailsjs/runtime/runtime.js";

  export let width;
  export let file;
  const upload = (file) => {
    if (!file.fileName && $currentHost === "Select a host") {
      LogInfo(`file name or current host not set`);
      return;
    }

    LogInfo(`file name: ${file.fileName}`);
    switch ($currentHost) {
      case "Catbox":
        LogInfo("catbox");
        UploadCatbox(file.fileName);
        break;
      case "Lobfile":
        LogInfo("lobfile");
        UploadLobfile(file.fileName);
        break;
      case "Pomf":
        LogInfo("pomf");
        UploadPomf(file.fileName);
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
    <svg
      class="icon"
      style="display: inline-block; vertical-align: middle;"
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
