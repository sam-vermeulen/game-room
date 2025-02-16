package room

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/sam-vermeulen/go-poker/internal/types"
	"github.com/sam-vermeulen/go-poker/pkg/utils"
)

type RoomManager struct {
	mu         sync.RWMutex
	rooms      map[string]*Room
	joinTokens map[string]JoinToken
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms:      make(map[string]*Room),
		joinTokens: make(map[string]JoinToken),
	}
}

type JoinToken struct {
	PlayerName string
	RoomCode   string
	ExpiresAt  time.Time
}

func (rm *RoomManager) StoreJoinToken(roomCode string, token string, playerName string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.joinTokens[token] = JoinToken{
		PlayerName: playerName,
		RoomCode:   roomCode,
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	}
}

func (rm *RoomManager) VerifyJoinToken(roomCode, token string) (string, bool) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	JoinToken, exists := rm.joinTokens[token]
	if !exists {
		return "", false
	}

	if time.Now().After(JoinToken.ExpiresAt) {
		delete(rm.joinTokens, token)
		return "", false
	}

	if JoinToken.RoomCode != roomCode {
		return "", false
	}

	delete(rm.joinTokens, token)

	return JoinToken.PlayerName, true
}

func (rm *RoomManager) CreateRoom() (string, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for i := 0; i < 10; i++ {
		code := utils.GenerateCode()
		if _, exists := rm.rooms[code]; !exists {
			room := &Room{
				Code:      code,
				Players:   make(map[string]*types.Player),
				createdAt: time.Now(),
			}

			room.OnEmpty = func(code string) {
				rm.removeRoom(code)
			}

			rm.rooms[code] = room
			room.resetCleanupTime(20 * time.Second)

			log.Printf("Created new room with code (%s)", code)
			return code, nil
		}
	}

	return "", errors.New("failed to generated unique room code")
}

func (rm *RoomManager) removeRoom(code string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, exists := rm.rooms[code]; exists {
		delete(rm.rooms, code)
		log.Printf("Room %s removed due to inactivity", code)
	}
}

func (rm *RoomManager) GetRoom(code string) (*Room, error) {
	room := rm.rooms[code]

	if room == nil {
		return nil, errors.New("room does not exist")
	}

	return room, nil
}
