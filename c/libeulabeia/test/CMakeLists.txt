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

# see /opt/local/share/cmake-3.16/Modules/FindGTest.cmake
find_package(cgreen 1.1.0)

# tracking of allocated memory only possible with CGREEN_NO_FORK and NUMA
if(DEFINED ENV{CGREEN_NO_FORK})
    add_definitions (-DCGREEN_NO_FORK="1")
find_library (NUMA numa)
message (STATUS "Looking for numa... ${NUMA}")
if (NOT NUMA)
  message (STATUS "Optional numa library is required for memory allocation checks.")
else (NOT NUMA)
    add_definitions (-DHAVE_NUMA="1")
    set (TEST_LINKER_WRAP_OPTIONS
    "-Wl,-wrap,calloc,-wrap,free")
endif (NOT NUMA)
endif()

# Now simply link against gtest as needed. Eg
add_executable(all_tests src/all.c src/eulabeia_json_tests src/start_scan.c src/scan_porgress.c src/crud_progress.c)
target_link_libraries(all_tests ${NUMA} ${CGREEN_LIBRARIES} ${TEST_LINKER_WRAP_OPTIONS} Eulabeia::eulabeia)

add_test(NAME start_scan_fail
    COMMAND all_tests start_scan_returns_error_on_publish_fail)
add_test(NAME start_scan_success
    COMMAND all_tests start_scan_success)
add_test(NAME scan_progress_fail
    COMMAND all_tests scan_progress_failures)
add_test(NAME scan_progress_success
    COMMAND all_tests scan_progress_success)
add_test(NAME json_create_object_success
	COMMAND all_tests create_object_success)
add_test(NAME json_create_message_success
	COMMAND all_tests create_message_success)
add_test(NAME json_create_failure_success
	COMMAND all_tests create_failure_success)
add_test(NAME json_create_hosts_success
	COMMAND all_tests create_hosts_success)
add_test(NAME json_plugins_create_success
	COMMAND all_tests plugins_create_success)
add_test(NAME json_ports_create_success
	COMMAND all_tests ports_create_success)
add_test(NAME got_target 
	COMMAND all_tests got_target)
add_test(NAME got_scan 
	COMMAND all_tests got_scan)
add_test(NAME crud_got_wrong_id
	COMMAND all_tests got_wrong_id)
add_test(NAME crud_got_got_plugin
	COMMAND all_tests got_got_plugin)
