package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hawx/serve"
)

type Any interface{}
type Json map[string]Any

func createResponseBody(r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		log.Println("form:", err)
	}

	form := Json{}
	for k, v := range r.Form {
		form[k] = v
	}

	headers := Json{}
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("body:", err)
	}

	j := Json{
		"method":  r.Method,
		"url":     r.Host + r.URL.String(),
		"version": r.Proto,
		"headers": headers,
		"body":    string(body),
		"form":    form,
	}

	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		log.Println("json:", err)
		return ""
	}

	return string(b) + "\n"
}

func main() {
	http.HandleFunc("/delay/", func(w http.ResponseWriter, r *http.Request) {
		ms, err := strconv.ParseInt(r.URL.Path[7:], 10, 64)
		if err != nil {
			log.Println(err)
			return
		}

		time.Sleep(time.Duration(ms) * time.Millisecond)
		fmt.Fprintf(w, createResponseBody(r))
	})

	http.HandleFunc("/code/", func(w http.ResponseWriter, r *http.Request) {
		code, err := strconv.ParseInt(r.URL.Path[6:], 10, 0)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(int(code))
		fmt.Fprint(w, createResponseBody(r))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, createResponseBody(r))
	})

	port := flag.String("port", "8080", "")
	socket := flag.String("socket", "", "")

	flag.Parse()

	serve.Serve(*port, *socket, http.DefaultServeMux)
}
