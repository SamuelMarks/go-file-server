package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

// Edited from: http://golang.org/pkg/net/http/#example_FileServer
func main() {
    if len(os.Args) > 3 || len(os.Args) < 2 {
        fmt.Printf("Usage: %s <folder_path> <port:8080>\n\tExample: %s /tmp 8090\n", os.Args[0], os.Args[0])
        os.Exit(1)
    }
    var port string
    if len(os.Args) > 2 {
        port = os.Args[2]
    } else {
        port = "8080"
    }
    log.Print(fmt.Sprintf("Serving %s on port %s", os.Args[1], port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), http.FileServer(http.Dir(os.Args[1]))))
}

