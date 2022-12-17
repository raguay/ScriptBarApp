<script>
  import { createEventDispatcher } from "svelte";
  import AnsiUp from "ansi_up";
  import { styles } from "../stores/styles.js";
  import { timer } from "../stores/timer.js";
  import * as App from '../../wailsjs/go/main/App.js';

  export let name;
  export let config = {
    script: "",
    env: "",
    envVar: "",
    commandLine: "",
    type: "",
    scriptImage: "",
    appBackground: $styles.appBackground,
    highlight: $styles.highlight,
  };
  export let index;
  export let height;

  let dom;

  let value = "loading...";
  let bodyHTML = "<h1>Loading...</h1>";
  let vHTML = false;
  let bodyConfig = {
    changeColor: false,
    width: 300,
    height: 300,
    showButton: true,
  };

  const dispatch = createEventDispatcher();

  $: updateWidget(index);

  async function getData() {
    //
    // Get the current value instead of waiting for the next update.
    //
    var callBody = JSON.stringify({
      script: config.script,
      text: config.commandLine,
      env: config.env,
      envVar: config.envVar,
      commandLine: config.commandLine,
    });
    let data = await App.SendHTTPQuery("PUT","http://localhost:9978/api/script/run", callBody)
    processScriptData(JSON.parse(data));
  }

  async function updateWidget(index) {
    $timer.AddTimer(config.script, getData);
    await getData();
    index = index;
  }

  function middleButton() {
    dispatch("middleButton", {
      index: index,
      dom: dom,
    });
  }

  function sendWebView() {
    if (typeof value === "string" && !value.includes("loading...")) {
      dispatch("changeView", {
        name: "webview",
        body: {
          html: bodyHTML,
          config: bodyConfig,
        },
      });
    }
  }

  function processScriptData(data) {
    bodyConfig.showButton = true;
    if (config.type === "generic") {
      //
      // It is a generic script. No processing other than
      // creating the HTML as text paragraph.
      //
      bodyHTML = "<p>" + data + "</p>";
    } else if (config.type === "bitbar") {
      //
      // Process it as a BitBar script.
      //
      bodyHTML = data;
    } else if (config.type === "textbar") {
      //
      // Process it as a TextBar Script.
      //
      bodyHTML = processTextBarData(data);
    } else if (config.type === "web") {
      //
      // The script output is html. Send it directly.
      //
      bodyHTML = data;
    }
  }

  function processTextBarData(data) {
    var resultText = "";
    var resultHTML = "";
    var maxWidth = 10;
    var result = null;
    bodyConfig = {};
    bodyConfig.height = 10;
    var count = 1;
    var viewSet = false;
    var textbarDirectives = false;
    var longLine = "";
    var ansi_up = new AnsiUp();
    bodyConfig.showButton = true;
    globalThis.ScriptClick = async (count, line, script, commandLine, env) => {
      var newEnv = [];
      newEnv.push(`TEXTBAR_INDEX=${count}`);
      newEnv.push(`TEXTBAR_TEXT="${line}"`);
      var callBody = {
        script: script.trim(),
        env: env.trim(),
        envVar: newEnv,
      };
      if(typeof config.envVar !== 'undefined') {
        callBody.envVar = [...config.envVar, ...callBody.envVar];
      }
      if (
        typeof commandLine !== "undefined" &&
        !commandLine.includes("undefined")
      ) {
        commandLine = commandLine.trim();
        if (commandLine !== "") {
          callBody.commandLine = commandLine;
          callBody.text = commandLine;
        }
      }
      const data = await App.SendHTTPQuery("PUT","http://localhost:9978/api/script/run/", JSON.stringify(callBody));
      globalThis.closeWebView();
    };
    if(data[0] === '"') {
      data = data.slice(1,-1);
    }
    data = data.replaceAll("\\n","\n").split(/\r?\n/);
    value = config.scriptImage + data[0];
    data.slice(1).forEach((line) => {
      if (line.trim() !== "") {
        line = new String(line);
        var banner = line.includes("----TEXTBAR----");
        if (banner || textbarDirectives) {
          if (banner) {
            //
            // All rest of the lines are part of the directives.
            //
            textbarDirectives = true;
          } else {
            //
            // Process each directive.
            //
            var parts = line.split("=");
            switch (parts[0]) {
              case "REFRESH":
                config.period = parseInt(parts[1]);
                break;
              case "VIEWTYPE":
                switch (parts[1]) {
                  case "HTML":
                    result = resultHTML;
                    bodyConfig.showButton = true;
                    break;
                  default:
                    result = resultText;
                    bodyConfig.showButton = false;
                    break;
                }
                break;
              case "VIEWSIZE":
                var sizes = parts[1].split(",");
                bodyConfig.width = parseInt(sizes[0], 10);
                bodyConfig.height = parseInt(sizes[1], 10);
                viewSet = true;
                break;
              case "IMAGE":
                if (parts[1][0] === ":") {
                  value = '<img src="' + parts[1].slice(1) + '" />';
                  vHTML = true;
                } else {
                  value = parts[1];
                }
                bodyConfig.showButton = true;
                break;
              default:
                break;
            }
          }
        } else {
          //
          // It is the main output lines and not the directive lines.
          //
          resultHTML += line;
          var nline = "";
          var pline = "";
          if (line.match(/\[\d+(\:\d+)*m/) !== null) {
            //
            // This line has ASNI codes. Process it as such.
            //
            var tline = line.replace(/\\e/g, "\x1b");

            //
            // Convert ANSI color codes to span colors.
            //
            nline = ansi_up.ansi_to_html(tline);

            //
            // Clean the line of unknown escape codes.
            //
            nline = nline.replace(/\[\d+(\:\d+)*m/g, "");

            //
            // Remove all escapte codes for the plain line.
            //
            pline = line
              .replace(/\\e/g, "")
              .replace(/\[\d+((\:|\;)\d+)*m/g, "")
              .replace("\t","");
          } else {
            //
            // No ANSI codes, process as a plain text line.
            //
            nline = line.replace(/ /g, "&nbsp;");

            //
            // Just use the line for the pline to determine line length.
            //
            pline = line;
          }

          //
          // Remove everything but the visible text.
          //
          pline = new String(pline);
          if (maxWidth < pline.length) {
            maxWidth = pline.length;
            longLine = pline;
          }
          if (line[line.length - 1] === ":") {
            resultText +=
              '<p style="margin: 0px; padding: 0px 0px 0px 0px; height: 20px; white-space: nowrap;">' +
              nline +
              "</p>";
          } else {
            resultText += `<p style="margin: 0px; padding: 0px 0px 0px 5px; height: 20px; border: 0px solid transparent; border-radius: 5px; white-space: nowrap;" 
                              onclick="globalThis.ScriptClick(${count},'${pline}','${config.script}', '${config.commandLine}', '${config.env}')"
                              onmouseenter="window.setLineColorNew(this)"
                              onmouseleave="window.setLineColor(this)"
                              onblur="window.setLineColor(this)"
                            >
                              ${nline}
                            </p>`;
          }
        }
        count += 1;
      }
    });

    //
    // Setup the height and width of the script output.
    //
    if (!viewSet) {
      bodyConfig.changeColor = true;
      //
      // These two colors should be set by the configuration.
      //
      bodyConfig.backgroundColor = config.appBackground;
      bodyConfig.highlightColor = config.highlight;
    } else {
      bodyConfig.changeColor = false;
    }
    if (result === null) result = resultText;
    return result;
  }
</script>

<div
  class="script"
  style="height: {Math.floor(height)}px;"
  data-index={index}
  bind:this={dom}
  on:mouseenter={() => {
    dom.style.color = $styles.highlight;
  }}
  on:mouseleave={() => {
    dom.style.color = $styles.textcolor;
  }}
  on:mouseout={() => {
    dom.style.color = $styles.textcolor;
  }}
  on:blur={() => {
    dom.style.color = $styles.textcolor;
  }}
  on:contextmenu={(e) => {
    e.preventDefault();
    middleButton();
  }}
>
  <span
    class="scriptName"
    on:dblclick={(e) => {
      dispatch("dblclick", {});
    }}
  >
    {name}
  </span>
  <span
    class="scriptValue"
    on:click={(e) => {
      sendWebView();
    }}
  >
    {#if vHTML}
      {@html value}
    {:else}
      {value}
    {/if}
  </span>
</div>

<style>
  .script {
    width: 100%;
    display: flex;
    flex-direction: row;
    margin: auto 0px auto 0px;
    padding: 0px;
  }

  .scriptName {
    margin: 0px;
    padding: 0px 10px 0px 0px;
  }

  .scriptValue {
    margin: 0px 0px 0px auto;
    padding: 0px;
    cursor: pointer;
  }
</style>
