#include "foo.h"

int test_c() {
  std::cout << "test c" << std::endl;
}

Foo::Foo() {
  m_var = 10;
  std::cout << "Foo constructor" << std::endl;
}

Foo::~Foo() {
  std::cout << "Foo destructor" << std::endl;
}

void Foo::test() {
  m_var = 5;
  std::cout << "Foo" << std::endl;
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
