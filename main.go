package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Привет интернет!")
	})

	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.URL.Query().Get("message"))
	})

	http.HandleFunc("/calculate-area", func(writer http.ResponseWriter, request *http.Request) {
		radius, err := strconv.ParseFloat(request.URL.Query().Get("radius"), 64)
		if err != nil {
			panic(err)
		}
		area := math.Pi * math.Pow(radius, 2.0)
		fmt.Fprint(writer, fmt.Sprintf("Площадь окружности с радиусом %s равна %f", request.URL.Query().Get("radius"), area))
	})

	http.ListenAndServe(":80", nil)
}
