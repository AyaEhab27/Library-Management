package main

import (
    "library/server"
    "library/cli"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "server" {
        go server.StartMainServer("8080")
        go server.StartBackupServer("8080","8081")
        
        select {}
    } else {
        var serverURL string
        if _, err := http.Get("http://localhost:8080"); err == nil {
            serverURL = "http://localhost:8080"
            println("Connected to main server")
        } else {
            serverURL = "http://localhost:8081"
            println("Connected to backup server")
        }
        
        cli.RunCLI(serverURL)

		
    }
}