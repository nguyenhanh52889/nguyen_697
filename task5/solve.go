package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/couchbase/gocb"
    "github.com/gorilla/mux"
)

type MyUrl struct {
    ID       string `json:"id,omitempty"`
    LongUrl  string `json:"longUrl,omitempty"`
    ShortUrl string `json:"shortUrl,omitempty"`
}

var bucket *gocb.Bucket
var bucketName string

func ExpandEndpoint(w http.ResponseWriter, req *http.Request) {
    var n1qlParams []interface{}
    query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE shortUrl = $1")
    params := req.URL.Query()
    n1qlParams = append(n1qlParams, params.Get("shortUrl"))
    rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
    var row MyUrl
    rows.One(&row)
    json.NewEncoder(w).Encode(row)
}


func RootEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var url MyUrl
    bucket.Get(params["id"], &url)
    http.Redirect(w, req, url.LongUrl, 301)
}

func main() {
    router := mux.NewRouter()
    cluster, _ := gocb.Connect("couchbase://localhost")
    bucketName = "example"
    bucket, _ = cluster.OpenBucket(bucketName, "")
    router.HandleFunc("/{id}", RootEndpoint).Methods("GET")
    router.HandleFunc("/expand/", ExpandEndpoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":8082", router))
}
