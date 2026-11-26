package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Account struct {
	Name    string  `json:"name"`
	Number  int     `json:"number"`
	Balance int     `json:"balance"`
}

func main() {
	var accounts []Account
	http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Recebendo uma requisição GET")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)

	} else if r.Method == http.MethodPost {
		fmt.Println("Recebendo uma requisição POST")

		var account Account
		err := json.NewDecoder(r.Body).Decode(&account)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		accounts = append(accounts, account)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}

	})

	http.ListenAndServe(":8080", nil)
}