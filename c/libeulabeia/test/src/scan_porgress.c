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
#include <cgreen/legacy.h>
#include <cgreen/unit.h>
#include <eulabeia/client.h>
#include <stdio.h>

Describe(Progress);
BeforeEach(Progress) {}
AfterEach(Progress) {}

#define SUCCESS 1
#define FAILURE 0

int progress(const char *topic, const char *message, void *context)
{
	int *c = (int *)context;
	return *c == SUCCESS ? 0 : -42;
}

Ensure(Progress, scan_progress_failures)
{
	struct EulabeiaScanProgress progress;
	int rc;
	rc = eulabeia_scan_progress(NULL, "", &progress);
	assert_equal_with_message(
	    rc, -1, "expected [%d] to be -1 because payload is NULL", rc);
	rc = eulabeia_scan_progress("", "", NULL);
	assert_equal_with_message(
	    rc, -1, "expected [%d] to be -1 because progress is NULL", rc);
	rc = eulabeia_scan_progress("", NULL, &progress);
	assert_equal_with_message(
	    rc, -1, "expected [%d] to be -1 because id is NULL", rc);
	rc = eulabeia_scan_progress("{ not a json", "wanted", &progress);
	assert_equal_with_message(
	    rc,
	    -2,
	    "expected [%d] to be -2 because paylod is invalid json",
	    rc);
	rc = eulabeia_scan_progress(
	    "\"not a json object\"", "wanted", &progress);
	assert_equal_with_message(
	    rc,
	    -3,
	    "expected [%d] to be -3 because payload is not a json-object",
	    rc);
	rc =
	    eulabeia_scan_progress("{\"id\":\"invalid\"}", "wanted", &progress);
	assert_equal_with_message(
	    rc,
	    -4,
	    "expected [%d] to be -4 because payload is not EulabeiaMessage",
	    rc);
	rc = eulabeia_scan_progress("{"
				    "\"message_id\": \"1\","
				    "\"message_type\":\"status.scan\","
				    "\"group_id\":null,"
				    "\"created\": 42,"
				    "\"id\": \"wanted\","
				    "\"status\": null"
				    "}",
				    "wanted",
				    &progress);
	assert_equal_with_message(
	    rc, -5, "expected [%d] to be -5 because status is NULL", rc);
	rc = eulabeia_scan_progress("{"
				    "\"message_id\": \"1\","
				    "\"message_type\":\"status.scan\","
				    "\"group_id\":null,"
				    "\"created\": 42,"
				    "\"id\": \"wanted\","
				    "\"status\": \"unknown\""
				    "}",
				    "wanted",
				    &progress);
	assert_equal_with_message(
	    rc, -5, "expected [%d] to be -5 because status is NULL", rc);
	rc = eulabeia_scan_progress("{"
				    "\"message_id\": \"1\","
				    "\"message_type\":\"status.scan\","
				    "\"group_id\":null,"
				    "\"created\": 42,"
				    "\"id\": null,"
				    "\"status\": \"unknown\""
				    "}",
				    "wanted",
				    &progress);
	assert_equal_with_message(rc,
				  3,
				  "expected [%d] to be 2 because without an ID "
				  "is not a known status message",
				  rc);
}

Ensure(Progress, scan_progress_success)
{
	struct EulabeiaScanProgress progress;
	int rc;
	char *j;
	j = calloc(1, 1024);

#define X(a, b)                                                                \
	snprintf(j,                                                            \
		 1024,                                                         \
		 "{"                                                           \
		 "\"message_id\": \"1\","                                      \
		 "\"message_type\":\"status.scan\","                           \
		 "\"group_id\":null,"                                          \
		 "\"created\": 42,"                                            \
		 "\"id\": \"wanted\","                                         \
		 "\"status\": \"%s\""                                          \
		 "}",                                                          \
		 (#b));                                                        \
	rc = eulabeia_scan_progress(j, "wanted", &progress);                   \
	assert_equal_with_message(                                             \
	    rc, 0, "expected [%d] to be 0 on %s", rc, j);                      \
	assert_equal_with_message(progress.status,                             \
				  a,                                           \
				  "expected [%d] to be %d",                    \
				  progress.status,                             \
				  (a));
	EULABEIA_SCAN_STATES
#undef X
	snprintf(j,
		 1024,
		 "{"
		 "\"message_id\": \"1\","
		 "\"message_type\":\"failure.start.scan\","
		 "\"group_id\":null,"
		 "\"created\": 42,"
		 "\"id\": \"wanted\","
		 "\"error\": \"%s\""
		 "}",
		 "scan id not found");
	rc = eulabeia_scan_progress(j, "wanted", &progress);
	assert_equal_with_message(
	    progress.status,
	    EULABEIA_SCAN_FAILED,
	    "expected %s (%d) to be %s (%d)",
	    eulabeia_scan_state_to_str(progress.status),
	    progress.status,
	    eulabeia_scan_state_to_str(EULABEIA_SCAN_FAILED),
	    EULABEIA_SCAN_FAILED);

	free(j);
}

TestSuite *progress_tests()
{
	TestSuite *suite = create_test_suite();
	add_test_with_context(suite, Progress, scan_progress_failures);
	add_test_with_context(suite, Progress, scan_progress_success);
	return suite;
}
