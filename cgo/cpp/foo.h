#ifndef FOO_H
#define FOO_H

#include <stdlib.h>
#include <iostream>

int test_c();

class Foo
{
private:
  uint8_t m_var;

public:
  Foo();
  virtual ~Foo();

  void test();
  int add(int a, int b);
  int subtract(int a, int b);
};

#endif /* FOO_H */
