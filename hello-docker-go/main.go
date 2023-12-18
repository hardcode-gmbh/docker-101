package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/fortune", FortuneServer)
	http.HandleFunc("/health", Health)
	http.ListenAndServe(":1234", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name, ok := os.LookupEnv("DEFAULT_NAME")
	if !ok {
		panic("DEFAULT_NAME is not set")
	}
	if len(r.URL.Query().Get("name")) > 0 {
		name = r.URL.Query().Get("name")
	}
	log.Println("/hello was requested")
	fmt.Fprintf(w, "Hallo, %s!", name)
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("/health was requested")
	fmt.Fprintf(w, "UP!")
}

func FortuneServer(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/bin/fortune")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s\n", stdout.String())
	log.Println("/fortune was requested")
}
