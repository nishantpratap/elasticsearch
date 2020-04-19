package main

import (
        "io/ioutil"
        "encoding/json"
        "context"
        "fmt"
        "strings"
        "strconv"
        "regexp"
        "log"
        "github.com/elastic/go-elasticsearch/v7"
  esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)

var r []map[string]interface{}

func CreateTemplate(s string,es *elasticsearch.Client) {
jsonRequestString :=`{
		  "index_patterns": [
		    "*"
		  ],
		  "settings": {
		    "number_of_shards": 5
		  }
		}`
tty := strings.ReplaceAll(jsonRequestString,"*",s)		
res1, err1 := es.Indices.PutTemplate(
			s,
			strings.NewReader(tty),
		)
		fmt.Println(res1, err1)
		if err1 != nil { // SKIP
			fmt.Println("Error getting the response") // SKIP
		} // SKIP
		defer res1.Body.Close() // SKIP    
}

func IndexSizeFilter(s string,t int,d int) string {
var str string
   if t >=d {
       str = s
     }
return str     
}
 

func CreateDefaultTemplate(s string,es *elasticsearch.Client) {
jsonRequestString :=`{
		  "index_patterns": [
		    "*"
		  ],
		  "settings": {
		    "number_of_shards": 2
		  }
		}`
tty := strings.ReplaceAll(jsonRequestString,"*",s)		
res1, err1 := es.Indices.PutTemplate(
			s,
			strings.NewReader(tty),
		)
		fmt.Println(res1, err1)
		if err1 != nil { // SKIP
			fmt.Println("Error getting the response") // SKIP
		} // SKIP
		defer res1.Body.Close() // SKIP    
}


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
     
     
var insize [] int
var s [] string
var s1 [] string
json.Unmarshal([]byte(body),&r)
for v, r := range r {
                       s = append(s,r["index"].(string))
                       rf := regexp.MustCompile(`-\d{4}.\d{2}.\d{2}`) 
                       rk := rf.ReplaceAllString(s[v], "${1}")
                       s1 = append(s1,rk)
                       converter := r["store.size"].(string)
                       h, _  := strconv.Atoi(converter)
                       insize = append(insize,h)                              
                       }
for i := 0 ; i < len(s1) ; i++ {
       res, err := es.Indices.GetTemplate(
		es.Indices.GetTemplate.WithName(s1[i]),
		es.Indices.GetTemplate.WithFilterPath("*.version"),
	)
	if err != nil {
         log.Fatalf("Error getting response : %s", err)
              } 
	fmt.Println(res.StatusCode)
  if res.StatusCode == 404 {
       
       CreateDefaultTemplate(s1[i],es)
       
      }
  if res.StatusCode == 200 {
       
       str := IndexSizeFilter(s1[i],insize[i],283)
       CreateTemplate(str,es)
   }
  }
   
}
 
 
 
     
     
