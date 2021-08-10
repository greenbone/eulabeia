#include "eulabeia/types.h"
#include <eulabeia/client.h>
#include <eulabeia/json.h>

#include <gvm/util/mqtt.h>
#include <json-glib/json-glib.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

volatile int already_connected =
    0; // is used to verify if mqtt connection should be handled

#define EULABEIA_INFO "#"

static int
default_publish(const char *topic, const char *message, void *context)
{
	return mqtt_publish(topic, message);
}

static int default_retrieve(char **topic,
			    int *topic_len,
			    char **payload,
			    int *payload_len,
			    void *context)
{
	return mqtt_retrieve_message(topic, topic_len, payload, payload_len);
}

struct EulabeiaClient *eulabeia_initialize(const char *broker_address,
					   void *context)
{
	struct EulabeiaClient *ec;
	int err;
	if ((ec = calloc(1, sizeof(struct EulabeiaClient))) == NULL)
		goto failure;
	already_connected = mqtt_is_initialized();
	if (!already_connected && (err = mqtt_init(broker_address)) != 0)
		goto failure;
	if ((err = mqtt_subscribe(EULABEIA_INFO)) != 0)
		goto failure;
	ec->publish = &default_publish;
	ec->retrieve = &default_retrieve;
	ec->context = context;

	return ec;
failure:
	if (ec != NULL)
		free(ec);
	return NULL;
}

void eulabeia_destroy(struct EulabeiaClient *ec)
{
	if (ec == NULL)
		return;
	mqtt_unsubscribe(EULABEIA_INFO);
	if (!already_connected) {
		g_info("closing mqtt connection");
		mqtt_reset();
	}
	free(ec);
}

int eulabeia_json_object(const char *payload,
			 JsonNode **j_node,
			 JsonObject **j_obj)
{
	if (payload == NULL)
		return -1;
	if ((*j_node = json_from_string(payload, NULL)) == NULL) {
		return -2;
	}
	if (!JSON_NODE_HOLDS_OBJECT(*j_node) ||
	    (*j_obj = json_node_get_object(*j_node)) == NULL) {
		json_node_free(*j_node);
		*j_node = NULL;
		return -3;
	}
	return 0;
}

int eulabeia_scan_progress(const char *payload,
			   struct EulabeiaScanProgress *progress)
{
	JsonNode *j_node;
	JsonObject *j_obj;
	struct EulabeiaStatus *status = NULL;
	struct EulabeiaFailure *failure = NULL;
	struct EulabeiaMessage *msg = NULL;
	int rc;

	if (payload == NULL || progress == NULL || progress->id == NULL) {
		rc = -1;
		goto clean_exit;
	}
	if ((rc = eulabeia_json_object(payload, &j_node, &j_obj)) != 0)
		goto clean_exit;

	if ((rc = eulabeia_json_message(j_obj, &msg)) < 0) {
		rc = -4;
		goto clean_exit;
	}
	if (eulabeia_json_status(j_obj, msg, &status) == 0) {
		if (strcmp(progress->id, status->id) != 0) {
			rc = 1;
		} else {
			rc = 0;
			if (status->status == NULL) {
				rc = -5;
				g_warning("status is null.");
			}
#define X(a, b)                                                                \
	else if (strcmp(status->status, #b) == 0) { progress->status = a; }
			EULABEIA_SCAN_RESULT_STATES
#undef X
			else
			{
				rc = -5;
				g_warning("Unknown status: %s", status->status);
			}
		}
	} else if (eulabeia_json_failure(j_obj, msg, &failure) == 0) {
		if (strcmp(progress->id, failure->id) == 0) {
			rc = 0;
			g_warning("scan (%s) failed with: %s",
				  progress->id,
				  failure->error ? failure->error : "N/A");
			progress->status = EULABEIA_SCAN_RESULT_FAILED;
		} else {
			rc = 2;
		}
	} else {
		rc = 3;
	}

clean_exit:
	if (msg)
		eulabeia_message_destroy(&msg);
	if (j_node)
		json_node_free(j_node);
	if (status)
		eulabeia_status_destroy(&status);
	if (failure)
		eulabeia_failure_destroy(&failure);
	return rc;
}

// @brief verify_data is used within publish message to verify given data
//
// @param[in] data the data struct to verify
// @return -1 if data is NULL, -2 if data is invalid or 0 if data is valid.
typedef int verify_data(void *data);

// @brief converts given data to json
//
// @param[in] em is included to the json to identify the message
// @param[in] data the business data to be included to the json
// @return json string on success or NULL on failure
typedef char *to_json(struct EulabeiaMessage *em, void *data);

// @brief calulates the topic to send the message into
//
// @param[in] mt is used to identify the event type
// @param[in] aggregate is used as the aggregate
// @param[in] context is used as the context; if NULL EULABEIA_SCANNER_CONTEXT
// is used.
// @param[in] destination is used to set the destination; if NULL then
// destination part will be skipped.
//
// @return the topic to send the message into. The result must be freed by the
// caller.
char *eulabeia_calulate_topic(enum eulabeia_message_type mt,
			      enum eulabeia_aggregate aggregate,
			      const char *context,
			      const char *destination)
{
	const char *c, *e, *a;
	char *result;
	unsigned int len;
	c = context ? context : EULABEIA_SCANNER_CONTEXT;
	e = eulabeia_message_type_to_event_type(mt);
	a = eulabeia_aggregate_to_str(aggregate);
	len = strlen(context) + 1 + strlen(a) + 1 + strlen(e) +
	      (destination ? strlen(destination) + 1 : 0) + 1;
	result = calloc(1, len);
	// <context>/<aggregate>/<event>/<destination>
	snprintf(result,
		 len,
		 "%s/%s/%s/%s",
		 c,
		 a,
		 e,
		 destination ? destination : "");
	return result;
}

// @brief skeleton method to publish a message.
//
// @return 0 on success, -1 when either ec or data is null, -2 when data in
// invalid; -3 when data could not be published
int publish_message(const struct EulabeiaClient *ec,
		    enum eulabeia_message_type mt,
		    enum eulabeia_aggregate a,
		    char *group_id,
		    void *data,
		    const char *destination,
		    verify_data verifier,
		    to_json tj)
{
	char *json, *topic;
	int rc;
	struct EulabeiaMessage *message;

	if (ec == NULL) {
		rc = -1;
		goto exit;
	}
	if ((rc = verifier(data)) != 0) {
		goto exit;
	}
	message = eulabeia_initialize_message(mt, a, group_id);
	json = tj(message, data);
	topic = eulabeia_calulate_topic(
	    mt, a, EULABEIA_SCANNER_CONTEXT, destination);
	if (ec->publish(topic, json, ec->context) != 0) {
		g_warning("unable to send %s to %s", json, topic);
		rc = -3;
	} else {
		rc = 0;
	}
	free(topic);
	free(message);
	free(json);
exit:
	return rc;
}

// @brief verifies scan_data according to @see verify_data.
int verify_scan_data(struct EulabeiaScan *scan)
{
	if (scan == NULL) {
		return -1;
	}
	if (scan->id == NULL) {
		return -2;
	}
	return 0;
}

// @brief verifies target_data according to @see verify_data.
int verify_target_data(struct EulabeiaTarget *target)
{
	if (target == NULL) {
		return -1;
	}
	if (target->id == NULL) {
		return -2;
	}
	return 0;
}

int eulabeia_start_scan(const struct EulabeiaClient *eulabeia_client,
			const struct EulabeiaScan *scan,
			const char *group_id)
{
	return publish_message(eulabeia_client,
			       EULABEIA_CMD_START,
			       EULABEIA_SCAN,
			       (char *)group_id,
			       (void *)scan,
			       EULABEIA_DIRECTOR,
			       (verify_data *)verify_scan_data,
			       (to_json *)eulabeia_scan_message_to_json);
}

int eulabeia_modify_scan(const struct EulabeiaClient *eulabeia_client,
			 const struct EulabeiaScan *scan,
			 const char *group_id)
{
	return publish_message(eulabeia_client,
			       EULABEIA_CMD_MODIFY,
			       EULABEIA_SCAN,
			       (char *)group_id,
			       (void *)scan,
			       EULABEIA_DIRECTOR,
			       (verify_data *)verify_scan_data,
			       (to_json *)eulabeia_scan_message_to_json);
}

int eulabeia_modify_target(const struct EulabeiaClient *eulabeia_client,
			   const struct EulabeiaTarget *target,
			   const char *group_id)
{
	return publish_message(eulabeia_client,
			       EULABEIA_CMD_MODIFY,
			       EULABEIA_TARGET,
			       (char *)group_id,
			       (void *)target,
			       EULABEIA_DIRECTOR,
			       (verify_data *)verify_target_data,
			       (to_json *)eulabeia_target_message_to_json);
}

int eulabeia_scan_finished(const struct EulabeiaScanProgress *progress)
{
	if (progress == NULL)
		return 0;
	switch (progress->status) {
	case EULABEIA_SCAN_RESULT_STOPPED:
	case EULABEIA_SCAN_RESULT_INTERRUPTED:
	case EULABEIA_SCAN_RESULT_FAILED:
	case EULABEIA_SCAN_RESULT_FINISHED:
		return 1;
	default:
		return 0;
	}
}
