package models

type Book struct {
    ID             string `json:"id"`
    Title          string `json:"title"`
    PublicationDate string `json:"publication_date"`
    Author         string `json:"author"`
    Genre          string `json:"genre"`
    Publisher      string `json:"publisher"`
    Language       string `json:"language"`
}