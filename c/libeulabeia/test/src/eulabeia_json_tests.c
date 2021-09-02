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

#include "eulabeia/types.h"
#include <cgreen/cgreen.h>
#include <cgreen/internal/assertions_internal.h>
#include <cgreen/legacy.h>
#include <cgreen/mocks.h>
#include <cgreen/unit.h>

#include <eulabeia/json.h>
#include <stdlib.h>

#if defined(HAVE_NUMA) && defined(CGREEN_NO_FORK)
#include <numa.h>
#endif

Describe(Eulabeia_json);
BeforeEach(Eulabeia_json) {}
AfterEach(Eulabeia_json) {}

/* Json Strings for testing */
#define HOSTS                                                                  \
	"[\"localhost\",\"123.123.123.123\",\"example.com\",\"192.158.56.2/"   \
	"18\"]"
#define MESSAGE_ID "normalID-e07f-11eb-99c4-6b7f958f017"
#define MESSAGE_TYPE "someMsgTypeString"
#define MESSAGE_CREATED "123456790"
#define MESSAGE_GROUP_ID "263cde52-e07f-11eb-99c4-6b7f958f017"
static const gchar *g_json_str = "{\
\"hosts\":" HOSTS ",\
\"message_id\":\"" MESSAGE_ID "\",\
\"message_type\":\"" MESSAGE_TYPE "\",\
\"group_id\":\"" MESSAGE_GROUP_ID "\",\
\"created\":" MESSAGE_CREATED "\
}";

#define ERROR_MESSAGE_TYPE "failure.modify.target"
#define ERROR_MESSAGE_DESCR "nsufficient space"
#define ERROR_MESSAGE_ID "error1e0-768d-4a86-ab6a-dd28e3f45776"
static const gchar *g_json_error_str = "{\
\"id\": \"" ERROR_MESSAGE_ID "\",\
\"created\": " MESSAGE_CREATED ",\
\"message_type\": \"" ERROR_MESSAGE_TYPE "\",\
\"message_id\": \"" MESSAGE_ID "\",\
\"group_id\": \"" MESSAGE_GROUP_ID "\",\
\"error\": \"" ERROR_MESSAGE_DESCR "\"\
}";

#define PLUGINS "[\"pluginOID1\",\"pluginOID2\"]"
static const gchar *g_json_plugins_str = "{\
\"plugins\": " PLUGINS "\
}";

#define PORTS "[\"22\",\"magic/tcp:44:..\",\"\"]"
static const gchar *g_json_ports_str = "{\
\"ports\": " PORTS "\
}";

#define RESULT                                                                 \
	"{\"message_id\":\"fa022daa-1d78-4d02-80b5-83af3086d7d0\","            \
	"\"message_type\":\"result.scan\","                                    \
	"\"group_id\":\"e069f31d-7047-4afb-b31a-65c821c98bad\","               \
	"\"created\":0,"                                                       \
	"\"id\":\"classic_scan_1\","                                           \
	"\"result_type\":\"LOG\","                                             \
	"\"host_ip\":\"127.0.0.1\","                                           \
	"\"host_name\":\"localhost\","                                         \
	"\"port\":\"general/tcp\","                                            \
	"\"oid\":\"1.3.6.1.4.1.25623.1.0.90022\","                             \
	"\"value\":\"this is a log message\n\","                               \
	"\"uri\":\"\"}"

#if defined(HAVE_NUMA) && defined(CGREEN_NO_FORK)
/* Track total size of allocated mem */
int g_alloc_cnt = 0;
/* Track total size of freed mem */
int g_free_cnt = 0;
/* Union for tracking size of allocated mem */
union track_size {
	size_t size;
	max_align_t p;
};

/**
 * @brief Alloc same amount of mem as call to calloc but save mem size in union
 * add it to g_alloc_cnt.
 *
 * @param nmemb	How many elements to alloc
 * @param size	Size of single element
 *
 * @return pointer to allocated memory.
 */
void *my_calloc(size_t nmemb, size_t size)
{
	size_t save_size = nmemb * size;
	union track_size *p =
	    numa_alloc_local(sizeof(union track_size *) + nmemb * size);
	if (p) {
		p->size = save_size;
		g_alloc_cnt += save_size;
		p++;
	}
	return p;
}

/**
 * @brief Free block of allocated mem and add size of freed mem to g_free_cnt.
 *
 * @param p	Pointer to allocated mem
 */
void my_free(void *p)
{
	if (p) {
		union track_size *ptr = p;
		ptr--;
		g_free_cnt += ptr->size;
		numa_free(ptr, sizeof(union track_size *) + ptr->size);
	}
}

/* Wrap calloc */
__attribute__((weak)) void *__real_calloc(size_t nmemb, size_t size);
gboolean g_calloc_use_real = TRUE;
void *__wrap_calloc(size_t nmemb, size_t size)
{
	if (g_calloc_use_real)
		return __real_calloc(nmemb, size);

	return my_calloc(nmemb, size);
}

/* Wrap free */
__attribute__((weak)) void __real_free(void *p);
gboolean g_free_use_real = TRUE;
void __wrap_free(void *p)
{
	if (g_free_use_real)
		return __real_free(p);

	return my_free(p);
}
#endif /* HAVE_NUMA && CGREEN_NO_FORK */

/**
 * @brief init global memory size counters and use mem wrapper
 */
static void activate_mem_tracking(void)
{
#if defined(HAVE_NUMA) && defined(CGREEN_NO_FORK)
	g_alloc_cnt = 0;
	g_free_cnt = 0;
	g_calloc_use_real = FALSE;
	g_free_use_real = FALSE;
#endif
}

/**
 * @brief init global memory size counters and use mem wrapper
 *
 */
static void deactivate_mem_tracking(void)
{
#if defined(HAVE_NUMA) && defined(CGREEN_NO_FORK)
	g_alloc_cnt = 0;
	g_free_cnt = 0;
	g_calloc_use_real = TRUE;
	g_free_use_real = TRUE;
#endif
}

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wunused-parameter"
static void ensure_alloc_equal_free(const char *call_func)
{
#if defined(HAVE_NUMA) && defined(CGREEN_NO_FORK)
	g_warning("%s: (%d:%d)", call_func, g_alloc_cnt, g_free_cnt);
	assert_equal_with_message(
	    g_alloc_cnt,
	    g_free_cnt,
	    "%s: Expected alloced mem (%d) to be equal to freed mem (%d)",
	    call_func,
	    g_alloc_cnt,
	    g_free_cnt);
#endif
}
#pragma GCC diagnostic pop

Ensure(Eulabeia_json, create_object_success)
{
	JsonNode *j_node = NULL;
	JsonObject *j_obj;

	assert_equal(eulabeia_json_object(g_json_str, &j_node, &j_obj), 0);
	assert_not_equal(j_node, NULL);
	assert_not_equal(j_obj, NULL);

	if (j_node)
		json_node_free(j_node);
}

Ensure(Eulabeia_json, create_message_success)
{
	int err;
	JsonNode *j_node = NULL;
	JsonObject *j_obj;
	struct EulabeiaMessage *eulabeia_message = NULL;

	assert_equal(err = eulabeia_json_object(g_json_str, &j_node, &j_obj),
		     0);
	if (err)
		goto clean_exit;
	assert_not_equal(j_node, NULL);
	assert_not_equal(j_obj, NULL);

	assert_equal(err = eulabeia_json_message(j_obj, &eulabeia_message), 0);
	if (err)
		goto clean_exit;

	assert_string_equal(eulabeia_message->type, MESSAGE_TYPE);
	assert_string_equal(eulabeia_message->id, MESSAGE_ID);
	assert_string_equal(eulabeia_message->group_id, MESSAGE_GROUP_ID);
	assert_equal(eulabeia_message->created,
		     (unsigned long)atoll(MESSAGE_CREATED));

clean_exit:
	if (j_node)
		json_node_free(j_node);
	if (eulabeia_message)
		eulabeia_message_destroy(&eulabeia_message);
}

Ensure(Eulabeia_json, create_failure_success)
{
	int err;
	JsonNode *j_node = NULL;
	JsonObject *j_obj;
	struct EulabeiaMessage *eulabeia_message = NULL;
	struct EulabeiaFailure *eulabeia_failure = NULL;

	assert_equal(
	    err = eulabeia_json_object(g_json_error_str, &j_node, &j_obj), 0);
	if (err)
		goto clean_exit;
	assert_not_equal(j_node, NULL);
	assert_not_equal(j_obj, NULL);

	assert_equal(err = eulabeia_json_message(j_obj, &eulabeia_message), 0);
	if (err)
		goto clean_exit;
	assert_string_equal(eulabeia_message->type, ERROR_MESSAGE_TYPE);
	assert_string_equal(eulabeia_message->id, MESSAGE_ID);
	assert_string_equal(eulabeia_message->group_id, MESSAGE_GROUP_ID);
	assert_equal(eulabeia_message->created,
		     (unsigned long)atoll(MESSAGE_CREATED));

	assert_equal(err = eulabeia_json_failure(
			 j_obj, eulabeia_message, &eulabeia_failure),
		     0);
	if (err)
		goto clean_exit;

	assert_string_equal(eulabeia_failure->message->type,
			    ERROR_MESSAGE_TYPE);
	assert_string_equal(eulabeia_failure->message->id, MESSAGE_ID);
	assert_string_equal(eulabeia_failure->message->group_id,
			    MESSAGE_GROUP_ID);
	assert_equal(eulabeia_failure->message->created,
		     (unsigned long)atoll(MESSAGE_CREATED));
	assert_string_equal(eulabeia_failure->id, ERROR_MESSAGE_ID);
	assert_string_equal(eulabeia_failure->error, ERROR_MESSAGE_DESCR);

clean_exit:
	if (j_node)
		json_node_free(j_node);
	if (eulabeia_message)
		eulabeia_message_destroy(&eulabeia_message);
	if (eulabeia_failure)
		eulabeia_failure_destroy(&eulabeia_failure);
}

Ensure(Eulabeia_json, create_hosts_success)
{
	int err;
	JsonNode *j_node = NULL;
	JsonObject *j_obj;
	JsonArray *host_arr;
	struct EulabeiaHosts *eulabeia_hosts = NULL;

	assert_equal(err = eulabeia_json_object(g_json_str, &j_node, &j_obj),
		     0);
	if (err)
		goto clean_exit;

	host_arr = json_object_get_array_member(j_obj, "hosts");
	assert_not_equal(host_arr, NULL);
	assert_equal(json_array_get_length(host_arr), 4);

	activate_mem_tracking();
	eulabeia_json_hosts(host_arr, &eulabeia_hosts);
	assert_equal(eulabeia_hosts->len, 4);
	assert_equal(eulabeia_hosts->cap, 4);
	assert_string_equal(eulabeia_hosts->hosts[0].address, "localhost");
	assert_string_equal(eulabeia_hosts->hosts[3].address,
			    "192.158.56.2/18");
	if (eulabeia_hosts) {
		eulabeia_hosts_destroy(&eulabeia_hosts);
	}
	ensure_alloc_equal_free(__func__);
	deactivate_mem_tracking();

clean_exit:
	if (j_node)
		json_node_free(j_node);
}

Ensure(Eulabeia_json, plugins_create_success)
{
	int err;
	JsonObject *j_obj;
	JsonArray *plugin_arr;

	JsonNode *j_node = NULL;
	struct EulabeiaPlugins *eulabeia_plugins = NULL;

	assert_equal(
	    err = eulabeia_json_object(g_json_plugins_str, &j_node, &j_obj), 0);
	if (err)
		goto clean_exit;

	plugin_arr = json_object_get_array_member(j_obj, "plugins");
	assert_not_equal(plugin_arr, NULL);
	assert_equal(json_array_get_length(plugin_arr), 2);

	activate_mem_tracking();
	eulabeia_json_plugins(plugin_arr, &eulabeia_plugins);
	assert_equal(eulabeia_plugins->len, 2);
	assert_equal(eulabeia_plugins->cap, 2);
	assert_string_equal(eulabeia_plugins->plugins[0].oid, "pluginOID1");
	assert_string_equal(eulabeia_plugins->plugins[1].oid, "pluginOID2");
	if (eulabeia_plugins)
		eulabeia_plugins_destroy(&eulabeia_plugins);
	ensure_alloc_equal_free(__func__);
	deactivate_mem_tracking();

clean_exit:
	if (j_node)
		json_node_free(j_node);
}

Ensure(Eulabeia_json, ports_create_success)
{
	int err;
	int rc;
	JsonNode *j_node = NULL;
	JsonObject *j_obj;
	JsonArray *port_arr;

	struct EulabeiaPorts *eulabeia_ports = NULL;

	assert_equal(
	    err = eulabeia_json_object(g_json_ports_str, &j_node, &j_obj), 0);
	if (err)
		goto clean_exit;

	assert_true(rc = json_object_has_member(j_obj, "ports"));
	if (!rc)
		goto clean_exit;

	port_arr = json_object_get_array_member(j_obj, "ports");
	assert_not_equal(port_arr, NULL);
	assert_equal(json_array_get_length(port_arr), 3);

	activate_mem_tracking();
	eulabeia_json_ports(port_arr, &eulabeia_ports);
	assert_equal(eulabeia_ports->len, 3);
	assert_equal(eulabeia_ports->cap, 3);
	assert_string_equal(eulabeia_ports->ports[0].port, "22");
	assert_string_equal(eulabeia_ports->ports[1].port, "magic/tcp:44:..");
	assert_string_equal(eulabeia_ports->ports[2].port, "");
	if (eulabeia_ports)
		eulabeia_ports_destroy(&eulabeia_ports);
	ensure_alloc_equal_free(__func__);
	deactivate_mem_tracking();

clean_exit:
	if (j_node)
		json_node_free(j_node);
}

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-prototypes"
TestSuite *eulabeia_json_tests()
{
	TestSuite *suite = create_test_suite();
	add_test_with_context(suite, Eulabeia_json, create_object_success);
	add_test_with_context(suite, Eulabeia_json, create_message_success);
	add_test_with_context(suite, Eulabeia_json, create_failure_success);
	add_test_with_context(suite, Eulabeia_json, create_hosts_success);
	add_test_with_context(suite, Eulabeia_json, plugins_create_success);
	add_test_with_context(suite, Eulabeia_json, ports_create_success);
	return suite;
}
#pragma GCC diagnostic pop
