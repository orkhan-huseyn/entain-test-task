package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/{userId}/transaction", app.createTransaction)
	mux.HandleFunc("GET /user/{userId}/balance", app.getUserBalance)

	return mux
}
