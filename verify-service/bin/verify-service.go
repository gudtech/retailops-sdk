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

type schemaExample struct {
  schemaPath string
  examplePath string
}

func main() {
  // var err error

  schemaPathPtr := flag.String("schema-path", "", "path to JSON or directory with JSON")
  baseURLPtr := flag.String("base-url", "http://localhost:5000/api/channel", "base url for sending requests")
  stopOnError := flag.Bool("stop-on-error", false, "stop immediately on error")
  flag.Parse()

  if len(*baseURLPtr) == 0 {
    fmt.Println("base-url cannot be empty")
    os.Exit(1)
  }

  fmt.Println(*schemaPathPtr)
  if isDir(*schemaPathPtr) {
    verPairs,err := allExamples(*schemaPathPtr)
    if err != nil {
      fmt.Println("failed:", err.Error())
      os.Exit(1)
    } else if len(verPairs) == 0 {
      fmt.Println("no verification files found. try `--help` for more information")
      os.Exit(1)
    }

    fmt.Println(len(verPairs),"REQUESTS TO BE GENERATED")
    for _,verPair := range verPairs {
      fmt.Println("-----------")
      err = request(*baseURLPtr, verPair.schemaPath, verPair.examplePath)
      if err != nil {
        fmt.Println("FAILURE:", err.Error())
        if *stopOnError {
          os.Exit(1)
        }
      } else {
        fmt.Println("SUCCESS")
      }
    }
    fmt.Println("-----------")

  } else {
    examples,err := examples(*schemaPathPtr)
    if err != nil {
      fmt.Println("failed:", err.Error())
      os.Exit(1)
    } else if len(examples) == 0 {
      fmt.Println("no verification files found. try `--help` for more information",*schemaPathPtr)
      os.Exit(1)
    }

    for _,examplePath := range examples {
      // fmt.Println(examplePath)
      fmt.Println("1 REQUEST TO BE GENERATED")
      err = request(*baseURLPtr, *schemaPathPtr, examplePath)
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
    }
  }
}

func allExamples(dirname string) (verifications []schemaExample, err error) {
  verifications = make([]schemaExample,0)
  allSchemasGlob := p.Join(dirname, "*v1.json")
  allSchemaPaths,err := fp.Glob(allSchemasGlob)
  if err != nil {
    return
  }
  if len(allSchemaPaths) == 0 {
    err = fmt.Errorf("`%s` did not contain schemas", dirname)
    return
  }

  for _,schemaPath := range allSchemaPaths {
    // fmt.Println(schemaPath)
    exs,err := examples(schemaPath)
    if err != nil {
      return nil,err
    }

    for _,ex := range exs {
      verifications = append(verifications, schemaExample {
        schemaPath: schemaPath,
        examplePath: ex,
      })
    }
  }

  return
}

func examples(path string) (examples []string, err error) {
  dirname,filename := p.Split(path)
  exampleFilename := strings.Replace(filename, ".json", "", -1)

  exampleFilenameGlob := fmt.Sprintf("%s_ex_*.json", exampleFilename)
  pathGlob := p.Join(dirname, exampleFilenameGlob)

  return fp.Glob(pathGlob)
}

func isDir(path string) (bool) {
  info, err := os.Stat(path)
  return err == nil && info.IsDir()
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