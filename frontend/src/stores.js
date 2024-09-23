import { writable } from "svelte/store";

export const currentHost = writable("Select a host")
export const currentTimeoutDuration = writable("0")
export const currentFiles = writable([])

export const updateFileName = (id, fileName) => {
  currentFiles.update(files => files.map(file => file.id === id ? { ...file, fileName } : file))
}
