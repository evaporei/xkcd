package comic

import (
  "os"
  "path/filepath"
)

type ComicInfo struct {
  Num int `json:"num"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Title string `json:"title"`
}

func GetXkcdFolder() (string, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }

  return filepath.Join(homeDir, ".xkcd"), nil
}
