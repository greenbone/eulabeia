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
	signal.Notify(ic, signs...)
	<-ic
	exec()
}
