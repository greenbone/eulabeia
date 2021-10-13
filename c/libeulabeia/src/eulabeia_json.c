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
#include <eulabeia/client.h>
#include <eulabeia/json.h>

#include <json-glib/json-glib.h>
#include <stdlib.h>

volatile int pretty_print = 0;

void eulabeia_json_set_pretty_print(int i) { pretty_print = i; }

static int json_object_get_and_assign_string(JsonObject *obj,
					     const char *key,
					     char **value)
{
	const char *jvalue;

	if ((jvalue = json_object_get_string_member(obj, key)) == NULL) {
		return 1;
	}

	if ((*value = g_strdup(jvalue)) == NULL)
		return -1;

	return 0;
}

int eulabeia_json_message(JsonObject *obj, struct EulabeiaMessage **msg)
{
	if (!(json_object_has_member(obj, "message_id") &&
	      json_object_has_member(obj, "message_type") &&
	      json_object_has_member(obj, "message_created") &&
	      json_object_has_member(obj, "group_id"))) {
		return -1;
	}
	if (*msg == NULL &&
	    (*msg = calloc(1, sizeof(struct EulabeiaMessage))) == NULL) {
		return -2;
	}
	(*msg)->created = json_object_get_int_member(obj, "message_created");
	if (json_object_get_and_assign_string(obj, "message_id", &(*msg)->id) <
	    0)
		return -3;
	if (json_object_get_and_assign_string(
		obj, "message_type", &(*msg)->type) < 0)
		return -4;
	if (json_object_get_and_assign_string(
		obj, "group_id", &(*msg)->group_id) < 0)
		return -5;
	return 0;
}

int eulabeia_json_failure(JsonObject *obj,
			  struct EulabeiaMessage *msg,
			  struct EulabeiaFailure **failure)
{
	enum eulabeia_message_type failure_type;
	if (msg == NULL || msg->type == NULL)
		return -1;
	failure_type = eulabeia_message_to_message_type(msg);
	switch (failure_type) {
	case EULABEIA_INFO_FAILURE:
	case EULABEIA_INFO_MODIFY_FAILURE:
	case EULABEIA_INFO_START_FAILURE:
	case EULABEIA_INFO_STOP_FAILURE:
		break;
	default:
		return -2;
	}

	if (!(json_object_has_member(obj, "id") &&
	      json_object_has_member(obj, "error"))) {

		return -3;
	}
	if (*failure == NULL &&
	    (*failure = calloc(1, sizeof(struct EulabeiaFailure))) == NULL) {
		return -4;
	}
	(*failure)->message = msg;
	// even when explicitely set to id is not allowed to be null
	if (json_object_get_and_assign_string(obj, "id", &(*failure)->id) != 0)
		return -5;
	if (json_object_get_and_assign_string(
		obj, "error", &(*failure)->error) < 0)
		return -6;
	return 0;
}

int eulabeia_json_id_message(JsonObject *obj,
			     enum eulabeia_message_type type,
			     struct EulabeiaMessage *msg,
			     struct EulabeiaIDMessage **idmessage)
{
	if (msg == NULL || msg->type == NULL)
		return -1;
	if (eulabeia_message_to_message_type(msg) != type)
		return -2;

	if (!json_object_has_member(obj, "id")) {

		return -3;
	}
	if (*idmessage == NULL &&
	    (*idmessage = calloc(1, sizeof(struct EulabeiaIDMessage))) ==
		NULL) {
		return -4;
	}
	(*idmessage)->message = msg;
	// even when explicitely set to id is not allowed to be null
	if (json_object_get_and_assign_string(obj, "id", &(*idmessage)->id) !=
	    0)
		return -5;
	return 0;
}
int eulabeia_json_scan_result(JsonObject *obj,
			      struct EulabeiaMessage *msg,
			      struct EulabeiaScanResult **scan_result)
{
	if (msg == NULL || msg->type == NULL)
		return -1;
	if (eulabeia_message_to_message_type(msg) != EULABEIA_INFO_SCAN_RESULT)
		return -2;
	if (!(json_object_has_member(obj, "oid") &&
	      json_object_has_member(obj, "host_ip") &&
	      json_object_has_member(obj, "host_name") &&
	      json_object_has_member(obj, "port") &&
	      json_object_has_member(obj, "value") &&
	      json_object_has_member(obj, "uri") &&
	      json_object_has_member(obj, "result_type"))) {
		return -3;
	}
	if (*scan_result == NULL &&
	    (*scan_result = calloc(1, sizeof(struct EulabeiaScanResult))) ==
		NULL)
		return -4;

	(*scan_result)->message = msg;
	if (json_object_get_and_assign_string(obj, "id", &(*scan_result)->id) !=
	    0)
		return -5;
	if (json_object_get_and_assign_string(
		obj, "oid", &(*scan_result)->oid) != 0)
		return -6;
	if (json_object_get_and_assign_string(
		obj, "result_type", &(*scan_result)->result_type) < 0)
		return -7;
	if (json_object_get_and_assign_string(
		obj, "host_ip", &(*scan_result)->host_ip) < 0)
		return -8;
	if (json_object_get_and_assign_string(
		obj, "host_name", &(*scan_result)->host_name) < 0)
		return -9;
	if (json_object_get_and_assign_string(
		obj, "port", &(*scan_result)->port) < 0)
		return -10;
	if (json_object_get_and_assign_string(
		obj, "value", &(*scan_result)->value) < 0)
		return -11;
	if (json_object_get_and_assign_string(
		obj, "uri", &(*scan_result)->uri) < 0)
		return -12;

	return 0;
}
int eulabeia_json_status(JsonObject *obj,
			 struct EulabeiaMessage *msg,
			 struct EulabeiaStatus **status)
{
	if (msg == NULL || msg->type == NULL)
		return -1;
	if (eulabeia_message_to_message_type(msg) != EULABEIA_INFO_STATUS)
		return -2;

	if (!(json_object_has_member(obj, "id") &&
	      json_object_has_member(obj, "status"))) {

		return -3;
	}
	if (*status == NULL &&
	    (*status = calloc(1, sizeof(struct EulabeiaStatus))) == NULL) {
		return -4;
	}
	(*status)->message = msg;
	// even when explicitely set to id is not allowed to be null
	if (json_object_get_and_assign_string(obj, "id", &(*status)->id) != 0)
		return -5;
	if (json_object_get_and_assign_string(
		obj, "status", &(*status)->status) < 0)
		return -6;
	return 0;
}

int eulabeia_json_hosts(JsonArray *arr, struct EulabeiaHosts **hosts)
{
	const char *a;
	unsigned int arr_len;
	if (*hosts == NULL) {
		*hosts = calloc(1, sizeof(struct EulabeiaHosts));
	}
	arr_len = json_array_get_length(arr);
	if (((*hosts)->hosts = calloc(arr_len, sizeof(struct EulabeiaHost))) ==
	    NULL) {
		return -1;
	}
	(*hosts)->cap = arr_len;
	(*hosts)->len = arr_len;
	for (unsigned int index = 0; index < arr_len; index++) {
		a = json_array_get_string_element(arr, index);
		(*hosts)->hosts[index].address = g_strdup(a);
	}

	return 0;
}

// TODO credentials
static int json_plugin_dependencies(JsonArray *ja,
				    struct EulabeiaPluginDependencies **r)
{
	unsigned int arr_len, i;
	struct EulabeiaPluginDependency *ri;
	if (ja == NULL || r == NULL)
		return -1;
	if (*r == NULL)
		*r = g_malloc0(sizeof(**r));

	arr_len = json_array_get_length(ja);
	if (((*r)->dependency = calloc(arr_len, sizeof(*(*r)->dependency))) ==
	    NULL) {
		return -1;
	}
	(*r)->cap = arr_len;
	(*r)->len = arr_len;
	for (i = 0; i < arr_len; i++) {
		ri = (*r)->dependency + i;
		ri->filename = g_strdup(json_array_get_string_element(ja, i));
	}

	return 0;
}

static int json_plugin_references(JsonArray *ja,
				  struct EulabeiaPluginReferences **r)
{
	unsigned int arr_len, i;
	struct EulabeiaPluginReference *ri;
	JsonObject *jo;
	if (ja == NULL || r == NULL)
		return -1;
	if (*r == NULL)
		*r = g_malloc0(sizeof(**r));

	arr_len = json_array_get_length(ja);
	if (((*r)->reference = calloc(arr_len, sizeof(*(*r)->reference))) ==
	    NULL) {
		return -1;
	}
	(*r)->cap = arr_len;
	(*r)->len = arr_len;
	for (i = 0; i < arr_len; i++) {
		ri = (*r)->reference + i;
		jo = json_array_get_object_element(ja, i);
		if (json_object_has_member(jo, "id"))
			json_object_get_and_assign_string(jo, "id", &ri->id);
		if (json_object_has_member(jo, "type"))
			json_object_get_and_assign_string(
			    jo, "type", &ri->type);
	}

	return 0;
}

static int json_plugin_parameters(JsonArray *ja,
				  struct EulabeiaPluginParameters **r)
{
	unsigned int arr_len, i;
	struct EulabeiaPluginParameter *ri;
	JsonObject *jo;
	if (ja == NULL || r == NULL)
		return -1;
	if (*r == NULL)
		*r = g_malloc0(sizeof(**r));

	arr_len = json_array_get_length(ja);
	if (((*r)->parameter = calloc(arr_len, sizeof(*(*r)->parameter))) ==
	    NULL) {
		return -1;
	}
	(*r)->cap = arr_len;
	(*r)->len = arr_len;
	for (i = 0; i < arr_len; i++) {
		ri = (*r)->parameter + i;
		jo = json_array_get_object_element(ja, i);
		if (json_object_has_member(jo, "id"))
			json_object_get_and_assign_string(jo, "id", &ri->id);
		if (json_object_has_member(jo, "name"))
			json_object_get_and_assign_string(
			    jo, "name", &ri->name);
		if (json_object_has_member(jo, "value"))
			json_object_get_and_assign_string(
			    jo, "value", &ri->value);
		if (json_object_has_member(jo, "type"))
			json_object_get_and_assign_string(
			    jo, "type", &ri->type);
		if (json_object_has_member(jo, "description"))
			json_object_get_and_assign_string(
			    jo, "description", &ri->description);
		if (json_object_has_member(jo, "default"))
			json_object_get_and_assign_string(
			    jo, "default", &ri->defaultvalue);
	}

	return 0;
}

int eulabeia_json_plugin(JsonObject *jo, struct EulabeiaPlugin **p)
{
	JsonObject *severity;
	if (p == NULL || jo == NULL)
		return -1;
	if (*p == NULL)
		*p = g_malloc0(sizeof(**p));

	// a plugin needs at least an oid or a family; when both are null we
	// cannot identify it.
	if (!json_object_has_member(jo, "id") &&
	    !json_object_has_member(jo, "oid") &&
	    !json_object_has_member(jo, "family"))
		return -2;

	if (json_object_has_member(jo, "id"))
		json_object_get_and_assign_string(jo, "id", &(*p)->oid);
	if (json_object_has_member(jo, "oid"))
		json_object_get_and_assign_string(jo, "oid", &(*p)->oid);
	if (json_object_has_member(jo, "name"))
		json_object_get_and_assign_string(jo, "name", &(*p)->name);
	if (json_object_has_member(jo, "filename"))
		json_object_get_and_assign_string(
		    jo, "filename", &(*p)->filename);
	if (json_object_has_member(jo, "required_keys"))
		json_object_get_and_assign_string(
		    jo, "required_keys", &(*p)->required_keys);
	if (json_object_has_member(jo, "mandatory_keys"))
		json_object_get_and_assign_string(
		    jo, "mandatory_keys", &(*p)->mandatory_keys);
	if (json_object_has_member(jo, "excluded_keys"))
		json_object_get_and_assign_string(
		    jo, "excluded_keys", &(*p)->excluded_keys);
	if (json_object_has_member(jo, "required_ports"))
		json_object_get_and_assign_string(
		    jo, "required_ports", &(*p)->required_ports);
	if (json_object_has_member(jo, "required_udp_ports"))
		json_object_get_and_assign_string(
		    jo, "required_udp_ports", &(*p)->required_udp_ports);
	if (json_object_has_member(jo, "category"))
		json_object_get_and_assign_string(
		    jo, "category", &(*p)->category);
	if (json_object_has_member(jo, "family"))
		json_object_get_and_assign_string(jo, "family", &(*p)->family);
	if (json_object_has_member(jo, "created"))
		json_object_get_and_assign_string(
		    jo, "created", &(*p)->created);
	if (json_object_has_member(jo, "modified"))
		json_object_get_and_assign_string(
		    jo, "modified", &(*p)->modified);
	if (json_object_has_member(jo, "summary"))
		json_object_get_and_assign_string(
		    jo, "summary", &(*p)->summary);
	if (json_object_has_member(jo, "solution"))
		json_object_get_and_assign_string(
		    jo, "solution", &(*p)->solution);
	if (json_object_has_member(jo, "solution_method"))
		json_object_get_and_assign_string(
		    jo, "solution_method", &(*p)->solution_method);
	if (json_object_has_member(jo, "solution_type"))
		json_object_get_and_assign_string(
		    jo, "solution_type", &(*p)->solution_type);
	if (json_object_has_member(jo, "impact"))
		json_object_get_and_assign_string(jo, "impact", &(*p)->impact);
	if (json_object_has_member(jo, "insight"))
		json_object_get_and_assign_string(
		    jo, "insight", &(*p)->insight);
	if (json_object_has_member(jo, "affected"))
		json_object_get_and_assign_string(
		    jo, "affected", &(*p)->affected);
	if (json_object_has_member(jo, "vuldetect"))
		json_object_get_and_assign_string(
		    jo, "vuldetect", &(*p)->vuldetect);
	if (json_object_has_member(jo, "qod_type"))
		json_object_get_and_assign_string(
		    jo, "qod_type", &(*p)->qod_type);
	if (json_object_has_member(jo, "qod"))
		json_object_get_and_assign_string(jo, "qod", &(*p)->qod);
	// TODO add references, dependencies
	if (json_object_has_member(jo, "vt_parameters")) {
		if (json_plugin_parameters(
			json_object_get_array_member(jo, "vt_parameters"),
			&(*p)->parameters) != 0)
			return -2;
	}
	if (json_object_has_member(jo, "vt_dependencies")) {
		if (json_plugin_dependencies(
			json_object_get_array_member(jo, "vt_dependencies"),
			&(*p)->dependencies) != 0)
			return -2;
	}
	if (json_object_has_member(jo, "references")) {
		if (json_plugin_references(
			json_object_get_array_member(jo, "references"),
			&(*p)->references) != 0)
			return -2;
	}
	if (json_object_has_member(jo, "severety")) {
		severity = json_object_get_object_member(jo, "severety");
		(*p)->severity = g_malloc0(sizeof(*(*p)->severity));
		if (json_object_has_member(severity, "severity_vector"))
			json_object_get_and_assign_string(
			    severity,
			    "severity_vector",
			    &(*p)->severity->vector);
		if (json_object_has_member(severity, "severity_type"))
			json_object_get_and_assign_string(
			    severity, "severity_type", &(*p)->severity->type);
		if (json_object_has_member(severity, "severity_date"))
			json_object_get_and_assign_string(
			    severity, "severity_date", &(*p)->severity->date);
		if (json_object_has_member(severity, "severity_origin"))
			json_object_get_and_assign_string(
			    severity,
			    "severity_origin",
			    &(*p)->severity->origin);
	}

	return 0;
}

int eulabeia_json_plugins(JsonArray *arr, struct EulabeiaPlugins **plugins)
{
	JsonObject *jo;
	struct EulabeiaPlugin *p;
	unsigned int arr_len, i;
	if (*plugins == NULL) {
		*plugins = calloc(1, sizeof(struct EulabeiaPlugins));
	}
	arr_len = json_array_get_length(arr);
	if (((*plugins)->plugins =
		 calloc(arr_len, sizeof(struct EulabeiaPlugin))) == NULL) {
		return -1;
	}
	(*plugins)->cap = arr_len;
	(*plugins)->len = arr_len;
	for (i = 0; i < arr_len; i++) {
		p = (*plugins)->plugins++;
		jo = json_array_get_object_element(arr, i);
		if (eulabeia_json_plugin(jo, &p) != 0)
			return -2;
	}
	// jump back to index 0
	(*plugins)->plugins = (*plugins)->plugins - arr_len;

	return 0;
}

int eulabeia_json_ports(JsonArray *arr, struct EulabeiaPorts **ports)
{
	const char *a;
	unsigned int arr_len;
	if (*ports == NULL) {
		*ports = calloc(1, sizeof(struct EulabeiaPorts));
	}
	arr_len = json_array_get_length(arr);
	if (((*ports)->ports = calloc(arr_len, sizeof(struct EulabeiaPort))) ==
	    NULL) {
		return -1;
	}
	(*ports)->cap = arr_len;
	(*ports)->len = arr_len;
	for (unsigned int index = 0; index < arr_len; index++) {
		a = json_array_get_string_element(arr, index);
		(*ports)->ports[index].port = g_strdup(a);
	}

	return 0;
}

int eulabeia_json_alive_test(JsonObject *jo,
			     struct EulabeiaAliveTest **alive_test)
{
	JsonArray *arr;
	*alive_test = calloc(1, sizeof(struct EulabeiaAliveTest));

	if (json_object_has_member(jo, "test_alive_hosts_only"))
		(*alive_test)->test_alive_hosts_only =
		    json_object_get_boolean_member(jo, "alive_test");
	if (json_object_has_member(jo, "methods_bitflag"))
		(*alive_test)->methods_bitflag =
		    json_object_get_boolean_member(jo, "methods_bitflag");
	if (json_object_has_member(jo, "ports")) {
		arr = json_object_get_array_member(jo, "ports");
		if (eulabeia_json_ports(arr, &(*alive_test)->ports) != 0)
			return -3;
	}

	return 0;
}

int eulabeia_json_target(JsonObject *jo, struct EulabeiaTarget **t)
{
	struct EulabeiaTarget *target;
	JsonArray *arr;
	if (t == NULL || jo == NULL)
		return -1;
	if (!json_object_has_member(jo, "id"))
		return -2;
	if (*t == NULL)
		*t = g_malloc0(sizeof(**t));
	target = *t;

	json_object_get_and_assign_string(jo, "id", &target->id);
	if (json_object_has_member(jo, "sensor"))
		json_object_get_and_assign_string(
		    jo, "sensor", &target->sensor);
	if (json_object_has_member(jo, "alive_test")) {
		JsonObject *alive_test_jo;
		alive_test_jo = json_object_get_object_member(jo, "alive_test");
		if (eulabeia_json_alive_test(alive_test_jo,
					     &target->alive_test) != 0)
			return -3;
	}
	if (json_object_has_member(jo, "parallel"))
		target->parallel = json_object_get_int_member(jo, "parallel");
	if (json_object_has_member(jo, "hosts")) {
		arr = json_object_get_array_member(jo, "hosts");
		if (eulabeia_json_hosts(arr, &target->hosts) != 0)
			return -3;
	}
	if (json_object_has_member(jo, "exclude")) {
		arr = json_object_get_array_member(jo, "exclude");
		if (eulabeia_json_hosts(arr, &target->exclude) != 0)
			return -3;
	}
	if (json_object_has_member(jo, "plugins")) {
		arr = json_object_get_array_member(jo, "plugins");
		if (eulabeia_json_plugins(arr, &target->plugins) != 0)
			return -3;
	}
	if (json_object_has_member(jo, "ports")) {
		arr = json_object_get_array_member(jo, "ports");
		if (eulabeia_json_ports(arr, &target->ports) != 0)
			return -3;
	}
	//	if (json_object_has_member(jo, "credentials")) {
	//		arr = json_object_get_array_member(jo, "credentials");
	//		if (eulabeia_json_credentials(arr, &target->credentials)
	//!= 0) 			return -2;
	//	}
	return 0;
}

int eulabeia_json_scan(JsonObject *jo, struct EulabeiaScan **s)
{
	struct EulabeiaScan *scan;
	JsonArray *arr;
	if (jo == NULL || s == NULL)
		return -1;
	if (*s == NULL)
		*s = g_malloc0(sizeof(**s));
	if (!json_object_has_member(jo, "id"))
		return -2;
	scan = *s;
	json_object_get_and_assign_string(jo, "id", &scan->id);
	if (json_object_has_member(jo, "target_id"))
		json_object_get_and_assign_string(
		    jo, "target_id", &scan->target_id);

	if (json_object_has_member(jo, "finished")) {
		arr = json_object_get_array_member(jo, "finished");
		if (eulabeia_json_hosts(arr, &scan->finished) != 0)
			return -3;
	}
	if (json_object_has_member(jo, "temporary"))
		scan->temporary =
		    json_object_get_boolean_member(jo, "temporary");
	else
		scan->temporary = FALSE;
	// we ignore failure since a scan just may include a target
	eulabeia_json_target(jo, &scan->target);
	return 0;
}

static void builder_add_plugins(JsonBuilder *builder,
				const struct EulabeiaPlugins *plugins)
{
	unsigned int i;
	json_builder_begin_object(builder);
	json_builder_set_member_name(builder, "single_vts");
	json_builder_begin_array(builder);
	for (i = 0; i < plugins->len; i++) {
		json_builder_begin_object(builder);
		json_builder_set_member_name(builder, "oid");
		json_builder_add_string_value(builder, plugins->plugins[i].oid);
		json_builder_end_object(builder);
	}
	json_builder_end_array(builder);
	json_builder_end_object(builder);
}

static void builder_add_ports(JsonBuilder *builder,
			      const struct EulabeiaPorts *ports)
{
	unsigned int i;
	json_builder_begin_array(builder);
	for (i = 0; i < ports->len; i++) {
		json_builder_add_string_value(builder, ports->ports[i].port);
	}
	json_builder_end_array(builder);
}

static void builder_add_hosts(JsonBuilder *builder,
			      const struct EulabeiaHosts *hosts)
{
	unsigned int i;
	json_builder_begin_array(builder);
	for (i = 0; i < hosts->len; i++) {
		json_builder_add_string_value(builder, hosts->hosts[i].address);
	}
	json_builder_end_array(builder);
}

static void builder_add_result(JsonBuilder *builder,
			       const struct EulabeiaScanResult *result)
{
	json_builder_set_member_name(builder, "result_type");
	json_builder_add_string_value(builder, result->result_type);
	json_builder_set_member_name(builder, "host_ip");
	json_builder_add_string_value(builder, result->host_ip);
	json_builder_set_member_name(builder, "host_name");
	json_builder_add_string_value(builder, result->host_name);
	json_builder_set_member_name(builder, "port");
	json_builder_add_string_value(builder, result->port);
	json_builder_set_member_name(builder, "id");
	json_builder_add_string_value(builder, result->id);
	json_builder_set_member_name(builder, "oid");
	json_builder_add_string_value(builder, result->oid);
	json_builder_set_member_name(builder, "value");
	json_builder_add_string_value(builder, result->value);
	json_builder_set_member_name(builder, "uri");
	json_builder_add_string_value(builder, result->uri);
}

static void builder_add_host_status(JsonBuilder *builder,
				    const struct EulabeiaHostStatus *status)
{
	json_builder_set_member_name(builder, "status_type");
	json_builder_add_int_value(builder, status->host_status_type);
	json_builder_set_member_name(builder, "host_ip");
	json_builder_add_string_value(builder, status->host_ip);
	json_builder_set_member_name(builder, "id");
	json_builder_add_string_value(builder, status->id);
	json_builder_set_member_name(builder, "value");
	json_builder_add_string_value(builder, status->value);
}

static void builder_add_status(JsonBuilder *builder,
			       const struct EulabeiaStatus *status)
{

	json_builder_set_member_name(builder, "id");
	json_builder_add_string_value(builder, status->id);

	json_builder_set_member_name(builder, "status");
	json_builder_add_string_value(builder, status->status);
}

static void builder_add_alive_test(JsonBuilder *builder,
				   const struct EulabeiaAliveTest *alive_test)
{
	json_builder_begin_object(builder);

	json_builder_set_member_name(builder, "test_alive_hosts_only");
	json_builder_add_boolean_value(builder,
				       alive_test->test_alive_hosts_only);

	json_builder_set_member_name(builder, "methods_bitflag");
	json_builder_add_int_value(builder, alive_test->methods_bitflag);

	json_builder_set_member_name(builder, "ports");
	builder_add_ports(builder, alive_test->ports);

	json_builder_end_object(builder);
}

// expects a builder with an open object, internal use
static void builder_add_target(JsonBuilder *builder,
			       const struct EulabeiaTarget *target,
			       const int modify)
{
	if (target->id) {
		json_builder_set_member_name(builder, "id");
		json_builder_add_string_value(builder, target->id);
	}
	if (modify) {
		json_builder_set_member_name(builder, "values");
		json_builder_begin_object(builder);
	}
	if (target->sensor) {
		json_builder_set_member_name(builder, "sensor");
		json_builder_add_string_value(builder, target->sensor);
	}
	if (target->alive) {
		json_builder_set_member_name(builder, "alive_test");
		builder_add_alive_test(builder, target->alive_test);
	}
	if (target->parallel) {
		json_builder_set_member_name(builder, "parallel");
		json_builder_add_boolean_value(builder, target->parallel);
	}
	if (target->hosts) {
		json_builder_set_member_name(builder, "hosts");
		builder_add_hosts(builder, target->hosts);
	}
	if (target->plugins) {
		json_builder_set_member_name(builder, "plugins");
		builder_add_plugins(builder, target->plugins);
	}
	if (target->ports) {
		json_builder_set_member_name(builder, "ports");
		builder_add_ports(builder, target->ports);
	}
	if (target->exclude) {
		json_builder_set_member_name(builder, "exclude");
		builder_add_hosts(builder, target->exclude);
	}
	if (modify) {
		json_builder_end_object(builder);
	}
}

static void builder_add_scan(JsonBuilder *builder,
			     const struct EulabeiaScan *scan,
			     const int modify)
{
	if (scan->id) {
		json_builder_set_member_name(builder, "id");
		json_builder_add_string_value(builder, scan->id);
	}
	if (modify) {
		json_builder_set_member_name(builder, "values");
		json_builder_begin_object(builder);
	}
	json_builder_set_member_name(builder, "temporary");
	json_builder_add_boolean_value(builder, scan->temporary);
	if (scan->target_id) {
		json_builder_set_member_name(builder, "target_id");
		json_builder_add_string_value(builder, scan->target_id);
	}
	if (scan->target) {
		// if there is a target-id it will override the scan id due to
		// flat json design. Therefore we ignore the target id.
		builder_add_target(builder, scan->target, 0);
	}
	if (modify) {
		json_builder_end_object(builder);
	}
}

static void builder_add_message(JsonBuilder *builder,
				const struct EulabeiaMessage *msg)
{
	json_builder_set_member_name(builder, "message_id");
	json_builder_add_string_value(builder, msg->id);
	json_builder_set_member_name(builder, "message_type");
	json_builder_add_string_value(builder, msg->type);
	json_builder_set_member_name(builder, "group_id");
	json_builder_add_string_value(builder, msg->id);
	json_builder_set_member_name(builder, "message_created");
	json_builder_add_int_value(builder, msg->created);
}

static void builder_add_failure(JsonBuilder *builder,
				const struct EulabeiaFailure *failure)
{
	json_builder_set_member_name(builder, "id");
	json_builder_add_string_value(builder, failure->id);
	json_builder_set_member_name(builder, "error");
	json_builder_add_string_value(builder, failure->error);
}

static char *json_builder_to_str(JsonBuilder *builder)
{
	char *json_str;
	JsonGenerator *gen;
	JsonNode *root;
	gen = json_generator_new();
	json_generator_set_pretty(gen, pretty_print);
	root = json_builder_get_root(builder);
	json_generator_set_root(gen, root);
	json_str = json_generator_to_data(gen, NULL);
	json_node_free(root);
	return json_str;
}

char *eulabeia_scan_message_to_json(const struct EulabeiaMessage *msg,
				    const struct EulabeiaScan *scan,
				    const int modify)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_scan(b, scan, modify);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}

char *eulabeia_failure_message_to_json(const struct EulabeiaMessage *msg,
				       const struct EulabeiaFailure *failure)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_failure(b, failure);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}
char *eulabeia_target_message_to_json(const struct EulabeiaMessage *msg,
				      const struct EulabeiaTarget *target,
				      const int modify)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_target(b, target, modify);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}
char *
eulabeia_scan_result_message_to_json(const struct EulabeiaMessage *msg,
				     const struct EulabeiaScanResult *result)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_result(b, result);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}

char *
eulabeia_host_status_message_to_json(const struct EulabeiaMessage *msg,
				     const struct EulabeiaHostStatus *status)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_host_status(b, status);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}

char *eulabeia_status_message_to_json(const struct EulabeiaMessage *msg,
				      const struct EulabeiaStatus *status)
{
	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	builder_add_status(b, status);
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}

char *eulabeia_id_message_to_json(const struct EulabeiaMessage *msg,
				  const char *id)
{

	JsonBuilder *b;
	char *json_str;
	b = json_builder_new();

	json_builder_begin_object(b);
	builder_add_message(b, msg);
	if (id != NULL) {
		json_builder_set_member_name(b, "id");
		json_builder_add_string_value(b, id);
	}
	json_builder_end_object(b);

	json_str = json_builder_to_str(b);
	return json_str;
}
