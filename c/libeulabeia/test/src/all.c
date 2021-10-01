/* Copyright (C) 2021 Greenbone Networks GmbH
 *
 * SPDX-License-Identifier: GPL-2.0-or-later
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA.
 */

#include <cgreen/cgreen.h>
TestSuite *publish_tests(void);
TestSuite *progress_tests(void);
TestSuite *eulabeia_json_tests(void);
TestSuite *crud_progress_tests(void);

int main(int argc, char **argv)
{
	TestSuite *suite = create_test_suite();
	add_suite(suite, publish_tests());
	add_suite(suite, progress_tests());
	add_suite(suite, eulabeia_json_tests());
	add_suite(suite, crud_progress_tests());
	if (argc > 1) {
		return run_single_test(suite, argv[1], create_text_reporter());
	}
	return run_test_suite(suite, create_text_reporter());
}
