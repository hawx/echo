package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Any interface{}

func createResponseBody(r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	form := map[string]Any{}
	for k, v := range r.Form {
		form[k] = v
	}

	headers := map[string]Any{}
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	j := map[string]Any{
		"method":  r.Method,
		"url":     r.Host + r.URL.String(),
		"version": r.Proto,
		"headers": headers,
		"body":    string(body),
		"form":    form,
	}

	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(b) + "\n"
}

// this is me being lazy
func isStatusCode(s string) bool {
	if len(s) != 3 {
		return false
	}

	switch {
	case s[0] < '1' || s[0] > '5':
		return false
	case s[1] < '0' || s[1] > '9':
		return false
	case s[2] < '0' || s[2] > '9':
		return false
	}

	return true
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isStatusCode(r.URL.Path[1:]) {
			code, err := strconv.ParseInt(r.URL.Path[1:], 10, 0)
			if err != nil {
				log.Println(err)
			}

			w.WriteHeader(int(code))
			fmt.Fprint(w, createResponseBody(r))
			return
		}

		fmt.Fprint(w, createResponseBody(r))
	})

	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
