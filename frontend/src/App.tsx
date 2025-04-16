import React, { useState, useCallback } from 'react';
import RoomCreation from './components/room/RoomCreation';
import RoomJoin from './components/room/RoomJoin';
import GameRoom from './components/games/GameRoom.tsx';
import { Message } from './types';
import { WebSocketMessage } from './types/api';
import { useWebSocket } from './hooks/useWebSocket';


function App() {
  const [currentRoom, setCurrentRoom] = useState<string | null>(null);
  const [username, setUsername] = useState<string>('');
  const [messages, setMessages] = useState<Message[]>([]);
  const [wsUrl, setWsUrl] = useState<string | null>(null);
  const [selectedGame, setSelectedGame] = useState<string | null>(null);

  const handleWebSocketMessage = useCallback((message: WebSocketMessage) => {
    if (message.type === "CHAT") {
      setMessages(prev => [...prev, {
        text: message.payload.text,
        sender: message.payload.sender,
        timestamp: message.payload.timestamp
      }]);
    }
  }, []);

  const { sendMessage, closeConnection } = useWebSocket({
    url: wsUrl,
    onMessage: handleWebSocketMessage
  });

  const handleSendMessage = (message: string) => {
    sendMessage("CHAT", {
        text: message
    });
  };

  const handleLeaveRoom = () => {
    setSelectedGame(null);
    closeConnection();
    setCurrentRoom(null);
    setUsername('');
    setMessages([]);
    setWsUrl(null);
  };

  const handleGameSelect = (gameId: string) => {
    setSelectedGame(gameId);
  };

  return (
    <div className="min-h-screen bg-gray-100">
      {!currentRoom ? (
        <div className="p-4">
          <div className="max-w-4xl mx-auto">
            <h1 className="text-3xl font-bold text-center mb-8 text-gray-800">
              Go Chat App 
            </h1>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <RoomCreation 
                onCreateRoom={(roomCode, user) => {
                  setCurrentRoom(roomCode);
                  setUsername(user);
                }}
                setWsUrl={setWsUrl}
              />
              <RoomJoin 
                onJoinRoom={(roomCode, user) => {
                  setCurrentRoom(roomCode);
                  setUsername(user);
                }}
                setWsUrl={setWsUrl}
              />
            </div>
          </div>
        </div>
      ) : (
        <GameRoom
          game={selectedGame}
          roomCode={currentRoom}
          username={username}
          messages={messages}
          onSendMessage={handleSendMessage}
          onLeaveRoom={handleLeaveRoom}
          onGameSelect={handleGameSelect}
        />
      )}
    </div>
  );
}

export default App;
