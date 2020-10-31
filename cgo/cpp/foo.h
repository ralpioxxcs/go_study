#ifndef FOO_H
#define FOO_H

#include <stdlib.h>
#include <iostream>

class Foo
{
private:
  uint8_t m_var;

public:
  Foo();
  virtual ~Foo();

  int add(int a, int b);
  int subtract(int a, int b);
};

#endif /* FOO_H */
