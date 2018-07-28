package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "time"
)
var lag time.Duration

func hitCpu(w http.ResponseWriter, r *http.Request) {
  wait := 9999999999
  for i:=0; i<wait; i++ {
    // CPU intensive wait
  }
  json.NewEncoder(w).Encode("{result: 200}")
}

// Endpoint that returns longer and longer wait times as you visit it.
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

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/lag", simulateLag).Methods("GET")
  router.HandleFunc("/lag/reset", resetLag).Methods("GET")
  router.HandleFunc("/cpu", hitCpu).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}
