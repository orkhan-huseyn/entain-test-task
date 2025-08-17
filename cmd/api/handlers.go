package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/orkhan-huseyn/entain-test-task/internal/data"
	"github.com/orkhan-huseyn/entain-test-task/internal/validator"
)

func (app *application) getUserBalance(w http.ResponseWriter, r *http.Request) {
	userId, err := validator.ValidateUserId(r.PathValue("userId"))
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	user, err := app.models.Users.Get(userId)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var output struct {
		UserID  uint64 `json:"userId"`
		Balance string `json:"balance"`
	}

	output.UserID = user.ID
	output.Balance = fmt.Sprintf("%.2f", user.Balance)

	err = app.writeJSON(w, http.StatusOK, output, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	userId, err := validator.ValidateUserId(r.PathValue("userId"))
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	sourceType, err := validator.ValidateSourceType(r.Header.Get("Source-Type"))
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

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

	err = validator.ValidateTransactionState(input.State)
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	amount, err := validator.ValidateTransactionAmount(input.Amount)
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	err = validator.ValidateTransactionId(input.TransactionID)
	if err != nil {
		app.badRequestResponse(w, r, err.Error())
		return
	}

	transaction := &data.Transaction{
		TransactionID: input.TransactionID,
		Amount:        amount,
		State:         input.State,
		UserID:        userId,
		SourceType:    sourceType,
	}

	txn, err := app.db.Begin()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	user, err := app.models.Users.GetForUpdate(txn, userId)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	defer txn.Rollback()

	err = app.models.Transactions.Insert(txn, transaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	balanceIncrement := amount
	if input.State == "lose" {
		balanceIncrement *= -1
	}

	if user.Balance+balanceIncrement < 0 {
		app.badRequestResponse(w, r, "user's balance cannot be negative")
		return
	}

	err = app.models.Users.Update(txn, uint64(userId), balanceIncrement)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	txn.Commit()

	app.writeJSON(w, 200, transaction, nil)
}
