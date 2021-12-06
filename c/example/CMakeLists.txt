# Copyright (C) 2021 Greenbone Networks GmbH
#
# SPDX-License-Identifier: GPL-2.0-or-later
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation; either version 2
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA.

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
