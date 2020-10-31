#include "foo.h"

Foo::Foo() {
  m_var = 10;
  std::cout << "Foo constructor" << std::endl;
}

Foo::~Foo() {
  std::cout << "Foo destructor" << std::endl;
}

int Foo::add(int a, int b) {
  std::cout << "a: "<<a << std::endl;
  std::cout << "b: "<<b << std::endl;
  std::cout << "a+b: "<<a+b << std::endl;
  return (a+b);
}

int Foo::subtract(int a, int b) {
  std::cout << "a: "<<a << std::endl;
  std::cout << "b: "<<b << std::endl;
  std::cout << "a-b: "<<a-b << std::endl;
  return (a-b);
}
