package main

import (
	"gateway/internal/app"
	"log"
	"os"
	"os/signal"
)

func main() {
	a, cleanup, err := app.InitApp()
	if err != nil {
		panic("failed to init app")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		if err := a.Run(); err != nil {
			log.Printf("app crashed: %s\n", err.Error())
			cleanup()
		}
	}()

	sig := <-sigChan
	cleanup()

	log.Printf("caught sig: %+v. graceful shutdown", sig)
	a.Shutdown()
}
