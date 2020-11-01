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

  test()

}


