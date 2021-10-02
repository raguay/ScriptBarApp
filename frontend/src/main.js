import * as Wails from '@wailsapp/runtime';
import Main from './Main.svelte';

//
// Create the tray icon. The path is from the main directory for the ScriptPad program.
//

//
// Launch the application.
//
var app;
Wails.Init(() => {
  app = new Main({
    target: document.body,
    props: {
    }
  });
});
/*
const app = new Main({
  target: document.body,
  props: {
  }
});
*/
globalThis.appScriptBar = app;

export default app;

