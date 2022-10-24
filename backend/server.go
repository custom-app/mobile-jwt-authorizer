package backend

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	port = 9990
)

func InitServer() {
	http.HandleFunc("/api/register", handleRegister)
	http.HandleFunc("/api/login", handleLogin)
	http.HandleFunc("/api/refresh_tokens", handleRefreshTokens)
	http.HandleFunc("/api/get_devs", handleGetDevs)
	fmt.Println("initiating db")
	if err := initDB(); err != nil {
		panic(err)
	}
	fmt.Println("starting server on " + strconv.Itoa(port) + " port")
	err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}