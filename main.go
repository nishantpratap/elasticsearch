package main

import (
        "context"
        "fmt"
        "log"
        "github.com/elastic/go-elasticsearch/v7"
  esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)

func main() {
//initializing my es client

cfg := elasticsearch.Config{
  Addresses: []string{
    "http://172.17.0.2:9200",
  },
  // ...
}
es, _ := elasticsearch.NewClient(cfg)
//getting cluster info and printing it
log.Println(es.Info())
//
req := esapi.CatIndicesRequest{
Pretty:     true,
Human:      true,
Format:    "json",
}
//performing request with client
res, err := req.Do(context.Background(),es)
if err != nil {
log.Fatalf("Error getting response : %s", err)
}
//printing my required result
fmt.Println(res)

}


