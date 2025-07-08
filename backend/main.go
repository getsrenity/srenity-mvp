// File: srenity-mvp/backend/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// MVPStatus represents the data sent to the frontend.
type MVPStatus struct {
	DetectedProblems []Problem `json:"detected_problems"`
	ActionsTaken     []Action  `json:"actions_taken"`
}

// Problem defines a detected issue.
type Problem struct {
	ID        string `json:"id"`
	Resource  string `json:"resource"`
	Issue     string `json:"issue"`
	Timestamp string `json:"timestamp"`
}

// Action defines a remediation step taken.
type Action struct {
	ID          string `json:"id"`
	ProblemID   string `json:"problem_id"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

// This function simulates detecting problems in a cluster.
func getSimulatedStatus() MVPStatus {
	problem := Problem{
		ID:        "p-123",
		Resource:  "pod/auth-service-7d5b",
		Issue:     "CrashLoopBackOff",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	action := Action{
		ID:          "a-456",
		ProblemID:   "p-123",
		Description: "Automated restart initiated.",
		Timestamp:   time.Now().Add(5 * time.Second).Format(time.RFC3339),
	}

	return MVPStatus{
		DetectedProblems: []Problem{problem},
		ActionsTaken:     []Action{action},
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := getSimulatedStatus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	r := mux.NewRouter()

	// API endpoint
	r.HandleFunc("/api/status", statusHandler).Methods("GET")

	// CORS handler to allow requests from the frontend
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) // For MVP; be more specific in production
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})

	log.Println("SREnity MVP Backend starting on port 8000")

	// Start server with CORS middleware
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
