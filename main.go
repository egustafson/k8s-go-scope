package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	listenAddr = ":8080"
)

func scopeHandler(w http.ResponseWriter, r *http.Request) {
	details := make(map[string]any)
	details["request-method"] = r.Method
	details["request-uri"] = r.RequestURI
	details["request-proto"] = r.Proto
	details["request-host"] = r.Host
	details["remote-addr"] = r.RemoteAddr
	details["cookies"] = stringCookies(r.Cookies())
	details["headers"] = stringHeaders(r.Header)
	details["referer"] = r.Referer()
	details["hostname"], _ = os.Hostname()
	details["environment"] = os.Environ()

	json_val, err := json.MarshalIndent(details, "", "  ")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(json_val)
}

func stringHeaders(hdr http.Header) []string {
	r := make([]string, 0, 0)
	for k, v := range hdr {
		for _, hv := range v {
			r = append(r, fmt.Sprintf("%s: %s", k, hv))
		}
	}
	return r
}

func stringCookies(cookies []*http.Cookie) []string {
	r := make([]string, len(cookies))
	for i, c := range cookies {
		r[i] = c.String()
	}
	return r
}

func main() {
	http.HandleFunc("/", scopeHandler)
	log.Printf("listenening on %s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
