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
#include <cgreen/constraint_syntax_helpers.h>
#include <cgreen/internal/assertions_internal.h>
#include <cgreen/legacy.h>
#include <cgreen/mocks.h>
#include <cgreen/unit.h>
#include "got_vts_json.h"
#include <eulabeia/client.h>
#include <stdlib.h>

Describe(CRUDProgress);
BeforeEach(CRUDProgress) {}
AfterEach(CRUDProgress) {}

#define GOT_TARGET_ID "target"
#define GOT_TARGET                                                             \
	"{\"message_id\":\"fa022daa-1d78-4d02-80b5-83af3086d7d0\","            \
	"\"message_type\":\"got.target\","                                     \
	"\"group_id\":\"e069f31d-7047-4afb-b31a-65c821c98bad\","               \
	"\"message_created\":0,"                                                       \
	"\"id\":\""GOT_TARGET_ID"\"}"

#define GOT_SCAN_ID "scan"
#define GOT_SCAN\
	"{\"message_id\":\"fa022daa-1d78-4d02-80b5-83af3086d7d0\","            \
	"\"message_type\":\"got.scan\","                                     \
	"\"group_id\":\"e069f31d-7047-4afb-b31a-65c821c98bad\","               \
	"\"message_created\":0,"                                                       \
	"\"id\":\""GOT_SCAN_ID"\"}"
Ensure(CRUDProgress, got_target)
{
	int rc;
	struct EulabeiaCRUDProgress *progress = calloc(1, sizeof(*progress));
	rc = eulabeia_crud_progress(GOT_TARGET, GOT_TARGET_ID, EULABEIA_INFO_GOT, progress);
	assert_equal(rc, 0);
	assert_equal(progress->status, EULABEIA_CRUD_SUCCESS);
	assert_that(progress->target, is_non_null);
	assert_that(progress->target->id, is_equal_to_string(GOT_TARGET_ID));

	free(progress->target);
	free(progress);
}

Ensure(CRUDProgress, got_scan)
{
	int rc;
	struct EulabeiaCRUDProgress *progress = calloc(1, sizeof(*progress));
	rc = eulabeia_crud_progress(GOT_SCAN, GOT_SCAN_ID, EULABEIA_INFO_GOT, progress);
	assert_equal(rc, 0);
	assert_equal(progress->status, EULABEIA_CRUD_SUCCESS);
	assert_that(progress->scan, is_non_null);
	assert_that(progress->scan->id, is_equal_to_string(GOT_SCAN_ID));

	free(progress->scan);
	free(progress);
}

Ensure(CRUDProgress, got_plugin)
{
	int rc;
	struct EulabeiaCRUDProgress *progress = calloc(1, sizeof(*progress));
	rc = eulabeia_crud_progress(GOT_VT, GOT_VT_ID, EULABEIA_INFO_GOT, progress);
	assert_equal(rc, 0);
	assert_equal(progress->status, EULABEIA_CRUD_SUCCESS);
	assert_that(progress->plugin, is_non_null);
	assert_that(progress->plugin->oid, is_equal_to_string(GOT_VT_ID));

	free(progress->plugin);
	free(progress);
}
Ensure(CRUDProgress, got_wrong_id)
{
	
	int rc;
	struct EulabeiaCRUDProgress *progress = calloc(1, sizeof(*progress));
	rc = eulabeia_crud_progress(GOT_SCAN, "waldbusch", EULABEIA_INFO_GOT, progress);
	assert_that(rc, is_equal_to(1));
	assert_that(progress->scan, is_null);
	free(progress);
}

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-prototypes"
TestSuite *crud_progress_tests()
{
	TestSuite *suite = create_test_suite();
	add_test_with_context(suite, CRUDProgress, got_target);
	add_test_with_context(suite, CRUDProgress, got_scan);
	add_test_with_context(suite, CRUDProgress, got_wrong_id);
	add_test_with_context(suite, CRUDProgress, got_plugin);
	return suite;
}
#pragma GCC diagnostic pop
