package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "net/http"
)

type ComicInfo struct {
  Num int `json:"num"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Title string `json:"title"`
}

func main() {
  lastComicUrl := "https://xkcd.com/info.0.json"

  res, err := http.Get(lastComicUrl)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get last xkcd comic: %s\n", err)
    os.Exit(1)
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to read the body of the last xkcd comic: %s\n", err)
    os.Exit(1)
  }

  lastComicInfo := ComicInfo{}
  err = json.Unmarshal(body, &lastComicInfo)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to unmarshal the JSON body of the last xkcd comic: %s\n", err)
    os.Exit(1)
  }

  fmt.Println("last comic number:", lastComicInfo.Num)
}
