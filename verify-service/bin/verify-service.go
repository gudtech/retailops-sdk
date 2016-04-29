package main

import (
  "flag"
  "fmt"
  "strings"

  // "net/http"
  // "encoding/json"
  "os"
  p "path"
  fp "path/filepath"

  "github.com/gudtech/retailops-sdk/verify-service/verify"
)

func main() {
  var err error

  schemaPathPtr := flag.String("schema-path", "", "path to JSON or directory with JSON")
  baseURLPtr := flag.String("base-url", "http://localhost:5000/api/channel", "base url for sending requests")
  flag.Parse()

  if len(*baseURLPtr) == 0 {
    fmt.Println("base-url cannot be empty")
    os.Exit(1)
  }

  fmt.Println(*schemaPathPtr)

  examples,err := examples(*schemaPathPtr)
  if err != nil {
    fmt.Println("failed:", err.Error())
    os.Exit(1)
  }

  for _,examplePath := range examples {
    err = request(*baseURLPtr, *schemaPathPtr, examplePath)
    if err != nil {
      panic(err.Error())
      os.Exit(1)
    }
  }
}

func examples(path string) (examples []string, err error) {
  dirname,filename := p.Split(path)
  exampleFilename := strings.Replace(filename, ".json", "", -1)

  exampleFilenameGlob := fmt.Sprintf("%s_ex_*.json", exampleFilename)
  pathGlob := p.Join(dirname, exampleFilenameGlob)

  return fp.Glob(pathGlob)
}

func request(baseUrl, schemaPath, examplePath string) (err error) {
  f,err := os.Open(schemaPath)
  if err != nil {
    return
  }

  exampleF,err := os.Open(examplePath)
  if err != nil {
    return
  }

  err = verify.Request(baseUrl, f,exampleF)
  if err != nil {
    return
  }

  return
}