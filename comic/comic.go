package comic

import (
  "encoding/json"
  "errors"
  "fmt"
  "os"
  "path/filepath"
  "strings"

  d "github.com/evaporei/xkcd/dir"
  h "github.com/evaporei/xkcd/http"
)

type Comic struct {
  Num int `json:"num"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Title string `json:"title"`
}

func FetchLastComic() (*Comic, error) {
  return FetchComic(0)
}

// get comic from the web
func FetchComic(num int) (*Comic, error) {
  comicUrl := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)

  // fetch last comic
  if num == 0 {
    comicUrl = "https://xkcd.com/info.0.json"
  }

  if num == 404 {
    return nil, errors.New("Not found page")
  }

  body, err := h.FetchBody(comicUrl)
  if err != nil {
    return nil, err
  }

  comic := Comic{}
  err = json.Unmarshal(body, &comic)
  if err != nil {
    return nil, err
  }

  return &comic, nil
}

// store comic into file system
func (comic *Comic) Save() error {
  contentBytes, err := json.Marshal(comic)
  if err != nil {
    return err
  }

  xkcdFolder, err := d.GetXkcdFolder()
  if err != nil {
    return err
  }

  comicFilePath := filepath.Join(xkcdFolder, fmt.Sprintf("%d.json", comic.Num))

  err = os.WriteFile(comicFilePath, contentBytes, 0644)

  return err
}

func DownloadComic(num int) (*Comic, error) {
  comic, err := FetchComic(num)
  if err != nil {
    return nil, err
  }

  err = comic.Save()
  if err != nil {
    return nil, err
  }

  return comic, nil
}

// get comic from index/storage (~/.xkcd)
func GetComic(name string) (*Comic, error) {
  xkcdFolder, err := d.GetXkcdFolder()
  if err != nil {
    return nil, err
  }

  fullPath := filepath.Join(xkcdFolder, name)
  contents, err := os.ReadFile(fullPath)
  if err != nil {
    return nil, err
  }

  comic := Comic{}
  err = json.Unmarshal(contents, &comic)
  if err != nil {
    return nil, err
  }

  return &comic, nil
}

func (comic *Comic) ContainsTerm(term string) bool {
  return strings.Contains(comic.SafeTitle, term) ||
    strings.Contains(comic.Transcript, term) ||
    strings.Contains(comic.Title, term)
}
