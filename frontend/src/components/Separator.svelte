<div class='main'  
     style="height: {Math.floor(height)}px;"
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
     on:dblclick={(e) => { e.preventDefault(); dispatch('dblclick', {}); }}
     bind:this={dom}>
  <hr 
    id={name}
  />
</div>

<style>
  .main {
    width: 100%;
    height: 20px;
    display: flex;
    flex-direction: row;
    margin: auto 0px auto 0px;
    padding: 0px;
  }

  .main hr {
    width: 100%;
    margin: 0px;
    padding: 0px;
    align-self: center;
  }
</style>

<script>
  import { createEventDispatcher } from 'svelte';
  
  export let config;
  export let name;
  export let index;
  export let height;
  
  let dom;
  
  const dispatch = createEventDispatcher();
  
  function middleButton() {
    config = config;
    dispatch('middleButton', {
      index: index,
      dom: dom
    })
  }
</script>
