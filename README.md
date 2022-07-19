![ScriptBarApp](https://socialify.git.ci/raguay/scriptbarapp/image?description=1&font=Raleway&forks=1&issues=1&language=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Dark)

[![Richard's GitHub stats](https://github-readme-stats.vercel.app/api?username=raguay)](https://github.com/anuraghazra/github-readme-stats)

# Script Bar

![Script Bar](https://github.com/raguay/ScriptBarApp/blob/main/img/ScriptBar.jpg?raw=truehttps)

An application to show results from scripts and Node-Red in the menubar. It is currently in the Alpha stage, but very usable. You can display results from the Node-Red server in [EmailIt Server](https://github.com/raguay/EmailItServer), command line scripts for either [TextBar](http://richsomerfield.com/apps/textbar/) (mostly working), [xBar](https://github.com/matryer/xbar) (not working - still in development), or some of the builtin components include web sites and IP addresses both local and external. You have to have the [EmailIt Server](https://github.com/raguay/EmailItServer) running to use this program.

The program is built with a [Svelte](https://svelte.dev/) JavaScript frontend and [Wails](https://wails.app/) backend and packager. The goal is to make this project usable on Windows, macOS, and Linux. I'm currently developing it on macOS and will be working of the other versions after the macOS is stable and Wails version 2 is ready for that platform. 

## Note:

This program isn't a replacement for xBar. xBar puts script outputs into individual menubar items. ScriptBar places all script output into a single interface for easy viewing when activated. I have a very cluttered menubar with many application that I use. So, I wanted some scripts to show in it's own menubar using xBar, and some grouped together in a single interface for convience. Pick the program that best suites your needs or use both like I do! This program also get output from a Node-Red server in the EmailIt Server program.

## Building

To build the application, first go to the `frontend` directory and build it using:

```sh
npm install
npm run build
```

The first line installs the libraries used to build the application. The second line compiles the Svelte front end. Next, you need to build the Wails app. First, install Wails by installing the prerequisites mentioned on their webiste and then run:

```sh
go get -u github.com/wailsapp/wails/cmd/wails
```

Then build the application or your platform with:

```sh 
wails build
```

or, to build a macOS universal program, use this command line:

```sh
wails build --platform "darwin/universal"
```

To run the program, simply run the executable in the `build/bin` directory:

```sh 
open ./build/bin/ScriptBar.app
```

You must have the [EmailIt Server](https://github.com/raguay/EmailItServer) running first.

## Documentation

This is coming. I haven't had time to work on this yet!
