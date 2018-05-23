package main

import (
	"agones.dev/agones/sdks/go"
	"fmt"
	"github.com/gtaylor/factorio-rcon"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"time"
)

var (
	host     = kingpin.Flag("host", "Host of the Factorio server").Short('h').Default("127.0.0.1").String()
	port     = kingpin.Flag("port", "RCON port of the Factorio server").Short('p').Default("27015").Int()
	password = kingpin.Flag("password", "RCON password of the Factorio server").String()
)

func main() {
	kingpin.Parse()

	// Create SDK
	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Factorio
	r := NewRcon(fmt.Sprintf("%v:%d", *host, *port))

	// Mark the game server as ready
	check(s.Ready())

	// Log into the game server
	check(r.Authenticate(*password))

	// Every 2 seconds run an RCON and send a health check
	tick := time.Tick(2 * time.Second)
	for {
		ExecuteCommand(r, "/version")
		check(s.Health())
		select {
		case <-tick:
		}
	}
}

// NewRcon establishes a new RCON connection to the factorio server.
// We purposely retry until we are able to connect so we are connected as soon as possible.
// Eventually Agones will restart the game server if it isn't ready.
func NewRcon(address string) *rcon.RCON {
	r, err := rcon.Dial(address)
	if err != nil {
		time.Sleep(time.Second)
		return NewRcon(address)
	}
	return r
}

// ExecuteCommand sends a command to the Factorio server to be executed.
// If it fails, the SDK will exit and the server won't respond Healthy
func ExecuteCommand(r *rcon.RCON, command string) {
	_, err := r.Execute(command)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
