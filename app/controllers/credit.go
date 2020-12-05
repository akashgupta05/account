package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/account/app/models"
)

const credit = "credit"

// Credit transaction for credit
func (ac *AccountsController) Credit(rw http.ResponseWriter, r *http.Request) {
	bodyBytes, err := readBodyBytes(r)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	transactionPayload := &TransactionPayload{}
	err = json.Unmarshal(bodyBytes, transactionPayload)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateCreditRequest(transactionPayload); err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	transaction := &models.Transaction{
		UserID:   transactionPayload.Payload.UserID,
		Amount:   transactionPayload.Payload.Amount,
		Expiry:   transactionPayload.Payload.Expiry,
		Priority: transactionPayload.Payload.Priority,
		Type:     transactionPayload.Payload.Type,
	}

	err = ac.accountRepo.Credit(transaction)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(rw)
}

func validateCreditRequest(tp *TransactionPayload) error {
	if tp.Activity != credit {
		return errors.New("invalid activity")
	}

	if tp.Payload == nil {
		return errors.New("missing payload")
	}

	if tp.Payload.Amount <= 0 {
		return errors.New("invalid amount")
	}

	if tp.Payload.Expiry == 0 {
		return errors.New("invalid expiry")
	}

	if tp.Payload.Priority == 0 {
		return errors.New("invalid priority")
	}

	if tp.Payload.Type == "" {
		return errors.New("invalid type")
	}

	if tp.Payload.UserID == "" {
		return errors.New("invalid userId")
	}

	return nil
}
