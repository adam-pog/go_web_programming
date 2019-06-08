package main

import (
  "net/http"
  "fmt"
)

func main(){
  server := http.Server{
    Addr: "localhost:8080",
    Handler: nil,
  }

  fmt.Println("Listening on 8080 over ssl...")
  server.ListenAndServeTLS("cert.pem", "key.pem")
}
