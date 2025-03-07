package main

import (
    "log"
	"net/http"
)

func main(){
    const rootFilePath = "."
    const port = "8080"
    const endpoint = "/healthz"

    mux := http.NewServeMux()
    mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(rootFilePath))))
    mux.HandleFunc("/healthz", readinessHandler)
    server := &http.Server{
        Handler: mux,
        Addr: ":" + port,
    }

    log.Printf("Serving files from %s on port: %s\n", rootFilePath, port)
	log.Fatal(server.ListenAndServe())
}

func readinessHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(http.StatusText(http.StatusOK)))
}
