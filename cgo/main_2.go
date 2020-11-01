package main

import(
  "fmt"
  "github.com/rainycape/dl"
)

func main() {
  fmt.Println("main")

  lib, err := dl.Open("/usr/local/lib/libfoo.so",0)
  if err != nil {
    panic(err)
  }

  defer lib.Close()
  //var snprintf func([]byte, uint, string, ...interface{}) int
  var test func()
  if err := lib.Sym("test",&test); err != nil {
    panic(err)
  }
  var add func(int, int, ...interface{}) int
  if err := lib.Sym("Foo_add",&add); err != nil {
    panic(err)
  }

  test()

  result := add(2,3)
  fmt.Println(result)

}


