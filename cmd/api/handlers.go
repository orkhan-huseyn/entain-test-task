package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/orkhan-huseyn/entain-test-task/internal/data"
)

func (app *application) getUserBalance(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		app.notFoundResponse(w, r)
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
	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	sourceType := r.Header.Get("Source-Type")

	var input struct {
		State         string `json:"state"`
		Amount        string `json:"amount"`
		TransactionID string `json:"transactionId"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	amount, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// TODO: validate userId, sourceType, state, amount and transactionId
	transaction := &data.Transaction{
		TransactionID: input.TransactionID,
		Amount:        amount,
		State:         input.State,
		UserID:        int64(userId),
		SourceType:    sourceType,
	}

	// TODO: wrap transaction insert and user balance update into DB TRANSACTION
	err = app.models.Transactions.Insert(transaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	balanceIncrement := amount
	if input.State == "lose" {
		balanceIncrement *= -1
	}

	// TODO: validate user balance to it doesn't go negative
	err = app.models.Users.Update(uint64(userId), balanceIncrement)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, 200, transaction, nil)
}
