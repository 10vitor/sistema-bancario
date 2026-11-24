package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Recebendo uma requisição GET")
		w.Write([]byte("Bem-Vindo ao sistema bancario!"))
	} else {

	}
	
	})

	http.ListenAndServe(":8080", nil)
}