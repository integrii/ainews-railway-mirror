package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "os"
)

type post struct {
    Title string `json:"title"`
    Date  string `json:"date"`
}

var posts = []post{{Title: "AINews restored", Date: "May 10, 2026"}}

var homeTmpl = template.Must(template.New("home").Parse(`<!doctype html><title>AI News</title><body><h1>AI News Briefs</h1>{{range .}}<h2>{{.Title}}</h2><p>{{.Date}}</p>{{end}}<a href="/healthz">/healthz</a> / <a href="/api/posts">/api/posts</a></body>`))

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        _ = homeTmpl.Execute(w, posts)
    })
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })
    mux.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        _ = json.NewEncoder(w).Encode(posts)
    })

    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    log.Printf("listen :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}
