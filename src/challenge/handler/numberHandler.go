/*======================Travel Audience Number Server===========================

	Author(s):
		Sugath Fernando

==============================================================================*/
package handler

import (
	"challenge/manager"
	"challenge/response"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"
)

// formats the numbers from different end points, and out put as a JSON.
func NumberFormat(res http.ResponseWriter, req *http.Request) {

	start := time.Now()

	// get the url  parameters from url
	param1, param2 := GetUrlParams(req.URL.RawQuery)

	err := CheckUrlErrors(param1, param2)

	if err != nil {
		response.Error(res, "Url is not valid")
		return
	}

	urlParams := manager.GetNumbersResult(param1, param2)

	sort.Ints(urlParams)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(map[string]interface{}{"Numbers": urlParams})
	elapsed := time.Now()

	log.Printf("Result took %s", elapsed.Sub(start))
}

// get url parameters from  url.
func GetUrlParams(requestUrl string) (string, string) {
	urlRawQuery, _ := url.ParseQuery(requestUrl)
	urlParam1 := urlRawQuery["u"][0]
	urlParam2 := urlRawQuery["u"][1]

	return urlParam1, urlParam2
}

// check url for errors
func CheckUrlErrors(param1 string, param2 string) error {
	_, err := url.ParseRequestURI(param1)
	url.ParseRequestURI(param2)

	return err
}
