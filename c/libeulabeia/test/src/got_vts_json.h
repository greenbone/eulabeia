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

#ifndef GOT_VTS_JSON_H
#define GOT_VTS_JSON_H

#define GOT_VT_ID "1.3.6.1.4.1.25623.1.0.90022"
#define GOT_VT                                                                 \
	"{"                                                                    \
	"\"message_created\": 646146362,"                                      \
	"\"message_type\": \"got.vt\","                                        \
	"\"message_id\": \"842a184a-cabc-4ac2-bc5c-1091e318b5f2\","            \
	"\"group_id\": \"4188bbf8-235e-4645-a653-7a01b723bb55\","              \
	"\"id\": \"1.3.6.1.4.1.25623.1.0.90022\","                             \
	"\"name\": \"mqtt test\","                                             \
	"\"filename\": \"test.nasl\","                                         \
	"\"required_keys\": \"test/key2\","                                    \
	"\"mandatory_keys\": \"test/key1\","                                   \
	"\"excluded_keys\": \"1, 2\","                                         \
	"\"required_ports\": \"\","                                            \
	"\"required_udp_ports\": \"\","                                        \
	"\"category\": \"0\","                                                 \
	"\"family\": \"my test family\","                                      \
	"\"created\": \"1427454000\","                                         \
	"\"modified\": \"1573399828\","                                        \
	"\"summary\": \"A short description of the problem\","                 \
	"\"solution\": \"Solution description\","                              \
	"\"solution_type\": \"Type of solution (e.g. mitigation, vendor "      \
	"fix)\","                                                              \
	"\"solution_method\": \"how to solve it (e.g. debian apt upgrade)\","  \
	"\"impact\": \"Some detailed about what is impacted\","                \
	"\"insight\": \"Some detailed insights of the problem\","              \
	"\"affected\": \"Affected programs, operation system, ...\","          \
	"\"vuldetect\": \"Describes what this plugin is doing to detect a "    \
	"vulnerability.\","                                                    \
	"\"qod_type\": \"package\","                                           \
	"\"qod\": \"0\","                                                      \
	"\"references\": ["                                                    \
	"{"                                                                    \
	"\"type\": \"CVE\","                                                   \
	"\"id\": \"CVE-0000-0000\""                                            \
	"},"                                                                   \
	"{"                                                                    \
	"\"type\": \"CVE\","                                                   \
	"\"id\": \"CVE-0000-0001\""                                            \
	"},"                                                                   \
	"{"                                                                    \
	"\"type\": \"Example\","                                               \
	"\"id\": \"GB-Test-1\""                                                \
	"},"                                                                   \
	"{"                                                                    \
	"\"type\": \"URL\","                                                   \
	"\"id\": \"https://www.greenbone.net\""                                \
	"}"                                                                    \
	"],"                                                                   \
	"\"vt_parameters\": ["                                                 \
	"{"                                                                    \
	"\"id\": 1,"                                                           \
	"\"name\": \"example\","                                               \
	"\"value\": \"\","                                                     \
	"\"type\": \"entry\","                                                 \
	"\"description\": \"\","                                               \
	"\"default\": \"a default string value\""                              \
	"}"                                                                    \
	"],"                                                                   \
	"\"vt_dependencies\": ["                                               \
	"\"keys.nasl\""                                                        \
	"],"                                                                   \
	"\"severety\": {"                                                      \
	"\"severity_vector\": \"AV:N/AC:L/Au:N/C:N/I:N/A:N\","                 \
	"\"severity_type\": \"cvss_base_v2\","                                 \
	"\"severity_date\": \"1427454000\","                                   \
	"\"severity_origin\": \"NVD\""                                         \
	"}"                                                                    \
	"}"

#endif
