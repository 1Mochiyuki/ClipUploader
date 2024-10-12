<script>
  import { push } from "svelte-spa-router";
  import {
    GetHost,
    GetTimeout,
    HostList,
    SaveHost,
    SaveTimeoutDuration,
  } from "../../wailsjs/go/main/App.js";
  import HostInput from "./HostInput.svelte";
  import TimeoutDuration from "./TimeoutDuration.svelte";
  import { currentHost, currentTimeoutDuration } from "../stores";
  import { LogInfo } from "../../wailsjs/runtime/runtime.js";

  const getPlaceholders = async () => {
    const host = await GetHost();
    LogInfo(`got value: ${host}`);
    const duration = await GetTimeout();
    LogInfo(`got value: ${duration}`);
    currentHost.set(host);
    currentTimeoutDuration.set(duration);
  };

  getPlaceholders();

  const handleSaveSettings = async () => {
    let timeout = Number($currentTimeoutDuration);
    let host = $currentHost;
    let savedHost = await GetHost();
    let savedTimeout = await GetTimeout();

    if (savedTimeout === timeout && savedHost === host) {
      LogInfo("values are the same !!");
      return;
    }
    if (savedHost !== host) {
      LogInfo(`saving host: ${$currentHost} type: ${typeof $currentHost}`);
      await SaveHost(host);
    }

    if (savedTimeout !== timeout) {
      LogInfo(`saving timeout: ${timeout} type: ${typeof timeout}`);
      await SaveTimeoutDuration(timeout);
    }
  };

  async function testThing() {
    const hosts = await HostList();
    let currHost = $currentHost;
    LogInfo(`curr host: ${currHost}`);
    const catbox = Object.values(hosts[currHost])[1];

    LogInfo(`keys: ${Object.keys(hosts)} | values: ${catbox}`);
  }
</script>

<div>
  <button
    on:click={() => push("/")}
    class="back"
    style="background-color: rgba(0,0,0,0); border: 0;  box-shadow: none; float: left; "
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="30px"
      viewBox="0 -960 960 960"
      width="30px"
      fill="#FFFFFF"
      ><path
        d="M400-240 160-480l240-240 56 58-142 142h486v80H314l142 142-56 58Z"
      /></svg
    >
  </button>
  <div style="justify-content: space-between; padding-top: 20px;">
    <div>
      <HostInput></HostInput>

      <TimeoutDuration></TimeoutDuration>
    </div>
    <button id="save-btn" class="button-4" on:click={handleSaveSettings}
      >Save</button
    >
    <button class="button-4" on:click={testThing}>Test</button>
  </div>
</div>

<style>
  .back:hover {
    cursor: pointer;
  }
</style>
