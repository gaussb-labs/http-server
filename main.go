package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var requestCounter = 0
var response200Counter = 0
var response500Counter = 0

type Metrics struct {
	TotalHttpRequests int `json:"total_http_requests,omitempty"`
	HttpResponse200   int `json:"http_response_200,omitempty"`
	HttpResponse500   int `json:"http_response_500,omitempty"`
}

func hiHandler(w http.ResponseWriter,r *http.Request) {

	requestCounter +=1
	if requestCounter%5 == 0 {
		response500Counter+=1
		fmt.Fprintf(w,"Something went wrong.")
	} else {
		response200Counter+=1
		fmt.Fprintf(w, "hello")
	}

}

func metricsHandler(w http.ResponseWriter, r *http.Request) {

	metrics := &Metrics{TotalHttpRequests: requestCounter, HttpResponse200: response200Counter, HttpResponse500: response500Counter}

	b,e := json.Marshal(metrics)
	if e != nil {
		fmt.Fprintf(w,"%s",e)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func main() {
	http.HandleFunc("/hi",hiHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":8080",nil)

}
