cmake_minimum_required(VERSION 3.0)
project(cpplib CXX)
add_library(foo SHARED export.cpp foo.cpp)
