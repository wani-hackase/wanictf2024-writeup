cmake_minimum_required(VERSION 3.22)
project(toybox C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_FLAGS "-Wall -Wextra")
set(CMAKE_C_FLAGS_RELEASE "-O3")

add_executable(server server.c)

add_executable(sandbox sandbox.c)
target_link_libraries(sandbox seccomp)
