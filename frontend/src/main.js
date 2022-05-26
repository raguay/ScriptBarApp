import Main from './Main.svelte';

//
// Create the tray icon. The path is from the main directory for the ScriptPad program.
//

//
// Launch the application.
//
var app;
app = new Main({
  target: document.body,
  props: {
  }
});

export default app;

