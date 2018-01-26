package main

import (
	"io"
	"net/http"
	"strconv"
)

type Job struct {
	k    string
	w    http.ResponseWriter
	done chan bool
}

var JobQueue chan Job

func HandleReq(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	keys := r.URL.Query()

	if len(keys) < 1 {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for k, v := range keys {
		if k == "key" {

			done := make(chan bool)
			j := Job{k: v[0], w: w, done: done}

			// Write the Task to Job Queue
			JobQueue <- j

			// Wait for the response send signal
			<-done

			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	return
}

func SendGetRsp(j Job, value string) {

	// Write headers and data
	j.w.WriteHeader(http.StatusOK)
	j.w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(j.w, value)

	// Signal that the response has been sent
	j.done <- true
}

func InitProxy(port int, queue chan Job) {

	JobQueue = queue

	http.HandleFunc("/proxy", HandleReq)
	http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)

}
