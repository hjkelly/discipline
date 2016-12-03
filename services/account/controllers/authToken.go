package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hjkelly/discipline/config"
)

// Turn an email and password into a JWT using Auth0's database. Return as
// helpful an error as possible.
func GetAuthToken(email, password string) (string, error) {
	// TODO: Error handling

	c := config.GetConfig()
	requestData := map[string]interface{}{
		"client_id":  c.Auth0ClientID,
		"username":   email,
		"password":   password,
		"connection": "Username-Password-Authentication",
		"grant_type": "password",
		"scope":      "openid",
	}
	// Turn the data into JSON or give up.
	requestBodyStr, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}
	// Send a request that will ask Auth0 for this user's token.
	resp, err := http.Post(
		"https://passthellama.auth0.com/oauth/ro",
		"application/json",
		bytes.NewReader(requestBodyStr),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Success, grab the JWT!
	tokenBody := &struct {
		Token string `json:"id_token"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(tokenBody)
	if err != nil {
		return "", err
	}

	return tokenBody.Token, nil
}
