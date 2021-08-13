#include <eulabeia/types.h>
#include <gvm/util/uuidutils.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

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
	int i;
	if (hosts == NULL || *hosts == NULL)
		return;
	for (i = 0; i < (*hosts)->len; i++) {
		free((*hosts)->hosts->address);
		free((*hosts)->hosts);
	}
	free(*hosts);
	*hosts = NULL;
}

void eulabeia_plugins_destroy(struct EulabeiaPlugins **plugins)
{
	int i;
	if (plugins == NULL || *plugins == NULL)
		return;
	for (i = 0; i < (*plugins)->len; i++) {
		free((*plugins)->plugins->oid);
		free((*plugins)->plugins);
	}
	free(*plugins);
	*plugins = NULL;
}

void eulabeia_ports_destroy(struct EulabeiaPorts **ports)
{
	int i;
	if (ports == NULL || *ports == NULL)
		return;
	for (i = 0; i < (*ports)->len; i++) {
		free((*ports)->ports->port);
		free((*ports)->ports);
	}
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
