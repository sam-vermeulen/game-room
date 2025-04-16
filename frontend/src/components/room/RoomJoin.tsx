import React, { useState } from 'react';
import Button from '../ui/Button';
import Input from '../ui/Input';
import { RoomJoinProps } from '../../types';
import { joinRoom, constructWsUrl } from '../../services/api';

interface ExtendedRoomJoinProps extends RoomJoinProps {
  setWsUrl: (url: string) => void;
}

const RoomJoin: React.FC<ExtendedRoomJoinProps> = ({ onJoinRoom, setWsUrl }) => {
  const [roomCode, setRoomCode] = useState('');
  const [username, setUsername] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleJoin = async (e: React.FormEvent) => {
    e.preventDefault();
    if (roomCode.trim() && username.trim()) {
      setIsLoading(true);
      setError(null);
      try {
        const joinResponse = await joinRoom(roomCode.toUpperCase(), username);

        const wsUrl = constructWsUrl(roomCode, joinResponse.token);

        setWsUrl(wsUrl);

        onJoinRoom(roomCode.toUpperCase(), username);
      } catch (err) {
        setError('Room does not exists.');
        console.log(err);
      } finally {
        setIsLoading(false);
      }
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-lg">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Join a Room</h2>
      {error && (
        <div className="mb-4 p-2 bg-red-100 text-red-700 rounded">
          {error}
        </div>
      )}
      <form onSubmit={handleJoin} className="space-y-4">
        <Input
          type="text"
          value={roomCode}
          onChange={(e) => setRoomCode(e.target.value)}
          placeholder="Enter room code"
          required
          disabled={isLoading}
        />
        <Input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter your username"
          required
          disabled={isLoading}
        />
        <Button 
          type="submit" 
          className="w-full"
          disabled={isLoading}
        >
          {isLoading ? 'Joining...' : 'Join Room'}
        </Button>
      </form>
    </div>
  );
};

export default RoomJoin;
