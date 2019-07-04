package main

import (
  "net/http"
  "fmt"
  "reflect"
  "runtime"
)

func main(){
  server := http.Server{
    Addr: "localhost:8080",
  }

  http.HandleFunc("/hello", log(hello))
  http.HandleFunc("/world", log(world))
  fmt.Println("Listening on 8080 over ssl...")
  server.ListenAndServeTLS("cert.pem", "key.pem")
}

func log(handler http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
    fmt.Println("Handler function called - " + name)
    handler(w, r)
  }
}

func hello(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  str := fmt.Sprintf("body: %s", r.Form["page"])
  fmt.Fprintln(w,str)
}

func world(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "world")
}
