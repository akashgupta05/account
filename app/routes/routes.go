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
		fmt.Fprint(w, "{ \"message\":\"Hello world!. I am Image Catalog.\",\"success\":true,\"api_version\": 1 }")
	})
	router.GET("/elb-check", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "{ \"message\":\"Hello world AWS ALB!. I am Image Catalog.\",\"success\":true,\"api_version\": 1 }")
	})
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(404)
		fmt.Fprint(rw, "{ \"message\":\"Not Found.\",\"success\":true,\"api_version\": 1 }")
	})

	accounts := controllers.NewAccountsController()
	router.POST("/credit", serveEndpoint(accounts.Credit))
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
				w.Write(respBuytes)
			}
		}()
		nextHandler(w, request)
	}
}
