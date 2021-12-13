// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package configuration

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	var out io.Writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	var service string
	var level zerolog.Level = zerolog.TraceLevel
	var warning string
	if s, ok := os.LookupEnv("LOG_OUTPUT"); ok {
		switch s {
		case "JSON":
			out = os.Stdout
		default:
			warning = fmt.Sprintf("Unknown LOG_OUTPUT (%s); using console", s)
		}
	}
	if s, ok := os.LookupEnv("LOG_SERVICE_NAME"); ok {
		service = s
	} else {
		if p, err := os.Executable(); err != nil {
			service = "unknown"
		} else {
			_, f := path.Split(p)
			service = f
		}
	}
	if s, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if l, err := zerolog.ParseLevel(s); err != nil {
			warning = fmt.Sprintf(
				"Unable to identify log level (%s) fallback to TraceLevel",
				s,
			)
		} else {
			level = l
		}
	}
	log.Logger = zerolog.New(out).
		With().
		Str("service", service).
		Caller().
		Timestamp().
		Logger().
		Level(level)
	if warning != "" {
		log.Warn().Msg(warning)
	}

}
