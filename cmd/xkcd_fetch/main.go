package main

import (
  "fmt"
  "os"
  "time"

  c "github.com/evaporei/xkcd/comic"
  d "github.com/evaporei/xkcd/dir"
)

func main() {
  lastComic, err := c.FetchLastComic()
  if err != nil {
    fmt.Printf("xkcd_fetch: failed to get last comic: %s\n", err)
    os.Exit(1)
  }

  fmt.Println("last comic number:", lastComic.Num)

  err = d.SetupXkcdFolder()
  if err != nil {
    fmt.Printf("xkcd_fetch: failed create ~/.xkcd folder: %s\n", err)
    os.Exit(1)
  }

  for i := 1; i <= lastComic.Num; i++ {
    comic, err := c.DownloadComic(i);
    if err != nil {
      fmt.Printf("xkcd_fetch: failed to download comic %d: %s\n", i, err)
      continue
    }

    fmt.Printf("downloaded comic num: %d, title: %s\n", comic.Num, comic.Title)

    // throttle
    if i % 10 == 0 {
      time.Sleep(time.Second)
    }
  }
}
