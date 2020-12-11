package main

import (
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"

    "fmt"
    "os"
)

// TODO: Replace with the absolute import path
var test int
func main() {

	go start()
	println("t")
	println("t")
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.Bind("goCode", handle);
	w.Bind("stuff", test);
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8181/index.html")
	w.Run()
	println("m")
}

func start() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":8181", nil)
}

func handle(myData string) string{
        f, err := os.Create("test.txt")
        if err != nil {
                fmt.Println(err)
                return "err"
        }
        l, err := f.WriteString("Hello World")
        if err != nil {
                fmt.Println(err)
                f.Close()
                return "done"
        }
        println(l)
        return "bad?"
}
