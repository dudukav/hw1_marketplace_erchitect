package internal

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    user := User{ID: "1", Name: "Данил Колбасенко"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}