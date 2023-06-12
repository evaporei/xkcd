package dir

import (
  "errors"
  "os"
  "path/filepath"
)

// create ~/.xkcd folder
func SetupXkcdFolder() error {
  xkcdFolder, err := GetXkcdFolder()
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

func GetXkcdFolder() (string, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }

  return filepath.Join(homeDir, ".xkcd"), nil
}

func OpenXkcdDir() (*os.File, error) {
  xkcdFolder, err := GetXkcdFolder()
  if err != nil {
    return nil, err
  }

  dir, err := os.Open(xkcdFolder)
  if err != nil {
    return nil, err
  }

  return dir, nil
}
