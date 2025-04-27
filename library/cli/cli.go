package cli

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "library/models"
)
// Check Server Status
type ServerStatus struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}
func checkServerStatus(serverURL string) bool {
    resp, err := http.Get(serverURL + "/status")
    if err != nil {
        fmt.Println("\033[31mServer is not responding:", err, "\033[0m")
        return false
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        var status ServerStatus
        json.NewDecoder(resp.Body).Decode(&status)
        fmt.Println("\033[32m✓ Server is running:", status.Message, "\033[0m")
        return true
    }
    return false
}

// RunCLI 
func RunCLI(serverURL string) {
	if !checkServerStatus(serverURL) {
        fmt.Println("\033[31mCannot proceed without server connection\033[0m")
        return
    }
    reader := bufio.NewReader(os.Stdin)
    
    for {
        displayMainMenu()
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)
        
        switch choice {
        case "1":
            bookOperations(serverURL, reader)
        case "2":
            readerOperations(serverURL, reader)
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            
			fmt.Println("\033[31mInvalid choice, please try again\033[0m")
        }
    }
}

// main menu
func displayMainMenu() {
	fmt.Println("-----------------------------")
    fmt.Println("| Library Management System |")
    fmt.Println("|1. Book Options            |")
    fmt.Println("|2. Reader Options          |")
    fmt.Println("|3. Exit                    |")
    // fmt.Println("\033[3. Exit                |\033[0m")
    fmt.Println("|Choose an option:          |")
	fmt.Println("-----------------------------")
}

// 1- Book Operations Menu
func bookOperations(serverURL string, reader *bufio.Reader) {
    for {
		fmt.Println("********************************************")
        fmt.Println("\n Book Options")
        fmt.Println("1. add book ")
        fmt.Println("2. search for a book by ( ID , Name)  ")
        fmt.Println("3. load books info")
		fmt.Println("4. save books info")
		fmt.Println("5. sort books by( Title, Publication Date )")
        fmt.Println("6. back to main menu")
        fmt.Println("Choose an option: ")
		fmt.Println("********************************************")
        
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)
        
        switch choice {
		case "1":
            addBook(serverURL, reader)
        case "2":
            searchBooks(serverURL, reader)
        case "3":
            getAllBooks(serverURL)
        case "4":
            saveBooks(serverURL)
        case "5":
            sortBooks(serverURL, reader)
        case "6":
            return
        default:
            
			fmt.Println("\033[31mInvalid choice, please try again\033[0m")
        }
    }
}

// 2- Reader Operations Menu
func readerOperations(serverURL string, reader *bufio.Reader) {
    for {
		fmt.Println("######################################")
        fmt.Println("\n Reader Options")
        fmt.Println("1. add a reader")
        fmt.Println("2. remove a reader")
        fmt.Println("3. search for a reader by (ID , Name)")
        fmt.Println("4. get readers info")
        fmt.Println("5. back to main menu")
        fmt.Println("Choose an option: ")
		fmt.Println("######################################")
        
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)
        
        switch choice {
        case "1":
            addReader(serverURL, reader)
        case "2":
            removeReader(serverURL, reader)
        case "3":
            searchReaders(serverURL, reader)
        case "4":
            getAllReaders(serverURL)
        case "5":
            return
        default:
			
			fmt.Println("\033[31mInvalid choice, please try again\033[0m")
        }
    }
}

// 1- add book
func addBook(serverURL string, reader *bufio.Reader) {
    var book models.Book
    
    fmt.Print("Enter Book ID : ")
    book.ID, _ = reader.ReadString('\n')
    book.ID = strings.TrimSpace(book.ID)
    
    fmt.Print(" Enter Title: ")
    book.Title, _ = reader.ReadString('\n')
    book.Title = strings.TrimSpace(book.Title)
    
    fmt.Print(" Enter Publication Date (YYYY-MM-DD): ")
    book.PublicationDate, _ = reader.ReadString('\n')
    book.PublicationDate = strings.TrimSpace(book.PublicationDate)
    
    fmt.Print(" Enter Author : ")
    book.Author, _ = reader.ReadString('\n')
    book.Author = strings.TrimSpace(book.Author)
    
    fmt.Print(" Enter Genre: ")
    book.Genre, _ = reader.ReadString('\n')
    book.Genre = strings.TrimSpace(book.Genre)
    
    fmt.Print("Enter Publisher : ")
    book.Publisher, _ = reader.ReadString('\n')
    book.Publisher = strings.TrimSpace(book.Publisher)
    
    fmt.Print("Enter Language : ")
    book.Language, _ = reader.ReadString('\n')
    book.Language = strings.TrimSpace(book.Language)
    
    jsonData, err := json.Marshal(book)
    if err != nil {
        fmt.Println("\033[31mError creating JSON:", err, "\033[0m")
        return
    }

    resp, err := http.Post(serverURL+"/books/add", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("\033[31mConnection error:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println("\033[32mBook added successfully\033[0m")
    } else {
        body, _ := io.ReadAll(resp.Body)
        fmt.Println("\033[31mFailed to add book. Server response:", resp.Status, string(body), "\033[0m")
    }
}

// 2- search book
func searchBooks(serverURL string, reader *bufio.Reader) {
    fmt.Print("Enter search term (ID or Title): ")
    term, _ := reader.ReadString('\n')
    term = strings.TrimSpace(term)
    
    resp, err := http.Get(serverURL + "/books/search?q=" + term)
    if err != nil {
        fmt.Println("\033[31m Search error:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\033[31m Search failed (Status: %d)\033[0m\n", resp.StatusCode)
		return
	}

	var books []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		fmt.Println("\033[31m Invalid response format:", err, "\033[0m")
		return
	}

	fmt.Println("\n\033[32mSearch Results:\033[0m")
	for _, book := range books {
		fmt.Printf("\033[33mID:\033[0m %s \033[33mTitle:\033[0m %s \033[33mAuthor:\033[0m %s\n", 
			book.ID, book.Title, book.Author)
	}
}

// 3-  load all books 
func getAllBooks(serverURL string) {
    resp, err := http.Get(serverURL + "/books/all")
    if err != nil {
		fmt.Println("\033[31m Error getting books:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()
    
    var books []models.Book
    json.NewDecoder(resp.Body).Decode(&books)
    
	fmt.Println("\n\033[32mAll Books:\033[0m")

    for _, book := range books {
		fmt.Printf("\033[33mID:\033[0m %s \033[33mTitle:\033[0m %s \033[33mAuthor:\033[0m %s\n", 
			book.ID, book.Title, book.Author)
    }
}
//4- save books info
func saveBooks(serverURL string) {
    resp, err := http.Post(serverURL+"/books/save", "application/json", nil)
    if err != nil {
		fmt.Println("\033[31m Error saving books:", err, "\033[0m")

        return
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == http.StatusOK {
		fmt.Println("\033[32mBooks saved successfully\033[0m")

    } else {
		fmt.Println("\033[31mFailed to save books\033[0m")
    }
}
//5- sort books
func sortBooks(serverURL string, reader *bufio.Reader) {
    fmt.Println("\nSort by:")
    fmt.Println("1. Title")
    fmt.Println("2. Publication Date")
    fmt.Print("Choose sorting option: ")
    
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)
    
    var sortBy string
    switch choice {
    case "1":
        sortBy = "title"
    case "2":
        sortBy = "date"
    default:
		fmt.Println("\033[31mInvalid choice\033[0m")
        return
    }
    
    resp, err := http.Get(serverURL + "/books/sort?by=" + sortBy)
    if err != nil {
		fmt.Println("\033[31mError sorting books\033[0m")
        return
    }
    defer resp.Body.Close()
    
    var books []models.Book
    json.NewDecoder(resp.Body).Decode(&books)
    
	fmt.Println("\n\033[32mSorted Books:\033[0m")
    for _, book := range books {
		fmt.Printf("\033[33mTitle:\033[0m %s \033[33mDate:\033[0m %s\n", 
		book.Title, book.PublicationDate)
    }
}

// 1- add reader
func addReader(serverURL string, reader *bufio.Reader) {
    var r models.Reader
    
    fmt.Print("Enter Reader ID: ")
    r.ID, _ = reader.ReadString('\n')
    r.ID = strings.TrimSpace(r.ID)
    
    fmt.Print("Enter Name: ")
    r.Name, _ = reader.ReadString('\n')
    r.Name = strings.TrimSpace(r.Name)
    
    fmt.Print("Enter Gender: ")
    r.Gender, _ = reader.ReadString('\n')
    r.Gender = strings.TrimSpace(r.Gender)
    
    fmt.Print("Enter Birthday (YYYY-MM-DD): ")
    r.Birthday, _ = reader.ReadString('\n')
    r.Birthday = strings.TrimSpace(r.Birthday)
    
    fmt.Print("Enter Height: ")
    r.Height, _ = reader.ReadString('\n')
    r.Height = strings.TrimSpace(r.Height)
    
    fmt.Print("Enter Weight: ")
    r.Weight, _ = reader.ReadString('\n')
    r.Weight = strings.TrimSpace(r.Weight)
    
    fmt.Print("Enter Employment: ")
    r.Employment, _ = reader.ReadString('\n')
    r.Employment = strings.TrimSpace(r.Employment)
    
    jsonData, _ := json.Marshal(r)
    resp, err := http.Post(serverURL+"/readers/add", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
		fmt.Println("\033[31mError adding reader:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == http.StatusOK {
		fmt.Println("\033[32mReader added successfully:\033[0m")
    } else {
		fmt.Println("\033[31mFailed to add reader:\033[0m")
    }
}

//2- remove reader
func removeReader(serverURL string, reader *bufio.Reader) {
    fmt.Print("Enter Reader ID : ")
    id, _ := reader.ReadString('\n')
    id = strings.TrimSpace(id)
    
    req, _ := http.NewRequest("DELETE", serverURL+"/readers/remove?id="+id, nil)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Println("\033[31mError removing reader:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == http.StatusOK {
		fmt.Println("\033[32mReader removed successfully:\033[0m")
    } else {
		fmt.Println("\033[31mFailed to remove reader (ID not found)\033[0m")
    }
}

// 3- search reader
func searchReaders(serverURL string, reader *bufio.Reader) {
    fmt.Print("Enter search term (ID or Name): ")
    term, _ := reader.ReadString('\n')
    term = strings.TrimSpace(term)
    
    resp, err := http.Get(serverURL + "/readers/search?q=" + term)
    if err != nil {
		fmt.Println("\033[31mError searching readers:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()
    
    var readers []models.Reader
    json.NewDecoder(resp.Body).Decode(&readers)
    
	fmt.Println("\n\033[32mSearch Results:\033[0m")
    for _, reader := range readers {
		fmt.Printf("\033[33mID:\033[0m %s \033[33mName:\033[0m %s \033[33mGender:\033[0m %s\n", 
		reader.ID, reader.Name, reader.Gender)
    }
}

// 4- get readers info
func getAllReaders(serverURL string) {
    resp, err := http.Get(serverURL + "/readers/all")
    if err != nil {
		fmt.Println("\033[31mError getting readers:", err, "\033[0m")
        return
    }
    defer resp.Body.Close()
    
    var readers []models.Reader
    json.NewDecoder(resp.Body).Decode(&readers)
    
	fmt.Println("\n\033[32mAll Readers:\033[0m")
    for _, reader := range readers {
        fmt.Printf("\033[33mID:\033[0m %s \033[33mName:\033[0m %s \033[33mGender:\033[0m %s\n", 
		reader.ID, reader.Name, reader.Gender)
    }
}