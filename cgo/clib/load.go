package clib

import (
  "github.com/rainycape/dl"
  "github.com/ralpioxxcs/go_study/cgo/mylogger"
)

// cpp symbol function variables in "libfoo.so"
var add func(int, int, ...interface{}) int
var subtract func(int, int, ...interface{}) int

func Add_cpp(a int, b int) int {
  return add(a,b)
}

func Subtract_cpp(a int, b int) int {
  return subtract(a,b)
}

func LoadLib() {
  lib, err := dl.Open("./libfoo.so", 0)
  if err != nil {
    panic(err)
  }

  //defer lib.Close()

  logger := mylogger.GetInstance()

  logger.Println("Get symbols")
  if err := lib.Sym("Foo_add", &add); err != nil {
    panic(err)
  }
  if err := lib.Sym("Foo_subtract", &subtract); err != nil {
    panic(err)
  }
}
