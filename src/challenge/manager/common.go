package manager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func doRequest(endpoint string) (interface{}, error) {
	resp, err := http.Get("http://" + endpoint)

	defer resp.Body.Close()

	if resp == nil {
		log.Println("ERROR executing end point request:")
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}

	errBody := json.Unmarshal([]byte(body[:]), &data)
	if err != nil {
		log.Panic("Internal issue: ", errBody)
	}
	return data["Numbers"], err
}
