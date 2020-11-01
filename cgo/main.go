package main

import(
  "fmt"
  "github.com/rainycape/dl"
)

func main() {
  fmt.Println("main")

  lib, err := dl.Open("./libfoo.so", 0)
  if err != nil {
    panic(err)
  }

  defer lib.Close()

  var add func(int, int, ...interface{}) int
  if err := lib.Sym("Foo_add",&add); err != nil {
    panic(err)
  }
  var subtract func(int, int, ...interface{}) int
  if err := lib.Sym("Foo_subtract", &subtract); err != nil {
    panic(err)
  }

  num1 := 5
  num2 := 10

  result := add(num1, num2)
  fmt.Println(result)

  result2 := subtract(num2, num1)
  fmt.Println(result2)
}


