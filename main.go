package main

import (
       _"unsafe"
       _"reflect"
       _"bytes"
        "io/ioutil"
        "encoding/json"
        "context"
        "fmt"
        "strings"
        "strconv"
       _"regexp"
        "log"
        "github.com/elastic/go-elasticsearch/v7"
  esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)

func IndexAllInfo(b []uint8) {
var r []map[string]interface{}
json.Unmarshal([]byte(b),&r)
for key, r := range r {
		fmt.Println("Reading Value for INDEX :", key)
		//Reading each value by its key
		fmt.Println("IndexName :", r["index"],
			"|| Health :", r["health"],"|| Size :",r["store.size"])

                       }
}

func IndexDateFilter(b []uint8,d string) {
var r []map[string]interface{}
json.Unmarshal([]byte(b),&r)
var s [] string
for _, r := range r {		
                  s = append(s,r["index"].(string),r["store.size"].(string))
                 }
var dfilter [] string
for i := 0; i < len(s) ; i++ {
if strings.Contains(s[i], d) {
   ra := "Indexname:"+s[i]+"||size="+s[i+1]
   dfilter = append(dfilter,ra)
       }
   }
fmt.Println("================================================================================================================")
for n :=0; n < len(dfilter) ; n++ {
   fmt.Println(dfilter[n])
  }
}

func IndexSizeFilter(b []uint8,d int) {
var r []map[string]interface{}
json.Unmarshal([]byte(b),&r)
var s [] string
var size [] int
for _, r := range r {		
                  s = append(s,r["index"].(string),r["store.size"].(string))
                  converter := r["store.size"].(string)
                  h, _  := strconv.Atoi(converter)
                  size = append(size,h)
                 }
fmt.Println("================================================================================================================")
for j :=0; j < len(size) ; j++ {
   if size[j]>=d {
   fmt.Println(size)
     }
  }
}

func main() {
//initializing my es client

cfg := elasticsearch.Config{
  Addresses: []string{
    "http://172.17.0.3:9200",
  },
  // ...
}
es, _ := elasticsearch.NewClient(cfg)
//getting cluster info and printing it
log.Println(es.Info())
//hitting cat api
req := esapi.CatIndicesRequest{

Bytes:        "b",
Format:       "json",
}
//performing request with client
res, err := req.Do(context.Background(),es)
if err != nil {
log.Fatalf("Error getting response : %s", err)
    } 
//fmt.Println(res)
defer res.Body.Close()
body, err := ioutil.ReadAll(res.Body)
if err != nil {
    fmt.Println("oops!you have something missing")
     }
//IndexAllInfo(body)
fmt.Println("Enter Date in this format YYYY.MM.DD")
var date string
fmt.Scanln(&date)
IndexDateFilter(body,date)
fmt.Println("Please enter size in bytes")
var indexsize int
fmt.Scanln(&indexsize)
IndexSizeFilter(body,283)

}

