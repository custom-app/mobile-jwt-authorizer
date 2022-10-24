package backend

import (
	"encoding/json"
	"net/http"
	"time"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(405, "", w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var registerBody authBody
	err := decoder.Decode(&registerBody)
	if err != nil {
		sendError(400, err.Error(), w)
		return
	}
	if len(registerBody.Password) < 4 {
		sendError(400, ERR_INVALID_PASSWORD, w)
		return
	}
	user, err := registerUser(registerBody.Login, registerBody.Password)
	if err != nil {
		sendError(400, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	send, err := json.Marshal(user)
	if err != nil {
		sendError(500, "error marshalling: " + err.Error(), w)
		return
	}
	writeBytes(w, &send, 200)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(405, "", w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var authBody authBody
	err := decoder.Decode(&authBody)
	if err != nil {
		sendError(400, err.Error(), w)
		return
	}
	user, err := getUser(authBody.Login)
	if err != nil || user == nil || user.Password != authBody.Password {
		sendError(401, ERR_WRONG_CREDENTIALS, w)
		return
	}
	token, tokenExpire, refreshToken, refreshTokenExpire, err := generateNewTokens(user.Login)
	err = updateUserTokens(user.Id, token, tokenExpire, refreshToken, refreshTokenExpire)
	if err != nil {
		sendError(500, err.Error(), w)
		return
	}
	user.AccessToken = token
	user.AccessTokenExpire = tokenExpire
	user.RefreshToken = refreshToken
	user.RefreshTokenExpire = refreshTokenExpire
	send, err := json.Marshal(user)
	if err != nil {
		sendError(500, "error marshalling: " + err.Error(), w)
		return
	}
	writeBytes(w, &send, 200)
}

func handleRefreshTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(405, "", w)
		return
	}
	refreshToken, err := getAuthToken(r)
	if err != nil {
		sendError(401, err.Error(), w)
		return
	}
	login, expire, err := validateRefreshToken(refreshToken)
	if err != nil {
		sendError(401, err.Error(), w)
		return
	}
	user, err := getUser(login)
	if err != nil || user == nil || user.RefreshToken != refreshToken {
		sendError(401, ERR_INVALID_REFRESH_TOKEN, w)
		return
	}
	if time.Now().Unix() * 1000 > expire {
		sendError(401, ERR_REFRESH_TOKEN_EXPIRED, w)
		return
	}
	token, tokenExpire, refreshToken, refreshTokenExpire, err := generateNewTokens(user.Login)
	err = updateUserTokens(user.Id, token, tokenExpire, refreshToken, refreshTokenExpire)
	if err != nil {
		sendError(500, err.Error(), w)
		return
	}
	user.AccessToken = token
	user.AccessTokenExpire = tokenExpire
	user.RefreshToken = refreshToken
	user.RefreshTokenExpire = refreshTokenExpire
	send, err := json.Marshal(user)
	if err != nil {
		sendError(500, "error marshalling: " + err.Error(), w)
		return
	}
	writeBytes(w, &send, 200)
}

func handleGetDevs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(405, "", w)
		return
	}
	authToken, err := getAuthToken(r)
	if err != nil {
		sendError(401, err.Error(), w)
		return
	}
	login, expire, err := validateAccessToken(authToken)
	if err != nil {
		sendError(401, err.Error(), w)
		return
	}
	if time.Now().Unix() * 1000 > expire {
		sendError(401, ERR_ACCESS_TOKEN_EXPIRED, w)
		return
	}
	user, err := getUser(login)
	if err != nil || user == nil || user.AccessToken != authToken {
		sendError(401, ERR_INVALID_ACCESS_TOKEN, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	send, err := json.Marshal(DEVS)
	if err != nil {
		sendError(500, "error marshalling: " + err.Error(), w)
		return
	}
	writeBytes(w, &send, 200)
}