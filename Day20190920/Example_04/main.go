package main

import (
	"io/ioutil"
	"net/http"
)

type tsBuffer struct {
	url     string
	content []byte
}

func httpGet(url string) []byte {
	rsp, err := http.Get(url)
	defer rsp.Body.Close()
	if err != nil {
		return []byte{}
	}
	content, _ := ioutil.ReadAll(rsp.Body)
	return content
}

func fetchURLContent(data chan *tsBuffer, prefix string, urls []string) {
	for _, url := range urls {
		go func(d chan *tsBuffer, u string) {
			tsB := &tsBuffer{
				url:     u,
				content: httpGet(prefix + u),
			}
			d <- tsB
		}(data, url)
	}
	close(data)
}

func main() {

}
