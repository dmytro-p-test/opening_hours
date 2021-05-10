package opening_hours

import (
"encoding/json"
"flag"
"fmt"
"github.com/gorilla/mux"
log "github.com/sirupsen/logrus"
// "log"
"net/http"
"time"
)

func main() {
	var port = flag.Int("port", 0, "listener port 1025 - 64000")

	flag.Parse()

	if *port < 1025 || *port > 64000 {
		log.Fatalf("listener port must be between 1025 - 64000, got %d", *port)
	}

	r := mux.NewRouter()
	r.Host("localhost").Methods(http.MethodPost).HandlerFunc("/", HomeHandler)

	http.Handle("/", r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", *port),
		WriteTimeout: 15 * time.Minute,
		ReadTimeout:  15 * time.Minute,
		IdleTimeout:  15 * time.Minute,
	}

	log.Fatal(srv.ListenAndServe())

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body.Read()

	//todo:
	//parse bytes into string
	//call ParseOpenHours
	//marshal result into JSON for the Ruleset struct
	// send it into writer
}
