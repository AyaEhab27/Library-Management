package server

import (
    "net"
    "time"
    "encoding/json"
    "net/http"
    "library/storage"
    "fmt"
	
)
//check main server is down or not
func StartBackupServer(mainPort, backupPort string) {
    for {
        if !isServerRunning(mainPort) {
            println("\033[33mMain server is down, starting backup on port", backupPort, "\033[0m")
            startBackupInstance(backupPort)
            return
        }
        time.Sleep(5 * time.Second)
    }
}

func isServerRunning(port string) bool {
    conn, err := net.Dial("tcp", "localhost:"+port)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}
// start backup server
func startBackupInstance(port string) {
    store := storage.NewStorage()
    store.LoadData()

    server := LibraryServer{storage: store}
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintf(w, `<h1 style="color:red">404 Page Not Found</h1>
                <p>Try these valid endpoints:</p>
                <ul>
                    <li><a href="/status">/status</a> - Check server status</li>
                    <li><a href="/books/all">/books/all</a> - List all books</li>
                    <li><a href="/readers/all">/readers/all</a> - List all readers</li>
                </ul>`)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `<h1 style="color:orange">Library Management System (Backup Server)</h1>
            <p>Backup server is running on port %s</p>
            <p>Available endpoints:</p>
            <ul>
                <li><a href="/status">/status</a></li>
                <li><a href="/books/all">/books/all</a></li>
                <li><a href="/readers/all">/readers/all</a></li>
            </ul>`, port)
    })

    mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        status := ServerStatus{
            Status:  "backup",
            Message: "Backup server is operational (Main server is down)",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(status)
    })

    mux.HandleFunc("/books/add", server.addBookHandler)
    mux.HandleFunc("/books/all", server.getAllBooksHandler)
    mux.HandleFunc("/books/search", server.searchBooksHandler)
    mux.HandleFunc("/books/save", server.saveBooksHandler)
    mux.HandleFunc("/books/sort", server.sortBooksHandler)

    mux.HandleFunc("/readers/add", server.addReaderHandler)
    mux.HandleFunc("/readers/remove", server.removeReaderHandler)
    mux.HandleFunc("/readers/all", server.getAllReadersHandler)
    mux.HandleFunc("/readers/search", server.searchReadersHandler)

    fmt.Println("\033[33mBackup server is fully operational on port", port, "\033[0m")
    fmt.Println("\033[33mAvailable at http://localhost:"+port, "\033[0m")
    err := http.ListenAndServe(":"+port, mux)
    if err != nil {
        fmt.Println("\033[31mBackup server failed:", err, "\033[0m")
    }
}