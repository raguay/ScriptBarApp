<div id='addComponentDialog'
>
  <label class="componentLabel" 
         for="componentName">
    What's the name?
  </label>
  <input type="text"
         class="componentInput"
         id="componentName"
         color="{$styles.appBackground}"
         bind:value={componentName} />
  <label for="componentTypeInput"
         class="componentLabel" >
    What Type of Component?
  </label>
  <select name="componentTypeInput" 
          class="componentInput"
          color="{$styles.appBackground}"
          id="componentTypeInput"
          bind:value={componentType}
          on:change={typeChanged}
  >
    {#each types as type}
      <option>{type.moduleName}</option>
    {/each}
  </select>
  <label for="componentHeight"
         class="componentLabel" >
    Height: 
  </label>
  <input type="text"
         class="componentInput"
         color="{$styles.appBackground}"
         id="componentHeight"
         bind:value={componentHeight} />
  {#if componentConfigWidget !== null}
    <svelte:component
        this={componentConfigWidget.config} 
        config={componentConfig}
        on:change={(e) => { componentConfig = e.detail; }}
    />
  {/if}
  <button id="addComponentButton"
          type="button"
          color="{$styles.appBackground}"
          on:click="{() => { addComponent(); }}" >
    Add Component
  </button>
</div>

<style>
  #addComponentDialog {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    margin: 5px;
    padding: 0px;
  }

  #addComponentButton {
    border-radius: 5px;
    font-size: 15px;
    height: 30px;
    outline: none;
    margin: 10px;
    padding: 6px 6px 6px 6px;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    -o-user-select: none;
    user-select: none;
    -webkit-tap-highlight-color: transparent;
    outline-style:none;
    background-color: rgba(255, 255, 255, 0.3);
  }

  .componentLabel {
    margin: 5px 10px 5px 0px;
  }

  .componentInput {
    margin: 5px 10px 10px 0px;
    border-radius: 10px;
    background-color: rgba(255, 255, 255, 0.3);
  }
</style>

<script>
  import { createEventDispatcher, onMount } from 'svelte';
  import { styles } from '../stores/styles.js';
  import { resizeWindow } from '../stores/resizeWindow.js';
  
  export let types;

  const dispatch = createEventDispatcher();
  
  let componentType = null;
  let componentHeight = 30;  
  let componentName = "";
  let componentConfig = {};
  let componentConfigWidget = null;
  let addDialogHeight = 280;
  
  $: componentConfigWidget = getConfigWidget(componentType);

  onMount(() => {
    componentType = types[0].moduleName;
    componentConfigWidget = getConfigWidget(componentType);
  });

  function getConfigWidget(type) {
    var comtype = {};
    if(type !== null) {
      comtype = types.find((comt) => { return comt.moduleName === type});
    }
    return comtype;
  }

  function addComponent() {
    if(componentConfig === null) {
      componentConfig = {};
    }
    dispatch('addNewComponent', {
      module: {
        name: componentName,
        widgetName: componentType,
        height: parseInt(componentHeight, 10),
        config: componentConfig
      }
    });
    componentName = '';
    componentType = null;
    componentHeight = 30;
    componentConfig = {};
  }

  function typeChanged() {
    setTimeout(()=>{
      $resizeWindow();
    }, 50);
  }
</script>
