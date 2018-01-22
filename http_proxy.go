package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func handleReq(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()

	if len(keys) < 1 {
		log.Println("Url Param 'key' is empty")
		return
	}

	for k, v := range keys {
		if k == "key" {
			//log.Println("Url Param ", k, v[0])
			data := ProxyCacheGet(v[0])

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			io.WriteString(w, data)

			return
		}
	}
}

func InitProxy(port int) {
	http.HandleFunc("/proxy", handleReq)
	http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)
}
