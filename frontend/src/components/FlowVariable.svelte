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
  import { variables } from '../stores/variables.js';
  
  export let name;
  export let config;
  export let index;
  export let height;
  
  let dom;

  const dispatch = createEventDispatcher();
  
  let value = 'loading...';

  $: updateWidget(index);
  
  onMount(() => {
    getData();
    $variables.setUpVar(config.flowVar, null, () => { getData(); });
    return(()=>{
    });
  })
 
  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    var data = $variables.getVar(config.flowVar);
    if(data !== null) {
      if((typeof data.text !== 'undefined')&&(data.text !== null)) {
        value = data.text;
        $resizeWindow();
      }
    }
  }

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
