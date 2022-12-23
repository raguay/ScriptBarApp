<div id='webview' 
     style='background-color: {body.config.backgroundColor};
            color: {$styles.textcolor};
            font-family: {$styles.fontFamily};
            font-size: {$styles.fontSize};
            width: {nwidth}px;
            height: {nheight}px;
            max-width: {$config.mwidth}px;
            max-height: {$config.mheight}px;'
    bind:this={mainDOM}
>
  {@html body.html}
  {#if body.config.showButton}
    <div id='buttonRow'>
      <button id="closeButton"
              type="button"
              style="background-color: {$styles.appBackground}; color: {$styles.textcolor}; font-size: {$styles.fontSize};"
              on:click="{() => { close(); }}" >
        Close
      </button>
    </div>
  {/if}
</div>

<style>
  #webview {
    display: flex;
    flex-direction: column;
    margin: 0px;
    padding: 5px;
    border: 10px solid transparent;
    border-radius: 10px;
    overflow: scroll;
    overflow-wrap: auto;
    --wails-draggable: drag;
  }

  #buttonRow {
    display: flex;
    flex-direction: row;
    margin: 0px;
    padding: 0px;
  }

  #closeButton {
    border-radius: 5px;
    height: 30px;
    outline: none;
    margin: 10px auto 10px auto;
    user-select: none;
    outline-style: none;
  }
</style>

<script>
  import { createEventDispatcher, onMount, tick } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { headerPosition } from '../stores/headerPosition.js';
  import { width } from '../stores/width.js';
  import { config } from '../stores/config.js';
  import * as rt from "../../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  export let body;

  let mainDOM;
  let nwidth;
  let nheight;

  const dispatch = createEventDispatcher();

  $: $width = bodyChange(body);

  onMount(() => {
    globalThis.closeWebView = close;
    $width = bodyChange(body);
    resizeWindow();
    window.setLineColor = setLineColor;
    window.setLineColorNew = setLineColorNew;
  });

  function setLineColorNew(e) {
    if(body.config.changeColor) {
      e.style.backgroundColor = body.config.highlightColor;
    }
  }

  function setLineColor(e) {
    if(body.config.changeColor) {
      e.style.backgroundColor = body.config.backgroundColor;
    }
  }

  function bodyChange(bod) {
    resizeWindow();
    return($width);
  }

  function close() {
    dispatch('changeView',{
      name: 'scriptbar',
      body: {
      }
    });
  }

  async function resizeWindow() {
    console.log("resizeWindow: ",$config);
    await tick();
    if(typeof body.config.width !== 'undefined') {
      nwidth = body.config.width;
      nheight = body.config.height;
      $headerPosition = Math.floor(nwidth/2);
      $width = nwidth;
    } else {
      if(mainDOM !== undefined) {
        nwidth = mainDOM.clientWidth;
        nheight = mainDOM.clientHeight;
        console.log("resizeWindow: ", mainDOM.clientWidth, mainDOM.clientHeight);
        if(nwidth < 100) {
          nwidth = 400;
        }
        if(nheight < 40) {
          nheight = 40;
        }
        if(nheight > $config.mheight) {
          nheight = $config.mheight;
        }
        if(nwidth > $config.mwidth) {
          nwidth = $config.mwidth;
        }
        $width = nwidth + 6;
        console.log("resizeWindow: ", nwidth, nheight, $width);
        rt.WindowSetSize(nwidth+6, nheight+6);
        $headerPosition = Math.floor($width/2);
      }
    }
  }
</script>

