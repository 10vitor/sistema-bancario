package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

type Account struct {
	Name    string  `json:"name"`
	Number  int     `json:"number"`
	Balance int     `json:"balance"`
}

type Deposit struct {
	Amount int `json:"amount"`
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

	http.HandleFunc("GET /accounts/{number}", func(w http.ResponseWriter, r *http.Request) {
		numero, err := strconv.Atoi(r.PathValue("number"))

	if err != nil {
    	http.Error(w, "Número da conta inválido", http.StatusBadRequest)
    	return
}
	for _, account := range accounts {
		if account.Number == numero {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(account)
			return
		}
	}

	http.Error(w, "Conta não encontrada", http.StatusNotFound)

})

	http.HandleFunc("POST /accounts/{number}/deposit", func(w http.ResponseWriter, r *http.Request) {
		numero := r.PathValue("number")
})
	http.ListenAndServe(":8080", nil)
}