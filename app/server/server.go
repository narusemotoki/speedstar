package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Part struct {
	Num      float64 `json:"num"`
	Operator string  `json:"operator"`
}

type Request struct {
	Parts []Part `json:"parts"`
}

type Result struct {
	Num float64 `json:"num"`
}

type Response struct {
	Result Result `json:"result"`
}

func calc(r Request) float64 {
	result := 1.0
	for _, part := range r.Parts {
		switch part.Operator {
		case "add":
			result += part.Num
		case "sub":
			result -= part.Num
		case "multi":
			result *= part.Num
		case "div":
			result /= part.Num
		case "mod":
			result = math.Mod(result, part.Num)
		}
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request Request
	decoder.Decode(&request)

	response := Response{Result{calc(request)}}
	j, _ := json.Marshal(response)
	fmt.Fprintf(w, string(j))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
