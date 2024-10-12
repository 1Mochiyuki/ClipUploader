<script>
  import { onMount } from "svelte";
  import { LogInfo } from "../../wailsjs/runtime/runtime.js";
  import { currentHost } from "../stores";
  import { GetHost, HostList } from "../../wailsjs/go/main/App.js";

  export let options = [
    { value: "Catbox" },
    { value: "Pomf" },
    { value: "other" },
  ];

  export let value = "";
  let label = "Select a host";

  let isOpen = false;
  let selectContainer;
  let containerWidth = "auto";
  let id = `select-${Math.random().toString(36).slice(2, 9)}`;

  const toggleDropdown = () => {
    isOpen = !isOpen;
  };
  /**
   * @param {Object} option - Currently selected option
   * */
  const selectOption = (option) => {
    value = option.value;
    isOpen = false;
    currentHost.set(value);
    LogInfo(`selected host: ${value}`);
  };

  onMount(() => {
    const tempElement = document.createElement("span");
    tempElement.style.visibility = "hidden";
    tempElement.style.position = "absolute";
    tempElement.style.whiteSpace = "nowrap";
    document.body.appendChild(tempElement);

    let maxWidth = 0;
    options.forEach((option) => {
      tempElement.textContent = option.value;
      const width = tempElement.offsetWidth;
      if (width > maxWidth) {
        maxWidth = width;
      }
    });

    tempElement.textContent = $currentHost;
    const placeholderWidth = tempElement.offsetWidth;
    if (placeholderWidth > maxWidth) {
      maxWidth = placeholderWidth;
    }

    document.body.removeChild(tempElement);

    // Add some padding to the width
    containerWidth = `${maxWidth + 60}px`;
  });
</script>

<div class="select-wrapper" style="width: {containerWidth};">
  <label for={id} class="select-label">{label}</label>
  <div class="select-container" bind:this={selectContainer}>
    <button
      type="button"
      {id}
      on:click={toggleDropdown}
      class="select-button"
      aria-haspopup="listbox"
      aria-expanded={isOpen}
    >
      <span class="select-value">
        {value ? value : $currentHost}
      </span>
      <span class="select-arrow">
        <svg viewBox="0 0 20 20" fill="none" stroke="currentColor">
          <path
            d="M7 7l3-3 3 3m0 6l-3 3-3-3"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </span>
    </button>

    {#if isOpen}
      <div class="select-dropdown" role="listbox">
        <ul class="select-options">
          {#each options as option (option.value)}
            <!-- svelte-ignore a11y-click-events-have-key-events </-->
            <li
              on:click={() => selectOption(option)}
              class="select-option"
              role="option"
              aria-selected={option.value === value}
            >
              {option.value}
              {#if option.value === value}
                <span class="select-check">
                  <svg viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </span>
              {/if}
            </li>
          {/each}
        </ul>
      </div>
    {/if}
  </div>
</div>

<style>
  .select-wrapper {
    font-family: Arial, sans-serif;
    margin: 0 auto;
  }

  .select-label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: bold;
    color: #ffffff;
  }

  .select-container {
    position: relative;
  }

  .select-button {
    width: 100%;
    padding: 10px 15px;
    background-color: #ffffff;
    border: 1px solid #d1d5db;
    border-radius: 4px;
    cursor: pointer;
    text-align: left;
    font-size: 16px;
    line-height: 1.5;
    color: #111827;
    transition:
      border-color 0.15s ease-in-out,
      box-shadow 0.15s ease-in-out;
  }

  .select-button:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.5);
  }

  .select-value {
    display: block;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }

  .select-arrow {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
    width: 20px;
    height: 20px;
    color: #6b7280;
    pointer-events: none;
  }

  .select-dropdown {
    position: absolute;
    top: 100%;
    left: 0;
    z-index: 10;
    width: 100%;
    margin-top: 4px;
    background-color: #ffffff;
    border: 1px solid #d1d5db;
    border-radius: 4px;
    box-shadow:
      0 4px 6px -1px rgba(0, 0, 0, 0.1),
      0 2px 4px -1px rgba(0, 0, 0, 0.06);
  }

  .select-options {
    max-height: 240px;
    overflow-y: auto;
    padding: 4px 0;
    margin: 0;
    list-style-type: none;
  }

  .select-option {
    padding: 8px 12px;
    cursor: pointer;
    user-select: none;
    color: black;
  }

  .select-option:hover {
    background-color: #e3e4e4;
  }

  .select-option[aria-selected="true"] {
    background-color: #e5e7eb;
  }

  .select-check {
    float: right;
    width: 20px;
    height: 20px;
    color: #3b82f6;
  }
</style>
