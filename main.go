package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RequestAt1S struct {
	Nomenclature string  `json:"nomenclature"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Total        float64 `json:"total"`
}
type ResponseAt1S struct {
	Message string        `json:"message"`
	Data    []RequestAt1S `json:"data"`
}

type RequestFromRN struct {
	Nomenclature string  `json:"nomenclature"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Total        float64 `json:"total"`
}
type ResponseInRN struct {
	Message string          `json:"message"`
	Data    []RequestFromRN `json:"data"`
}

func accountingHandlerFromRN(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Принимаем данные от RN")
	var request RequestFromRN
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Данные от RN:", request)
	response := ResponseInRN{
		Message: "Success",
		Data:    []RequestFromRN{request},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func accountingHandlerAt1S(w http.ResponseWriter, r *http.Request) {
	response := ResponseAt1S{
		Message: "Success",
		Data:    []RequestAt1S{},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := chi.NewRouter()
	r.Post("/api/accounting", accountingHandlerFromRN)
	r.Post("/api/1s", accountingHandlerAt1S)
	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
