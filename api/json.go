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

func Handler(w http.ResponseWriter, r *http.Request) {
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

	// Optional JSON Syntax

	// jsonResp, err := json.Marshal(response)

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(400)

	// 	errorRes := make(map[string]string)

	// 	errorRes["sucess"] = "false"
	// 	errorRes["error"] = "Encoding Problem"

	// 	jsonErrorResp, _ := json.Marshal(errorRes)

	// 	w.Write(jsonErrorResp)

	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(200)
	// w.Write(jsonResp)
}
