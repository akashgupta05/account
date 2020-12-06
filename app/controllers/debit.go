package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/account/app/models"
)

const debit = "debit"

// Debit utilises the credits
func (ac *AccountsController) Debit(rw http.ResponseWriter, r *http.Request) {
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

	if err = validateDebitRequest(transactionPayload); err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	debit := &models.Debit{
		UserID:      transactionPayload.Payload.UserID,
		Amount:      transactionPayload.Payload.Amount,
		UsedCredits: transactionPayload.Payload.Credits,
	}

	err = ac.accountRepo.Debit(debit)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(rw, &Response{Success: true, Data: debit})
}

func validateDebitRequest(tp *TransactionPayload) error {
	if tp.Activity != debit {
		return errors.New("invalid activity")
	}

	if tp.Payload == nil {
		return errors.New("missing payload")
	}

	if tp.Payload.Amount <= 0 {
		return errors.New("invalid amount")
	}

	if tp.Payload.UserID == "" {
		return errors.New("invalid userId")
	}

	if tp.Payload.Credits == 0 {
		return errors.New("invalid credits")
	}

	return nil
}
