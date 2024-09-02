<script>
  import { each } from "svelte/internal";
  import { ChooseFile } from "../wailsjs/go/main/App.js";
  import FilePicker from "./components/FilePicker.svelte";
  import Files from "./components/Files.svelte";

  export let files = [
    { id: 0, fileName: "balls.mp4" },
    { id: 1, fileName: "meow.mp4" },
    { id: 2, fileName: "lol.mp4" },
  ];
</script>

<Files {files}></Files>
