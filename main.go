package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "time"
  "sync"
)

// Used for increasingly slower lag
var lag time.Duration

// Used to simulate bottleneck
var mutex = &sync.Mutex{}

// Simulates high cpu utilization
func hitCpu(w http.ResponseWriter, r *http.Request) {
  wait := 9999999999
  for i:=0; i<wait; i++ {
    // CPU intensive wait
  }
  json.NewEncoder(w).Encode("{result: 200}")
}

// Simulates endpoint that becomes increasingly unresponsive as visited
func simulateLag(w http.ResponseWriter, r *http.Request) {
  wait := lag
  lag++
  time.Sleep(wait * time.Second)
  json.NewEncoder(w).Encode("{result: 200}")
}

// Resets lag to zero
func resetLag(w http.ResponseWriter, r *http.Request) {
  lag = 0
}

// Simulate bottleneck with locked resource
func lock(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  time.Sleep(2 * time.Second)
  mutex.Unlock()
  json.NewEncoder(w).Encode("{result: 200}")
}

// Simulates 1s wait
func wait(w http.ResponseWriter, r *http.Request) {
  time.Sleep(1 * time.Second)
  json.NewEncoder(w).Encode("{result: 200}")
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/cpu", hitCpu).Methods("GET")
  router.HandleFunc("/lag", simulateLag).Methods("GET")
  router.HandleFunc("/lag/reset", resetLag).Methods("GET")
  router.HandleFunc("/lock", lock).Methods("GET")
  router.HandleFunc("/1s-wait", wait).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}
