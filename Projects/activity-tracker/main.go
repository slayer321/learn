package main

import (
	"activitytracker/part1/internal/server"
	"fmt"
)

func main() {
	fmt.Printf("Starting listintening on port 8080\n")
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
