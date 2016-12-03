package views

import (
	"encoding/json"
	"net/http"

	"github.com/hjkelly/discipline/services/account/controllers"
	"github.com/julienschmidt/httprouter"
)

func RegisterV1Handlers(r *httprouter.Router) {
	r.POST("/v1/auth-token", v1AuthToken)
	//r.GET("/v1/account", v1GetAccount)
}

func v1AuthToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	creds := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	token, err := controllers.GetAuthToken(creds.Email, creds.Password)
	if err != nil {
		// TODO: Implement
		w.WriteHeader(500)
		return
	}

	// Otherwise, write the token.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

//func v1GetAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
