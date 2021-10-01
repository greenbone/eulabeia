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

#include <eulabeia/types.h>
#include <gvm/util/uuidutils.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include <glib.h>

char *eulabeia_scan_state_to_str(enum eulabeia_scan_state srs)
{
	switch (srs) {
#define X(a, b)                                                                \
	case a:                                                                \
		return #b;
		EULABEIA_SCAN_STATES
#undef X
	default:
		return NULL;
	}
}

char *eulabeia_message_type_to_event_type(enum eulabeia_message_type mt)
{
	switch (mt) {
#define X(a, b, c)                                                             \
	case a:                                                                \
		return #c;
		EULABEIA_MESSAGE_TYPES
#undef X
	default:
		return "info";
	}
}

char *eulabeia_message_type_to_str(enum eulabeia_message_type mt)
{
	switch (mt) {
#define X(a, b, c)                                                             \
	case a:                                                                \
		return #b;
		EULABEIA_MESSAGE_TYPES
#undef X
	default:
		return NULL;
	}
}

char *eulabeia_aggregate_to_str(enum eulabeia_aggregate a)
{
	switch (a) {
#define X(a, b)                                                                \
	case a:                                                                \
		return #b;
		EULABEIA_AGGREGATES
#undef X
	default:
		return NULL;
	}
}

enum eulabeia_aggregate eulabeia_aggregate_from_str(char *rt)
{
	if (rt == NULL)
		return EULABEIA_AGGREGATE_UNKNOWN;
#define X(a, b) else if (strncmp(rt, #b, strlen(#b)) == 0) return (a);
	EULABEIA_AGGREGATES
#undef X
	return EULABEIA_AGGREGATE_UNKNOWN;
}

char *eulabeia_result_type_to_str(enum eulabeia_result_type mt)
{
	switch (mt) {
#define X(a, b)                                                                \
	case a:                                                                \
		return b;
		EULABEIA_RESULT_TYPES
#undef X
	default:
		return NULL;
	}
}

char *eulabeia_host_status_type_to_str(enum eulabeia_host_status_type mt)
{
	switch (mt) {
#define X(a, b)                                                                \
	case a:                                                                \
		return b;
		EULABEIA_HOST_STATUS_TYPES
#undef X
	default:
		return NULL;
	}
}

void eulabeia_message_destroy(struct EulabeiaMessage **msg)
{
	if (*msg == NULL) {
		return;
	}
	g_free((*msg)->id);
	g_free((*msg)->type);
	g_free((*msg)->group_id);
	g_free(*msg);
	*msg = NULL;
}

void eulabeia_failure_destroy(struct EulabeiaFailure **failure)
{
	if (*failure == NULL)
		return;
	g_free((*failure)->id);
	g_free((*failure)->error);
	g_free(*failure);
	*failure = NULL;
}

void eulabeia_status_destroy(struct EulabeiaStatus **status)
{
	if (*status == NULL)
		return;
	g_free((*status)->id);
	g_free((*status)->status);
	g_free(*status);
	*status = NULL;
}

void eulabeia_hosts_destroy(struct EulabeiaHosts **hosts)
{
	struct EulabeiaHost *p_index, *p_orig;
	unsigned int i = 0;

	if (hosts == NULL || *hosts == NULL) {
		g_warning("hosts == NULL || *hosts == NULL");
		return;
	}

	p_index = (*hosts)->hosts;
	p_orig = (*hosts)->hosts;
	/* Free addresses of EulabeiaHost structs in array */
	for (; i < (*hosts)->len; p_index++, i++) {
		g_free(p_index->address);
	}
	/* Free EulabeiaHost array */
	g_free(p_orig);
	/* Free EulabeiaHosts struct */
	g_free(*hosts);
	*hosts = NULL;
}

static void
plugin_references_destroy(struct EulabeiaPluginReferences **references)
{
	unsigned int i = 0;
	struct EulabeiaPluginReference *pi, *po;

	if (references == NULL || *references == NULL)
		return;

	pi = (*references)->reference;
	po = (*references)->reference;
	for (; i < (*references)->len; pi++, i++) {
		g_free(pi->id);
		g_free(pi->type);
	}
	g_free(po);
	g_free(*references);
	*references = NULL;
}

static void
plugin_dependencies_destroy(struct EulabeiaPluginDependencies **dependencies)
{
	unsigned int i = 0;
	struct EulabeiaPluginDependency *di, *deo;

	if (dependencies == NULL || *dependencies == NULL)
		return;

	di = (*dependencies)->dependency;
	deo = (*dependencies)->dependency;
	for (; i < (*dependencies)->len; di++, i++) {
		g_free(di->filename);
	}
	g_free(deo);
	g_free(*dependencies);
	*dependencies = NULL;
}

static void
plugin_parameters_destroy(struct EulabeiaPluginParameters **parameters)
{
	unsigned int i = 0;
	struct EulabeiaPluginParameter *ri, *ro;

	if (parameters == NULL || *parameters == NULL)
		return;

	ri = (*parameters)->parameter;
	ro = (*parameters)->parameter;
	for (; i < (*parameters)->len; ri++, i++) {
		g_free(ri->id);
		g_free(ri->name);
		g_free(ri->value);
		g_free(ri->type);
		g_free(ri->description);
		g_free(ri->defaultvalue);
	}
	g_free(ro);
	g_free(*parameters);
	*parameters = NULL;
}

static void plugin_severity_destroy(struct EulabeiaPluginSeverity **severity)
{
	if (severity == NULL || *severity == NULL)
		return;
	g_free((*severity)->vector);
	g_free((*severity)->type);
	g_free((*severity)->date);
	g_free((*severity)->origin);
	g_free(*severity);
	*severity = NULL;
}

void eulabeia_plugins_destroy(struct EulabeiaPlugins **plugins)
{
	unsigned int i = 0;
	struct EulabeiaPlugin *p_index, *p_orig;

	if (plugins == NULL || *plugins == NULL)
		return;

	p_index = (*plugins)->plugins;
	p_orig = (*plugins)->plugins;
	for (; i < (*plugins)->len; p_index++, i++) {
		g_free(p_index->oid);
		g_free(p_index->affected);
		g_free(p_index->category);
		g_free(p_index->created);
		g_free(p_index->excluded_keys);
		g_free(p_index->family);
		g_free(p_index->filename);
		g_free(p_index->impact);
		g_free(p_index->insight);
		g_free(p_index->mandatory_keys);
		g_free(p_index->modified);
		g_free(p_index->name);
		g_free(p_index->qod);
		g_free(p_index->qod_type);
		g_free(p_index->required_keys);
		g_free(p_index->required_ports);
		g_free(p_index->required_udp_ports);
		g_free(p_index->solution);
		g_free(p_index->solution_method);
		g_free(p_index->summary);
		g_free(p_index->vuldetect);
		if (p_index->references != NULL)
			plugin_references_destroy(&p_index->references);
		if (p_index->parameters != NULL)
			plugin_parameters_destroy(&p_index->parameters);
		if (p_index->dependencies != NULL)
			plugin_dependencies_destroy(&p_index->dependencies);
		if (p_index->severity != NULL)
			plugin_severity_destroy(&p_index->severity);
	}
	/* Free EulabeiaPlugin array */
	g_free(p_orig);
	/* Free EulabeiaPlugins struct */
	g_free(*plugins);

	*plugins = NULL;
}

void eulabeia_ports_destroy(struct EulabeiaPorts **ports)
{
	unsigned int i = 0;
	struct EulabeiaPort *p_index, *p_orig;

	if (ports == NULL || *ports == NULL)
		return;

	p_index = (*ports)->ports;
	p_orig = (*ports)->ports;
	/* Free port of EulabeiaPort structs in array */
	for (; i < (*ports)->len; p_index++, i++) {
		g_free(p_index->port);
	}
	/* Free EulabeiaPort array */
	g_free(p_orig);
	/* Free EulabeiaPorts struct */
	g_free(*ports);

	*ports = NULL;
}

static void free_scan_result_data(struct EulabeiaScanResult *scan_result)
{
	g_free((scan_result)->result_type);
	g_free((scan_result)->host_ip);
	g_free((scan_result)->host_name);
	g_free((scan_result)->oid);
	g_free((scan_result)->id);
	g_free((scan_result)->uri);
	g_free((scan_result)->value);
	g_free((scan_result)->port);
}

void eulabeia_scan_result_destroy(struct EulabeiaScanResult **scan_result)
{
	if (scan_result == NULL || *scan_result == NULL)
		return;

	free_scan_result_data(*scan_result);

	g_free(*scan_result);
	*scan_result = NULL;
}

static void free_host_status_data(struct EulabeiaHostStatus *status)
{
	g_free((status)->host_ip);
	g_free((status)->id);
	g_free((status)->value);
}

void eulabeia_host_status_destroy(struct EulabeiaHostStatus **status)
{
	if (status == NULL || *status == NULL)
		return;

	free_host_status_data(*status);

	g_free(*status);
	*status = NULL;
}

void eulabeia_scan_progress_destroy(struct EulabeiaScanProgress **scan_progress)
{
	unsigned int i;
	struct EulabeiaScanResult *ptr;
	if (scan_progress == NULL || *scan_progress == NULL)
		return;
	if ((*scan_progress)->results != NULL &&
	    (*scan_progress)->results->results != NULL) {
		for (i = 0; i < (*scan_progress)->results->len; i++) {
			ptr = (*scan_progress)->results->results + i;
			free_scan_result_data(ptr);
		}
		g_free((*scan_progress)->results->results);
	}
	g_free(*scan_progress);
	*scan_progress = NULL;
}
char *eulabeia_message_type(enum eulabeia_message_type message_type,
			    enum eulabeia_aggregate aggregate,
			    char *destination)
{
	char *result;
	unsigned long len;
	len = destination == NULL ? 0 : strlen(destination);
	len = strlen(eulabeia_message_type_to_str(message_type)) +
	      strlen(eulabeia_aggregate_to_str(aggregate)) + 2;
	if ((result = calloc(1, len)) == NULL)
		return NULL;
	if (destination != NULL) {
		snprintf(result,
			 len,
			 "%s.%s.%s",
			 eulabeia_message_type_to_str(message_type),
			 eulabeia_aggregate_to_str(aggregate),
			 destination);

	} else {
		snprintf(result,
			 len,
			 "%s.%s",
			 eulabeia_message_type_to_str(message_type),
			 eulabeia_aggregate_to_str(aggregate));
	}
	return result;
}

struct EulabeiaMessage *
eulabeia_initialize_message(enum eulabeia_message_type message_type,
			    enum eulabeia_aggregate aggregate,
			    char *group_id,
			    char *destination)
{
	struct EulabeiaMessage *em;
	struct timespec spec;

	if ((em = calloc(1, sizeof(struct EulabeiaMessage))) == NULL) {
		return NULL;
	}
	em->id = gvm_uuid_make();
	if (group_id == NULL)
		em->group_id = gvm_uuid_make();
	else
		em->group_id = group_id;
	em->type = eulabeia_message_type(message_type, aggregate, destination);
	clock_gettime(CLOCK_REALTIME, &spec);
	em->created = (unsigned long)spec.tv_sec * 1000000000L +
		      (unsigned long)spec.tv_nsec;
	em->destination = destination;

	return em;
}

enum eulabeia_message_type
eulabeia_message_to_message_type(const struct EulabeiaMessage *message)
{
	if (message == NULL || message->type == NULL)
		return EULABEIA_UNKNOWN;
#define DOT_HACK(a) #a "."
#define X(a, b, c)                                                             \
	else if (strncmp(message->type, DOT_HACK(b), strlen(DOT_HACK(b))) ==   \
		 0) return (a);
	EULABEIA_MESSAGE_TYPES
#undef X
#undef DOT_HACK
	return EULABEIA_UNKNOWN;
}

enum eulabeia_result_type eulabeia_result_type_from_str(char *rt)
{
	if (rt == NULL)
		return EULABEIA_RESULT_TYPE_UNKNOWN;
#define X(a, b) else if (strncmp(rt, b, strlen(b)) == 0) return (a);
	EULABEIA_RESULT_TYPES
#undef X
	return EULABEIA_RESULT_TYPE_UNKNOWN;
}

enum eulabeia_host_status_type eulabeia_host_status_type_from_str(char *rt)
{
	if (rt == NULL)
		return EULABEIA_HOST_STATUS_TYPE_UNKNOWN;
#define X(a, b) else if (strncmp(rt, b, strlen(b)) == 0) return (a);
	EULABEIA_HOST_STATUS_TYPES
#undef X
	return EULABEIA_HOST_STATUS_TYPE_UNKNOWN;
}
