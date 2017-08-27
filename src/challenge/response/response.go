/*==========================Travel Audience Number Server====================================

	Author(s):
		Sugath Fernando

==============================================================================*/
package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	successMessage = "OK"
)

func Error(rw http.ResponseWriter, message string) {
	r := Response{Error: true, Message: message, Data: nil}
	resp, err := json.Marshal(r)
	if err != nil {
		fmt.Fprint(rw, "500: Critical Server Error")
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(400)
	fmt.Fprint(rw, string(resp))
	return
}
