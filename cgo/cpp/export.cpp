#include <stdio.h>

#include "export.h"

extern "C" int Foo_add(int a, int b) {
  Foo foo;
  int var = foo.add(a,b);
  return var;
}

extern "C" int Foo_subtract(int a, int b) {
  Foo foo;
  int var = foo.subtract(a,b);
  return var;
}

