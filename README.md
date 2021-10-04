# ScriptBarApp

![Script Bar with NodeRed Items](tmp/ScriptBar1.png)

An application to show results from scripts and Node-Red in the menubar. It is currently in the Alpha stage, but very usable. You can display results from the Node-Red server in [EmailIt+ Server](https://github.com/raguay/EmailItServer), command line scripts for either [TextBar](http://richsomerfield.com/apps/textbar/) (mostly working), [xBar](https://github.com/matryer/xbar) (not working - still in development), or some of the builtin components include web sites and IP addresses both local and external. You have to have the [EmailIt+ Server](https://github.com/raguay/EmailItServer) running to use this program.

![Script Bar with Web Links](tmp/ScriptBar2.png)

The program is built with a [Svelte](https://svelte.dev/) JavaScript frontend and [Wails](https://wails.app/) backend and packager. The goal is to make this project usable on Windows, macOS, and Linux. I'm currently developing it on macOS and will be working of the other versions after the macOS is stable. 
