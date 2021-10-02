<div class='main' 
     style="height: {Math.floor(height)}px;" 
     data-index="{index}"
     bind:this={dom}>
  <p class="name" 
     style="font-size: {$styles.fontSize};"
     on:dblclick={(e) => { dispatch('dblclick',{});}}
     data-index="{index}"
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
  >{name}</p>
  <p class="value" 
     style="font-size: {$styles.fontSize};"
     data-index="{index}">{value}</p>
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
  import { createEventDispatcher, onMount, tick } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { socket } from '../stores/socket.js';
  
  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;

  const dispatch = createEventDispatcher();
  
  let value = 'loading...';

  $: NewSocket($socket);
  $: updateWidget(index);
  
  //
  // I'm using a reactive function call due to the fact that 
  // on mounting it is null and then it get's updated.
  //
  function NewSocket(soc) {
    if(soc !== null) {
      soc.on(config.flowVar, (data) => {
        if(data !== null) {
          value = data;
        }
      });
    }
  }

  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    fetch('http://localhost:9978/api/nodered/var/' + config.flowVar)
                .then((resp) => { 
                  return resp.json();
                }).then((data) => {
                  if(data !== null) {
                    if((typeof data.text !== 'undefined')&&(data.text !== null)) {
                      value = data.text;
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

  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
