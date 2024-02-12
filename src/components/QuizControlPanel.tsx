import { useState, ChangeEvent, CSSProperties } from "react";

const MIN_PLAYERS = 4;
const MAX_PLAYERS = 12;

const QuizControlPanel = () => {
  const [quizConfig, setQuizConfig] = useState({
    title: "",
    code: "",
    players: 4,
  })
  
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const target = e.target;

    setQuizConfig((previous) => ({
      ...previous,
      [target.id]: target.value,
    }))
  }
  
  return (
    <div className="flex flex-col gap-4 col-span-3 bg-white my-8 mx-4 rounded-md p-4">
      <h2 className="text-5xl text-purple-500 font-bold mb-8">Quiz Options</h2>
      <label htmlFor="title" className="text-purple-500 text-xl font-semibold">Quiz title</label>
      <input
        id="title"
        className="py-6 rounded-md px-4 shadow-md"
        type="text"
        placeholder="Quiz Title"
        value={quizConfig.title}
        onChange={handleChange}
      />
      <label htmlFor="code" className="text-purple-500 text-xl font-semibold">Quiz code</label>
      <input
        id="code"
        className="py-6 rounded-md px-4 shadow-md"
        type="text"
        placeholder="Quiz Entry Code"
        value={quizConfig.code}
        onChange={handleChange}
      />
      <label htmlFor="players" className="text-purple-500 text-xl font-semibold">Nº of players</label>
      <div className="flex gap-4">
        <div className="flex gap-6 w-full shadow-md py-4 px-4 rounded-md">
          <input 
            style={{ "--percentage": `${((quizConfig.players - MIN_PLAYERS) / (MAX_PLAYERS - MIN_PLAYERS)) * 100}%`} as CSSProperties}
            className="w-4/5" 
            id="players"
            type="range" 
            min={MIN_PLAYERS} 
            max={MAX_PLAYERS} 
            value={quizConfig.players}
            onChange={handleChange}
          />
          <span className="text-2xl text-purple-500">{quizConfig.players < 10 ? "0" + quizConfig.players : quizConfig.players}</span>
        </div>
      </div>
      <button
        className="py-4 px-12 rounded-md bg-purple-500 text-white text-xl font-bold hover:bg-purple-700"
      >
        Create Quiz 
      </button>
    </div>
  );
};

export default QuizControlPanel;
