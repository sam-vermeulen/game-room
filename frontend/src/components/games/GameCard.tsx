import React from 'react';
import { GameCardProps } from '../../types';

const GameCard: React.FC<GameCardProps> = ({ game, onSelect }) => {
  return (
    <div 
      className="w-full bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
      onClick={() => onSelect(game.id)}
    >
      <div className="relative">
        <img 
          className="w-full h-48 object-cover" 
          src={game.imageUrl || "https://picsum.photos/400/320"}
          alt={""}
        />
      </div>
      
      <div className="p-4">
        <p className="text-gray-600 text-center item-center text-sm">{game.name}</p>
      </div>
    </div>
  );
};

export default GameCard;
