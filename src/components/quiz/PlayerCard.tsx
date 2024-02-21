import { Player } from "@/types/websocket";

type PlayerCardProps = {
  player: Player;
};

function PlayerCard({ player }: PlayerCardProps) {
  return (
    <div className="flex flex-col bg-white shadow-md rounded-md px-8 py-4 gap-4">
      <h4 className="text-gradient text-2xl font-bold uppercase">
        {player.name}
      </h4>
      <div className="flex items-center gap-4">
        <p className="text-xl text-gray-700 font-bold">SCORE:</p>
        <span className="text-xl text-blue-500 font-bold">{player.score}</span>
      </div>
    </div>
  );
}

export default PlayerCard;
