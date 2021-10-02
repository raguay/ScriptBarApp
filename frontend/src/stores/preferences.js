import { writable } from 'svelte/store';

export const preferences = writable( {
  port: 9978
} );

