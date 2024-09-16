<script>
  import { ChooseFile, RemovePathFromFile } from "../../wailsjs/go/main/App.js";
  import Navbar from "./Navbar.svelte";
  import { LogPrint } from "../../wailsjs/runtime/runtime.js";
  let files = [];
  $: newFileId = totalFiles ? Math.max(...files.map((f) => f.id)) + 1 : 1;
  $: totalFiles = files.length;
  $: placeholderText = "";
  $: placeHolderEmpty = !placeholderText;
  const SELECT_A_FILE_TEXT = "Select a file";
  /**
   * @param {Object} file the currently selected file.
   **/

  const removeSection = (file) => {
    LogPrint(`length: ${totalFiles} id: ${file.id}`);

    files = files.filter((x) => x.id !== file.id);
    LogPrint(`${files}`);
  };
  /**
   *
   * @param {Object} file the currently selected file.
   *
   * */
  const getFileName = (file) => {
    if (!file.fileName) return SELECT_A_FILE_TEXT;
    return file.fileName;
  };

  /**
   * @param {Object} file the currently selected file.
   **/

  const handleChooseFile = async (file) => {
    placeholderText = undefined;

    await ChooseFile().then((absPath) => {
      file.absPath = absPath;
    });
    await RemovePathFromFile(file.absPath).then((name) => {
      file.fileName = name;
      placeholderText = name;
    });
    LogPrint(`name outside promise: ${file.fileName}`);
  };

  const addFileSection = () => {
    LogPrint(`length: ${totalFiles} id: ${newFileId}`);
    files = [...files, { id: newFileId, fileName: "", filePath: "" }];
    LogPrint(`${files}`);
  };
</script>

<Navbar {addFileSection} />

<ul aria-labelledby="list-heading">
  {#each files as file}
    <li>
      <div class="btn-div">
        <button
          on:click={() => removeSection(file)}
          class="button-4"
          style="padding-bottom: 3px;"
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            height="24px"
            viewBox="0 -960 960 960"
            width="30px"
            fill="#000000"
            ><path
              d="m336-280 144-144 144 144 56-56-144-144 144-144-56-56-144 144-144-144-56 56 144 144-144 144 56 56ZM480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"
            /></svg
          ></button
        >
        <button
          class="button-4"
          style="padding-bottom: 3px;"
          on:click={() => handleChooseFile(file)}
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            height="24px"
            viewBox="0 -960 960 960"
            width="34px"
            fill="#000000"
            ><path
              d="M440-200h80v-167l64 64 56-57-160-160-160 160 57 56 63-63v167ZM240-80q-33 0-56.5-23.5T160-160v-640q0-33 23.5-56.5T240-880h320l240 240v480q0 33-23.5 56.5T720-80H240Zm280-520v-200H240v640h480v-440H520ZM240-800v200-200 640-640Z"
            /></svg
          ></button
        >
        <span class="fileName-text"
          >{placeHolderEmpty ? SELECT_A_FILE_TEXT : getFileName(file)}</span
        >
      </div>
    </li>
  {/each}
</ul>

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
