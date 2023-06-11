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

func getLastComic() (*ComicInfo, error) {
  lastComicUrl := "https://xkcd.com/info.0.json"

  body, err := getBody(lastComicUrl)
  if err != nil {
    return nil, err
  }

  lastComicInfo := ComicInfo{}
  err = json.Unmarshal(body, &lastComicInfo)
  if err != nil {
    return nil, err
  }

  return &lastComicInfo, nil
}

func getComic(num int) (*ComicInfo, error) {
  comicUrl := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)

  body, err := getBody(comicUrl)
  if err != nil {
    return nil, err
  }

  comicInfo := ComicInfo{}
  err = json.Unmarshal(body, &comicInfo)
  if err != nil {
    return nil, err
  }

  return &comicInfo, nil
}

func main() {
  lastComic, err := getLastComic()
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get last comic: %s\n", err)
    os.Exit(1)
  }

  fmt.Println("last comic number:", lastComic.Num)

  comic571, err := getComic(571)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get comic 571: %s\n", err)
    os.Exit(1)
  }

  fmt.Println("comic 571 title:", comic571.Title)
}
