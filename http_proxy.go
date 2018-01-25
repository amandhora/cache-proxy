package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type Job struct {
	k    string
	w    http.ResponseWriter
	done chan bool
}

var JobQueue chan Job

func handleReq(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	keys := r.URL.Query()

	if len(keys) < 1 {
		log.Println("Url Param 'key' is empty")

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
}

func SendGetRsp(j Job, value string) {

	j.w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(j.w, value)

	// Signal that the response has been sent
	j.done <- true
}

func InitProxy(port int, queue chan Job) {

	JobQueue = queue

	http.HandleFunc("/proxy", handleReq)
	http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)

}
