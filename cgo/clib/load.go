package clib

import (
  "fmt"
  "github.com/rainycape/dl"
  "github.com/ralpioxxcs/cgo/mylogger"
)

var add func(int, int, ...interface{}) int
var subtract func(int, int, ...interface{}) int

func Start() {
  num1 := 5
  num2 := 10

  result := add(num1, num2)
  fmt.Println(result)

  result2 := subtract(num2, num1)
  fmt.Println(result2)
}

func LoadLib() {
  lib, err := dl.Open("./libfoo.so", 0)
  if err != nil {
    panic(err)
  }

  defer lib.Close()

  logger := mylogger.GetInstance()

  logger.Println("Get symbols")
  if err := lib.Sym("Foo_add", &add); err != nil {
    panic(err)
  }
  if err := lib.Sym("Foo_subtract", &subtract); err != nil {
    panic(err)
  }
}
