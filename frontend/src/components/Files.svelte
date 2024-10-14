<script>
  import { LogInfo } from "../../wailsjs/runtime/runtime.js";
  import { currentHost, currentFiles as storedFiles } from "../stores";
  import RemoveSection from "./fileinput/buttons/RemoveSection.svelte";
  import SelectFile from "./fileinput/buttons/SelectFile.svelte";
  import UploadFile from "./fileinput/buttons/UploadFile.svelte";
  import Navbar from "./Navbar.svelte";

  $: newFileId = totalFiles
    ? Math.max(...currentFiles.map((f) => f.id)) + 1
    : 1;
  $: LogInfo(`[NEW FILE ID] = ${newFileId}`);
  $: totalFiles = currentFiles.length;
  $: LogInfo(`current host: ${$currentHost}`);
  let currentFiles = [];
  let width = 10;
  const addFileSection = () => {
    LogInfo(`length: ${totalFiles} id: ${newFileId}`);
    storedFiles.set([
      ...$storedFiles,
      { id: newFileId, fileName: "Select a file", filePath: "" },
    ]);
    storedFiles.update(() => (currentFiles = $storedFiles));
    LogInfo(`${$storedFiles}`);
  };
</script>

<Navbar {addFileSection} />
<div>
  <ul aria-labelledby="list-heading">
    {#each $storedFiles as file}
      <li id={file.id}>
        <div class="btn-div">
          <RemoveSection
            currentFiles={$storedFiles}
            {file}
            {totalFiles}
            {width}
          />
          <UploadFile {file} {width} />
          <SelectFile fileId={file.id} {width} />
          <span class="fileName-text">{file.fileName}</span>
        </div>
      </li>
    {/each}
  </ul>
</div>

<style>
  ul {
    display: inline-block;
    text-align: left;
    list-style-type: none;
  }
  .fileName-text {
    padding-left: 7.5px;
  }
</style>
