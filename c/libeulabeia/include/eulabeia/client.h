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

#ifndef _EULABEIA_CLIENT_H
#define _EULABEIA_CLIENT_H
#include <eulabeia/types.h>
#include <stdlib.h>
#define EULABEIA_SCANNER_CONTEXT "eulabeia"
#define EULABEIA_DIRECTOR "director"

/**
 * @brief defines an publish method used within eulabeia.
 *
 * For all parameter the caller is responsible.
 *
 * @param[in] topic the topic to send the message into
 * @param[in] message the actual message to send
 * @param[in] context user defined data passed through
 *
 * @return 0 on success otherwise it is a failure.
 */
typedef int
eulabeia_publish(const char *topic, const char *message, void *context);

/**
 * @brief defines an retrieve message method used within eulabeia.
 *
 * The implementation of this method must set payload and topic as well as the
 * length of those so that the caller can work the payload and topic
 * information.
 *
 * @param[out] topic must be allocated and set within the implementation. The
 * caller is responsible for cleaning up.
 * @param[out] topic_len must be set within the implementation. The caller is
 * responsible for cleaning up if necessary.
 * @param[out] payload must be allocated and set within the implementation. The
 * caller is responsible for cleaning up.
 * @param[out] payload_len must be set within the implementation. The caller is
 * responsible for cleaning up if necessary.
 * @param[in] context user defined data passed through
 *
 * @return 0 on success otherwise it is a failure.
 */
typedef int eulabeia_retrieve(char **topic,
			      int *topic_len,
			      char **payload,
			      int *payload_len,
			      void *context);

/*
 * @brief contains the publish and retrieve method.
 *
 * To have an easier way of integrating different messaging methods then mqtt
 * based on gvm the publish and retrieve methods are redefined and have the
 * possibility to pass though user set context data.
 *
 * This struct needs to be initialized before working with eulabeia.
 */
struct EulabeiaClient {
	eulabeia_publish *publish;
	eulabeia_retrieve *retrieve;
	void *context;
};

/*
 * @brief starts a scan.
 *
 * Starts a scan, when the given scan just contains an ID the scan including the
 * target must be created before starting a scan.
 *
 * If you want to start a scan with the containing information the target and at
 * least the hosts within it must be set. The director then creates
 * automatically the create and modify messages as it sees fit.
 *
 * @param[in] eulabeia_client, must be initialized before. Use
 * eulabeia_initialize to use the default mqtt setup.
 * @param[in] scan, contains the needed information to start a scan. With scan
 * you can pick if you want to use an already created scan by just setting the
 * id or if you want to cerate it within the same step by setting the target
 * within scan.
 * @param[in] group_id, set group_id if you want to mark the start scan message
 * to belong to a grouping of messages.
 *
 * @return 0 on success, -1 when eulabeia_client or scan is NULL, -2 when scan
 * does not contain an id, -3 when publishing failed.
 */
int eulabeia_start_scan(const struct EulabeiaClient *eulabeia_client,
			const struct EulabeiaScan *scan,
			const char *group_id);

/*
 * @brief set the scan progress into progress based on payload.
 *
 * This function should be called periodically after eulabeia_start_scan to
 * verify the retrieved payload if it is scan progress relevant and sets the
 * progress accordingly if it is relevant for a scan of progress.
 *
 * progress needs to be inititalized and it's scan_id must be set.
 *
 * @param[in] payload, the payload to verify.
 * @param[in] id, the id to look out for.
 * @param[out] progress, contains the id of the scan and is used to set progress
 *
 * @return 0 when progress got changed,
 * 1 when payload is a valid progress message but not for the given scan id,
 * 1 when payload is a EulabeiaMessage but not ,
 * -1 when either payload, progress or progress->id is NULL,
 * -2 when payload is not valid json, -3 when payload is not a json object,
 * -4 when the payload is not a valid EulabeiaMessage,
 * -5 when the status is not defined in EULABEIA_SCAN_STATES.
 */
int eulabeia_scan_progress(const char *payload,
			   const char *id,
			   struct EulabeiaScanProgress *progress);

/*
 * @brief modifies a scan.
 *
 * TBD
 */
int eulabeia_modify_scan(const struct EulabeiaClient *eulabeia_client,
			 const struct EulabeiaScan *scan,
			 const char *group_id);

/*
 * @brief modifies a target.
 *
 * TBD
 */
int eulabeia_modify_target(const struct EulabeiaClient *eulabeia_client,
			   const struct EulabeiaTarget *target,
			   const char *group_id);

/*
 * @brief set the CRUD progress into progress based on payload.
 *
 * This function should be called periodically after an create or modify to
 * verify the retrieved payload if it is progress relevant and sets the
 * progress accordingly if it is relevant.
 *
 * progress needs to be inititalized upfront.
 *
 * @param[in] payload, the payload to verify.
 * @param[in] id, the id to look out for.
 * @param[in] type, the message type to look for to set it to success
 * @param[out] progress, is used to set progress
 *
 * @return 0 when progress got changed,
 * 1 when payload is a valid progress message but not for the given scan id,
 * 2 when payload is a EulabeiaMessage but not progress relevant,
 * -1 when either payload, progress or progress->id is NULL,
 * -2 when payload is not valid json, -3 when payload is not a json object,
 * -4 when the payload is not a valid EulabeiaMessage,
 * -5 when the status is not defined in EULABEIA_SCAN_STATES.
 */
int eulabeia_crud_progress(const char *payload,
			   const char *id,
			   enum eulabeia_message_type type,
			   struct EulabeiaCRUDProgress *progress);
/*
 * @brief set the modify progress into progress based on payload.
 *
 * This function should be called periodically after an create or modify to
 * verify the retrieved payload if it is progress relevant and sets the
 * progress accordingly if it is relevant.
 *
 * progress needs to be inititalized upfront.
 *
 * @param[in] payload, the payload to verify.
 * @param[in] id, the id to look out for.
 * @param[out] progress, is used to set progress
 *
 * @return 0 when progress got changed,
 * 1 when payload is a valid progress message but not for the given scan id,
 * 2 when payload is a EulabeiaMessage but not progress relevant,
 * -1 when either payload, progress or progress->id is NULL,
 * -2 when payload is not valid json, -3 when payload is not a json object,
 * -4 when the payload is not a valid EulabeiaMessage,
 * -5 when the status is not defined in EULABEIA_SCAN_STATES.
 */
int eulabeia_modify_progress(const char *payload,
			     const char *id,
			     struct EulabeiaCRUDProgress *progress);
/*
 * @brief checks progress if the scan is finished.
 *
 * @return when the progress status is EULABEIA_SCAN_STOPPED,
 * EULABEIA_SCAN_INTERRUPTED or EULABEIA_SCAN_FINISHED then it
 * returns 1 otherwise 0.
 */
int eulabeia_scan_finished(const struct EulabeiaScanProgress *progress);

/*
 * @brief TBD
 *
 * @return the percentage in 0..100
 */
int eulabeia_scan_percent(const struct EulabeiaScan *scan,
			  const struct EulabeiaScanProgress *progress);

/*
 * @brief inititalizes EulabeiaClient to use the default mqtt definition.
 *
 * Init an EulabeiaClient to use gvm-libs#mqtt to publish and retrieve messages.
 * When the gvm-libs#mqtt is not yet inititalized then it is using the given
 * broker_address to init mqtt. When connected it subscribes to the topic
 * relevant for eulabeia_clients so that retrieve is capable of getting
 * messages.
 * @param[in] broker_address, the address of the broker
 * @param[in] context, context will be passed through to eulabeia_publish,
 * eulabeia_retrieve method.
 *
 * @return an initialized EulabeiaClient.
 */
struct EulabeiaClient *eulabeia_initialize(const char *broker_address,
					   void *context);

/*
 * @brief destroys the EulabeiaClient.
 *
 * Frees given EulabeiaClient and unsubscribes from the topic. If the
 * connection was established on eulabeia_initialize then it also resets mqtt
 * connection.
 *
 *	@param[out] client, the EulabeiaClient to be destroyed.
 */
void eulabeia_destroy(struct EulabeiaClient *client);

#endif
