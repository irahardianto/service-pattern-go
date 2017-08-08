package main

import (
	"net/http"
)

func init() {

}

func main() {

	r := ChiMuxRouter().InitRouter()

	http.ListenAndServe(":8080", r)
}
