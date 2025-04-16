package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sam-vermeulen/go-poker/internal/room"
	"github.com/sam-vermeulen/go-poker/internal/types"
	"github.com/sam-vermeulen/go-poker/pkg/utils"
)

type Server struct {
	router   *mux.Router
	rm       *room.RoomManager
	server   *http.Server
	upgrader websocket.Upgrader
}

func NewServer(rm *room.RoomManager) *Server {
	r := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Be more restrictive in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	s := &Server{
		router: r,
		rm:     rm,
		server: &http.Server{
			Addr:         ":8080",
			Handler:      corsHandler.Handler(r),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	s.setupRoutes()
	return s
}

func (s *Server) Start() error {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server starting on %s\n", s.server.Addr)
		serverErrors <- s.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %v", err)
	case sig := <-stop:
		log.Printf("Server shutting down due to %v signal\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			return fmt.Errorf("graceful shutdown failed: %v\n", err)
		}

		log.Printf("Server gracefully shutdown")
		return nil
	}
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/api/room/create", s.handleCreateRoom).Methods("POST")
	s.router.HandleFunc("/api/room/{code}/join", s.handleJoinRoom).Methods("POST")
	s.router.HandleFunc("/ws/room/{code}", s.handleWebSocket)
}

func (s *Server) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlayerName string `json:"playerName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	code, err := s.rm.CreateRoom()
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"code": code,
	})
}

func (s *Server) handleJoinRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var req struct {
		PlayerName string `json:"playerName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	rm, err := s.rm.GetRoom(code)
	if err != nil {
		http.Error(w, "Invalid room code", http.StatusNotFound)
		return
	}

	if rm.IsFull() {
		http.Error(w, "Room is full", http.StatusLocked)
	}

	joinToken := utils.GenerateToken()

	s.rm.StoreJoinToken(code, joinToken, req.PlayerName)

	json.NewEncoder(w).Encode(map[string]string{
		"code":  code,
		"wsURL": fmt.Sprintf("ws/room/%s", code),
		"token": joinToken,
	})
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	token := r.URL.Query().Get("token")

	playerName, valid := s.rm.VerifyJoinToken(code, token)

	if !valid {
		http.Error(w, "invalid or expired join token", http.StatusUnauthorized)
		return
	}

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	rm, err := s.rm.GetRoom(code)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": "Room not found"})
		conn.Close()
		return
	}

	player := &types.Player{
		Name: playerName,
		Conn: conn,
	}

	rm.AddPlayer(player)

	go player.HandleConnection(rm)
}
