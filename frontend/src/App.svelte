<script>
  import Files from "./components/Files.svelte";

  export let files = [
    { id: 0, fileName: "balls.mp4" },
    { id: 1, fileName: "meow.mp4" },
    { id: 2, fileName: "lol.mp4" },
  ];
</script>

<Files {files}></Files>

<style>
  :global(.visually-hidden) {
    position: absolute !important;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px 1px 1px 1px);
    clip: rect(1px, 1px, 1px, 1px);
    white-space: nowrap;
  }
</style>
