package main

import (
	"net/http"
)

func main() {

	r := ChiRouter().InitRouter()

	http.ListenAndServe(":8080", r)
}
