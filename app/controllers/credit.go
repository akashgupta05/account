package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/account/app/models"
	log "github.com/sirupsen/logrus"
)

const credit = "credit"

// Credit add credits for a user account
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

	credit := &models.Credit{
		UserID:          transactionPayload.Payload.UserID,
		CreditAmount:    transactionPayload.Payload.Amount,
		AvailableAmount: transactionPayload.Payload.Amount,
		Expiry:          transactionPayload.Payload.Expiry,
		Priority:        transactionPayload.Payload.Priority,
		Type:            transactionPayload.Payload.Type,
	}

	err = ac.accountRepo.Credit(credit)
	if err != nil {
		log.Warnf("Error while making credit transaction : %v", err)
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(rw, &Response{Success: true, Data: credit})
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
