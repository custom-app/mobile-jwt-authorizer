package backend

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendError(code int, msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var e errorResponse
	e.Code = code
	e.Message = msg
	send, _ := json.Marshal(e)
	writeBytes(w, &send, code)
}

func writeBytes(w http.ResponseWriter, bytes *[]byte, code int) {
	w.WriteHeader(code)
	_, err := w.Write(*bytes)
	if err != nil {
		log.Println("write err: ", err)
	}
}