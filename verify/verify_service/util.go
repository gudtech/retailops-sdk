package verify_service

import (
  "os"
  "fmt"
  "encoding/json"
)

func panicjson(thing interface{}) {
  thingBytes,_ := json.Marshal(thing)
  fmt.Printf("%s\n", thingBytes)
  os.Exit(1)
}