package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "time"
)
var lag time.Duration

func cpu(w http.ResponseWriter, r *http.Request) {
  wait := 9999999999
  for i:=0; i<wait; i++ {
    // CPU intensive wait
  }
  json.NewEncoder(w).Encode("{result: 200}")
}

func sleep(w http.ResponseWriter, r *http.Request) {
  wait := lag
  lag++
  time.Sleep(wait * time.Second)
  json.NewEncoder(w).Encode("{result: 200}")
}

func main() {

  router := mux.NewRouter()
  router.HandleFunc("/sleep", sleep).Methods("GET")
  router.HandleFunc("/cpu", cpu).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}
