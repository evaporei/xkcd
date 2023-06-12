package main

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  "github.com/evaporei/xkcd/comic"
)

type Args struct {
  SearchTerm string
}

func parseArgs() *Args {
  args := os.Args[1:]

  if len(args) != 1 {
    fmt.Println("xkcd: expects just one command-line argument, the search term")
    os.Exit(1)
  }

  return &Args{
    SearchTerm: args[0],
  }
}

func openXkcdDir() (*os.File, error) {
  xkcdFolder, err := comic.GetXkcdFolder()
  if err != nil {
    return nil, err
  }

  dir, err := os.Open(xkcdFolder)
  if err != nil {
    return nil, err
  }

  return dir, nil
}

// get comic from index/storage (~/.xkcd)
func getComic(name string) (*comic.ComicInfo, error) {
  xkcdFolder, err := comic.GetXkcdFolder()
  if err != nil {
    return nil, err
  }

  fullPath := filepath.Join(xkcdFolder, name)
  contents, err := os.ReadFile(fullPath)
  if err != nil {
    return nil, err
  }

  comic := comic.ComicInfo{}
  err = json.Unmarshal(contents, &comic)
  if err != nil {
    return nil, err
  }

  return &comic, nil
}

func main() {
  args := parseArgs()

  dir, err := openXkcdDir()
  if err != nil {
    fmt.Println("xkcd: failed to open ~/.xkcd folder:", err)
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
      comic, err := getComic(comicFile.Name())
      if err != nil {
        fmt.Printf("xkcd: failed to get comic file '%s'\n", comicFile.Name())
        continue
      }

      if strings.Contains(comic.SafeTitle, args.SearchTerm) ||
        strings.Contains(comic.Transcript, args.SearchTerm) ||
        strings.Contains(comic.Title, args.SearchTerm) {
          fmt.Printf("URL: https://xkcd.com/%d\n", comic.Num)
          fmt.Printf("Transcript: %s\n", comic.Transcript)
          fmt.Println()
        }
    }
  }
}
