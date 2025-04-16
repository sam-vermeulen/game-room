import React, { useState } from 'react';
import Button from '../ui/Button';
import Input from '../ui/Input';
import { RoomCreationProps } from '../../types';
import { createRoom, joinRoom, constructWsUrl } from '../../services/api';

interface ExtendedRoomCreationProps extends RoomCreationProps {
  setWsUrl: (url: string) => void;
}

const RoomCreation: React.FC<ExtendedRoomCreationProps> = ({ onCreateRoom, setWsUrl}) => {
  const [username, setUsername] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (username.trim()) {
      setIsLoading(true);
      setError(null);
      try {
        const createResponse = await createRoom(username);
        const roomCode = createResponse.code;

        const joinResponse = await joinRoom(roomCode, username);
       
        const wsUrl = constructWsUrl(roomCode, joinResponse.token);
        setWsUrl(wsUrl);

        onCreateRoom(joinResponse.code, username);
      } catch (err) {
        setError(`Failed to create room. Please try again: ${err}`);
      } finally {
        setIsLoading(false);
      }
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-lg">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Create a New Room</h2>
      {error && (
        <div className="mb-4 p-2 bg-red-100 text-red-700 rounded">
          {error}
        </div>
      )}
      <form onSubmit={handleCreate} className="space-y-4">
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
          variant="success" 
          className="w-full"
          disabled={isLoading}
        >
          {isLoading ? 'Creating...' : 'Create Room'}
        </Button>
      </form>
    </div>
  );
};

export default RoomCreation;
