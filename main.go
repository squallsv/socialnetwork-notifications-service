package main

import (
	"encoding/json"
	"net/http"
	"github.com/squallsv/socialnetwork-notifications-service/DbConnection"
	"time"

	//"log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


func main() {
	CassandraSession := DbConnection.Session
	defer CassandraSession.Close()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", heartbeat)

	http.ListenAndServe(":3000", r)
}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

