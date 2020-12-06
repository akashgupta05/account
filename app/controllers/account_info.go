package controllers

import (
	"errors"
	"net/http"
)

var userId = "userId"

// AccountInfo returns account info of a user
func (ac *AccountsController) AccountInfo(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	userID := queryValues.Get(userId)
	if userID == "" {
		respondWithError(rw, http.StatusBadRequest, errors.New("invalid userId"))
		return
	}

	accountInfo, err := ac.accountRepo.AccountInfo(userID)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(rw, &Response{
		Success: true,
		Data:    accountInfo,
	})
}

// CreditActivity return credit activities of a user
func (ac *AccountsController) CreditActivity(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	userID := queryValues.Get(userId)
	if userID == "" {
		respondWithError(rw, http.StatusBadRequest, errors.New("invalid userId"))
		return
	}

	creditActivity, err := ac.accountRepo.CreditHistory(userID)
	if err != nil {
		respondWithError(rw, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(rw, &Response{
		Success: true,
		Data:    creditActivity,
	})
}
