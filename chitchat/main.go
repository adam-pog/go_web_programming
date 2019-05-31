package main

import (
    "net/http"
    "html/template"
    "os"
    "log"
    // "fmt"
)

func main() {
    mux := http.NewServeMux()
    // files := http.FileServer(http.Dir("public"))
    // mux.Handle("/static", http.StripPrefix("/static/", files))

    mux.HandleFunc("/", index)
    logger := log.New(os.Stdout, "http: ", log.LstdFlags)
    logger.Printf("Server is starting...")

    server := &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: mux,
        ErrorLog: logger,
    }
    // fmt.Println("Serving...")
    server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
    files := []string{
        "templates/layout.html",
        "templates/navbar.html",
        "templates/index.html",
    }

    templates := template.Must(template.ParseFiles(files...))
    // threads, err := data.Threads();

    // if err == nil {
    templates.ExecuteTemplate(w, "layout", "World!")
    // }
}
