package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akashgupta05/account/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "{ \"message\":\"Hello world!. I am Account Service.\",\"success\":true }")
	})
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(404)
		fmt.Fprint(rw, "{ \"message\":\"Not Found.\",\"success\":true }")
	})

	accounts := controllers.NewAccountsController()
	router.GET("/accounts", serveEndpoint(accounts.AccountInfo))
	router.POST("/accounts/credit", serveEndpoint(accounts.Credit))
	router.POST("/accounts/debit", serveEndpoint(accounts.Debit))
	router.GET("/accounts/credit_activity", serveEndpoint(accounts.CreditActivity))
}

func serveEndpoint(nextHandler func(rw http.ResponseWriter, r *http.Request)) httprouter.Handle {
	return func(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		defer func() {
			if recvr := recover(); recvr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := map[string]interface{}{
					"success": false,
					"error":   fmt.Sprintf("%v", recvr),
				}
				respBuytes, _ := json.Marshal(resp)
				setCommonHeaders(w)
				w.Write(respBuytes)
			}
		}()
		setCommonHeaders(w)
		nextHandler(w, request)
	}
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
}
