#include <eulabeia/types.h>
#include <string.h>
#define EULABEIA_PRETTY_JSON 1
#include <eulabeia/client.h>
#include <eulabeia/json.h>
#include <gvm/util/mqtt.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define TARGET_ID "example.id.target"

int is_relevant(enum eulabeia_aggregate a, enum eulabeia_message_type mt)
{
	if (mt == EULABEIA_UNKNOWN)
		return 0;
	if (a == EULABEIA_SCAN)
		return 1;
	switch (mt) {
	case EULABEIA_CMD_GET:
	case EULABEIA_CMD_CREATE:
	case EULABEIA_INFO_MODIFIED:
	case EULABEIA_INFO_CREATED:
	case EULABEIA_INFO_GOT:
	case EULABEIA_CMD_MODIFY:
	case EULABEIA_INFO_CREATE_FAILURE:
	case EULABEIA_INFO_MODIFY_FAILURE:
	case EULABEIA_INFO_GET_FAILURE:
	case EULABEIA_INFO_FAILURE:
		return 1;

	default:
		return 0;
	}
}

static char *message_type_aggregate_link(enum eulabeia_message_type mt,
					 enum eulabeia_aggregate a)
{
	char *link, *clink;
	unsigned int i, j, link_len, a_len;
	a_len = strlen(eulabeia_aggregate_to_str(a));
	link_len = a_len + strlen(eulabeia_message_type_to_str(mt)) + +2;
	link = calloc(link_len, sizeof(*link));
	snprintf(link,
		 link_len,
		 "#%s%s",
		 eulabeia_message_type_to_str(mt),
		 eulabeia_aggregate_to_str(a));
	// slow but I don't care
	for (i = 0; i < link_len; i++)
		if (link[i] == '.')
			for (j = i; j < link_len; j++)
				link[j] = link[j + 1];
	link_len *= 2;
	link_len -= a_len - 2;
	clink = calloc(link_len, sizeof(*clink));
	snprintf(clink,
		 link_len,
		 "[%s](%s)",
		 eulabeia_message_type_to_str(mt),
		 link);
	fprintf(stderr, "%s: link %s\n", __func__, clink);
	free(link);
	return clink;
}

static struct EulabeiaTarget *example_target()
{
	struct EulabeiaTarget *target;
	target = calloc(1, sizeof(struct EulabeiaTarget));
	target->id = TARGET_ID;
	target->alive = 1;
	target->sensor = "example.sensor.1";
	struct EulabeiaHosts *t_hosts;
	struct EulabeiaPlugins *t_plugins;
	struct EulabeiaPorts *t_ports;
	t_hosts = calloc(1, sizeof(struct EulabeiaHosts));
	t_hosts->hosts = calloc(2, sizeof(struct EulabeiaHost));
	t_hosts->hosts[0].address = "example.host.to.scan.com";
	t_hosts->len = 1;
	t_hosts->cap = 2;
	target->hosts = t_hosts;
	t_plugins = calloc(1, sizeof(struct EulabeiaPlugins));
	t_plugins->cap = 1;
	t_plugins->len = 1;
	t_plugins->plugins = calloc(1, sizeof(struct EulabeiaPlugins));
	t_plugins->plugins[0].oid = "example.oid.1";
	target->plugins = t_plugins;
	t_ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->cap = 1;
	t_ports->len = 1;
	t_ports->ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->ports[0].port = "1337";
	target->ports = t_ports;
	return target;
}

static struct EulabeiaScan *example_scan(struct EulabeiaTarget *target)
{
	struct EulabeiaScan *scan;
	scan = calloc(1, sizeof(*scan));
	scan->id = "example.scan.id";
	if (target)
		scan->target = target;
	else
		scan->target_id = "example.target.id";
	return scan;
}

static void free_example_target(struct EulabeiaTarget *target)
{
	free(target->hosts->hosts);
	free(target->hosts);
	free(target->plugins->plugins);
	free(target->plugins);
	free(target->ports->ports);
	free(target->ports);
	free(target);
}

static void free_example_scan(struct EulabeiaScan *scan) { free(scan); }

struct AggregateTOC {
	char **links; // the actual link content
	enum eulabeia_message_type
	    *types; // which types are valid for aggregate
	unsigned int len;
	unsigned int cap;
};

struct TOC {
	enum eulabeia_aggregate
	    aggregate; // is used to build the actual content
	struct AggregateTOC *tocs;
	unsigned int len;
	unsigned int cap;
};

struct AggregateTOC *build_aggregate_toc(enum eulabeia_aggregate aggregate)
{
	struct AggregateTOC *toc;
	fprintf(stderr,
		"%s: build toc for %s\n",
		__func__,
		eulabeia_aggregate_to_str(aggregate));
	toc = calloc(1, sizeof(*toc));
#define X(a, b, c)                                                             \
	if (is_relevant(aggregate, a)) {                                       \
		if (toc->len == toc->cap) {                                    \
			toc->cap += 10;                                        \
			if ((toc->links = realloc(                             \
				 toc->links,                                   \
				 toc->cap * sizeof(*toc->links))) == NULL)     \
				exit(23);                                      \
			if ((toc->types = realloc(                             \
				 toc->types,                                   \
				 toc->cap * sizeof(*toc->types))) == NULL)     \
				exit(23);                                      \
		}                                                              \
		toc->links[toc->len] =                                         \
		    message_type_aggregate_link(a, aggregate);                 \
		toc->types[toc->len] = a;                                      \
		toc->len++;                                                    \
	}
	EULABEIA_MESSAGE_TYPES
#undef X
	return toc;
}

struct TOC *build_toc()
{
	struct TOC *toc;
	toc = calloc(1, sizeof(*toc));
#define X(a, b)                                                                \
	if (toc->len == toc->cap) {                                            \
		toc->cap += 2;                                                 \
		if ((toc->tocs = realloc(                                      \
			 toc->tocs, toc->cap * sizeof(*toc->tocs))) == NULL)   \
			exit(42);                                              \
	}                                                                      \
	toc->tocs[toc->len++] = *build_aggregate_toc((a));                     \
	toc->aggregate = (a);

	EULABEIA_AGGREGATES
#undef X
	return toc;
}

void print_single_response(enum eulabeia_message_type mt,
			   enum eulabeia_aggregate a)
{
	char *link;
	link = message_type_aggregate_link(mt, a);
	printf("- %s\n", link);
	free(link);
}

void print_response(enum eulabeia_message_type type, enum eulabeia_aggregate a)
{
	switch (type) {
	case EULABEIA_CMD_START:
		print_single_response(EULABEIA_INFO_STATUS, a);
		print_single_response(EULABEIA_INFO_START_FAILURE, a);
		break;
	case EULABEIA_CMD_STOP:
		print_single_response(EULABEIA_INFO_STOPPED, a);
		print_single_response(EULABEIA_INFO_STOP_FAILURE, a);
		break;
	case EULABEIA_CMD_MODIFY:
		print_single_response(EULABEIA_INFO_MODIFIED, a);
		print_single_response(EULABEIA_INFO_MODIFY_FAILURE, a);
		break;
	case EULABEIA_CMD_CREATE:
		print_single_response(EULABEIA_INFO_CREATED, a);
		print_single_response(EULABEIA_INFO_CREATE_FAILURE, a);
		break;
	case EULABEIA_CMD_GET:
		print_single_response(EULABEIA_INFO_GOT, a);
		print_single_response(EULABEIA_INFO_GET_FAILURE, a);
		break;
	default:
		break;
	}
}

char *build_id(enum eulabeia_aggregate aggregate)
{
	char *response = calloc(1, 128);
	snprintf(response,
		 128,
		 "example.id.%s",
		 eulabeia_aggregate_to_str(aggregate));
	return response;
}

char *destination(enum eulabeia_message_type mt)
{
	if (strncmp("info", eulabeia_message_type_to_event_type(mt), 4) == 0)
		return NULL;
	return EULABEIA_DIRECTOR;
}

static void print_entry(enum eulabeia_aggregate aggregate,
			struct AggregateTOC *toc)
{
	struct EulabeiaMessage *message = NULL;
	struct EulabeiaStatus *et = NULL;
	struct EulabeiaScanResult *esr = NULL;
	struct EulabeiaScan *scan = NULL;
	struct EulabeiaTarget *target = NULL;
	struct EulabeiaFailure *failure = NULL;
	char *id, *example, *response, *addition = NULL;
	enum eulabeia_message_type type;
	int i, modify;
	fprintf(stderr,
		"%s: for %s\n",
		__func__,
		eulabeia_aggregate_to_str(aggregate));

	for (i = 0; i < toc->len; i++) {
		type = toc->types[i];
		message = eulabeia_initialize_message(
		    type, aggregate, "optional_grouping_id", destination(type));
		id = type == EULABEIA_CMD_CREATE ? NULL : build_id(aggregate);
		switch (type) {
		case EULABEIA_UNKNOWN:
			break;
		case EULABEIA_CMD_START:
		case EULABEIA_CMD_STOP:
		case EULABEIA_INFO_STOPPED:
		case EULABEIA_CMD_GET:
		case EULABEIA_CMD_CREATE:
		case EULABEIA_INFO_MODIFIED:
		case EULABEIA_INFO_CREATED:
			example = eulabeia_id_message_to_json(message, id);
			break;
		case EULABEIA_INFO_GOT:
			modify = 0;
			if (aggregate == EULABEIA_SCAN) {
				target = example_target();
				target->id = "example.id.scan";
			}
		case EULABEIA_CMD_MODIFY:
			switch (aggregate) {
			case EULABEIA_SCAN:
				scan = example_scan(target);
				example = eulabeia_scan_message_to_json(
				    message, scan, modify);
				break;
			case EULABEIA_TARGET:
				target = example_target();
				example = eulabeia_target_message_to_json(
				    message, target, modify);
				break;
			}
			break;
		case EULABEIA_INFO_SCAN_RESULT:
			esr = calloc(1, sizeof(*esr));
			esr->id = id;
			esr->result_type = eulabeia_result_type_to_str(
			    EULABEIA_RESULT_TYPE_LOG);
			esr->host_ip = "192.168.1.1";
			esr->host_name = "example.host.domain";
			esr->oid = "example.oid.1";
			esr->port = "1337";
			esr->uri = "uri.to.oid.description";
			esr->value = "This an example log message";
			example =
			    eulabeia_scan_result_message_to_json(message, esr);
			addition = calloc(1, 1024);
			snprintf(addition, 1024, "Valid `result_type` are:\n");
#define X(a, b) strncat(addition, "- `" b "`\n", 1024);
			EULABEIA_RESULT_TYPES
#undef X
			break;

		case EULABEIA_INFO_STATUS:
			et = calloc(1, sizeof(*et));
			et->id = id;
			et->message = message;
			et->status =
			    eulabeia_scan_state_to_str(EULABEIA_SCAN_REQUESTED);
			example = eulabeia_status_message_to_json(message, et);
			addition = calloc(1, 1024);
			snprintf(addition, 1024, "Valid `status` are:\n");
#define X(a, b) strncat(addition, "- `" #b "`\n", 1024);
			EULABEIA_SCAN_STATES
#undef X
			break;
		case EULABEIA_INFO_CREATE_FAILURE:
		case EULABEIA_INFO_MODIFY_FAILURE:
		case EULABEIA_INFO_STOP_FAILURE:
		case EULABEIA_INFO_START_FAILURE:
		case EULABEIA_INFO_GET_FAILURE:
		case EULABEIA_INFO_FAILURE:
			failure = calloc(1, sizeof(*failure));
			failure->id = id;
			failure->error = "some error description";
			example =
			    eulabeia_failure_message_to_json(message, failure);
			break;
		}
		printf("## %s/%s\n\n",
		       eulabeia_message_type_to_str(type),
		       eulabeia_aggregate_to_str(aggregate));
		printf("Topic: %s\n",
		       eulabeia_calculate_topic(type,
						aggregate,
						EULABEIA_SCANNER_CONTEXT,
						message->destination));
		printf("```\n%s\n```\n", example);
		if (addition != NULL)
			printf("%s\n", addition);
		if (strncmp("cmd",
			    eulabeia_message_type_to_event_type(type),
			    3) == 0) {
			printf("Responses:\n\n");
			print_response(type, aggregate);
		}
		if (message != NULL) {
			free(message);
			message = NULL;
		}
		if (et != NULL) {
			free(et);
			et = NULL;
		}
		if (esr != NULL) {
			free(esr);
			esr = NULL;
		}
		if (scan != NULL) {
			free_example_scan(scan);
			scan = NULL;
		}
		if (target != NULL) {
			free_example_target(target);
			target = NULL;
		}
		if (failure != NULL) {
			free(failure);
			failure = NULL;
		}
		if (addition != NULL) {
			free(addition);
			addition = NULL;
		}
		if (id != NULL) {
			free(id);
			id = NULL;
		}
	}
}

void print_message_output(struct TOC *toc)
{
	int i, j;
	for (i = 0; i < toc->len; i++) {
		printf("# %s\n\n", eulabeia_aggregate_to_str(i));
		print_entry(i, &toc->tocs[i]);
	}
}

int main()
{
	eulabeia_json_set_pretty_print(1);
	int i, j;
	char *a;
	struct TOC *toc;
	printf("<!-- DON'T EDIT THIS FILE; INSTEAD RUN: generate_md -->\n");
	printf("# Table of content\n");
	toc = build_toc();
	for (i = 0; i < toc->len; i++) {
		a = eulabeia_aggregate_to_str(i);
		printf("- [%s](#%s)\n", a, a);
		for (j = 0; j < toc->tocs[i].len; j++) {
			printf("  - %s\n", toc->tocs[i].links[j]);
		}
	}
	printf("\n\n");
	print_message_output(toc);
	for (i = 0; i < toc->len; i++) {
		for (j = 0; j < toc->tocs[i].len; j++) {
			free(toc->tocs[i].links[j]);
		}
		// realloc hence just once
		free(toc->tocs[i].links);
		free(toc->tocs[i].types);
	}
	free(toc->tocs);

	return 0;
}
