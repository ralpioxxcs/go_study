#include "foo.h"

Foo::Foo() {
  m_var = 10;
  //std::cout << "Foo constructor" << std::endl;
}

Foo::~Foo() {
  //std::cout << "Foo destructor" << std::endl;
}

int Foo::add(int a, int b) {
  return (a+b);
}

int Foo::subtract(int a, int b) {
  return (a-b);
}
