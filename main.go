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
        "regexp"
        "log"
        "github.com/elastic/go-elasticsearch/v7"
  esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)

var r []map[string]interface{}

func IndexAllInfo(b []uint8) {

json.Unmarshal([]byte(b),&r)
for key, r := range r {
		fmt.Println("Reading Value for INDEX :", key)
		//Reading each value by its key
		fmt.Println("IndexName :", r["index"],
			"|| Health :", r["health"],"|| Size :",r["store.size"])

                       }
}

func IndexDateFilter(b []uint8,d string) {

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
for n := 0; n < len(dfilter) ; n++ {
   fmt.Println(dfilter[n])
  }
}

func IndexSizeFilter(b []uint8,d int,es *elasticsearch.Client) {

json.Unmarshal([]byte(b),&r)
var s [] string
var sout [] string
var sin [] int
for _, r := range r {		
                  converter := r["store.size"].(string)
                  h, _  := strconv.Atoi(converter)
                  sin = append(sin,h,0)
                  s = append(s,r["index"].(string),r["store.size"].(string))
                  
                 }
fmt.Println("================================================================================================================")
for j := 0; j < len(sin) ; j++ {
   if sin[j]>=d {
      rs := "Indexname:"+s[j]+"||size="+s[j+1]
       fmt.Println(rs)
        rf := regexp.MustCompile(`-\d{4}.\d{2}.\d{2}`) 
        rk := rf.ReplaceAllString(s[j], "${1}")
        sout = append(sout,rk)
     }
   }
fmt.Println("================================================================================================================")
for k := 0 ; k < len(sout) ; k++ {
CreateTemplate(sout[k],es)
   }
   
} 

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

func CreateDefaulttEz
func GetTemplate(s string,es *elasticsearch.Client) {
res, err := es.Indices.GetTemplate(
	es.Indices.GetTemplate.WithName(s),
	es.Indices.GetTemplate.WithFilterPath("*.version"),
              )
       fmt.Println(res, err)
      }  
      
func RemoveAllDates( b []uint8) [] string {
json.Unmarshal([]byte(b),&r)
var s [] string
var s1 [] string

for v, r := range r {		
                  s = append(s,r["index"].(string))
                  rf := regexp.MustCompile(`-\d{4}.\d{2}.\d{2}`) 
                  rk := rf.ReplaceAllString(s[v], "${1}")
                  s1 = append(s1,rk)          
                 }
return s1
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
//fmt.Printf("%T",es)
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


IndexAllInfo(body)
//fmt.Println("Enter Date in this format YYYY.MM.DD")
//var date string
//fmt.Scanln(&date)
//IndexDateFilter(body,date)
fmt.Println("Please enter size in bytes")
var indexsize int
fmt.Scanln(&indexsize)
IndexSizeFilter(body,indexsize,es)
//RemoveAllDates(body)
}

