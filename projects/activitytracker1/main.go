package main

import (
	"fmt"

	"github.com/slayer321/learn/projects/activitytracker1/internal/server"
)

func main() {
	fmt.Printf("Starting listintening on port 8080\n")
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
