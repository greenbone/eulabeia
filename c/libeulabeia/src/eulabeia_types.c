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
	if ((*msg)->id)
		free((*msg)->id);
	if ((*msg)->type)
		free((*msg)->type);
	if ((*msg)->group_id)
		free((*msg)->group_id);
	free(*msg);
	*msg = NULL;
}

void eulabeia_failure_destroy(struct EulabeiaFailure **failure)
{
	if (*failure == NULL)
		return;
	free((*failure)->id);
	if ((*failure)->error)
		free((*failure)->error);
	free(*failure);
	*failure = NULL;
}

void eulabeia_status_destroy(struct EulabeiaStatus **status)
{
	if (*status == NULL)
		return;
	free((*status)->id);
	if ((*status)->status)
		free((*status)->status);
	free(*status);
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
		free(p_index->address);
	}
	/* Free EulabeiaHost array */
	free(p_orig);
	/* Free EulabeiaHosts struct */
	free(*hosts);
	*hosts = NULL;
}

static void
plugin_parameters_destroy(struct EulabeiaPluginParameters **parameters)
{
	unsigned int i = 0;
	struct EulabeiaPluginParameter *pi, *po;

	if (parameters == NULL || *parameters == NULL)
		return;

	pi = (*parameters)->parameter;
	po = (*parameters)->parameter;
	for (; i < (*parameters)->len; pi++, i++) {
		if (pi->id != NULL)
			free(pi->id);
		if (pi->type != NULL)
			free(pi->type);
	}
	free(po);
	free(*parameters);
	*parameters = NULL;
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
		if (di->oid != NULL)
			free(di->oid);
	}
	free(deo);
	free(*dependencies);
	*dependencies = NULL;
}

static void
plugin_references_destroy(struct EulabeiaPluginReferences **references)
{
	unsigned int i = 0;
	struct EulabeiaPluginReference *ri, *ro;

	if (references == NULL || *references == NULL)
		return;

	ri = (*references)->reference;
	ro = (*references)->reference;
	for (; i < (*references)->len; ri++, i++) {
		if (ri->id != NULL)
			free(ri->id);
		if (ri->name != NULL)
			free(ri->name);
		if (ri->value != NULL)
			free(ri->value);
		if (ri->type != NULL)
			free(ri->type);
		if (ri->description != NULL)
			free(ri->description);
		if (ri->defaultvalue != NULL)
			free(ri->defaultvalue);
	}
	free(ro);
	free(*references);
	*references = NULL;
}

static void plugin_severity_destroy(struct EulabeiaPluginSeverity **severity)
{
	if (severity == NULL || *severity == NULL)
		return;
	if ((*severity)->vector != NULL)
		free((*severity)->vector);
	if ((*severity)->type != NULL)
		free((*severity)->type);
	if ((*severity)->date != NULL)
		free((*severity)->date);
	if ((*severity)->origin != NULL)
		free((*severity)->origin);

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
		if (p_index->oid != NULL)
			free(p_index->oid);
		if (p_index->affected != NULL)
			free(p_index->affected);
		if (p_index->category != NULL)
			free(p_index->category);
		if (p_index->created != NULL)
			free(p_index->created);
		if (p_index->excluded_keys != NULL)
			free(p_index->excluded_keys);
		if (p_index->family != NULL)
			free(p_index->family);
		if (p_index->filename != NULL)
			free(p_index->filename);
		if (p_index->impact != NULL)
			free(p_index->impact);
		if (p_index->insight != NULL)
			free(p_index->insight);
		if (p_index->mandatory_keys != NULL)
			free(p_index->mandatory_keys);
		if (p_index->modified != NULL)
			free(p_index->modified);
		if (p_index->name != NULL)
			free(p_index->name);
		if (p_index->qod != NULL)
			free(p_index->qod);
		if (p_index->qod_type != NULL)
			free(p_index->qod_type);
		if (p_index->required_keys != NULL)
			free(p_index->required_keys);
		if (p_index->required_ports != NULL)
			free(p_index->required_ports);
		if (p_index->required_udp_ports != NULL)
			free(p_index->required_udp_ports);
		if (p_index->solution != NULL)
			free(p_index->solution);
		if (p_index->solution_method != NULL)
			free(p_index->solution_method);
		if (p_index->summary != NULL)
			free(p_index->summary);
		if (p_index->vuldetect != NULL)
			free(p_index->vuldetect);
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
	free(p_orig);
	/* Free EulabeiaPlugins struct */
	free(*plugins);

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
		free(p_index->port);
	}
	/* Free EulabeiaPort array */
	free(p_orig);
	/* Free EulabeiaPorts struct */
	free(*ports);

	*ports = NULL;
}

static void free_scan_result_data(struct EulabeiaScanResult *scan_result)
{
	if ((scan_result)->result_type)
		free((scan_result)->result_type);
	if ((scan_result)->host_ip)
		free((scan_result)->host_ip);
	if ((scan_result)->host_name)
		free((scan_result)->host_name);
	if ((scan_result)->oid)
		free((scan_result)->oid);
	if ((scan_result)->id)
		free((scan_result)->id);
	if ((scan_result)->uri)
		free((scan_result)->uri);
	if ((scan_result)->value)
		free((scan_result)->value);
	if ((scan_result)->port)
		free((scan_result)->port);
}

void eulabeia_scan_result_destroy(struct EulabeiaScanResult **scan_result)
{
	if (scan_result == NULL || *scan_result == NULL)
		return;

	free_scan_result_data(*scan_result);

	free(*scan_result);
	*scan_result = NULL;
}

static void free_host_status_data(struct EulabeiaHostStatus *status)
{
	if ((status)->host_ip)
		free((status)->host_ip);
	if ((status)->id)
		free((status)->id);
	;
	if ((status)->value)
		free((status)->value);
}

void eulabeia_host_status_destroy(struct EulabeiaHostStatus **status)
{
	if (status == NULL || *status == NULL)
		return;

	free_host_status_data(*status);

	free(*status);
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
		free((*scan_progress)->results->results);
	}
	free(*scan_progress);
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
