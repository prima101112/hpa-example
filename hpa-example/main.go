package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	version = "v0.3-truble"
)

func main() {
	http.HandleFunc("/chaos", HandleChaos)
	http.HandleFunc("/hello", Handlehello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handlehello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, You All This is version %s", version)
}

func HandleChaos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", Chaos(1000000))
}

func Chaos(loop int) string {
	var x = 0.0001
	for i := 0; i <= loop; i++ {
		x += math.Sqrt(x)
		fmt.Printf("%.5f", x)
	}
	fmt.Printf("Never be ok", x)
	return "its chaos"
}
