import React from 'react';
import Button from '../ui/Button';
import { ChatHeaderProps } from '../../types';

const ChatHeader: React.FC<ChatHeaderProps> = ({ 
  roomCode, 
  username, 
  onLeaveRoom 
}) => {
  return (
    <div className="p-4 border-b border-gray-200 bg-white">
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-xl font-semibold text-gray-800">
            Room: {roomCode}
          </h2>
          <p className="text-gray-600">Username: {username}</p>
        </div>
        <Button variant="danger" onClick={onLeaveRoom}>
          Leave Room
        </Button>
      </div>
    </div>
  );
};

export default ChatHeader;
