package views

import "net/http"

func registerV1Handlers(r *httprouter.Router) {
	r.POST("/auth-token", v1AuthToken)
	r.GET("/account", v1GetAccount)
}

func v1AuthToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
