<div class='IPAddressConfig'>
  <label for='Variable'
         class='componentLabel' >
    What is the variable name?
  </label>
  <input id='Variable'
         type="text"
         class="componentInput"
         bind:value={config.flowVar}
         on:change={valueChanged}
  />
  <label for='Link'
         class='componentLabel' >
    Have a link?
  </label>
  <input id='Link'
         type="checkbox"
         class="componentInput"
         bind:checked={config.showLink}
         on:change={valueChanged}
  />
  <label for='Port'
         class='componentLabel' >
    What port to open on click?
  </label>
  <input id='Port'
         type="text"
         class="componentInput"
         bind:value={config.port}
         on:change={valueChanged}
  />
</div>

<style>
  .IPAddressConfig {
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

  function valueChanged(e) {
    console.log(config.showLink);
    dispatch('change',{
          flowVar: config.flowVar,
          showLink: config.showLink,
          port: parseInt(config.port, 10)
         });
  }
</script>

