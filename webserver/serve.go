package main

import (
  "net/http"
  "fmt"
)

type MyHandler struct{}

func main(){
  server := http.Server{
    Addr: "localhost:8080",
    Handler: &MyHandler{},
  }

  fmt.Println("Listening on 8080 over ssl...")
  server.ListenAndServeTLS("cert.pem", "key.pem")
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "hello World!")
}
