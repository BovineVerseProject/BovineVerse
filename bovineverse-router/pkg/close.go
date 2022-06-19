package pkg

import (
	"os"
	"os/signal"
	"syscall"
)

func ShutdownHook() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch
}
