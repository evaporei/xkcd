package main

import (
  "encoding/json"
  "fmt"
  "io"
  "os"
  "net/http"
)

type ComicInfo struct {
  Num int `json:"num"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Title string `json:"title"`
}

// Does a GET request and return its body
func getBody(url string) ([]byte, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  return body, nil
}

func main() {
  lastComicUrl := "https://xkcd.com/info.0.json"

  body, err := getBody(lastComicUrl)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get last comic: %s\n", err)
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
