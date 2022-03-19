<div id='ScriptConfig'>
  <label for='Name'
         class='componentLabel' >
    What is the name of the script?
  </label>
  <input id='Name'
         type="text"
         class="componentInput"
         bind:value={config.script}
         on:change={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           period: config.period,
           envVar: config.envVar,
           commandLine: config.commandLine,
           type: config.type,
           scriptImage: config.scriptImage, 
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  />
  <label for='CommandLine'
         class='componentLabel' >
    What is the command line?
  </label>
  <input id='CommandLine'
         type="text"
         class="componentInput"
         bind:value={config.commandLine}
         on:change={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           period: config.period,
           envVar: config.envVar,
           commandLine: config.commandLine,
           type: config.type,
           scriptImage: config.scriptImage, 
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  />
  <label for='Env'
         class='componentLabel' >
    What is the name of the environment?
  </label>
  <input id='Env'
         type="text"
         class="componentInput"
         bind:value={config.env}
         on:change={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           envVar: config.envVar,
           period: config.period,
           commandLine: config.commandLine,
           type: config.type,
           scriptImage: config.scriptImage, 
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  />
  <label for='EnvVar'
         class='componentLabel' >
    What environment variables?
  </label>
  <input id='EnvVar'
         type="text"
         class="componentInput"
         bind:value={config.envVar}
         on:change={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           envVar: config.envVar,
           period: config.period,
           commandLine: config.commandLine,
           type: config.type,
           scriptImage: config.scriptImage,
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  />
  <label for='Period'
         class='componentLabel' >
    How often to reload? (number of minutes)
  </label>
  <input id='Period'
         type="text"
         class="componentInput"
         bind:value={config.period}
         on:change={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           envVar: config.envVar,
           period: config.period,
           commandLine: config.commandLine,
           type: config.type,
           period: config.period,
           scriptImage: config.scriptImage,
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  />
 <label for='Type'
         class='componentLabel' >
    What type of script?
  </label>
  <select id='Type'
         class="componentInput"
         bind:value={config.type}
         on:blur={() => { dispatch('change',{
           script: config.script,
           env: config.env,
           envVar: config.envVar,
           period: config.period,
           commandLine: config.commandLine,
           type: config.type,
           scriptImage: config.scriptImage,
           highlight: config.highlight,
           appBackground: config.appBackground
         });}}
  >
    <option value={"textbar"}>TextBar</option>
    <option value={"bitbar"}>BitBar</option>
    <option value={"generic"}>Generic</option>
    <option value={"web"}>Web Page</option>
  </select>
  {#if config.type === "textbar"}
    <label for='ScriptIcon'
           class='componentLabel' >
      What emoji?
    </label>
    <input id='ScriptIcon'
           type="text"
           class="componentInput"
           bind:value={config.scriptImage}
           on:change={() => { dispatch('change',{
             script: config.script,
             env: config.env,
             envVar: config.envVar,
             commandLine: config.commandLine,
             period: config.period,
             type: config.type,
             scriptImage: config.scriptImage,
             highlight: config.highlight,
             appBackground: config.appBackground
           });}}
    />
    <label for='ScriptHigh'
           class='componentLabel' >
      What Highlight color?
    </label>
    <input id='ScriptHigh'
           type="text"
           class="componentInput"
           bind:value={config.highlight}
           on:change={() => { dispatch('change',{
             script: config.script,
             env: config.env,
             envVar: config.envVar,
             period: config.period,
             commandLine: config.commandLine,
             type: config.type,
             scriptImage: config.scriptImage,
             highlight: config.highlight,
             appBackground: config.appBackground
           });}}
    />
    <label for='ScriptBack'
           class='componentLabel' >
      What background color?
    </label>
    <input id='ScriptBack'
           type="text"
           class="componentInput"
           bind:value={config.appBackground}
           on:change={() => { dispatch('change',{
             script: config.script,
             env: config.env,
             envVar: config.envVar,
             period: config.period,
             commandLine: config.commandLine,
             type: config.type,
             scriptImage: config.scriptImage,
             highlight: config.highlight,
             appBackground: config.appBackground
           });}}
    />
  {/if}
</div>

<style>
  #ScriptConfig {
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
  import { styles } from '../stores/styles.js';
  
  const dispatch = createEventDispatcher();
  
  export let config = {
    script: "",
    env: "",
    envVar: "",
    period: 0,
    commandLine: "",
    scriptImage: "",
    highlight: $styles.highlight,
    appBackground: $styles.backgroundColor
  };
</script>

