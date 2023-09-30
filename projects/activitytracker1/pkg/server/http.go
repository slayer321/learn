package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Activities *Activities
}

type IDDocument struct {
	ID int `json:"id"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

func NewHTTPServer(add string) *http.Server {
	server := httpServer{
		Activities: &Activities{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.handlerGet).Methods("GET")
	r.HandleFunc("/", server.handlerPost).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	return srv
}

func (s httpServer) handlerGet(w http.ResponseWriter, r *http.Request) {

	var req IDDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	activity, err := s.Activities.Retrieve(req.ID)
	if err == ErrIDNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := ActivityDocument{Activity: activity}
	json.NewEncoder(w).Encode(res)
}

func (s httpServer) handlerPost(w http.ResponseWriter, r *http.Request) {

	var req ActivityDocument

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := s.Activities.Insert(req.Activity)
	res := IDDocument{ID: id}
	json.NewEncoder(w).Encode(res)
}
