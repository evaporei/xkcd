package main

import (
  "fmt"
  "os"

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

func printIfMatchesSearch(name string, searchTerm string) {
  comic, err := c.GetComic(name)
  if err != nil {
    fmt.Printf("xkcd: failed to get comic file '%s': %s\n", name, err)
    return
  }

  if comic.ContainsTerm(searchTerm) {
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
