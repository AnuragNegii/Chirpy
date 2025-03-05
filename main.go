package main

import (
    "log"
	"net/http"
)

func main(){
    const rootFilePath = "."
    const port = "8080"

    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir(rootFilePath)))

    server := &http.Server{
        Handler: mux,
        Addr: ":" + port,
    }

    log.Printf("Serving files from %s on port: %s\n", rootFilePath, port)
	log.Fatal(server.ListenAndServe())
}
