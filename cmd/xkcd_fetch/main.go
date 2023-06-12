package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net/http"
  "os"
  "path/filepath"
  "time"

  "github.com/evaporei/xkcd/comic"
)

// Does a GET request and return its body
func fetchBody(url string) ([]byte, error) {
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

func fetchLastComic() (*comic.ComicInfo, error) {
  return fetchComic(0)
}

// get comic from the web
func fetchComic(num int) (*comic.ComicInfo, error) {
  comicUrl := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)

  // fetch last comic
  if num == 0 {
    comicUrl = "https://xkcd.com/info.0.json"
  }

  if num == 404 {
    return nil, errors.New("Not found page")
  }

  body, err := fetchBody(comicUrl)
  if err != nil {
    return nil, err
  }

  comicInfo := comic.ComicInfo{}
  err = json.Unmarshal(body, &comicInfo)
  if err != nil {
    return nil, err
  }

  return &comicInfo, nil
}

// create ~/.xkcd folder
func setupXkcdFolder() error {
  xkcdFolder, err := comic.GetXkcdFolder()
  if err != nil {
    return err
  }

  // check if it exists already
  _, err = os.Stat(xkcdFolder)

  if errors.Is(err, os.ErrNotExist) {
    err = os.Mkdir(xkcdFolder, os.ModePerm)
  }

  return err
}

// store comic into file system
func saveComic(comic *comic.ComicInfo) error {
  contentBytes, err := json.Marshal(comic)
  if err != nil {
    return err
  }

  xkcdFolder, err := comic.GetXkcdFolder()
  if err != nil {
    return err
  }

  comicFilePath := filepath.Join(xkcdFolder, fmt.Sprintf("%d.json", comic.Num))

  err = os.WriteFile(comicFilePath, contentBytes, 0644)

  return err
}

func downloadComic(num int) {
  comic, err := fetchComic(num)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get comic %d: %s\n", num, err)
    return
  }

  err = saveComic(comic)
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to save comic %d: %s\n", comic.Num, err)
    return
  }

  fmt.Printf("downloaded comic num: %d, title: %s\n", comic.Num, comic.Title)
}

func main() {
  lastComic, err := fetchLastComic()
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get last comic: %s\n", err)
    os.Exit(1)
  }

  fmt.Println("last comic number:", lastComic.Num)

  err = setupXkcdFolder()
  if err != nil {
    fmt.Printf("xkcd_fetch: failed create ~/.xkcd folder: %s\n", err)
    os.Exit(1)
  }

  for i := 1; i <= lastComic.Num; i++ {
    go downloadComic(i)

    // throttle
    if i % 10 == 0 {
      time.Sleep(time.Second)
    }
  }
}
