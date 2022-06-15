package main

import (
	"fmt"
	"net/http"
)

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong")
}
