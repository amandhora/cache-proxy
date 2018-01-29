// concurrent.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MakeRequest(url string, val string, ch chan<- string) {

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln("FAIL:", err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != val {
		ch <- fmt.Sprintf("FAIL: url: %s rsp: %s expected:%s", url, string(body), val)
	} else {
		ch <- fmt.Sprintf("PASS: url: %s rsp: %s expected:%s", url, string(body), val)
	}

}

func main() {
	start := time.Now()
	ch := make(chan string)

	kv := []string{"one", "two", "three", "four", ""}

	for k, v := range kv {
		go MakeRequest("http://"+os.Getenv("PROXY_URL")+"/proxy?key="+strconv.Itoa(k+1), v, ch)

	}

	for range kv {
		fmt.Println(<-ch)

	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
