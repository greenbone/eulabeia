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
#include <cgreen/unit.h>
#include <eulabeia/client.h>
#include <stdio.h>

Describe(Publish);
BeforeEach(Publish) {}
AfterEach(Publish) {}

#define SUCCESS 1
#define FAILURE 0

int publish(const char *topic, const char *message, void *context)
{
	int *c = (int *)context;
	return *c == SUCCESS ? 0 : -42;
}

Ensure(Publish, start_scan_returns_error_on_publish_fail)
{

	struct EulabeiaClient *ec = calloc(1, sizeof(struct EulabeiaClient));
	ec->publish = publish;
	struct EulabeiaScan *scan = calloc(1, sizeof(struct EulabeiaScan));
	int fail = FAILURE;
	ec->context = &fail;
	int rc;
	rc = eulabeia_start_scan(ec, NULL, (void *)&fail);
	assert_equal_with_message(
	    rc,
	    -1,
	    "expected error code [%d] to be -1 because scan is NULL",
	    rc);
	rc = eulabeia_start_scan(ec, scan, (void *)&fail);
	assert_equal_with_message(rc,
				  -2,
				  "expected error code [%d] to be -2 because "
				  "scan is set but without an id",
				  rc);
	scan->id = "set";
	rc = eulabeia_start_scan(ec, scan, (void *)&fail);
	assert_equal_with_message(
	    rc,
	    -3,
	    "expected error code [%d] to be -3 because publish fails",
	    rc);
	free(ec);
	free(scan);
}
Ensure(Publish, start_scan_success)
{
	struct EulabeiaClient *ec = calloc(1, sizeof(struct EulabeiaClient));
	ec->publish = publish;
	struct EulabeiaScan *scan = calloc(1, sizeof(struct EulabeiaScan));
	int success = SUCCESS;
	ec->context = &success;
	int rc;
	scan->id = "set";
	rc = eulabeia_start_scan(ec, scan, NULL);
	assert_equal_with_message(
	    rc,
	    0,
	    "expected return code [%d] to be 0 because publish succeeded",
	    rc);
	free(ec);
	free(scan);
}

TestSuite *publish_tests()
{
	TestSuite *suite = create_test_suite();
	add_test_with_context(
	    suite, Publish, start_scan_returns_error_on_publish_fail);
	add_test_with_context(suite, Publish, start_scan_success);
	return suite;
}
