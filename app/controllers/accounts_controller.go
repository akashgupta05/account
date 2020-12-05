package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/akashgupta05/account/app/repository"
	log "github.com/sirupsen/logrus"
)

type AccountsController struct {
	accountRepo repository.AccountsRepositoryInterface
}

func NewAccountsController() *AccountsController {
	return &AccountsController{repository.NewAccountsRepository()}
}

type TransactionPayload struct {
	Activity string   `json:"activity"`
	Payload  *Payload `json:"payload"`
}

type Payload struct {
	UserID   string `json:"userId"`
	Amount   int64  `json:"amount"`
	Type     string `json:"type"`
	Priority int    `json:"priority"`
	Expiry   int64  `json:"expiry"`
}

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func readBodyBytes(r *http.Request) ([]byte, error) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn("Error reading request body", err.Error())
		return []byte(``), err
	}

	return bodyByte, nil
}

func respondWithError(rw http.ResponseWriter, statusCode int, err error) {
	responseBytes, err := json.Marshal(Response{Success: false, Error: err.Error()})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(statusCode)
	_, err = rw.Write(responseBytes)
	if err != nil {
		log.Errorf("failed to write response")
	}
}

func respondWithSuccess(rw http.ResponseWriter) {
	responseBytes, err := json.Marshal(Response{Success: true})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(responseBytes)
	if err != nil {
		log.Errorf("failed to write response")
	}
}
