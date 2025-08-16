package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/orkhan-huseyn/entain-test-task/internal/data"
)

func (app *application) getUserBalance(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	user, err := app.models.Users.Get(uint64(userId))
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, user, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {

}
