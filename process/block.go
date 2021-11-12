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

package process

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/greenbone/eulabeia/logging"
	"io"
)

var log = logging.Logger()

func Block(c ...io.Closer) {
	BlockUntil(func() {
		log.Info().Msg("Exiting")
		for _, cl := range c {
			if cl != nil {
				err := cl.Close()
				if err != nil {
					log.Error().Msgf("failed to send Disconnect: %s", err)
				}
			}
		}

	}, os.Interrupt, syscall.SIGTERM)
}

func BlockUntil(exec func(), signs ...os.Signal) {
	ic := make(chan os.Signal, 1)
	defer close(ic)
	signal.Notify(ic, signs...)
	<-ic
	exec()
}
