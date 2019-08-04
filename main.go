package main

import (
  "fmt"
  "log"

  "github.com/gobuffalo/packr/v2"
)

func main() {
  box := packr.New("myBox", "./templates")

  s, err := box.FindString("admin/index.html")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(s)
}
