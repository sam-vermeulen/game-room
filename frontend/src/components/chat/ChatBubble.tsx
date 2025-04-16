import React from 'react';
import { ChatBubbleProps } from '../../types';

const ChatBubble: React.FC<ChatBubbleProps> = ({ message, isOwn }) => {
  return (
    <div className={`flex ${isOwn ? 'justify-end' : 'justify-start'} mb-4`}>
      <div
        className={`max-w-[75%] px-4 py-2 rounded-lg break-words ${
          isOwn
            ? 'bg-blue-500 text-white rounded-br-none'
            : 'bg-gray-200 text-gray-800 rounded-bl-none'
        }`}
      >
        <p className="font-semibold text-sm">{message.sender}</p>
        <p className="mt-1 whitespace-pre-wrap">{message.text}</p>
        <p className="text-xs mt-1 opacity-75">
          {new Date(message.timestamp).toLocaleTimeString()}
        </p>
      </div>
    </div>
  );
};

export default ChatBubble;
