package play

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gernest/utron/router"
)

func TestService(t *testing.T) {
	r := router.NewRouter()
	r.Add(New)

	code := `
	package main
	func main(){
		fmt.Println("hello")
	}
	`
	ex, err := ioutil.ReadFile("expect_fmt.txt")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/fmt", strings.NewReader(code))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !bytes.Equal(w.Body.Bytes(), ex) {
		t.Fatal("wrong format response")
	}
}
