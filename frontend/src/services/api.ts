import { CreateRoomResponse, JoinRoomResponse } from '../types/api';

const API_BASE_URL = 'http://192.168.2.131:8080/api';
const WS_BASE_URL = 'http://192.168.2.131:8080/ws';

export const createRoom = async (playerName: string): Promise<CreateRoomResponse> => {
  const response = await fetch(`${API_BASE_URL}/room/create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ playerName }),
  });

  if (!response.ok) {
    throw new Error('Failed to create room');
  }
  
  return response.json();
};

export const joinRoom = async (roomCode: string, playerName: string): Promise<JoinRoomResponse> => {
  const response = await fetch(`${API_BASE_URL}/room/${roomCode}/join`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ playerName }),
  });

  if (!response.ok) {
    throw new Error('Failed to join room');
  }
  
  const data = await response.json();
  return data;
};

export const constructWsUrl = (roomCode: string, token: string): string => {
  return `${WS_BASE_URL}/room/${roomCode}?token=${token}`;
};
