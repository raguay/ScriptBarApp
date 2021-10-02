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
  import { socket } from '../stores/socket.js';
  
  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;
  
  let value = 'loading...';
  let ipaddress = '';
  
  $: NewSocket($socket);
  $: updateWidget(index);
   
  const dispatch = createEventDispatcher();
  
  //
  // I'm using a reactive function call due to the fact that 
  // on mounting it is null and then it get's updated.
  //
  function NewSocket(soc) {
    if(soc !== undefined) {
      soc.on(config.flowVar, (data) => {
        if(data !== undefined) {
          if(typeof data.ip !== 'undefined') {
            value = data.ip;
          }
          ipaddress = 'http://' + value;
          if(parseInt(config.port) !== 0) {
            ipaddress += ':' + config.port;
          }
        }
      });
    }
  }

  function ClickIP() {
    if(config.showLink) {
      backend.writeClipboard(ipaddress);
    }
  }

  function updateWidget(index) {
    getData();
  }

  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    fetch('http://localhost:9978/api/nodered/var/' + config.flowVar)
                .then((resp) => { 
                  return resp.json();
                }).then((data) => {
                  if((data !== null)&&(data.text !== null)) {
                    if(typeof data.text.ip !== 'undefined') {
                      value = data.text.ip;
                    }
                    ipaddress = 'http://' + value;
                    if(parseInt(config.port) !== 0) {
                      ipaddress += ':' + config.port;
                    }
                  } else {
                    value = 'loading...';
                    ipaddress = '';
                  }
                });
  }

  onMount(() => {
    getData();
  })
  
  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
