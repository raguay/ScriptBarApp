<div class='editSpanField'>
  {#if editH2Flag}
    <textarea class='eListName'
           bind:value={name}
           bind:this={editField}
           on:keydown={(e) => {if(e.code === 'Enter') nameChanged(); }}
           on:blur={() => { nameChanged(); }}
    />
    {:else}
      <span class='pListName'
        on:dblclick={() => { editName(); }}
        on:contextmenu={(e) => { e.preventDefault(); middleButton();}}>
        {name}
      </span>
  {/if}
</div>

<style>
  .eListName {
    background-color: rgba(255,255,255,0.3);
    margin: 5px 0px 0px 0px;
    width: 90px;
    padding: 0px;
    border-radius: 10px;
    min-height: 16px;
    max-height: 16px;
  }

  .pListName {
    margin: 5px;
    padding: 0px;
    min-height: 16px;
    max-height: 16px;
  }

  .editSpanField {
    cursor: pointer;
  }
</style>

<script>
  import { createEventDispatcher, tick } from 'svelte';
  
  export let name;

  let editField;
  let editH2Flag = false;

  const disbatch = createEventDispatcher();

  function middleButton() {
    disbatch('middleButton',{});
  }
  
  async function editName() {
    editH2Flag = true;
    await tick();
    editField.focus();
  }

  function nameChanged() {
    disbatch('nameChanged', {
      name: editField.value
    });
    editH2Flag = false;
  }
</script>

