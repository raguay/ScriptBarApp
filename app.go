package main

import (
	"bytes"
	"context"
	//	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"time"

	clip "github.com/atotto/clipboard"
	cp "github.com/otiai10/copy"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	err string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	//
	// Run the server for getting variable information.
	//
	go RunServer(&a, ctx)
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) Quit() {
	rt.Quit(a.ctx)
}

type FileParts struct {
	Dir       string
	Name      string
	Extension string
}

type FileInfo struct {
	Dir       string
	Name      string
	Extension string
	IsDir     bool
	Size      int64
	Modtime   string
}

type Environment struct {
	Name   string   `json:"name" binding:"required"`
	EnvVar []string `json:"envVar" binding:"required"`
}

type ServerResponce struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func (a *App) SendHTTPQuery(method string, uri string, body string) string {
	var result string
	switch method {
	case "GET":
		{
			resp, err := http.Get(uri)
			if err != nil {
				print(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			result = string(body)
		}

	case "POST":
		{
			resp, err := http.Post(uri, "application/json", bytes.NewBuffer([]byte(body+"\n")))
			if err != nil {
				print(err)
			}
			defer resp.Body.Close()
			repBody, errdec := ioutil.ReadAll(resp.Body)
			if errdec != nil {
				print(errdec)
			}
			result = string(repBody)
		}
	case "PUT":
		{
			fmt.Print("\n\nPUT request\n\n")
			client := http.Client{}

			// regardless of GET or POST, we can safely add the body
			fmt.Print("\n", body)
			req, err := http.NewRequest("PUT", uri, bytes.NewBuffer([]byte(body+"\n")))
			if err != nil {
				fmt.Print(err.Error())
				return err.Error()
			}
			fmt.Print("\nCreate the headers...")
			headers := map[string]string{}
			headers["Method"] = "PUT"
			headers["Content-Type"] = "application/json"

			// for each header passed, add the header value to the request
			fmt.Print("\nSet the headers...")
			for k, v := range headers {
				req.Header.Set(k, v)
			}

			// finally, do the request
			fmt.Print("\nSend the request...")
			res, err := client.Do(req)
			if err != nil {
				return err.Error()
			}

			if res == nil {
				return fmt.Sprint("error: calling %s returned empty response", uri)
			}

			fmt.Print("\nGet the response...")
			responseData, err := io.ReadAll(res.Body)
			if err != nil {
				return err.Error()
			}

			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return fmt.Sprintf("error calling %s:\nstatus: %s\nresponseData: %s", uri, res.Status, responseData)
			}

			result = string(responseData)
		}
	}
	fmt.Print("\nReturn result: ", result)
	return result
}

func (a *App) SystemOpenFile(prog string) {
	if a.FileExists(prog) {
		var ArgsArray [2]string
		var sar []string
		ArgsArray[1] = prog
		a.RunCommandLine("/usr/bin/open", ArgsArray[:], sar, "")
	}
}

func (a *App) GetExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		a.err = err.Error()
	}
	return ex
}

func (a *App) Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		a.err = err.Error()
	}
	return wd
}

func (a *App) ReadFile(path string) string {
	a.err = ""
	contents, err := os.ReadFile(path)
	if err != nil {
		a.err = err.Error()
	}
	return string(contents[:])
}

func (a *App) GetHomeDir() string {
	a.err = ""
	hdir, err := os.UserHomeDir()
	if err != nil {
		a.err = err.Error()
	}
	return hdir
}

func (a *App) WriteFile(path string, data string) {
	err := os.WriteFile(path, []byte(data), 0666)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) FileExists(path string) bool {
	a.err = ""
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (a *App) DirExists(path string) bool {
	a.err = ""
	dstat, err := os.Stat(path)
	if err != nil {
		a.err = err.Error()
		return false
	}
	return dstat.IsDir()
}

func (a *App) SplitFile(path string) FileParts {
	a.err = ""
	var parts FileParts
	parts.Dir, parts.Name = filepath.Split(path)
	parts.Extension = filepath.Ext(path)
	return parts
}

func (a *App) ReadDir(path string) []FileInfo {
	a.err = ""
	var result []FileInfo
	result = make([]FileInfo, 0, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		a.err = err.Error()
	} else {
		for _, file := range files {
			var fileInfo FileInfo
			fileInfo.Name = file.Name()
			fileInfo.Size = file.Size()
			fileInfo.IsDir = file.IsDir()
			fileInfo.Modtime = file.ModTime().Format(time.ANSIC)
			fileInfo.Dir = path
			fileInfo.Extension = filepath.Ext(file.Name())
			result = append(result, fileInfo)
		}
	}
	return result
}

func (a *App) MakeDir(path string) {
	a.err = ""
	err := os.MkdirAll(path, 0755)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) MakeFile(path string) {
	a.err = ""
	a.WriteFile(path, "")
}

func (a *App) MoveEntries(from string, to string) {
	a.err = ""
	err := os.Rename(from, to)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) RenameEntry(from string, to string) {
	a.err = ""
	err := os.Rename(from, to)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) GetError() string {
	return a.err
}

func (a *App) CopyEntries(from string, to string) {
	a.err = ""
	info, err := os.Stat(from)
	if os.IsNotExist(err) {
		a.err = err.Error()
		return
	}
	if info.IsDir() {
		//
		// It's a directory! Do a deap copy.
		//
		err := cp.Copy(from, to)
		if err != nil {
			a.err = err.Error()
			return
		}
	} else {
		//
		// It's a file. Just copy it.
		//
		source, err := os.Open(from)
		if err != nil {
			a.err = err.Error()
			return
		}
		defer source.Close()

		destination, err := os.Create(to)
		if err != nil {
			a.err = err.Error()
			return
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)

		if err != nil {
			a.err = err.Error()
		}
	}
}

func (a *App) DeleteEntries(path string) {
	a.err = ""
	err := os.RemoveAll(path)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) RunCommandLine(cmd string, args []string, env []string, dir string) string {
	a.err = ""
	cmdline := exec.Command(cmd)
	cmdline.Args = args
	cmdline.Env = env
	cmdline.Dir = dir
	result, err := cmdline.CombinedOutput()
	if err != nil {
		a.err = err.Error()
	}

	return string(result[:])
}

func (a *App) GetClip() string {
	result, err := clip.ReadAll()
	if err != nil {
		a.err = err.Error()
	}
	return result
}

func (a *App) SetClip(msg string) {
	err := clip.WriteAll(msg)
	if err != nil {
		a.err = err.Error()
	}
}

func (a *App) GetEnvironment() []string {
	return os.Environ()
}

func (a *App) GetEnv(name string) string {
	return os.Getenv(name)
}

func (a *App) AppendPath(dir string, name string) string {
	return filepath.Join(dir, name)
}

func (a *App) GetOSName() string {
	os := goruntime.GOOS
	result := ""
	switch os {
	case "windows":
		result = "windows"
		break
	case "darwin":
		result = "macos"
	case "linux":
		result = "linux"
	default:
		result = fmt.Sprintf("%s", os)
	}
	return result
}
