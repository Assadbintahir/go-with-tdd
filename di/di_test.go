package main

import (
	"bytes"
	"go-with-tdd/coverage"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
If we want to cover main function for a genuine application, we need to use os package to execute the main goroutine
Usually main should just invoke our domain package and that's it, but if we really want to cover it,
we can either use external process testing tool or use some trivia hacks. At this point it will no longer
remain mere unit test because we are invoking our main application and then testing the functionality
For further reading, check following:
1- https://stackoverflow.com/questions/38584033/how-to-test-main-routine-from-go-test/48674736
2- https://stackoverflow.com/questions/31352239/how-to-test-the-main-package-functions-in-golang
*/

func TestMain(t *testing.M) {
	coverage.EnforceCoverage(t, 0.5, "di")
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Asad")

	got := buffer.String()
	expected := "Hello Asad"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}

func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GreetHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d but expected %d", status, http.StatusOK)
	}

	expected := "Hello G"

	if rr.Body.String() != "Hello G" {
		t.Errorf("handler returned unexpected body: got %q but expected %q", rr.Body.String(), expected)
	}
}
