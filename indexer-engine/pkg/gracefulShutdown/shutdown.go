package graceful_shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SafeShutDown(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		fmt.Println("[SHUTDOWNING] ")
		fmt.Println("[SHUTDOWNING] Server is going to shutdown by any reason.")
		fmt.Println("[SHUTDOWNING] 10 seconds to all parts of the system are killed")
		cb()
		time.Sleep(time.Second * 10)
		fmt.Println("[SHUTDOWNING] BYE")
		fmt.Println("[SHUTDOWNING] ")
		os.Exit(0)
	}
}
