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
	int i;

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

void eulabeia_plugins_destroy(struct EulabeiaPlugins **plugins)
{
	int i;
	struct EulabeiaPlugin *p_index, *p_orig;

	if (plugins == NULL || *plugins == NULL)
		return;

	p_index = (*plugins)->plugins;
	p_orig = (*plugins)->plugins;
	/* Free oids of EulabeiaPlugin structs in array */
	for (; i < (*plugins)->len; p_index++, i++) {
		free(p_index->oid);
	}
	/* Free EulabeiaPlugin array */
	free(p_orig);
	/* Free EulabeiaPlugins struct */
	free(*plugins);

	*plugins = NULL;
}

void eulabeia_ports_destroy(struct EulabeiaPorts **ports)
{
	int i;
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

char *eulabeia_message_type(enum eulabeia_message_type message_type,
			    enum eulabeia_aggregate aggregate)
{
	char *result;
	unsigned long len;
	len = strlen(eulabeia_message_type_to_str(message_type)) +
	      strlen(eulabeia_aggregate_to_str(aggregate)) + 2;
	if ((result = calloc(1, len)) == NULL)
		return NULL;
	snprintf(result,
		 len,
		 "%s.%s",
		 eulabeia_message_type_to_str(message_type),
		 eulabeia_aggregate_to_str(aggregate));
	return result;
}

struct EulabeiaMessage *
eulabeia_initialize_message(enum eulabeia_message_type message_type,
			    enum eulabeia_aggregate aggregate,
			    char *group_id)
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
	em->type = eulabeia_message_type(message_type, aggregate);
	clock_gettime(CLOCK_REALTIME, &spec);
	em->created = (unsigned long)spec.tv_sec * 1000000000L +
		      (unsigned long)spec.tv_nsec;

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
