package backend

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

const (
	LOGIN_KEY = "login"
	EXPIRE_KEY = "expire"
)

func getAuthToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New(ERR_MISSING_AUTH_HEADER)
	}
	splittedHeader := strings.Split(authHeader, "Bearer ")
	if len(splittedHeader) < 2 {
		return "", errors.New(ERR_MISSING_AUTH_HEADER)
	}
	return splittedHeader[1], nil
}

func validateAccessToken(token string) (login string, expire int64, err error) {
	return validateToken(token, ERR_INVALID_ACCESS_TOKEN)
}

func validateRefreshToken(token string) (login string, expire int64, err error) {
	return validateToken(token, ERR_INVALID_REFRESH_TOKEN)
}

func validateToken(token, errorStr string) (login string, expire int64, err error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SIGN_KEY), nil
	})
	if parsedToken == nil {
		return "" , 0, errors.New(ERR_FAILED_PARSE_TOKEN)
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		login, ok := claims[LOGIN_KEY].(string)
		if !ok {
			return "", 0, errors.New(ERR_LOGIN_NOT_FOUND)
		}
		expire, ok := claims[EXPIRE_KEY].(float64)
		if !ok {
			return "", 0, errors.New(ERR_EXPIRE_NOT_FOUND)
		}
		return login, int64(expire),nil
	} else {
		return "", 0, errors.New(errorStr)
	}
}


func generateNewTokens(login string) (string, int64, string, int64, error) {
	accessTokenExpire := time.Now().Add(time.Millisecond * ACCESS_TOKEN_LIFE_MILLIS).Unix() * 1000
	refreshTokenExpire := time.Now().Add(time.Millisecond * REFRESH_TOKEN_LIFE_MILLIS).Unix() * 1000
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		LOGIN_KEY: login,
		EXPIRE_KEY: accessTokenExpire,
	})
	accessTokenString, err := accessToken.SignedString([]byte(SIGN_KEY))
	if err != nil {
		return "", 0, "", 0, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		LOGIN_KEY: login,
		EXPIRE_KEY: refreshTokenExpire,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(SIGN_KEY))
	if err != nil {
		return "", 0, "", 0, err
	}
	return accessTokenString, accessTokenExpire, refreshTokenString, refreshTokenExpire, nil
}