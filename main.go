package main

import (
	"encoding/json"
	"log"
	"net/http"

	repo "./repository"
	"./route"
)

func main() {

	CassandraSession := repo.Session
	defer CassandraSession.Close()

	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":9080", router))
}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
