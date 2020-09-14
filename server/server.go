package main

import (
	"net/http"

	"github.com/joshcarp/cloudfunctionsdemo"
)

func main() {
	http.HandleFunc("/", cloudfunctionsdemo.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
