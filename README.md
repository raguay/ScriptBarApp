![ScriptBarApp](https://socialify.git.ci/raguay/scriptbarapp/image?description=1&font=Raleway&forks=1&issues=1&language=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Dark)

[![Richard's GitHub stats](https://github-readme-stats.vercel.app/api?username=raguay)](https://github.com/anuraghazra/github-readme-stats)

# Script Bar

![Script Bar](https://github.com/raguay/ScriptBarApp/blob/main/img/ScriptBar.jpg?raw=truehttps)

An application to show results from scripts and Node-Red in the menubar. It is currently in the Alpha stage, but very usable. You can display results from the Node-Red server in [EmailIt Server](https://github.com/raguay/EmailItServer), command line scripts for either [TextBar](http://richsomerfield.com/apps/textbar/) (mostly working), [xBar](https://github.com/matryer/xbar) (not working - still in development), or some of the builtin components include web sites and IP addresses both local and external. You have to have the [EmailIt Server](https://github.com/raguay/EmailItServer) running to use this program. Currently, EmailItServer is now apart of [EmailIt](https://GitHub.com/raguay/EmailIt).

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

You must have the [EmailIt Server](https://github.com/raguay/EmailItServer) running first. Currently, EmailItServer is embedded inside of [EmailIt](https://GitHub.com/raguay/EmailIt).

## Documentation

The ScriptBar application is launched along with EmailIt. It relies upon the EmailIt Server for all of it's information. Just like EmailIt, it is a Wails application using Svelte for the frontend framework.

![ScriptBar Application](/img/ScriptBar.png)

This shows my ScriptBar. The `+` in the upper left corner allows you to add new items to each tab. The tab with a `+` in it allows you to add a new tab. By double clicking the tab names, you can change text shown for each tab. You can have as many tabs as you want and as many items in each tab as you want. But, if you get too many, it will go past your screen edges! You can click and drag the top area to move the application around.

![Double Click Entry](/img/doubleclickentry.png)

If you double click on the name of an entry, you can click `Up` to move it up the list or `Down` to move it down the list. By double clicking the entry name, you can edit it's parameters.

There are seven types of entries that you can use. They are described below:

#### Birthday Counter

![Birthday Counter](/img/birthdaycounter.png)

The birthday counter takes the output of a birthday counter flow in Node-Red. It will show the time to the person's next birthday and their current age. The script for the birthday node in Node-Red is:

```javascript
//
// Get the current date and the count down
// date.
//
var now = new Date();


var dmon = msg.payload.month - 1;
var dday = msg.payload.day;
var dyear = msg.payload.year;
var bd = new Date();
bd.setFullYear(now.getFullYear());
bd.setMonth(dmon);
bd.setDate(dday);

//
// Calculate the number of milliseconds until
// the date.
//
var diff = 0;
if(now.getMonth() <= dmon) {
    diff = bd.getTime() - now.getTime();
    if((now.getMonth() === dmon)&&(now.getDate() > dday)) {
        diff = 0;
    }
}

if(diff === 0) {
    var yearEnd = new Date();
    yearEnd.setMonth(11);
    yearEnd.setDate(31);
    diff = yearEnd.getTime() - now.getTime();
    bd.setFullYear(now.getFullYear()+1);
    diff += bd.getTime() - yearEnd.getTime();
}

//
// Convert the difference to months, weeks, days.
//
ddays = Math.floor(diff / (1000*60*60*24));
dweeks = Math.floor(ddays / 7);
ddays -= dweeks * 7;

//
// Create the answer and send it on.
//
var result = dweeks.toString() + ' weeks, ' + ddays.toString() + ' days';
var ans = {
    payload: result,
    topic: 'days'
}
return ans;
```

You put this code into a function node. Then create another function node with this code to calculate the current age:

```javascript
//
// Get the current date and the count down
// date.
//
var now = new Date();

var dmon = msg.payload.month - 1;
var dday = msg.payload.day;
var dyear = msg.payload.year;

//
// Create the answer and send it on.
//
var result = now.getFullYear() - dyear;
var ans = {
    payload: result,
    topic: 'age'
}
return ans;
```

You feed these two function nodes with an inject node with a payload like this:

```json
{"year":1969,"month":8,"day":1}
```

You need to change the `year`, `month`, and `day` fields to match the person's birthday. Then add a join node with these properties:

![Join Node for Birthday Counter](/img/joinnode.png)

The last node is a `SPVariables` node set to the name of the Node-Red variable you want to use. The full flow looks like this:

![Birthday Counter Flow](/img/birthdayflow.png)

#### Flow Variable

![Flow Variable](/img/flowvariable.png)

The flowvariable entry takes the text in a Node-Red flow variable and simply displays it.

#### Internal IP Address

![Internal IP Address](/img/intipaddress.png)

The Internal IP Address entry will display the internal IP address of your computer. That is the IP of your local network, not the IP used on the Internet. If you give a port number and check the link checkbox, you can click the IP address shown in ScriptBar and it will open a webpage to that address and port number. This entry type uses an API from the EmailItServer.

#### External IP Address

![External IP Address](/img/ipaddress.png)

The External IP Address is the same as the Internal IP address except that it takes a Node-Red flow variable name. You have to create a flow that fills in the external IP address of your system. There are several nodes that can supply that information, so I leave up to you. The one I use is an HTTP request node with the url set to `https://api.ipify.org?format=json`.

#### Script

![Script](/img/script.png)

The Script entry allows you to call an External Script from EmailIt and display it's output. The output types can be: Generic, TextBar, BitBar (which is now xBar), and HTML. The Generic setting just displays the output as plain text. The TextBar setting processes the output as a TextBar script. The BitBar setting processes the output as a BitBar script (not currently working), and the HTML setting displays the output as HTML. The `What is the command line?` isn't currently used. I plan to have it just run the command line and display the results. 

![Output of a Script](/img/scriptoutput.png)

The picture above shows an TextBar script output called TaskPaperTodo.rb. This is a ruby script that displays the tasks I have in the TaskPaper program.

#### Separator

![Separator](/img/separator.png)

This is a simple line separator to separate entry types.

#### Web Link

![Web Link](/img/weblink.png)

The Web Link entry type allows you to give a website address. By clicking on the glob on the ScriptBar program, it will open that web link in the default browser. I use this type a whole lot to keep up with websites that I go to often.

