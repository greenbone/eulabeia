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

#include <json-glib/json-glib.h>
#include <stdlib.h>

static int json_object_get_and_assign_string(JsonObject *obj,
					     const char *key,
					     char **value)
{
	const char *jvalue;
	unsigned int len;
	if ((jvalue = json_object_get_string_member(obj, key)) == NULL) {
		return 1;
	}
	len = strlen(jvalue);
	if ((*value = calloc(1, len)) == NULL)
		return -1;
	if (strncpy(*value, jvalue, len) == NULL)
		return -2;
	return 0;
}

int eulabeia_json_message(JsonObject *obj, struct EulabeiaMessage **msg)
{
	if (!(json_object_has_member(obj, "message_id") &&
	      json_object_has_member(obj, "message_type") &&
	      json_object_has_member(obj, "created") &&
	      json_object_has_member(obj, "group_id"))) {
		return -1;
	}
	if (*msg == NULL &&
	    (*msg = calloc(1, sizeof(struct EulabeiaMessage))) == NULL) {
		return -2;
	}
	(*msg)->created = json_object_get_int_member(obj, "created");
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
	int rc;
	if (msg == NULL || msg->type == NULL)
		return -1;
	if (eulabeia_message_to_message_type(msg) != EULABEIA_INFO_FAILURE)
		return -2;

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
	int rc;
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
int eulabeia_json_status(JsonObject *obj,
			 struct EulabeiaMessage *msg,
			 struct EulabeiaStatus **status)
{
	int rc;
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
	struct EulabeiaHost *h;
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
	for (int index = 0; index < arr_len; index++) {
		a = json_array_get_string_element(arr, index);
		(*hosts)->hosts[index].address = g_strdup(a);
	}

	return 0;
}

// TODO credentials

int eulabeia_json_plugins(JsonArray *arr, struct EulabeiaPlugins **plugins)
{
	struct EulabeiaPlugin *h;
	const char *a;
	unsigned int arr_len;
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
	for (int index = 0; index < arr_len; index++) {
		a = json_array_get_string_element(arr, index);
		(*plugins)->plugins[index].oid = g_strdup(a);
	}

	return 0;
}

int eulabeia_json_ports(JsonArray *arr, struct EulabeiaPorts **ports)
{
	struct EulabeiaPort *h;
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
	for (int index = 0; index < arr_len; index++) {
		a = json_array_get_string_element(arr, index);
		(*ports)->ports[index].port = g_strdup(a);
	}

	return 0;
}
void builder_add_plugins(JsonBuilder *builder,
			 const struct EulabeiaPlugins *plugins)
{
	int i;
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

void builder_add_ports(JsonBuilder *builder, const struct EulabeiaPorts *ports)
{
	int i;
	json_builder_begin_array(builder);
	for (i = 0; i < ports->len; i++) {
		json_builder_add_string_value(builder, ports->ports[i].port);
	}
	json_builder_end_array(builder);
}

void builder_add_hosts(JsonBuilder *builder, const struct EulabeiaHosts *hosts)
{
	int i;
	json_builder_begin_array(builder);
	for (i = 0; i < hosts->len; i++) {
		json_builder_add_string_value(builder, hosts->hosts[i].address);
	}
	json_builder_end_array(builder);
}

// expects a builder with an open object, internal use
void builder_add_target(JsonBuilder *builder,
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
		json_builder_set_member_name(builder, "alive");
		json_builder_add_boolean_value(builder, target->alive);
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

void builder_add_scan(JsonBuilder *builder,
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

void builder_add_message(JsonBuilder *builder,
			 const struct EulabeiaMessage *msg)
{
	json_builder_set_member_name(builder, "message_id");
	json_builder_add_string_value(builder, msg->id);
	json_builder_set_member_name(builder, "message_type");
	json_builder_add_string_value(builder, msg->type);
	json_builder_set_member_name(builder, "group_id");
	json_builder_add_string_value(builder, msg->id);
	json_builder_set_member_name(builder, "created");
	json_builder_add_int_value(builder, msg->created);
}

char *json_builder_to_str(JsonBuilder *builder)
{
	char *json_str;
	JsonGenerator *gen;
	JsonNode *root;
	gen = json_generator_new();
	root = json_builder_get_root(builder);
	json_generator_set_root(gen, root);
	json_str = json_generator_to_data(gen, NULL);
	json_node_free(root);
	g_object_unref(gen);
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
	g_object_unref(b);
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
	g_object_unref(b);
	return json_str;
}
