package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	api "github.com/slayer321/learn/projects/activitytracker1/pkg/server"
)

type Actvities struct {
	URL       string
	Actvities []api.Activity
}

func main() {
	r := mux.NewRouter()

	// server := Actvities{
	// 	Actvities: []api.Activity{},
	// }

	//r.HandleFunc("/").Methods(http.MethodGet)
	//r.HandleFunc("/", server.Insert).Methods(http.MethodPost)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logrus.Info("Starting webserver on ", srv.Addr)
	srv.ListenAndServe()
}

func (s Actvities) Insert(w http.ResponseWriter, r *http.Request) {
	var act api.ActivityDocument
	fmt.Printf("Inside the Insert function\n")
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v\n", err)
		return
	}

	err = json.Unmarshal(bodyBytes, &act)
	if err != nil {
		fmt.Fprintf(w, "Error unmarshalling request body: %v\n", err)
		return
	}
	act.Activity.ID = uint64(len(s.Actvities) + 1)
	s.Actvities = append(s.Actvities, act.Activity)

}
