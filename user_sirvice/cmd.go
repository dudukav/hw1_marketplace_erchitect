package main

import (
	"net/http"

	"gitlab.education.tbank.ru/v.kashurkina-37455/user_service/internal"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    http.HandleFunc("/users", internal.GetUser)
    http.ListenAndServe(":8080", nil)
}
k
rrr
