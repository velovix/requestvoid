package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type request struct {
	body         string
	recievedTime time.Time
}

type templatableRequest struct {
	TimeSinceRecieved time.Duration
	Body              string
}

func main() {
	requests := make([]request, 0)
	var requestsLock sync.Mutex

	http.HandleFunc("/void", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "could not read request body: "+err.Error(), 500)
			return
		}

		requestsLock.Lock()
		requests = append(requests, request{body: string(body), recievedTime: time.Now()})
		requestsLock.Unlock()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestsLock.Lock()
		defer requestsLock.Unlock()
		templInfo := make([]templatableRequest, 0, len(requests))

		for i := len(requests) - 1; i >= 0; i-- {
			templInfo = append(templInfo, templatableRequest{
				TimeSinceRecieved: time.Since(requests[i].recievedTime),
				Body:              requests[i].body})
		}

		inspectPageTemplate.Execute(w, templInfo)
	})

	log.Fatalln(http.ListenAndServe(":8090", nil))
}
