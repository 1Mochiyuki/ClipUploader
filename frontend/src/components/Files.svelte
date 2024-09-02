<script>
  import FilePicker from "./FilePicker.svelte";
  export let files = [];

  $: newFileId = totalFiles ? Math.max(...files.map((f) => f.id)) + 1 : 1;
  $: totalFiles = files.length;

  function addFileSection() {
    console.log(`length: ${totalFiles} id: ${newFileId}`);
    files = [
      ...files,
      { id: newFileId, fileName: `Choose a file ${newFileId}` },
    ];
    console.log(files);
  }
  function closeFileSection(file) {
    console.log(`length: ${totalFiles} id: ${file.id}`);

    files = files.filter((x) => x.id !== file.id);
    console.log(files);
  }
</script>

<div style="text-align: center;">
  <button on:click={addFileSection}>Add File Section Meow</button>
</div>

<ul aria-labelledby="list-heading">
  {#each files as file}
    <li>
      <FilePicker on:toggle={() => closeFileSection(file)} filePath={file} />
    </li>
  {/each}
</ul>

<style>
  ul {
    display: inline-block;
    text-align: left;
    list-style-type: none;
  }
</style>
