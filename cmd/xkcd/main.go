package main

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  "github.com/evaporei/xkcd/comic"
)

func main() {
  args := os.Args[1:]

  if len(args) != 1 {
    fmt.Println("xkcd: expects just one command-line argument, the search term")
    os.Exit(1)
  }

  searchTerm := args[0]

  xkcdFolder, err := comic.GetXkcdFolder()
  if err != nil {
    fmt.Println("xkcd: failed to get ~/.xkcd folder")
    os.Exit(1)
  }
  dir, err := os.Open(xkcdFolder)
  if err != nil {
    fmt.Println("xkcd: failed to open ~/.xkcd folder")
    os.Exit(1)
  }

  defer dir.Close()

  comicsIterator, err := dir.Readdir(-1)
  if err != nil {
    fmt.Println("xkcd: failed to readdir ~/.xkcd folder")
    os.Exit(1)
  }

  for _, comicFile := range comicsIterator {
    if comicFile.Mode().IsRegular() {
      fullPath := filepath.Join(xkcdFolder, comicFile.Name())
      contents, err := os.ReadFile(fullPath)
      if err != nil {
        fmt.Printf("xkcd: failed to read file '%s' folder\n", fullPath)
        os.Exit(1)
      }

      comicInfo := comic.ComicInfo{}
      err = json.Unmarshal(contents, &comicInfo)
      if err != nil {
        fmt.Printf("xkcd: failed to unmarshal file '%s' into ComicInfo structure\n", fullPath)
        os.Exit(1)
      }

      if strings.Contains(comicInfo.SafeTitle, searchTerm) ||
        strings.Contains(comicInfo.Transcript, searchTerm) ||
        strings.Contains(comicInfo.Title, searchTerm) {
          fmt.Printf("URL: https://xkcd.com/%d\n", comicInfo.Num)
          fmt.Printf("Transcript: %s\n", comicInfo.Transcript)
        }
    }
  }
}
