import React, { useEffect, useRef } from 'react';
import ChatInput from './ChatInput';
import ChatBubble from './ChatBubble';
import { X } from 'lucide-react';
import { ChatBarProps } from '../../types';

const ChatBar: React.FC<ChatBarProps> = ({ 
  username, 
  messages, 
  onSendMessage, 
  onClose
}) => {
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  return (
    <div className="bg-white shadow-lg overflow-hidden h-full flex flex-col">
      <div className="h-20 bg-gray-600 text-white flex items-center px-4">
        <h2 className="text-xl font-bold flex-1">Chat</h2>
        <button
          onClick={onClose}
          className="md:hidden p-2 hover:bg-blue-700 rounded-full transition-colors ml-auto"
        >
          <X size={24} />
        </button>
      </div>

      <div className="flex-1 overflow-y-auto p-4 bg-gray-50">
        {messages.map((message, index) => (
          <ChatBubble
            key={index}
            message={message}
            isOwn={message.sender === username}
          />
        ))}
        <div ref={messagesEndRef} />
      </div>

      <ChatInput onSendMessage={onSendMessage} />
    </div>
  );
};

export default ChatBar;
