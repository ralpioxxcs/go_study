package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"
import "fmt"

func main() {
  fmt.Println("Start main")

  handle := C.dlopen(C.CString("libfoo.so"), C.RTLD_LAZY)
  bar := C.dlsym(handle, C.CString("test_c"))
  fmt.Println("bar is at %p\n", bar)
  //p, err := plugin.Open("./cpp/lib/libfoo.so")
  //if err != nil {
  //  fmt.Println("fail")
  //  panic(err)
  //}

  //f, err := p.Lookup("foo")
  //if err != nil {
  //  panic(err)
  //}

  //f.(func())()

  //f1, err := p.Lookup("add")
  //if err != nil {
  //  panic(err)
  //}

  //sum := f1.(func(int, int) int)(10,15)

  //fmt.Println(sum)
}
