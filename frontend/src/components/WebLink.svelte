<div class='main'  
     style="height: {Math.floor(height)}px;"
     data-index="{index}"
     on:contextmenu={(e) => { e.preventDefault(); middleButton();}}
     on:mouseenter={() => { dom.style.color = $styles.highlight; }}
     on:mouseleave={() => { dom.style.color = $styles.textcolor; }}
     on:mouseout={() => { dom.style.color = $styles.textcolor; }}
     on:blur={() => { dom.style.color = $styles.textcolor; }}
     bind:this={dom}>
  <span class='name'
    on:dblclick={(e) => { e.preventDefault(); dispatch('dblclick', {}); }}
  >
    {name}
  </span>
  <span class='weblinklink'
     on:click={(e) => { openLink(); }}
   >
     🌎
  </span>
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

  .weblinklink {
    margin: 0px 0px 0px auto;
    cursor: pointer;
  }
</style>

<script>
  import { createEventDispatcher, onMount, tick } from 'svelte';
  import { styles } from '../stores/styles.js';
  
  export let name;
  export let config;
  export let index;
  export let height;

  let dom;

  const dispatch = createEventDispatcher();
  
  async function middleButton() {
    await tick();
    if((typeof dom === undefined) || (dom === {})){
      console.log("No dom");
    } else {
      dispatch('middleButton', {
        index: index,
        dom: dom
      });
    }
  }

  //
  // Function:     openLink
  //
  // Description:  This function opens the given link in the browser specified.
  //
  function openLink() {
    window.wails.Browser.OpenURL(config.link);
  }

  function basename(str) {
    var parts = str.split('/');
    return parts[parts.length - 1];
  }
</script>
