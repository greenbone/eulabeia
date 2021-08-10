#ifndef __EULABEIA_JSON_H
#define __EULABEIA_JSON_H
#include <eulabeia/types.h>
#include <json-glib/json-glib.h>

int eulabeia_json_object(const char *payload,
			 JsonNode **j_node,
			 JsonObject **j_obj);

/*
 * @brief parses already initialized JsonObject to EulabeiaMessage
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[out] msg, the EulabeiaMessage; this function will allocate the memory
 * within msg to set the variables. The caller is responsible for cleaning msg.
 *
 * @return 0 on success,
 *  -1 when the JsonObject does not contain all necessary fields,
 *  -2 when allocation of memory for msg fails
 *  -3 when setting message_id failed
 *  -4 when setting message_type failed
 *  -5 when setting group_id failed
 */
int eulabeia_json_message(JsonObject *obj, struct EulabeiaMessage **msg);

int eulabeia_json_failure(JsonObject *obj,
			  struct EulabeiaMessage *msg,
			  struct EulabeiaFailure **failure);
int eulabeia_json_status(JsonObject *obj,
			 struct EulabeiaMessage *msg,
			 struct EulabeiaStatus **status);
int eulabeia_json_hosts(JsonArray *arr, struct EulabeiaHosts **hosts);
int eulabeia_json_plugins(JsonArray *arr, struct EulabeiaPlugins **plugins);
int eulabeia_json_ports(JsonArray *arr, struct EulabeiaPorts **ports);
char *eulabeia_scan_message_to_json(const struct EulabeiaMessage *msg,
				    const struct EulabeiaScan *scan);
char *eulabeia_target_message_to_json(const struct EulabeiaMessage *msg,
				      const struct EulabeiaTarget *target);
#endif
