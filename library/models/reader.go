package models

type Reader struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    Gender     string `json:"gender"`
    Birthday   string `json:"birthday"`
    Height     string `json:"height"`
    Weight     string `json:"weight"`
    Employment string `json:"employment"`
}