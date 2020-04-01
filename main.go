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
       _"regexp"
        "log"
        "github.com/elastic/go-elasticsearch/v7"
  esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)
type Required struct {
	//Health       string `json:"health"`
	//Status       string `json:"status"`
	Index        string `json:"index"`
	//UUID         string `json:"uuid"`
	//Pri          string `json:"pri"`
	//Rep          string `json:"rep"`
	//DocsCount    string `json:"docs.count"`
	//DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	//PriStoreSize string `json:"pri.store.size"`
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
//
req := esapi.CatIndicesRequest{
//Index:        []string{"stark","lannister"},
Pretty:       true,
Human:        true,
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
var jsonData [] Required
err = json.Unmarshal([]byte(body), &jsonData) 
if err != nil {
                 panic(err)
         }
for _, value := range jsonData {
		fmt.Println(value)
	}
fmt.Println(len(jsonData))
out, err := json.Marshal(jsonData)
    if err != nil {
        panic (err)
    }
//fmt.Println(string(out))
a := string(out)
res1 := strings.Split(a,",")
for _, value1 := range res1 {
		fmt.Println(value1)
	}
								//output := make([]string, len(jsonData))
								//a := string(out)
								//strArray := strings.Fields(a)
								//fmt.Println(strArray)
								//fmt.Printf("%T",strArray)
								/*for index, each := range strArray {
										 fmt.Printf("strArray value [%d] is [%s]\n", index, each)
									 }
								*/
								/*re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
								submatchall := re.FindAllString(a,-1)
									for _, element := range submatchall {
										fmt.Println(element)
									}*/
								// Declared an empty interface of type Array
var result []map[string]interface{}
// Unmarshal or Decode the JSON to the interface.
json.Unmarshal([]byte(body), &result)

for key, result := range result {

		fmt.Println("Reading Value for INDEX :", key)
		//Reading each value by its key
		fmt.Println("IndexName :", result["index"],
			"|| Health :", result["health"],"|| Size :",result["store.size"])
       }
}




