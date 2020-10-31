#ifndef EXPORT_H
#define EXPORT_H

#include "foo.h"

typedef void (*test_type)();

extern "C" int Foo_add(int a, int b);
extern "C" int Foo_subtract(int a, int b);
extern "C" void test();

#endif
