export interface Message {
  text: string;
  sender: string;
  timestamp: string;
}

export interface Room {
  code: string;
  messages: Message[];
}

export type Rooms = {
  [key: string]: Message[];
};

export type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'success';

export interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  className?: string;
  children: React.ReactNode;
}

export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  className?: string;
}

export interface ChatBubbleProps {
  message: Message;
  isOwn: boolean;
}

export interface Game {
  id: string;
  name: string;
  description: string;
  imageUrl?: string;
  minPlayers?: number;
  maxPlayers?: number;
}

export interface GameCardProps {
  game: Game;
  onSelect: (gameId: string) => void;
}

export interface ChatHeaderProps {
  roomCode: string;
  username: string;
  onLeaveRoom: () => void;
}

export interface ChatInputProps {
  onSendMessage: (message: string) => void;
}

export interface ChatBarProps {
  roomCode: string;
  username: string;
  messages: Message[];
  onSendMessage: (message: string) => void;
  onLeaveRoom: () => void;
  onClose?: () => void;
}

export interface GameRoomProps {
  game: string;
  roomCode: string;
  username: string;
  onSendMessage: (message: string) => void;
  onLeaveRoom: () => void;
}

export interface RoomCreationProps {
  onCreateRoom: (roomCode: string, username: string) => void;
}

export interface RoomJoinProps {
  onJoinRoom: (roomCode: string, username: string) => void;
}
