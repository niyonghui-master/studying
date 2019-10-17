package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type tsBuffer struct {
	url     string
	content []byte
}

func main() {
	link := os.Args[1]
	tsList := readVideoBlockList(link)

	endIndex := strings.LastIndex(link, "/")
	playPrefix := link[:endIndex+1]

	createPath("ts")

	data := make(chan *tsBuffer)
	done := make(chan bool)

	go fetchURLContent(data, playPrefix, tsList)
	writeFiles(data, done)

	<-done
	fmt.Println("task has been complete")
}

func fetchURLContent(data chan *tsBuffer, prefix string, urls []string) {
	for _, url := range urls {
		tsB := &tsBuffer{
			url:     url,
			content: httpGet(prefix + url),
		}
		data <- tsB
	}

	close(data)
}

func writeFiles(data chan *tsBuffer, done chan bool) {
	for {
		if buf, isClose := <-data; !isClose {
			break
		} else {
			fmt.Println("URL:", buf.url)
			go writeFile("ts/"+buf.url+".ts", buf.content, 0755)
		}
	}

	done <- true
}

func writeFile(name string, content []byte, perm os.FileMode) {
	ioutil.WriteFile(name, content, perm)
}

func existsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func createPath(name string) string {
	if !existsPath(name) {
		os.Mkdir(name, os.ModePerm)
	}
	return name
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

func readVideoBlockList(url string) (contentList []string) {
	for _, item := range strings.Split(string(httpGet(url)), "\n") {
		if strings.Contains(item, ".ts") {
			contentList = append(contentList, item)
		}
	}
	return contentList
}
