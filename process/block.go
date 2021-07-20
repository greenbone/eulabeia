// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package process

import (
	"os"
	"os/signal"
	"syscall"

	"io"
	"log"
)

func Block(c io.Closer) {
	BlockUntil(func() {
		log.Println("Exiting")
		if c != nil {
			err := c.Close()
			if err != nil {
				log.Fatalf("failed to send Disconnect: %s", err)
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
