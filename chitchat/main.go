package main

import (
    "net/http"
    "html/template"
    "os"
    "log"
    "time"
)

func main() {
    // files := http.FileServer(http.Dir("public"))
    // mux.Handle("/static", http.StripPrefix("/static/", files))

    server := &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: handlerWithLog(),
    }
    // fmt.Println("Serving...")
    server.ListenAndServe()
}

func handlerWithLog() http.Handler {
    mux := http.NewServeMux()
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

    // if err == nil {
    templates.ExecuteTemplate(w, "layout", "World!")
    // }
}
