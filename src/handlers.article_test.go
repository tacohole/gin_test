package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//test that a GET request to the home page returns the home page with
//the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.Get("/", showIndexPage)

	//create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		//test that the HTTP status code is 200
		statusOK := w.Code == http.StatusOK

		//test that the tile page is Home Page
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK

	})
}
