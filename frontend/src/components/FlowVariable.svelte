<div class='main' 
     style="height: {Math.floor(height)}px;" 
     data-index="{index}"
     on:mouseenter={() => { dom.style.color = $styles.highlight; }}
     on:mouseleave={() => { dom.style.color = $styles.textcolor; }}
     on:mouseout={() => { dom.style.color = $styles.textcolor; }}
     on:blur={() => { dom.style.color = $styles.textcolor; }}
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
  import { createEventDispatcher, onMount } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { resizeWindow } from '../stores/resizeWindow.js';
  import { timer } from '../stores/timer.js';
  
  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;
  let count = 0;
  let tm;

  const dispatch = createEventDispatcher();
  
  let value = 'loading...';

  $: updateWidget(index);
  
  onMount(() => {
    getData();
    var unsubtimer = timer.subscribe((val) => {
      if((typeof config.period !== undefined)&&(config.period !== 0)) {
        if((val % config.period) === 0) {
          getData();
        }
      }
    });
    return(()=>{
      unsubtimer();
    });
  })
 
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
          count = 0;
          $resizeWindow();
        }
      }
    }).catch((err) => {
      value = "not reachable: " + count;
      setTimeout(()=>{ count++; getData(); }, 60000); 
    });
  }

  function updateWidget(index) {
    getData(config, value);
  }

  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
