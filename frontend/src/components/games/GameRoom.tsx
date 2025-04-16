import React, { useState } from 'react';
import ChatBar from '../chat/ChatBar';
import { GameRoomProps } from '../../types';
import { MessageCircle } from 'lucide-react';
import GameGrid from './GameGrid';

const AVAILABLE_GAMES: Game[] = [
  {
    id: 'blackjack',
    name: 'Blackjack',
    imageUrl: '/blackjack.png'
  },
];


const GameRoom: React.FC<GameRoomProps> = ({ 
  game,
  roomCode, 
  username, 
  messages, 
  onSendMessage, 
  onLeaveRoom,
  onGameSelect 
}) => {
  const [isChatOpen, setIsChatOpen] = useState(false);


  return (
    <div className="flex h-screen bg-gray-100 relative">
      <div className="flex-1 flex flex-col">
        <div className="h-20 p-4 bg-blue-600 text-white">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-xl font-bold">Room: {roomCode}</h2>
              <p className="text-sm opacity-90">Player: {username}</p>
            </div>
            <div>
              {game}
            </div>
            <button
              onClick={onLeaveRoom}
              className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
            >
              Leave Room
            </button>
          </div>
        </div>
        
        <div className="flex-1 p-4">
          <div className="w-full h-full bg-gray-200 rounded-lg">
            <GameGrid
              games={AVAILABLE_GAMES}
              onGameSelect={onGameSelect}
            />
          </div>
        </div>
      </div>

      <div className={`
        fixed md:relative top-0 right-0 h-full w-80 md:w-96
        transform transition-transform duration-300 ease-in-out
        ${isChatOpen ? 'translate-x-0' : 'translate-x-full md:translate-x-0'}
        bg-white shadow-lg md:shadow-none
      `}>
        <ChatBar
          roomCode={roomCode}
          username={username}
          messages={messages}
          onSendMessage={onSendMessage}
          onLeaveRoom={onLeaveRoom}
          hideHeader
          onClose={() => setIsChatOpen(false)}
        />
      </div>

      <button
        onClick={() => setIsChatOpen(true)}
        className={`
          md:hidden fixed bottom-4 right-4 p-4 bg-blue-600 text-white rounded-full 
          shadow-lg hover:bg-blue-700 transition-colors
          ${isChatOpen ? 'hidden' : 'flex'}
        `}
      >
        <MessageCircle size={24} />
      </button>
    </div>
  );
};

export default GameRoom;
