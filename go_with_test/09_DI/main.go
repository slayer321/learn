package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":12345", http.HandlerFunc(MyGreetHandler)))
}
