<script>
  import { createEventDispatcher, onMount, afterUpdate, tick } from "svelte";
  import AddComponent from "./AddComponent.svelte";
  import EditComponent from "./EditComponent.svelte";
  import EditSpanField from "./EditSpanField.svelte";
  import { styles } from "../stores/styles.js";
  import { headerPosition } from "../stores/headerPosition.js";
  import { resizeWindow } from "../stores/resizeWindow.js";
  import { width } from "../stores/width.js";
  import { timer } from "../stores/timer.js";
  import { config } from "../stores/config.js";
  import * as App from '../../wailsjs/go/main/App.js';
  import * as rt from '../../wailsjs/runtime/runtime.js';

  //
  // The following components are widgets in the application and
  // their configuration widget.
  //
  import FlowVariable from "./FlowVariable.svelte";
  import FlowVariableConfig from "./FlowVariableConfig.svelte";
  import Separator from "./Separator.svelte";
  import SeparatorConfig from "./SeparatorConfig.svelte";
  import BirthdayCounter from "./BirthdayCounter.svelte";
  import BirthdayCounterConfig from "./BirthdayCounterConfig.svelte";
  import IPAddress from "./IPAddress.svelte";
  import IPAddressConfig from "./IPAddressConfig.svelte";
  import Script from "./Script.svelte";
  import ScriptConfig from "./ScriptConfig.svelte";
  import WebLink from "./WebLink.svelte";
  import WebLinkConfig from "./WebLinkConfig.svelte";

  export let oldState;

  let containers = [
    {
      name: "Info",
      widgets: [],
    },
  ];
  let widgets = [];
  let widgetTypes;
  let showAdd = false;
  let showEdit = false;
  let widgetToEdit;
  let currentID = 0;
  let currentContainer;
  let currentComponent;
  let showContainerMenu = false;
  let containerRectBox;
  let showComponentMenu = false;
  let componentRectBox;
  let mainDOM;
  let currentContainerDOM;
  let newWidgets = null;

  const dispatch = createEventDispatcher();

  onMount(async () => {
    //
    // Create the widgetTypes structure.
    //
    widgetTypes = [
      {
        moduleName: "FlowVariable",
        module: FlowVariable,
        config: FlowVariableConfig,
        configHeight: 70,
      },
      {
        moduleName: "Separator",
        module: Separator,
        config: SeparatorConfig,
        configHeight: 70,
      },
      {
        moduleName: "BirthdayCounter",
        module: BirthdayCounter,
        config: BirthdayCounterConfig,
        configHeight: 70,
      },
      {
        moduleName: "IPAddress",
        module: IPAddress,
        config: IPAddressConfig,
        configHeight: 194,
      },
      {
        moduleName: "Script",
        module: Script,
        config: ScriptConfig,
        configHeight: 333,
      },
      {
        moduleName: "WebLink",
        module: WebLink,
        config: WebLinkConfig,
        configHeight: 133,
      },
    ];

    //
    // Set the old state if it's been set.
    //
    setOldState(oldState);

    //
    // Get the widget list from the user's configuration.
    //
    await getWidgets();

    //
    // Setup the window position.
    //
    window.moveTo(400, -20);

    $resizeWindow = resizeWindowFunction;

    //
    // Setup the timer.
    //
    $timer.StartTimer();

    //
    // Return a function to do cleanup when the component is removed.
    //
    return () => {
        $timer.StopTimer();
    };
  });

  afterUpdate(() => {
    if (newWidgets !== null) {
      widgets = addWidgets(newWidgets);
      newWidgets = null;
    }
    resizeWindowFunction();
  });

  async function resizeWindowFunction() {
    await tick();
    if ((mainDOM !== undefined)&&(mainDOM !== null)) {
      var nwidth = mainDOM.clientWidth;
      var nheight = mainDOM.clientHeight;
      if (nwidth < 100) {
        nwidth = 400;
      }
      if (nheight < 40) {
        nheight = 40;
      }
      if (showContainerMenu && widgets.length < 3) {
        nheight += 20 * (3 - widgets.length);
      }
      if (showComponentMenu && widgets.length - currentComponent < 3) {
        nheight += 20 * (widgets.length - currentComponent);
      }
      $headerPosition = Math.floor(nwidth / 2);
      $width = nwidth;
      rt.WindowSetSize(nwidth, nheight);
    }
  }

  function setOldState(old) {
    if (typeof old !== "undefined") {
      if (typeof old.container !== "undefined") {
        currentContainer = old.container;
        currentComponent = 0;
      } else {
        currentContainer = 0;
        currentComponent = 0;
      }
    } else {
      currentContainer = 0;
      currentComponent = 0;
    }
  }

  async function getWidgets() {
    //
    // Get the configuration.
    //
    var homedir = await App.GetHomeDir();
    var configPath = await App.AppendPath(homedir, ".scriptbar");
    var configfile = await App.ReadFile(configPath);
    $config = JSON.parse(configfile);
    containers = $config.containers;
    widgets = [];
    if (typeof containers[currentContainer].widgets !== "undefined") {
      containers[currentContainer].widgets = addWidgets(
        containers[currentContainer].widgets
      );
    }
  }

  function switchContainer(contNum) {
    widgets = [];
    showContainerMenu = false;
    showComponentMenu = false;
    currentContainer = contNum;
    $timer.ClearTimers();
    containers[contNum].widgets = addWidgets(containers[contNum].widgets);
  }

  async function saveWidgets() {
    //
    // Save the widgets for loading latter.
    //
    var homedir = await App.GetHomeDir();
    var configPath = await App.AppendPath(homedir, ".scriptbar");
    $config.containers = containers;
    await App.WriteFile(configPath, JSON.stringify($config));
  }

  function addWidgets(obj) {
    for (var i = 0; i < obj.length; i++) {
      var tmp = widgetTypes.find((mod) => obj[i].widgetName === mod.moduleName);
      if (typeof tmp !== "undefined") {
        //
        // Add the widget module definition.
        //
        obj[i].widget = tmp.module;
        obj[i].configWidget = tmp.config;
        obj[i].configHeight = tmp.configHeight;

        //
        // Add an ID.
        //
        obj[i].id = currentID;
        currentID = currentID + 1;

        //
        // Add it to the list of widgets to be displayed.
        //
        widgets.push(obj[i]);
      }
    }
    //
    // Adjust the window.
    //
    containers[currentContainer].widgets = widgets;
    return widgets;
  }

  function addNewComponent(e) {
    widgets = addWidgets([e.detail.module]);
    containers[currentContainer].widgets = widgets;
    closeAddDialog();
    saveWidgets();
  }

  function editWidget(widget) {
    widgetToEdit = widget;
    openEditDialog();
  }

  function editComponent(e) {
    widgets = widgets.map((item) => {
      if (item.id === e.detail.module.id) {
        item.name = e.detail.module.name;
        item.height = parseInt(e.detail.module.height, 10);
        item.config = e.detail.module.config;
      }
      return item;
    });
    closeEditDialog();
    containers[currentContainer].widgets = widgets;
    saveWidgets();
  }

  function deleteComponent(e) {
    widgets = widgets.filter((item) => {
      if (item.id !== e.detail.module.id) {
        return item;
      }
    });
    closeEditDialog();
    containers[currentContainer].widgets = widgets;
    saveWidgets();
  }

  function openEditDialog() {
    showEdit = true;
    showContainerMenu = false;
    showComponentMenu = false;
  }

  function closeEditDialog() {
    showEdit = false;
  }

  function openAddDialog() {
    showAdd = true;
    showContainerMenu = false;
    showComponentMenu = false;
  }

  function closeAddDialog() {
    showAdd = false;
  }

  function addTab() {
    containers.push({
      name: "New",
      widgets: [],
    });
    containers = containers;
    saveWidgets();
  }

  function changeContainerName(name) {
    containers[currentContainer].name = name;
    containers = containers;
    saveWidgets();
  }

  function toggleContainerMenu() {
    showContainerMenu = !showContainerMenu;
    if (showContainerMenu) {
      showComponentMenu = false;
      containerRectBox =
        containers[currentContainer].this.getBoundingClientRect();
    }
  }

  function moveTabLeft() {
    containerRectBox =
      containers[currentContainer - 1].this.getBoundingClientRect();
    if (currentContainer !== 0) {
      var tmp = containers[currentContainer];
      currentContainer = currentContainer - 1;
      containers[currentContainer + 1] = containers[currentContainer];
      containers[currentContainer] = tmp;
    }
    containers = containers;
    saveWidgets();
  }

  function moveTabRight() {
    containerRectBox =
      containers[currentContainer + 1].this.getBoundingClientRect();
    if (currentContainer !== containers.length - 1) {
      var tmp = containers[currentContainer];
      currentContainer = currentContainer + 1;
      containers[currentContainer - 1] = containers[currentContainer];
      containers[currentContainer] = tmp;
    }
    containers = containers;
    saveWidgets();
  }

  function deleteTab() {
    containers = containers.filter((item, index) => {
      if (index !== currentContainer) {
        return true;
      } else {
        return false;
      }
    });
    currentContainer = 0;
    widgets = [];
    containers[currentContainer].widgets = addWidgets(
      containers[currentContainer].widgets
    );
    widgets = containers[currentContainer].widgets;
    saveWidgets();
    showContainerMenu = false;
  }

  function toggleComponentMenu(info) {
    currentComponent = info.index;
    if (widgets.length > 1) {
      showComponentMenu = !showComponentMenu;
      if (showComponentMenu) {
        showContainerMenu = false;
        currentContainerDOM = info.dom;
        componentRectBox = info.dom.firstChild.getBoundingClientRect();
      } else {
        componentRectBox = {};
        currentComponent = 0;
        currentContainerDOM = undefined;
      }
    }
  }

  function moveComponentUp() {
    if (currentComponent !== 0) {
      componentRectBox = currentContainerDOM.firstChild.getBoundingClientRect();
      var tmp = widgets[currentComponent];
      currentComponent = currentComponent - 1;
      widgets[currentComponent + 1] = widgets[currentComponent];
      widgets[currentComponent] = tmp;
      containers[currentContainer].widgets = widgets;
      saveWidgets();
      newWidgets = widgets;
      widgets = [];
      showComponentMenu = false;
    }
  }

  function moveComponentDown() {
    if (currentComponent !== widgets.length - 1) {
      componentRectBox = currentContainerDOM.firstChild.getBoundingClientRect();
      var tmp = widgets[currentComponent];
      currentComponent = currentComponent + 1;
      widgets[currentComponent - 1] = widgets[currentComponent];
      widgets[currentComponent] = tmp;
      containers[currentContainer].widgets = widgets;
      saveWidgets();
      newWidgets = widgets;
      widgets = [];
      showComponentMenu = false;
    }
  }

  function changeView(data) {
    data.scriptbar = {
      container: currentContainer,
    };
    dispatch("changeView", data);
  }

  function deleteWidgets() {
    const size = widgets.length;
    for (var i = 0; i < size; i++) {
      delete widgets[i];
    }
    widgets = [];
  }

  function Quit() {
    window.go.main.App.Quit();
  }
</script>

<div
  id="main"
  style="background-color: {$styles.appBackground};
         color: {$styles.textcolor};
         font-family: {$styles.fontFamily};
         font-size: {$styles.fontSize};
        "
  bind:this={mainDOM}
>
  {#if showAdd}
    <div
      id="addDialog"
      style="width: 100%; 
                background-color: {$styles.appBackground};
                color: {$styles.textcolor};
                font-family: {$styles.fontFamily};
                font-size: {$styles.fontSize};"
    >
      <div id="addDialogControls">
        <p
          id="closeAddDialogControl"
          on:click={() => {
            closeAddDialog();
          }}
        >
          X
        </p>
      </div>
      <AddComponent types={widgetTypes} on:addNewComponent={addNewComponent} />
    </div>
  {:else if showEdit}
    <div
      id="editDialog"
      style="width: 100%; 
                background-color: {$styles.appBackground};
                color: {$styles.textcolor};
                font-family: {$styles.fontFamily};
                font-size: {$styles.fontSize};"
    >
      <div id="editDialogControls">
        <p
          id="closeEditDialogControl"
          on:click={() => {
            closeEditDialog();
          }}
        >
          X
        </p>
      </div>
      <EditComponent
        widget={widgetToEdit}
        on:editComponent={editComponent}
        on:deleteComponent={deleteComponent}
      />
    </div>
  {:else}
    {#if showContainerMenu}
      <div
        id="containerMenu"
        style="background-color: {$styles.appBackground};
                  top: {Math.floor(
          containerRectBox.height / 2 + containerRectBox.y
        )}px; 
                  left: {Math.floor(
          containerRectBox.x + containerRectBox.width / 2
        )}px;"
      >
        {#if currentContainer !== 0}
          <p
            class="ContainerMenuItem"
            on:click={(e) => {
              moveTabLeft();
            }}
          >
            Move Left
          </p>
        {/if}
        {#if currentContainer !== containers.length - 1}
          <p
            class="ContainerMenuItem"
            on:click={(e) => {
              moveTabRight();
            }}
          >
            Move Right
          </p>
        {/if}
        <p
          class="ContainerMenuItem"
          on:click={(e) => {
            deleteTab();
          }}
        >
          Delete
        </p>
      </div>
    {/if}
    {#if showComponentMenu}
      <div
        id="ComponentMenu"
        style="background-color: {$styles.appBackground};
                  top: {Math.floor(componentRectBox.y)}px; 
                  left: 20px;"
      >
        {#if currentComponent !== 0}
          <p
            class="ComponentMenuItem"
            on:click={(e) => {
              moveComponentUp();
            }}
          >
            Up
          </p>
        {/if}
        {#if currentComponent !== widgets.length - 1}
          <p
            class="ComponentMenuItem"
            on:click={(e) => {
              moveComponentDown();
            }}
          >
            Down
          </p>
        {/if}
      </div>
    {/if}
    <div id="controls">
      <p
        id="addControl"
        on:click={() => {
          openAddDialog();
        }}
      >
        +
      </p>
      <p
        id="closeControl"
        on:click={() => {
          Quit();
        }}
      >
        X
      </p>
    </div>
    <div id="tabs">
      {#each containers as container, index}
        {#if index === currentContainer}
          <div class="activeTab" bind:this={container.this}>
            <EditSpanField
              name={container.name}
              on:nameChanged={(e) => {
                changeContainerName(e.detail.name);
              }}
              on:middleButton={(e) => {
                toggleContainerMenu();
              }}
            />
          </div>
        {:else}
          <div
            class="tab"
            on:click={() => {
              switchContainer(index);
            }}
            bind:this={container.this}
          >
            {container.name}
          </div>
        {/if}
      {/each}
      <div
        id="addTab"
        class="tab"
        on:click={() => {
          addTab();
        }}
      >
        <span id="addTabText">+</span>
      </div>
    </div>
    <div id="componentList">
      {#if widgets.length > 0}
        {#each widgets as widget, index}
          <svelte:component
            this={widget.widget}
            {index}
            height={widget.height}
            name={widget.name}
            config={widget.config}
            on:dblclick={() => {
              editWidget(widget);
            }}
            on:middleButton={(e) => {
              e.preventDefault();
              toggleComponentMenu(e.detail);
            }}
            on:changeView={(e) => {
              changeView(e.detail);
            }}
          />
        {/each}
      {/if}
    </div>
  {/if}
</div>

<style>
  #main {
    position: absolute;
    left: 1px;
    top: 1px;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    margin: 0px;
    padding: 5px;
    border: 0px transparent;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
  }

  #tabs {
    display: flex;
    flex-direction: row;
    margin: 0px 0px 0px 0px;
    padding: 0px 0px 0px 0px;
    min-height: 30px;
    max-height: 30px;
    height: 30px;
  }

  #tabs::-webkit-scrollbar {
    height: 3px;
    background-color: rgba(255, 255, 255, 0.5);
    border-radius: 3px;
  }

  #tabs::-webkit-scrollbar-thumb {
    height: 3px;
    background-color: rgba(10, 10, 10, 0.3);
    border-radius: 3px;
  }

  #addTab {
    margin: 0px 0px 0px auto;
    cursor: pointer;
  }

  #addTabText {
    color: red;
    padding: 0px;
    margin: 0px;
  }

  #containerMenu {
    position: absolute;
    border: 2px solid rgba(255, 255, 255, 0.3) !important;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    padding: 5px 10px 5px 10px;
  }

  #closeEditDialogControl {
    color: red;
    cursor: pointer;
    margin: 0px 0px 0px auto;
  }

  #ComponentMenu {
    position: absolute;
    border: 2px solid rgba(255, 255, 255, 0.3) !important;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    padding: 5px 10px 5px 10px;
  }

  .ComponentMenuItem {
    padding: 0px;
    margin: 3px 0px 3px 0px;
  }

  .ComponentMenuItem:hover {
    color: white;
  }

  .ContainerMenuItem {
    cursor: pointer;
    margin: 3px 0px 3px 0px;
    padding: 0px;
  }

  .ContainerMenuItem:hover {
    color: white;
  }

  .tab {
    border: 2px solid rgba(255, 255, 255, 0.3) !important;
    border-bottom: 0px !important;
    border-radius: 50% 20% 0px 0px;
    cursor: pointer;
    user-select: none;
    padding: 2px 8px;
    margin: 0px 1px 0px 0px;
    height: 20px;
  }

  .activeTab {
    border: 2px solid rgba(255, 255, 255, 0.6) !important;
    border-bottom: 0px !important;
    border-radius: 50% 20% 0px 0px;
    cursor: pointer;
    user-select: none;
    padding: 2px 8px;
    margin: 0px 1px 0px 0px;
    height: 20px;
  }

  #editDialog {
    margin: 0px;
    padding: 0px;
    height: 100%;
    width: 100%;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
  }

  #editDialogControls {
    margin: 0px 5px 0px auto;
    padding: 0px;
    color: red;
    cursor: pointer;
  }

  #addDialog {
    margin: 0px;
    padding: 0px;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
  }

  #controls {
    display: flex;
    flex-direction: row;
    justify-content: flex;
    width: 100%;
    height: 10px;
    margin: 0px 0px 10px 0px;
    padding: 0px;
    --wails-draggable: drag;
  }

  #closeControl {
    margin: 0px 5px 0px auto;
    padding: 0px;
    color: red;
    cursor: pointer;
  }

  #addControl {
    margin: 0px 0px 0px 5px;
    padding: 0px;
    color: red;
    cursor: pointer;
  }

  #addDialogControls {
    margin: 0px 5px 0px auto;
    padding: 0px;
    color: red;
    cursor: pointer;
  }

  #closeAddDialogControl {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    width: 100%;
    height: 10px;
    margin: 0px 0px 10px 0px;
    padding: 0px;
  }

  #componentList {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: 0px;
    white-space: nowrap;
    padding: 0px;
  }
</style>
