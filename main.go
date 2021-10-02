package main

import (
    _ "embed"
	"github.com/wailsapp/wails"
  "github.com/atotto/clipboard"
)

func writeClipboard(txt string) {
  clipboard.WriteAll(txt)
}

func readClipboard() string {
	txt, _ := clipboard.ReadAll()
	return txt
}


//go:embed frontend/public/bundle.js
var js string

//go:embed frontend/public/bundle.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 590,
		JS:     js,
		CSS:    css,
		Colour: "rgba(0,0,0,0)",
		Resizable: false,
	})

	app.Bind(writeClipboard)
	app.Bind(readClipboard)
	app.Run()
}
