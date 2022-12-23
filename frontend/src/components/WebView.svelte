<div id='webview' 
     style='background-color: {body.config.backgroundColor};
            color: {$styles.textcolor};
            font-family: {$styles.fontFamily};
            font-size: {$styles.fontSize};
            max-width: {$config.mwidth}px;
            max-height: {$config.mheight}px;'
    bind:this={mainDOM}
  >
  <div
    id='htmlbody'
    bind:this={htmlDOM}
  >
    {@html body.html}
  </div>
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
    position: absolute;
    top: 0px;
    left: 0px;
    display: flex;
    flex-direction: column;
    margin: 0px;
    padding: 10px;
    border: 1px solid transparent;
    border-radius: 10px;
    overflow: hidden;
    --wails-draggable: drag;
  }

  #htmlbody {
    padding: 5px;
    margin: 0px;
    overflow: auto;
  }

  #buttonRow {
    display: flex;
    flex-direction: row;
    margin: 0px 0px 5px 0px;
    padding: 0px;
  }

  #closeButton {
    border-radius: 5px;
    height: 30px;
    outline: none;
    margin: 5px auto;
    user-select: none;
    outline-style: none;
  }
</style>

<script>
  import { createEventDispatcher, onMount, tick, afterUpdate } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { headerPosition } from '../stores/headerPosition.js';
  import { width } from '../stores/width.js';
  import { config } from '../stores/config.js';
  import * as rt from "../../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  export let body;

  let mainDOM;
  let htmlDOM;
  let nwidth;
  let nheight;
  let oldheight = 0;
  let oldwidth = 0;

  const dispatch = createEventDispatcher();

  onMount(() => {
    globalThis.closeWebView = close;
    window.setLineColor = setLineColor;
    window.setLineColorNew = setLineColorNew;
  });

  afterUpdate(async () => {
    await resizeWindow();
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

  function close() {
    dispatch('changeView',{
      name: 'scriptbar',
      body: {
      }
    });
  }

  async function resizeWindow() {
    await rt.WindowSetSize($config.mwidth, $config.mheight);
    await tick();
    if(typeof body.config.width !== 'undefined') {
      nwidth = body.config.width;
      nheight = body.config.height;
      $headerPosition = Math.floor(nwidth/2);
      $width = nwidth;
    } else {
      if(mainDOM !== undefined) {
        nwidth = mainDOM.offsetWidth;
        nheight = mainDOM.offsetHeight;
        if((oldheight !== nheight) || (oldwidth !== nwidth)) {
          oldheight = nheight;
          oldwidth = nwidth;
          if(nwidth < 100) {
            nwidth = 400;
          }
          if(nheight < 40) {
            nheight = 40;
          }
          if(nheight > $config.mheight) {
            nheight = $config.mheight;
            htmlDOM.style.height = `${nheight-80}px`;
          }
          if(nwidth > $config.mwidth) {
            nwidth = $config.mwidth;
            htmlDOM.style.width = `${nwidth-41}px`;
          }
          $width = nwidth + 6;
          rt.WindowSetSize(nwidth+6, nheight+6);
          $headerPosition = Math.floor($width/2);
        }
      }
    }
  }
</script>

