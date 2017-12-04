package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/justincampbell/url-shortener-go/urlstore"
)

var (
	port     = flag.String("port", "8082", "port to listen on")
	urlStore = urlstore.NewURLStore()
)

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", expandHandler)

	fmt.Println("Listening on", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}
}

func expandHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.RequestURI, request.RemoteAddr)

	if request.RequestURI == "/" {
		url := "http://github.com/justincampbell/url-shorteners"
		http.Redirect(response, request, url, 301)
		return
	}

	token := request.URL.Path[len("/"):]
	url := urlStore.Expand(token)

	if url == "" {
		http.NotFound(response, request)
		return
	}

	http.Redirect(response, request, url, 301)
}
