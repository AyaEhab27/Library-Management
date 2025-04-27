package storage

import (
    "encoding/json"
    "os"
	"strings"
    "path/filepath"
    "library/models"
)

type Storage struct {
    books   []models.Book
    readers []models.Reader
}

func NewStorage() Storage {
    return Storage{
        books:   make([]models.Book, 0),
        readers: make([]models.Reader, 0),
    }
}
// Load  Data
func (s *Storage) LoadData() error {
    booksData, err := os.ReadFile(filepath.Join("data", "books.json"))
    if err == nil {
        json.Unmarshal(booksData, &s.books)
    }

    readersData, err := os.ReadFile(filepath.Join("data", "readers.json"))
    if err == nil {
        json.Unmarshal(readersData, &s.readers)
    }

    return nil
}
// Save  Data
func (s *Storage) SaveData() error {
    // Save books data
    booksData, err := json.MarshalIndent(s.books, "", "  ")
    if err != nil {
        return err
    }
    os.WriteFile(filepath.Join("data", "books.json"), booksData, 0644)

    // Save readers data
    readersData, err := json.MarshalIndent(s.readers, "", "  ")
    if err != nil {
        return err
    }
    os.WriteFile(filepath.Join("data", "readers.json"), readersData, 0644)

    return nil
}

// Book operations
func (s *Storage) AddBook(book models.Book) {
    s.books = append(s.books, book)
	s.SaveData()
}

func (s *Storage) GetAllBooks() []models.Book {
    return s.books
}

func (s *Storage) SearchBooks(searchTerm string) []models.Book {
    var results []models.Book
    for _, book := range s.books {
        if strings.Contains(strings.ToLower(book.ID), strings.ToLower(searchTerm)) ||
            strings.Contains(strings.ToLower(book.Title), strings.ToLower(searchTerm)) {
            results = append(results, book)
        }
    }
    return results
}
func (s *Storage) SortBooksByTitle() []models.Book {
    sorted := make([]models.Book, len(s.books))
    copy(sorted, s.books)
    for i := 0; i < len(sorted)-1; i++ {
        for j := i + 1; j < len(sorted); j++ {
            if sorted[i].Title > sorted[j].Title {
                sorted[i], sorted[j] = sorted[j], sorted[i]
            }
        }
    }
    return sorted
}
func (s *Storage) SortBooksByDate() []models.Book {
    sorted := make([]models.Book, len(s.books))
    copy(sorted, s.books)
    for i := 0; i < len(sorted)-1; i++ {
        for j := i + 1; j < len(sorted); j++ {
            if sorted[i].PublicationDate > sorted[j].PublicationDate {
                sorted[i], sorted[j] = sorted[j], sorted[i]
            }
        }
    }
    return sorted
}

// Reader operations
func (s *Storage) AddReader(reader models.Reader) {
    s.readers = append(s.readers, reader)
	s.SaveData()
}

func (s *Storage) RemoveReader(id string) bool {
    for i, reader := range s.readers {
        if reader.ID == id {
            s.readers = append(s.readers[:i], s.readers[i+1:]...)
			s.SaveData()
            return true
        }
    }
    return false
}

func (s *Storage) GetAllReaders() []models.Reader {
    return s.readers
}

func (s *Storage) SearchReaders(searchTerm string) []models.Reader {
    var results []models.Reader
    for _, reader := range s.readers {
        if strings.Contains(strings.ToLower(reader.ID), strings.ToLower(searchTerm)) ||
            strings.Contains(strings.ToLower(reader.Name), strings.ToLower(searchTerm)) {
            results = append(results, reader)
        }
    }
    return results
}