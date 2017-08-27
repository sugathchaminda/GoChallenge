/*======================Travel Audience Number Server===========================

	Author(s):
		Sugath Fernando

==============================================================================*/
package main

import (
	"challenge/config"
	"challenge/handler"
	"log"
	"net/http"
)

// main function that servers the end point and runs in the port
func main() {
	http.HandleFunc("/numbers", handler.NumberFormat)
	log.Fatal(http.ListenAndServe(":"+config.App.Server.Port, nil))
}
