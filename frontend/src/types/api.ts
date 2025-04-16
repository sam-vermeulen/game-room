export interface CreateRoomResponse {
  code: string;
}

export interface JoinRoomResponse {
  token: string;
  wsUrl: string;
}

export interface WebSocketMessage {
  type: string;
  payload: string;
}
