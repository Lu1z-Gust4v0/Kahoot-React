import { WebsocketScoresMessage } from "@/types/websocket"
import PlayerCard from "./PlayerCard"

type PlayersScoreProps = {
  scores: WebsocketScoresMessage
}

const PlayersScore = ({ scores }: PlayersScoreProps) => {
  const players = scores.players.sort((playerOne, playerTwo) => -(playerOne.score - playerTwo.score))

  return (
    <div className="flex flex-col h-fit col-start-3 col-span-8 bg-white p-8 gap-8 rounded-md shadow-md">
      <h2 className="text-5xl text-gradient font-bold mb-4">
        Players Scores
      </h2>
      <div className="grid grid-cols-2 w-full gap-4 max-h-[30rem] overflow-y-auto p-8 bg-gray-100 rounded-md">
        {players.map((player) => {
          return <PlayerCard player={player} key={player.id} />;
        })}
      </div>
    </div>
  )
}

export default PlayersScore;
