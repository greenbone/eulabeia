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

cmake_minimum_required(VERSION 3.13...3.19 FATAL_ERROR)
project(libeulabeia VERSION 1.0.0 LANGUAGES C)
# define library before find packages to allow import of autotools projects (glib)
add_library(eulabeia
	src/eulabeia_types.c
    src/eulabeia_json.c
    src/eulabeia_client.c
)

find_package(PkgConfig REQUIRED)
pkg_search_module (LIBGVM_BASE REQUIRED libgvm_base>=21.10)
pkg_search_module (LIBGVM_UTIL REQUIRED libgvm_util>=21.10)
pkg_search_module(JSONGLIB REQUIRED json-glib-1.0)
target_include_directories(eulabeia PUBLIC 
	${LIBGVM_BASE_INCLUDE_DIRS}
	${LIBGVM_UTIL_INCLUDE_DIRS}
	${JSONGLIB_INCLUDE_DIRS}
)
target_link_libraries(eulabeia PUBLIC 
	${JSONGLIB_LDFLAGS}
	${LIBGVM_UTIL_LDFLAGS}
	${LIBGVM_BASE_LDFLAGS}
	)


#Add an alias so that library can be used inside the build tree, e.g. when testing
add_library(Eulabeia::eulabeia ALIAS eulabeia)

target_include_directories(eulabeia
	PUBLIC
	$<INSTALL_INTERFACE:include>
	$<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
    PRIVATE
        {CMAKE_CURRENT_SOURCE_DIR}/src
)

target_compile_options(eulabeia PRIVATE
    $<$<OR:$<CXX_COMPILER_ID:Clang>,$<CXX_COMPILER_ID:AppleClang>,$<CXX_COMPILER_ID:GNU>>:
    -Wall -Wextra -Wpedantic>)


# Installation instructions

include(GNUInstallDirs)
set(INSTALL_CONFIGDIR ${CMAKE_INSTALL_LIBDIR}/cmake/Eulabeia)

install(TARGETS eulabeia
    EXPORT eulabeia-targets
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}
)

#This is required so that the exported target has the name Eulabeia and not eulabeia
set_target_properties(eulabeia PROPERTIES EXPORT_NAME Eulabeia)

install(DIRECTORY include/ DESTINATION ${CMAKE_INSTALL_INCLUDEDIR})

#Export the targets to a script
install(EXPORT eulabeia-targets
    FILE
        EulabeiaTargets.cmake
    NAMESPACE
        Eulabeia::
    DESTINATION
        ${INSTALL_CONFIGDIR}
)

#Create a ConfigVersion.cmake file
include(CMakePackageConfigHelpers)
write_basic_package_version_file(
    ${CMAKE_CURRENT_BINARY_DIR}/EulabeiaConfigVersion.cmake
    VERSION ${PROJECT_VERSION}
    COMPATIBILITY AnyNewerVersion
)

configure_package_config_file(${CMAKE_CURRENT_LIST_DIR}/cmake/EulabeiaConfig.cmake.in
    ${CMAKE_CURRENT_BINARY_DIR}/EulabeiaConfig.cmake
    INSTALL_DESTINATION ${INSTALL_CONFIGDIR}
)

#Install the config, configversion and custom find modules
install(FILES
    ${CMAKE_CURRENT_BINARY_DIR}/EulabeiaConfig.cmake
    ${CMAKE_CURRENT_BINARY_DIR}/EulabeiaConfigVersion.cmake
    DESTINATION ${INSTALL_CONFIGDIR}
)

export(EXPORT eulabeia-targets
    FILE ${CMAKE_CURRENT_BINARY_DIR}/EulabeiaTargets.cmake
    NAMESPACE Eulabeia::)

#Register package in the User Package Registry
export(PACKAGE Eulabeia)

enable_testing()
add_subdirectory(test)
