package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type ResponseData struct {
	Name string
}

type Response struct {
	Success bool
	Data    ResponseData
	Error string
}

func Handler(w http.ResponseWriter, r http.Request) {
	var response Response

	if r.Method != "GET" {
		response.Success = false
		response.Error = "Method not allowed!!"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	var names []string = []string{
		"Ankan",
		"Aditya",
		"Dilpreet",
		"Pragati",
		"Manish",
		"Sweety",
	}

	rand.Seed(time.Now().UnixNano())
	generatedName := names[rand.Intn(len(names)-0)+0]

	response.Success = true
	response.Data.Name = generatedName

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}