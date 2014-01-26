package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
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

func handleCode(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.ParseInt(r.URL.Path[1:], 10, 0)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(int(code))
	fmt.Fprint(w, createResponseBody(r))
	return
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isStatusCode(r.URL.Path[1:]) {
			handleCode(w, r)
			return
		}

		fmt.Fprint(w, createResponseBody(r))
	})

	port := flag.String("port", "8080", "")
	socket := flag.String("socket", "", "")

	flag.Parse()

	if *socket == "" {
		log.Println("serving on :" + *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))

	} else {
		l, err := net.Listen("unix", *socket)

		if err != nil {
			log.Fatal("%s\n", err)
		} else {
			err := http.Serve(l, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
