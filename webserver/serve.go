package main

import (
  "net/http"
  "fmt"
)

func main(){
  server := http.Server{
    Addr: "localhost:8080",
  }

  http.HandleFunc("/hello", hello)
  http.HandleFunc("/world", world)
  fmt.Println("Listening on 8080 over ssl...")
  server.ListenAndServeTLS("cert.pem", "key.pem")
}

func hello(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "world")
}
