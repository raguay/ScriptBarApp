<div class='main' 
     style="height: {Math.floor(height)}px;" 
     data-index="{index}"
     on:mouseenter={() => { dom.style.color = $styles.highlight; }}
     on:mouseleave={() => { dom.style.color = $styles.textcolor; }}
     on:mouseout={() => { dom.style.color = $styles.textcolor; }}
     on:blur={() => { dom.style.color = $styles.textcolor; }}
     bind:this={dom}>
  <p class='name' 
     on:dblclick={(e) => { dispatch('dblclick',{});}}
     style='font-size: {$styles.fontSize};'
     data-index="{index}"
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
   >{name}</p>
    {#if config.showLink}
      <p class='value' 
        style='font-size: {$styles.fontSize}; cursor: pointer;'
        on:click="{(event) => { ClickIP(); }}"
        data-index="{index}">{value}</p>
    {:else}
       <p class='value' 
        style='font-size: {$styles.fontSize};'
        data-index="{index}">{value}</p>
   {/if}
  </div>

<style>
  .main {
    width: 100%;
    display: flex;
    flex-direction: row;
    margin: auto 0px auto 0px;
    padding: 0px;
  }

  .name {
    margin: 0px;
    padding: 0px 10px 0px 0px;
  }

  .value {
    margin: 0px 0px 0px auto;
    padding: 0px;
  }
</style>

<script>
  import { createEventDispatcher, onMount } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { variables } from '../stores/variables.js';
  import * as rt from "../../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;
  
  let value = 'loading...';
  let ipaddress = '';
  
  $: updateWidget(index);
   
  const dispatch = createEventDispatcher();
  
  async function ClickIP() {
    if(config.showLink) {
      await rt.SetClip(ipaddress);
    }
  }

  function updateWidget(index) {
    getData();
  }

  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    var data = $variables.getVar(config.flowVar);
    if((data !== null)&&(data.ip !== "")) {
      value = data.ip;
      ipaddress = 'http://' + value;
      if(parseInt(config.port) !== 0) {
        ipaddress += ':' + config.port;
      }
    } else {
      value = 'loading...';
      ipaddress = '';
    }
  }

  onMount(() => { 
    $variables.setUpVar(config.flowVar, null, () => { getData(); });
    getData();
  })
  
  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
