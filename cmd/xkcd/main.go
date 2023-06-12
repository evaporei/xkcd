package main

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  c "github.com/evaporei/xkcd/comic"
  d "github.com/evaporei/xkcd/dir"
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

// get comic from index/storage (~/.xkcd)
func getComic(name string) (*c.Comic, error) {
  xkcdFolder, err := d.GetXkcdFolder()
  if err != nil {
    return nil, err
  }

  fullPath := filepath.Join(xkcdFolder, name)
  contents, err := os.ReadFile(fullPath)
  if err != nil {
    return nil, err
  }

  comic := c.Comic{}
  err = json.Unmarshal(contents, &comic)
  if err != nil {
    return nil, err
  }

  return &comic, nil
}

func containsSearchTerm(comic *c.Comic, searchTerm string) bool {
  return strings.Contains(comic.SafeTitle, searchTerm) ||
    strings.Contains(comic.Transcript, searchTerm) ||
    strings.Contains(comic.Title, searchTerm)
}

func printIfMatchesSearch(name string, searchTerm string) {
  comic, err := getComic(name)
  if err != nil {
    fmt.Printf("xkcd: failed to get comic file '%s': %s\n", name, err)
    return
  }

  if containsSearchTerm(comic, searchTerm) {
    fmt.Printf("URL: https://xkcd.com/%d\n", comic.Num)
    fmt.Printf("Transcript: %s\n", comic.Transcript)
    fmt.Println()
  }
}

func main() {
  args := parseArgs()

  dir, err := d.OpenXkcdDir()
  if err != nil {
    fmt.Println("xkcd: failed to open ~/.xkcd folder:", err)
    os.Exit(1)
  }

  defer dir.Close()

  comicsIterator, err := dir.Readdir(-1)
  if err != nil {
    fmt.Println("xkcd: failed to readdir ~/.xkcd folder", err)
    os.Exit(1)
  }

  for _, comicFile := range comicsIterator {
    if comicFile.Mode().IsRegular() {
      printIfMatchesSearch(comicFile.Name(), args.SearchTerm)
    }
  }
}
