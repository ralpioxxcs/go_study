# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.10

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build

# Include any dependencies generated for this target.
include CMakeFiles/foo.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/foo.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/foo.dir/flags.make

CMakeFiles/foo.dir/foo.cpp.o: CMakeFiles/foo.dir/flags.make
CMakeFiles/foo.dir/foo.cpp.o: ../foo.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/foo.dir/foo.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/foo.dir/foo.cpp.o -c /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/foo.cpp

CMakeFiles/foo.dir/foo.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/foo.dir/foo.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/foo.cpp > CMakeFiles/foo.dir/foo.cpp.i

CMakeFiles/foo.dir/foo.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/foo.dir/foo.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/foo.cpp -o CMakeFiles/foo.dir/foo.cpp.s

CMakeFiles/foo.dir/foo.cpp.o.requires:

.PHONY : CMakeFiles/foo.dir/foo.cpp.o.requires

CMakeFiles/foo.dir/foo.cpp.o.provides: CMakeFiles/foo.dir/foo.cpp.o.requires
	$(MAKE) -f CMakeFiles/foo.dir/build.make CMakeFiles/foo.dir/foo.cpp.o.provides.build
.PHONY : CMakeFiles/foo.dir/foo.cpp.o.provides

CMakeFiles/foo.dir/foo.cpp.o.provides.build: CMakeFiles/foo.dir/foo.cpp.o


# Object files for target foo
foo_OBJECTS = \
"CMakeFiles/foo.dir/foo.cpp.o"

# External object files for target foo
foo_EXTERNAL_OBJECTS =

libfoo.so: CMakeFiles/foo.dir/foo.cpp.o
libfoo.so: CMakeFiles/foo.dir/build.make
libfoo.so: CMakeFiles/foo.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX shared library libfoo.so"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/foo.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/foo.dir/build: libfoo.so

.PHONY : CMakeFiles/foo.dir/build

CMakeFiles/foo.dir/requires: CMakeFiles/foo.dir/foo.cpp.o.requires

.PHONY : CMakeFiles/foo.dir/requires

CMakeFiles/foo.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/foo.dir/cmake_clean.cmake
.PHONY : CMakeFiles/foo.dir/clean

CMakeFiles/foo.dir/depend:
	cd /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build /home/hong/gowork/src/github.com/ralpioxxcs/cgo/cpp/build/CMakeFiles/foo.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/foo.dir/depend
