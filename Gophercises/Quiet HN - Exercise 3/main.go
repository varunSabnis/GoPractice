package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strconv"
  "time"
)

/*
To fetch json respones concurrently using multiple goroutines
Create a page to display the urls
*/


//Structure format of required fields from json responses
type HackerNews struct {
  By string
  Title string
  Url string
  Type string
}

func ExampleJSONFetch(){
  cli := &http.Client{}
  resp,err := cli.Get("https://hacker-news.firebaseio.com/v0/item/3.json")
  if err!=nil{
    fmt.Println("ERROR in fetching data")
  }

  defer resp.Body.Close()

  bodyBytes,_ := ioutil.ReadAll(resp.Body)

  //Unmarshalling example
  news := HackerNews{}
  json.Unmarshal(bodyBytes,&news)

  //Marshalling example
  res,_ := json.Marshal(&news)
  print(string(res))
}

func FetchNewsWithThreading() []HackerNews {
  cli := &http.Client{}
  count := 0
  var news HackerNews
  var NewsList []HackerNews
  //var resp *http.Response
  count_ch := make(chan int)
  fetch_ch := make(chan *http.Response)

  //if 11th item is fetched before 10th how will it be handled?
  //We can't push 11th item before 10th, need to maintain order

  //Handle only to fetch 30 items first, ordering handle it next
  problemloop:
  for i:=0;i<50;i++{

    go func(){
      resp,_ := cli.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(i) + ".json")
      //if(err!=nil){
      //  fmt.Println("ERROR in fetching data inside function and index is " + strconv.Itoa(i))
      // }
      //defer resp.Body.Close()
      fetch_ch <- resp
      if(count == 30){
          count_ch <- 1
        }
    }()

    select{
    case  resp := <- fetch_ch:
      defer resp.Body.Close()
      bodyBytes,_ := ioutil.ReadAll(resp.Body)
      news = HackerNews{}
      json.Unmarshal(bodyBytes,&news)
      //fmt.Println("hi inside case fetch_ch", news)
      if(news.Type=="story" && count<30){
             NewsList = append(NewsList,news)
             count = count + 1
            // fmt.Println("count inside select",count)
         }
    case <- count_ch:
               break problemloop
    }
   //fmt.Println("count",count)
 }
  return(NewsList)
}


func FetchNewsWithoutThreading() []HackerNews {
cli := &http.Client{}
var NewsList []HackerNews
count := 0

for i:=0;i<50;i++{
  resp,err := cli.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(i) + ".json")
  if(err!=nil){
    fmt.Println("ERROR in fetching data inside function and index is " + strconv.Itoa(i))
  }
  defer resp.Body.Close()

  bodyBytes,_ := ioutil.ReadAll(resp.Body)
  news := HackerNews{}
  json.Unmarshal(bodyBytes,&news)

  if(news.Type=="story" && count<30){
    NewsList = append(NewsList,news)
    count = count + 1
  }

  if(count == 30){
    break
  }
}
  return(NewsList)
}


func main(){

    fmt.Println("Example of functionality provided by http,json,ioutil package")

    //Using http client to make request,ioutil to read response and use json unmarshal to get structure object
    ExampleJSONFetch()

     /*
      Create 50 threads running parllely to fetch json, but we will write data to a
      shared list object sequentially in correct order and compare the list fetched
      without using a thread.
    */
    start1 := time.Now()
    NewsList1 := FetchNewsWithoutThreading()
    elapsed1 := time.Since(start1)
    fmt.Println("\nTime taken to execute FetchNewsWithoutThreading() is ", elapsed1)
    fmt.Println("\nLength of news list 1",len(NewsList1))

    start2 := time.Now()
    NewsList2 := FetchNewsWithThreading()
    elapsed2 := time.Since(start2)
    fmt.Println("\nTime taken to execute FetchNewsWithThreading() is ",elapsed2)
    fmt.Println("\nLength of news list 2",len(NewsList2))

    for _,news := range(NewsList2){
      fmt.Println("Type of news ", news.Type)
    }
    for i:=0;i<30;i++{
      if(NewsList1[i]==NewsList2[i]){
        fmt.Println("same",i)
      }else{
        fmt.Println("not same",i)
      }
    }
// Handler to display the single page having news url, time of execution
  http.HandleFunc("/newsposts", func(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":3001", nil))
}
