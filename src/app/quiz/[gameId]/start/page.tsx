import { QuizStatus } from "@/types/quiz";

type StartGamePageProps = {
  params: { gameId: string };
};

const MOCK_GAME_DATA = {
  title: "test game 1",
  code: "TEST_GAME",
  current_question: 0,
  question_count: 10,
  max_players: 10,
  active_players: 5,
  players: [
    { id: "player1", name: "player 1", score: 0 },
    { id: "player2", name: "player 2", score: 0 },
    { id: "player3", name: "player 3", score: 0 },
    { id: "player4", name: "player 4", score: 0 },
    { id: "player5", name: "player 5", score: 0 },
  ],
  status: QuizStatus.WAITING,
};

function QuizStatusBagde({ status }: { status: QuizStatus }) {
  const badgeStyle =
    "text-white px-6 py-2 text-xl font-bold rounded-md shadow-md";

  switch (status) {
    case QuizStatus.WAITING:
      return <span className={`${badgeStyle} bg-yellow-400`}>WAITING</span>;

    case QuizStatus.STARTED:
      return <span className={`${badgeStyle} bg-green-400`}>STARTED</span>;

    case QuizStatus.FINISHED:
      return <span className={`${badgeStyle} bg-red-500`}>FINISHED</span>;

    default:
      return <span className={`${badgeStyle} bg-gray-400`}>UNKNOWN</span>;
  }
}

function PlayerCard({
  player,
}: {
  player: { id: string; name: string; score: number };
}) {
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

export default function StartGamePage({ params }: StartGamePageProps) {
  console.log(params.gameId);
  return (
    <main className="grid grid-cols-12 h-screen w-full bg-gradient py-12">
      <div className="flex flex-col h-fit col-start-3 col-span-8 bg-white p-8 gap-8 rounded-md shadow-md">
        <h2 className="text-5xl text-gradient font-bold mb-4">
          {MOCK_GAME_DATA.title.toUpperCase()}
        </h2>
        <div className="grid grid-cols-2 w-2/3 gap-8">
          <h3 className="text-3xl text-gray-700 font-bold">
            CODE: <span className="text-blue-500">{MOCK_GAME_DATA.code}</span>
          </h3>
          <div className="flex items-center gap-4">
            <h3 className="text-3xl text-gray-700 font-bold">STATUS:</h3>
            <QuizStatusBagde status={MOCK_GAME_DATA.status} />
          </div>
          <h3 className="text-3xl text-gray-700 font-bold">
            PLAYERS:{" "}
            <span className="text-blue-500">
              {MOCK_GAME_DATA.active_players} / {MOCK_GAME_DATA.max_players}
            </span>
          </h3>
          <h3 className="text-3xl text-gray-700 font-bold">
            QUESTIONS:{" "}
            <span className="text-blue-500">
              {MOCK_GAME_DATA.current_question + 1} /{" "}
              {MOCK_GAME_DATA.question_count}
            </span>
          </h3>
        </div>
        <div className="grid grid-cols-2 w-full gap-8 max-h-96 overflow-y-auto p-8 bg-gray-100 rounded-md">
          <h3 className="col-span-2 text-3xl text-gray-700 font-bold">
            PLAYERS' SCORES
          </h3>
          {MOCK_GAME_DATA.players.map((player) => {
            return <PlayerCard player={player} key={player.id} />;
          })}
        </div>
        <button 
          className={`
            self-start py-4 px-16 bg-purple-500 
            text-white text-2xl font-bold rounded-md 
            hover:bg-purple-700 disabled:bg-gray-400
          `}
          disabled={MOCK_GAME_DATA.status !== QuizStatus.WAITING}
        >
          Start Quiz
        </button>
      </div>
    </main>
  );
}
