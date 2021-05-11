package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseOpenHours(t *testing.T) {
	tests := []struct {
		input string
		want  int //http status code
	}{
		{"Mo-Fr 10:00-20:00; Sa,Su 09:00-12:00; PH off", http.StatusOK},
		{"Mo-Fr 09:00-12:00,13:00-19:00; PH off", http.StatusOK},
		{"Mo 09:00+; PH off", http.StatusOK},
		{"Mo 09:00+; \n PH off", http.StatusBadRequest},
	}

	for _, test := range tests {
		handler := http.HandlerFunc(ParseOpenHours)
		w := httptest.NewRecorder()

		payload := strings.NewReader(test.input)
		req := httptest.NewRequest("POST", "http://localhost:8080", payload)
		req.Header.Add("Content-Type", "text/plain")

		handler(w, req)
		resp := w.Result()
		// fmt.Errorf("Input %v: want %d, got %d", test.input, test.want, resp.StatusCode)
		if status := resp.StatusCode; status != test.want {
			t.Errorf("Input %v: want %d, got %d", test.input, test.want, status)
		}

	}

}
