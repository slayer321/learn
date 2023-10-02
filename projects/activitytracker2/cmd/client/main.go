package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	client "github.com/slayer321/learn/projects/activitytracker1/pkg/server"
	"github.com/slayer321/learn/projects/activitytracker2/pkg/server"
)

var DefaultURL = "http://localhost:8080/"

func main() {

	add := flag.Bool("add", false, "Add Activity")
	get := flag.Bool("get", false, "Get Activity")

	flag.Parse()
	srv := &server.Actvities{URL: DefaultURL}

	//go webserver(&sy)
	switch {
	case *add:
		fmt.Println("Inside the add case")
		a := client.Activity{Time: time.Now(), Description: os.Args[2]}

		srv.Insert(a)
	case *get:
		fmt.Println("Inside the get case")
		id, _ := strconv.Atoi(os.Args[2])
		act, _ := srv.Retrieve(uint64(id))
		fmt.Printf("Activity: %v\n", act)

	default:
		fmt.Println("No flag specified")
		flag.Usage()
		os.Exit(1)
	}

}

func webserver(sy *sync.WaitGroup) {
	r := mux.NewRouter()

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logrus.Info("Starting webserver on ", DefaultURL)
	err := srv.ListenAndServe()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Exiting webserver\n")
	sy.Done()
}
