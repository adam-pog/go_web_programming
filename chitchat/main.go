package main

import (
    "net/http"
    "html/template"
    "os"
    "log"
    "time"
    "fmt"
    "github.com/adam-pog/go_web_programming/chitchat/data"
)

func main() {
    server := &http.Server{
        Addr: "0.0.0.0:8081",
        Handler: handlerWithLog(),
    }
    // fmt.Println("Serving...")
    server.ListenAndServe()
}

func handlerWithLog() http.Handler {
    mux := http.NewServeMux()
    files := http.FileServer(http.Dir("public"))
    mux.Handle("/static/", http.StripPrefix("/static/", files))

    mux.HandleFunc("/", index)

    logger := log.New(os.Stdout, "http: ", log.LstdFlags)
    logger.Printf("Server is starting...")

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        mux.ServeHTTP(w, r)

        // log request by who(IP address)
        requesterIP := r.RemoteAddr

        log.Printf(
                "%s\t\t%s\t\t%s\t\t%v",
                r.Method,
                r.RequestURI,
                requesterIP,
                time.Since(start),
        )
    })
}

func index(w http.ResponseWriter, r *http.Request) {
    files := []string{
        "templates/layout.html",
        "templates/navbar.html",
        "templates/index.html",
    }

    templates := template.Must(template.ParseFiles(files...))
    // threads, err := data.Threads();

    cookie := http.Cookie{
        Name: "_cooki2e",
        Value: "UNIQUEVALUECharlie",
        HttpOnly: true,
    }
    // cookie must be set before body starts being written. Headers cannot be changed after body
    http.SetCookie(w, &cookie)

    threads, err := data.Threads(); if err == nil {
      templates.ExecuteTemplate(w, "layout", threads)
    } else {
      fmt.Println(err)
    }
}
