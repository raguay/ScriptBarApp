<div class='main' 
     style="height: {Math.floor(height)}px;" 
     data-index="{index}"
     on:mouseenter={() => { dom.style.color = $styles.highlight; }}
     on:mouseleave={() => { dom.style.color = $styles.textcolor; }}
     on:mouseout={() => { dom.style.color = $styles.textcolor; }}
     on:blur={() => { dom.style.color = $styles.textcolor; }}
     bind:this={dom}>
  <p on:dblclick={(e) => { dispatch('dblclick',{});}} 
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
     class='name' style='font-size: {$styles.fontSize};' 
     data-index="{index}"
  >
    {name}
  </p>
  {#if config.showLink}
    <p class='value' 
       style='font-size: {$styles.fontSize}; cursor: pointer;' 
       on:click="{(event) =>{ clickIP();}}" data-index="{index}"
    >
      {value}
    </p>
  {:else}
     <p class='value' 
       style='font-size: {$styles.fontSize};' 
    >
      {value}
    </p>
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
  
  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;

  let value = 'loading...';
  let ipaddress = '';
  
  const dispatch = createEventDispatcher();
  
  $: updateWidget(index);
  
  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    fetch('http://localhost:9978/api/getip')
                .then((resp) => { 
                  return resp.json();
                }).then((data) => {
                  if(data !== null) {
                    if(typeof data.ip !== 'undefined') {
                      value = data.ip;
                    }
                    ipaddress = 'http://' + value
                    if(parseInt(config.port) !== 0) {
                      ipaddress += ':' + config.port;
                    }
                  }
                });
  }

  onMount(() => {
    getData();
  })

  function updateWidget(index) {
    getData();
  }

  function clickIP() {
    if(config.showLink) {
      backend.writeClipboard(ipaddress);
    }
  }
  
  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
