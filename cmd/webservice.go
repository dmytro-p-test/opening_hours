package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
	"unicode"

	log "github.com/sirupsen/logrus"

	"opening_hours/parser"
)

// Serve is the main function that spin-offs the webservice
// called from the main()
func Serve() {
	// New http-multiplexer with the POST method attached to parseOpenHours()
	r := mux.NewRouter()
	r.HandleFunc("/", ParseOpenHours).Methods(http.MethodPost)
	http.Handle("/", r)

	srv := http.Server{
		Addr:         "localhost:8080", // hard-coded port, can be parametrized with an input flag
		WriteTimeout: 15 * time.Minute,
		ReadTimeout:  15 * time.Minute,
		IdleTimeout:  15 * time.Minute,
	}

	log.Fatal(srv.ListenAndServe())
}

func ParseOpenHours(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("error while reading request: %v", err)
	}

	if !isValidString(body) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request, expect only printable symbols\n")
		return
		// log.Error("invalid request, expect only printable symbols")
	}

	//log.Infof("request: %v", string(body))

	var openHours parser.OpenHours
	err = (&openHours).Parse(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot parse the string %v \ngot:%v", string(body), err)
		return
	}
	log.Infof("openHours: %v", openHours)
	json, err := (&openHours).ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON conversion error: %v", err)
	}
	fmt.Printf("json: %s", json)
	fmt.Fprintf(w, "%s", json)
	w.WriteHeader(http.StatusOK)
}

// isValidString validates slice of bytes and return boolean result true/false
// only printable symbols are allowed
func isValidString(body []byte) bool {
	for _, c := range string(body) {
		if !unicode.IsPrint(c) {
			return false
		}
	}
	return true
}
