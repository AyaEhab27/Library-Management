package server

import (
    "encoding/json"
    "net/http"
    //"sort"
    "library/storage"
    "library/models"
	"fmt"
)

type LibraryServer struct {
    storage storage.Storage
}
// to return Status for Server
type ServerStatus struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func StartMainServer(port string) {
    store := storage.NewStorage()
    //load  data from storage
    store.LoadData()

    server := LibraryServer{storage: store}
	mux := http.NewServeMux()
    //main page
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
		fmt.Fprintf(w, `<h1 style="color:green">Library Management System</h1>
			<p>Server is running on port %s</p>
			<p>Available endpoints:</p>
			<ul>
            <li><a href="/status">/status</a></li>
            <li><a href="/books/all">/books/all</a></li>
            <li><a href="/readers/all">/readers/all</a></li>
            </ul>`, port)
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        status := ServerStatus{
            Status:  "running",
            Message: "Server is operational",
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

	fmt.Println("\033[32mServer is running on port", port, "\033[0m")
	fmt.Println("\033[32mAvailable at http://localhost:"+port, "\033[0m")
	err := http.ListenAndServe(":"+port, mux)
    if err != nil {
        fmt.Println("\033[31mServer failed to start:", err, "\033[0m")
    }
}

func (s *LibraryServer) addBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    var book models.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    s.storage.AddBook(book)
    w.WriteHeader(http.StatusOK)
}

func (s *LibraryServer) getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
    books := s.storage.GetAllBooks()
    json.NewEncoder(w).Encode(books)
}

func (s *LibraryServer) searchBooksHandler(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("q")
    books := s.storage.SearchBooks(searchTerm)
    json.NewEncoder(w).Encode(books)
}

func (s *LibraryServer) saveBooksHandler(w http.ResponseWriter, r *http.Request) {
    err := s.storage.SaveData()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (s *LibraryServer) sortBooksHandler(w http.ResponseWriter, r *http.Request) {
    sortBy := r.URL.Query().Get("by")
    var books []models.Book
    
    switch sortBy {
    case "title":
        books = s.storage.SortBooksByTitle()
    case "date":
        books = s.storage.SortBooksByDate()
    default:
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    json.NewEncoder(w).Encode(books)
}

func (s *LibraryServer) addReaderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    var reader models.Reader
    err := json.NewDecoder(r.Body).Decode(&reader)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    s.storage.AddReader(reader)
    w.WriteHeader(http.StatusOK)
}

func (s *LibraryServer) removeReaderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "DELETE" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    id := r.URL.Query().Get("id")
    success := s.storage.RemoveReader(id)
    if success {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

func (s *LibraryServer) getAllReadersHandler(w http.ResponseWriter, r *http.Request) {
    readers := s.storage.GetAllReaders()
    json.NewEncoder(w).Encode(readers)
}

func (s *LibraryServer) searchReadersHandler(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("q")
    readers := s.storage.SearchReaders(searchTerm)
    json.NewEncoder(w).Encode(readers)
}