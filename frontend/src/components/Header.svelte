<div 
  id="HeaderOutside" 
  style="width: {$width}px;"
  bind:this={headerDOM}
>
    <div 
      id="HeaderInside" 
      style="left: {$headerPosition}px; height: 100%; border-bottom-color: {$styles.appBackground};"
    >
    </div>
</div>

<style>
    #HeaderOutside {
      position: absolute;
      top: 0px;
      width: 100%;
      height: 10px;
      align-content: center;
      margin: 0px;
      padding: 0px;
    }

    #HeaderInside {
      position: absolute;
      color: transparent;
      top: -9px;
      margin: 0px;
      padding: 0px;
      background-color: transparent;
      border-top: 0px;        
      border-left: 5px solid transparent;
      border-right: 5px solid transparent;
      border-bottom: 10px solid;
    }
</style>

<script>
  import { styles } from '../stores/styles.js';
  import { headerPosition } from '../stores/headerPosition.js';
  import { width } from '../stores/width.js';
  
  let headerDOM;
</script>
