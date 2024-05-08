package main

import (
  "github.com/renlord/go-template/pkg/foo"
  "log"
)
func main() {
  s := foo.Foo()
  log.Printf("%s", s)
}
