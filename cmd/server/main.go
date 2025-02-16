package main

import (
	"github.com/sam-vermeulen/go-poker/internal/api"
	"github.com/sam-vermeulen/go-poker/internal/room"
)

func main() {
	rm := room.NewRoomManager()
	server := api.NewServer(rm)

	server.Start()
}
