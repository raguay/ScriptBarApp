<div class='BirthdayConfig'>
  <label for='Name'
         class='componentLabel' >
    Who's birthday? (Name of Variable)
  </label>
  <input id='Name'
         type="text"
         class="componentInput"
         bind:value={config.flowVar}
         on:change={() => { dispatch('change',{
          flowVar: config.flowVar
         });}}
  />
</div>

<style>
  .BirthdayConfig {
    display: flex;
    flex-direction: column;
    width: 100%;
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
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher();
  
  export let config;
</script>

