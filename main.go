package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	basicAuth = flag.Bool("basicAuth", false, "Request BasicAuth")
	body      = flag.Bool("body", true, "Request body")
	cookies   = flag.Bool("cookies", false, "Request cookies")
	headers   = flag.Bool("headers", false, "Request headers")
	quiet     = flag.Bool("q", false, "Removes log prefixes")
	std       = flag.Bool("std", false, "Prints the request response in the terminal")
	uri       = flag.Bool("uri", false, "Request URI")

	host = flag.String("host", "", "Host [Required]")
	port = flag.String("port", "8080", "Port")

	serverEventsLogger *log.Logger
	errorLogger        *log.Logger
	consoleLogger      *log.Logger
)

type response struct {
	fields map[string]interface{}
}

type basicAuthStruct struct {
	usr  string
	pswd string
	ok   bool
}

func init() {

	var logVoid bytes.Buffer

	serverEventsLogger = log.New(os.Stdout,
		"LOG: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	consoleLogger = log.New(&logVoid, "", 0)
	errorLogger = log.New(os.Stdout, "", 0)

	flag.Parse()

	if *quiet {
		serverEventsLogger.SetOutput(&logVoid)
	}
	if *std {
		consoleLogger.SetOutput(os.Stdout)
	}
}

func main() {

	if *host == "" {
		flag.PrintDefaults()
		os.Exit(3)
	}

	resp := response{}
	resp.fields = make(map[string]interface{})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if *headers {
			resp.fields["headers"] = req.Header
		}
		if *uri {
			resp.fields["uri"] = req.RequestURI
		}
		if *cookies {
			resp.fields["cookies"] = req.Cookies()
		}

		if *body {

			var bodyHolder interface{}
			err := json.NewDecoder(req.Body).Decode(&bodyHolder)
			if err != nil {
				errorLogger.Println(err)
			}
			resp.fields["body"] = bodyHolder
		}

		if *basicAuth {
			usr, pswd, ok := req.BasicAuth()
			resp.fields["basicAuth"] = basicAuthStruct{usr, pswd, ok}
		}

		value, err := json.Marshal(resp.fields)

		if err != nil {
			panic(err)
		}

		valueString := string(value)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, valueString)

		consoleLogger.Println(valueString)
	})

	var result bytes.Buffer
	fmt.Fprintf(&result, "%s:%s", *host, *port)
	serverString := result.String()

	serverEventsLogger.Printf("Listening and serving on %s", serverString)
	serverEventsLogger.Fatal(http.ListenAndServe(serverString, nil))
}
