package main

// #include <dlfcn.h>
// #cgo LDFLAGS: -ldl
import "C"

import (
        "fmt"
      )
const (
	// dlopen() flags. See man dlopen.
	RTLD_LAZY     = int(C.RTLD_LAZY)
	RTLD_NOW      = int(C.RTLD_NOW)
	RTLD_GLOBAL   = int(C.RTLD_GLOBAL)
	RTLD_LOCAL    = int(C.RTLD_LOCAL)
	RTLD_NODELETE = int(C.RTLD_NODELETE)
	RTLD_NOLOAD   = int(C.RTLD_NOLOAD)
)

func main() {
  fmt.Println("Start main")

  export_name := "test"
  lib_path := "/usr/local/lib/libfoo.so"

  handle := C.dlopen(C.CString(lib_path), C.RTLD_LAZY)
  if handle == nil {
    fmt.Println(lib_path+":not found")
    return
  } else {
    fmt.Println(lib_path+":SUCCESS")
  }

  func_ptr := C.dlsym(handle, C.CString(export_name))
  if func_ptr == nil {
    fmt.Println(export_name+":not found")
    return
  } else {
    fmt.Println(export_name+":SUCCESS")
  }


  fmt.Printf("func_ptr is at %p\n",func_ptr)

  var vars func(d *C.char)
  



}
