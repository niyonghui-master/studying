package main

func httpGet(url string) string {
	return ""
}

type tsBuffer struct {
	url     string
	content string
}

func fetchURLContent(data chan tsBuffer, prefix string, urls []string) {
	for _, url := range urls {
		go func(u string) {
			tsB := tsBuffer{
				url:     u,
				content: httpGet(prefix + u),
			}
			data <- tsB
		}(url)
	}

	close(data)
}



func writeFiles(data chan tsBuffer) {
	for {
		if buf, isClose := <-data; !isClose {
			break
		} else {
			go writeFile("ts/"+buf.url+".ts", buf.content)
		}
	}
}

func writeFile(path, content string) {

}