/*========================== ====================================

	Author(s):
		Sugath Fernando

==============================================================================*/
package manager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// wait group initialize variable
var wg sync.WaitGroup

// timeout variable for the end points
var timeout = time.Duration(100 * time.Second)

type Response struct {
	Numbers []int `json:"Numbers"`
}

// get the from both the end points, combine the result, sort and out put.
func GetNumbersResult(param1 string, param2 string) []int {

	ch := make(chan []int)

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go GetEndPointNumbers(ch, param1)
	wg.Done()
	go GetEndPointNumbers(ch, param2)
	wg.Done()

	// append and remove the duplicates.
	result := Ints(append(<-ch, <-ch...))
	wg.Wait()

	return result
}

func GetEndPointNumbers(ch1 chan []int, param1 string) {

	// set the transport dial for the end point.
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Get("http://" + param1)

	defer resp.Body.Close()

	if err != nil {
		log.Println("ERROR executing end point request:")
	}

	// read the response body.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("ERROR reading response body")
	}

	var data Response
	errUnmarshall := json.Unmarshal(body, &data)
	if err != nil {
		log.Panic("ERROR unmarshall data ", errUnmarshall)
	}

	ch1 <- data.Numbers
}

// remove the duplicate values from slice
func Ints(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	// range for input slice
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}
