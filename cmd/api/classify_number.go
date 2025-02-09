package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

type NumberResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

func (app *application) classifyNumber(w http.ResponseWriter, r *http.Request) {
	// Get the number from query parameters
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Number: "", Error: true})
		return
	}

	// Validate that the input is an integer
	number, convErr := strconv.Atoi(numberStr)
	if convErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberStr, Error: true})
		return
	}

	// Fetch fun fact from Numbers API
	numbersAPIURL := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	resp, apiErr := http.Get(numbersAPIURL)
	if apiErr != nil {
		http.Error(w, "Failed to fetch number trivia", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	funFactBytes, _ := io.ReadAll(resp.Body)
	funFact := strings.TrimSpace(string(funFactBytes))

	// Determine properties
	properties := []string{}
	if IsArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	// Prepare response
	response := NumberResponse{
		Number:     number,
		IsPrime:    IsPrime(number),
		IsPerfect:  IsPerfect(number),
		Properties: properties,
		DigitSum:   DigitSum(number),
		FunFact:    funFact,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
