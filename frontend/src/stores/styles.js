import { writable } from 'svelte/store';

let defaultStyles = {
  appBackground: '#2A212C',
  textcolor: '#80FFEA',
  fontSize: '14px',
  fontFamily: 'Arial',
  highlight: '#9580FF'
};

export const styles = writable( defaultStyles );

