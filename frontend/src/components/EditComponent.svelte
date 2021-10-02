<div id='editComponentDialog'>
  <label class="componentLabel" 
         for="componentName">
    What's the name?
  </label>
  <input type="text"
         class="componentInput"
         id="componentName"
         color="{$styles.appBackground}"
         bind:value={componentName} />
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
        this={componentConfigWidget} 
        config={componentConfig}
        on:change={(e) => { componentConfig = e.detail; }}
    />
  {/if}
  <div id="buttonRow">
    <button id="editComponentButton"
            type="button"
            color="{$styles.appBackground}"
            on:click="{() => { editComponent(); }}" >
      Save Component
    </button>
    <button id="deleteComponentButton"
            type="button"
            color="{$styles.appBackground}"
            on:click="{() => { deleteComponent(); }}" >
      Delete Component
    </button>
  </div>
</div>

<style>
  #editComponentDialog {
    display: flex;
    flex-direction: column;
    margin: 5px;
    padding: 0px;
  }

  #editComponentButton {
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
  
  #deleteComponentButton {
    border-radius: 5px;
    font-size: 15px;
    height: 30px;
    outline: none;
    margin: 10px 10px 10px auto;
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

  #buttonRow {
    display: flex;
    flex-direction: row;
    margin: 0px;
    padding: 0px;
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
  import { createEventDispatcher, onMount, tick } from 'svelte';
  import { styles } from '../stores/styles.js';
  
  export let widget;

  const dispatch = createEventDispatcher();
  
  let componentType = null;
  let componentHeight = 30;  
  let componentName = "";
  let componentConfig = {};
  let componentConfigWidget = null;
  let editDialogHeight = 250;

  $: widget = assignVars(widget);

  onMount(() => {
    widget = assignVars(widget);
  })

  function assignVars(wid) {
    componentConfigWidget = wid.configWidget;
    componentName = wid.name;
    componentType = wid.widgetName;
    componentHeight = wid.height;
    componentConfig = wid.config;
    return(wid);
  }
  
  function editComponent() {
    if(componentConfig === null) {
      componentConfig = {};
    }
    dispatch('editComponent', {
      module: {
        id: widget.id,
        name: componentName,
        widgetName: componentType,
        height: parseInt(componentHeight, 10),
        config: componentConfig
      }
    });
  }
  
  function deleteComponent() {
    dispatch('deleteComponent', {
      module: {
        id: widget.id
      }
    });
  }
</script>
