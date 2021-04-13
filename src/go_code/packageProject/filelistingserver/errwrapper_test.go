package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
 * @author  CoreyChen Zhang
 * @version  2021/4/13 15:46
 * @modified
 */

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}


func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h appHandler
		code int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error\nInternal Server Error"},
	}

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
		f(response, request)
		all, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(all), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s);  got (%d, %s)",
				tt.code, tt.message,
				response.Code, body)
		}
	}
}
