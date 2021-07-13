package process

import (
	"os"
	"os/signal"
	"syscall"

	"io"
	"log"
)

// Block the caller until interrupted by e.g. SIGINT or SIGTERM
// Iff the channel has been notified by such an interrupt, the socket/file c will be closed, to aviod any I/O failures.
func BlockAndClose(c io.Closer) {
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
