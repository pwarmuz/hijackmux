package main

import (
	"fmt"
	hijackmux "hijackmux/hijackmux" // uses the hijackmux module in the seperate folder, pretend this is an external api
	"io"
	"log"
	"net/http"
)

const port = "8000"

func defaultMux() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hijackmux.Exportable()
		// interesting side note: I'm using io.WriteString because it uses a similar implementation to how Sam recommended the writer check filenames
		// https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/io/io.go;l=310
		io.WriteString(w, "default mux\n")
	})
	fmt.Println("Listen and serving default mux:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func localMux() {
	// mux is local scope so global default is never used
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hijackmux.Exportable()
		io.WriteString(w, "hello\n")
	})
	fmt.Println("Listen and serving local mux:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func main() {
	// swap defaultMux and localMux comments to utilize
	// defaultMux represents default implementation which allows endpoint hijacking
	defaultMux()

	// localMux prevents endpoint hijacking because the global default mux is never used
	// localMux()
}
