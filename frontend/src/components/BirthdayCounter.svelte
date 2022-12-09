<div class='main' 
     style="height: {Math.floor(height)}px;" 
     data-index="{index}"
     bind:this={dom}
     on:mouseenter={() => { dom.style.color = $styles.highlight; }}
     on:mouseleave={() => { dom.style.color = $styles.textcolor; }}
     on:mouseout={() => { dom.style.color = $styles.textcolor; }}
     on:blur={() => { dom.style.color = $styles.textcolor; }}
>
  <p on:dblclick={(e) => { dispatch('dblclick',{});}} 
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
     class='name' style="font-size: {$styles.fontSize};"
     data-index="{index}"
  >
    {disName}
  </p>
  <p class='value' 
     style="font-size: {$styles.fontSize};"
     data-index="{index}"
  >
    {value}
  </p>
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
  import { resizeWindow } from '../stores/resizeWindow.js';
  import { styles } from '../stores/styles.js';
  import { variables } from '../stores/variables.js';
  
  export let name;
  export let config;
  export let index;
  export let height;

  let dom;
  let value = 'loading...';
  let disName;

  const dispatch = createEventDispatcher();

  $: updateWidget(index);
  
  onMount(() => {
    disName = name;
    $variables.setUpVar(config.flowVar, null, () => { getData(); });
  });
  
  function updateWidget(index) {
    getData();
  }

  function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    const data = $variables.getVar(config.flowVar);
    if((data !== null) && (typeof data.text !== 'undefined')) {
      if(typeof data.days !== 'undefined') {
        value = data.days;
      }
      if(typeof data.age !== 'undefined') {
        disName = name + ': ' + data.age;
      } else {
        disName = name;
      }
      $resizeWindow();
    }
  }
  
  function middleButton() {
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
