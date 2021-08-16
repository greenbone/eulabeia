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

#ifndef __EULABEIA_JSON_H
#define __EULABEIA_JSON_H
#include <eulabeia/types.h>
#include <json-glib/json-glib.h>
/*
 * @brief inititalizes JsonNode and JsonObject.
 *
 * @param[out] j_node; the JsonNode to be initialized. The caller is responsible
 * for j_node.
 * @param[out] j_obj; the JsonObject to be initialized based on the j_node. The
 * caller is responsible for j_obj.
 * @return 0 on success; -1 when payload is empty; -2 when payload is not a
 * json; -3 when payload is not a json object.
 */
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

/*
 * @brief parses already initialized JsonObject to EulabeiaFailure
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[in] msg, the EulabeiaMessage to be included into failure
 * @param[out] failure, the EulabeiaFailure this function will allocate the
 * memory. The caller is responsible for cleaning.
 *
 * @return 0 on success,
 *  -1 on msg is NULL
 *  -2 on msg type is incorrect
 *  -3 on JsonObject is incorrect
 *  -4 on allocation failure
 *  -5 on missing ID
 *  -6 on failure to set value
 */
int eulabeia_json_failure(JsonObject *obj,
			  struct EulabeiaMessage *msg,
			  struct EulabeiaFailure **failure);

/*
 * @brief parses already initialized JsonObject to EulabeiaIdMessage
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[in] type, the eulabeia_message_type to be expected
 * @param[in] msg, the EulabeiaMessage to be included into id_message
 * @param[out] id_message, the EulabeiaIdMessage this function will allocate the
 * memory. The caller is responsible for cleaning.
 *
 * @return 0 on success,
 * @return 0 on success,
 *  -1 on msg is NULL
 *  -2 on msg type is incorrect
 *  -3 on JsonObject is incorrect
 *  -4 on allocation id_message
 *  -5 on missing ID
 */
int eulabeia_json_id_message(JsonObject *obj,
			     enum eulabeia_message_type type,
			     struct EulabeiaMessage *msg,
			     struct EulabeiaIDMessage **id_message);

/*
 * @brief parses already initialized JsonObject to EulabeiaStatus
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[in] msg, the EulabeiaMessage to be included into status
 * @param[out] status, the EulabeiaStatus this function will allocate the
 * memory. The caller is responsible for cleaning.
 *
 * @return 0 on success,
 *  -1 on msg is NULL
 *  -2 on msg type is incorrect
 *  -3 on JsonObject is incorrect
 *  -4 on allocation failure
 *  -5 on missing ID
 *  -6 on failure to set value
 */
int eulabeia_json_status(JsonObject *obj,
			 struct EulabeiaMessage *msg,
			 struct EulabeiaStatus **status);
/*
 * @brief parses already initialized JsonObject to EulabeiaHosts
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[out] hosts, the EulabeiaHosts this function will allocate the memory.
 * The caller is responsible for cleaning.
 *
 * @return 0 on success,
 *  -1 on invalid JsonObject
 *  -2 on host allocation failure
 *  -3,-4 on setting value failure
 */
int eulabeia_json_hosts(JsonArray *arr, struct EulabeiaHosts **hosts);
/*
 * @brief parses already initialized JsonObject to EulabeiaPlugins
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[out] plugins, the EulabeiaPlugins this function will allocate the
 * memory. The caller is responsible for cleaning.
 *
 * @return 0 on success,
 *  -1 on invalid JsonObject
 *  -2 on plugin allocation failure
 *  -3,-4 on setting value failure
 */
int eulabeia_json_plugins(JsonArray *arr, struct EulabeiaPlugins **plugins);
/*
 * @brief parses already initialized JsonObject to EulabeiaPorts
 *
 * This function expects an already initialized JsonObject. To initialize one
 * you can call eulabeia_json_jsonobject.
 *
 * @param[in] obj, the JsonObject to be parsed
 * @param[out] ports, the EulabeiaPorts this function will allocate the memory.
 * The caller is responsible for cleaning.
 *
 * @return 0 on success,
 *  -1 on invalid JsonObject
 *  -2 on port allocation failure
 *  -3,-4 on setting value failure
 */
int eulabeia_json_ports(JsonArray *arr, struct EulabeiaPorts **ports);
/*
 * @brief transforms EulabeiaScan to json string.
 *
 * @param[in] msg, the EulabeiaMessage to include
 * @param[in] scan, the scan to transform to json string.
 * @param[in] modify, 0 for not a modify message and 1 for it is a modify
 * message
 * @return a json char array or NULL on failure.
 */
char *eulabeia_scan_message_to_json(const struct EulabeiaMessage *msg,
				    const struct EulabeiaScan *scan,
				    const int modify);
/*
 * @brief transforms EulabeiaTarget to json string.
 *
 * @param[in] msg, the EulabeiaMessage to include
 * @param[in] target, the target to transform to json string.
 * @param[in] modify, 0 for not a modify message and 1 for it is a modify
 * message
 * @return a json char array or NULL on failure.
 */
char *eulabeia_target_message_to_json(const struct EulabeiaMessage *msg,
				      const struct EulabeiaTarget *target,
				      const int modify);
#endif
