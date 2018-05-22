package main

import (
	"agones.dev/agones/sdks/go"
	"github.com/prometheus/common/log"
	"time"
)

func main() {
	// Create SDK
	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatal(err)
	}

	// Mark server ready
	err = s.Ready()
	if err != nil {
		log.Fatal(err)
	}

	// Every 2 seconds send a health check
	tick := time.Tick(2 * time.Second)
	for {
		err := s.Health()
		if err != nil {
			log.Fatal(err)
		}
		select {
		case <-tick:
		}
	}
}