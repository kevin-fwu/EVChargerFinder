package util

import (
	"os"
	"os/signal"
	"syscall"
)

var chSignaled chan os.Signal
var chStop chan interface{}

func InitSignalHandler() error {
	chStop = make(chan interface{})

	chSignaled = make(chan os.Signal)

	signal.Notify(chSignaled, syscall.SIGTERM)
	signal.Notify(chSignaled, syscall.SIGINT)
	// signal.Notify(chSignaled, syscall.SIGUSR1)
	// signal.Notify(chSignaled, syscall.SIGUSR2)

	go handleSignal(chSignaled)
	return nil
}

func SignalStop() {
	chSignaled <- syscall.SIGTERM
}

func handleSignal(chSig chan os.Signal) {
	for {
		// When I find a need for SIGUSR1/SIGUSR2, handle it here.
		<-chSig
		// if sig == syscall.SIGUSR1 {

		// } else if sig == syscall.SIGUSR2 {

		// } else {
		close(chStop)
		os.Exit(0)
		break
		// }
	}
}
