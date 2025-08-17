package validator

import (
	"errors"
	"slices"
	"strconv"
)

var (
	validSourceTypes       = []string{"game", "server", "payment"}
	validTransactionStates = []string{"win", "lose"}
)

func ValidateUserId(userId string) (uint64, error) {
	parsedUserId, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}

	if parsedUserId < 1 {
		return 0, errors.New("userId cannot be less than 1")
	}

	return uint64(parsedUserId), nil
}

func ValidateSourceType(sourceType string) (string, error) {
	if slices.Contains(validSourceTypes, sourceType) {
		return sourceType, nil
	}

	return "", errors.New("source type is not valid")
}

func ValidateTransactionAmount(amount string) (float64, error) {
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0.0, err
	}

	if parsedAmount < 0 {
		return 0.0, errors.New("transaction amount cannot be negative number")
	}

	return parsedAmount, nil
}

func ValidateTransactionState(transactionState string) error {
	if slices.Contains(validTransactionStates, transactionState) {
		return nil
	}

	return errors.New("transaction state is not valid")
}

func ValidateTransactionId(transactionID string) error {
	if transactionID == "" {
		return errors.New("transaction id cannot be empty")
	}

	return nil
}
