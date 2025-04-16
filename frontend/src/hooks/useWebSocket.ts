import { useEffect, useRef, useCallback, useState } from 'react';
import { WebSocketMessage } from '../types/api';

interface UseWebSocketProps {
  url: string | null;
  onMessage: (message: WebSocketMessage) => void;
  onClose?: () => void;
}

export const useWebSocket = ({ url, onMessage }: UseWebSocketProps) => {
  const ws = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);

  const sendMessage = useCallback((type: string, payload: string) => {
    if (ws.current?.readyState === WebSocket.OPEN) {
      const message: WebSocketMessage = {
        type: type,
        payload: payload 
      };
      console.log(message)
      ws.current.send(JSON.stringify(message));
    }
  }, []);

  const closeConnection = useCallback(() => {
    if (ws.current) {
      ws.current.close();
      ws.current = null;
      setIsConnected(false);
    }
  }, []);

  useEffect(() => {
    if (url) { 
      const socket = new WebSocket(url);
      ws.current = socket;

      socket.onopen = () => {
        console.log('WebSocket connected');
        setIsConnected(true);
      }

      socket.onmessage = (event) => {
        const message: WebSocketMessage = JSON.parse(event.data); 
        onMessage(message);
      };

      socket.onclose = () => { 
        setIsConnected(false);
        closeConnection();
      };
    }

    return () => {
      closeConnection();
    };
  }, [url, onMessage, closeConnection]);

  return { sendMessage, isConnected, closeConnection };
};
