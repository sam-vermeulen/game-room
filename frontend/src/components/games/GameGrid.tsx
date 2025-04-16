import React from 'react';
import GameCard from './GameCard';
import { Game } from '../../types';

interface GameGridProps {
  games: Game[];
  onGameSelect: (gameId: string) => void;
}

const GameGrid: React.FC<GameGridProps> = ({ games, onGameSelect }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
      {games.map((game) => (
        <GameCard
          key={game.id}
          game={game}
          onSelect={onGameSelect}
        />
      ))}
    </div>
  );
};

export default GameGrid;
