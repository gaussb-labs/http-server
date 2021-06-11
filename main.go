package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

func hiHandler(w http.ResponseWriter,r *http.Request) {

	httpRequestReceived.Inc()

	if getRandomNumber() % 7 == 0 {
		httpResponse500.Inc()
		fmt.Fprintf(w,"Something went wrong.")
	} else {
		httpResponse200.Inc()
		fmt.Fprintf(w, "hello")
	}
}

func getRandomNumber() int{
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source)
	return randomNumber.Intn(1000)
}

var (
	httpRequestReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_http_server_requests_total",
		Help: "Total number of Http Requests received",
	})
)

var (
	httpResponse500 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_http_server_response_500_total",
		Help: "Total number of 500 Http Responses",
	})
)

var (
	httpResponse200 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_http_server_response_200_total",
		Help: "Total number of 200 Http Responses",
	})
)

func main() {
	http.HandleFunc("/hi",hiHandler)
	http.Handle("/metrics",promhttp.Handler())
	http.ListenAndServe(":8080",nil)

}
