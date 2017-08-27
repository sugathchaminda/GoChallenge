package handler

import (
	"testing"
)

// Tests url parameters are correct
func TestUrlParams(t *testing.T) {
	requesturl := "u=localhost:8090/primes&u=localhost:8090/fibo"
	param1, param2 := GetUrlParams(requesturl)

	expectedParam1 := "localhost:8090/primes"
	expectedParam2 := "localhost:8090/fibo"

	if param1 != expectedParam1 || param2 != expectedParam2 {
		t.Errorf("Result url params are incorrect, got: %d %d, want: %d %d.", param1, param2, expectedParam1, expectedParam2)
	}
}

// check the url for errors
func TestUrlErrors(t *testing.T) {
	param1 := "localhost:8090/primes"
	param2 := "localhost:8090/fibo"
	err := CheckUrlErrors(param1, param2)
	if err != nil {
		t.Errorf("Url is not valid, got: %d, want: %d.", err, nil)
	}
}
