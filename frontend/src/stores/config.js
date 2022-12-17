import { writable } from "svelte/store";

export const config = writable({
  containers: [],
  mheight: 400,
  mwidth: 700,
});
