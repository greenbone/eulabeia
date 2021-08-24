cmake_minimum_required(VERSION 3.13...3.16 FATAL_ERROR)
project(
  eulabeia_example_c
  VERSION 0.0.1
  LANGUAGES C)

add_executable(start_mega_scan ./start_mega_scan.c)
add_executable(start_scan ./start_scan.c)
add_executable(message-json-overview-md ./generate_description.c)
find_package(Eulabeia REQUIRED)
target_link_libraries(start_mega_scan Eulabeia::Eulabeia)
target_link_libraries(start_scan Eulabeia::Eulabeia)
target_link_libraries(message-json-overview-md Eulabeia::Eulabeia)
target_compile_options(
  message-json-overview-md
  PRIVATE
    $<$<OR:$<CXX_COMPILER_ID:Clang>,$<CXX_COMPILER_ID:AppleClang>,$<CXX_COMPILER_ID:GNU>>:
    -Wall
    -Wextra
    -Werror
    -Wsequence-point
    -Wstrict-prototypes
    -Wshadow
    -Wmissing-prototypes
    -Wpedantic>)
