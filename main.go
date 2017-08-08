package main

import (
	"net/http"
)

func init() {

}

func main() {

	r := GorillaMuxRouter().InitRouter()

	http.ListenAndServe(":8080", r)
}
