package main

import (
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"
)

// TODO: Replace with the absolute import path

func main() {

	go start()
	println("t")
	println("t")
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
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

	println("ta")
	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":8181", nil)
}
