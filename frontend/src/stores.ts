import { writable } from "svelte/store";


export const currentHost = writable("")
export const currentUploadUrl = writable("")
export const currentTimeoutDuration = writable(0)
export const currentFiles = writable([])


/** 
 *
 * @param {Number} id 
 * @param {String} fileName 
 *
 * **/
export function updateFileName(id: Number, fileName: String): void {

  currentFiles.update(files => files.map(file => file.id === id ? { ...file, fileName } : file))

}


