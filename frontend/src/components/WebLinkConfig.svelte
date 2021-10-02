<div class='WebLinkConfig'>
  <label for='LinkAddress'
         class='componentLabel' >
    What is the link address?
  </label>
  <input id='LinkAddress'
         type="text"
         class="componentInput"
         bind:value={config.link}
         on:change={() => { dispatch('change',{
           link: config.link,
           browser: config.browser
         });}}
  />
</div>

<style>
  .WebLinkConfig {
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
  
  export let config;
  
  const dispatch = createEventDispatcher();
  
  let browserList = [];
</script>
