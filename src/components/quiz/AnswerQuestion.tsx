"use client";

import { WebsocketMessage } from "@/types/quiz";
import { WebsocketQuestionMessage } from "@/types/websocket";
import { useState } from "react";

type AnswerQuestionProps = {
  question: WebsocketQuestionMessage;
  sendMessage: (message: string) => void;
};

type OptionProps = {
  id: string;
  value: string;
  handleClick: (id: string) => void;
  className?: string;
};

const Option = ({ id, value, handleClick, className }: OptionProps) => {
  return (
    <button
      className={`w-full py-8 px-4 pr-12 rounded-md text-white text-2xl placeholder:text-white ${className}`}
      onClick={() => handleClick(id)}
    >
      {value}
    </button>
  );
};

const AnswerQuestion = ({ question, sendMessage }: AnswerQuestionProps) => {
  const [selected, setSelected] = useState("");
  const [answered, setAnswered] = useState(false);

  const handleClick = (id: string) => setSelected(id);
  const handleAnswer = () => {
    sendMessage(JSON.stringify({
      type: WebsocketMessage.ANSWER,
      answer: selected,
    }))
    setAnswered(true)
  };

  const selectedStyle = "bg-purple-500";
  const correctStyle = "bg-green-500";
  const incorrectStyle = "bg-red-500";

  const getStyle = (id: string, defaultStyle: string) => {
    if (answered && id === question.correct) return correctStyle;
    if (answered && id !== question.correct) return incorrectStyle;
    if (selected === id) return selectedStyle;
    return defaultStyle;
  };

  return (
    <div className="flex flex-col h-fit col-start-3 col-span-8 bg-white p-8 gap-8 rounded-md shadow-md">
      <h2 className="text-5xl text-gradient font-bold mb-4">
        {question.title.toUpperCase()}
      </h2>
      <p className="text-3xl text-gray-700 font-bold mb-4">{question.body}</p>
      <div className="grid grid-cols-2 gap-4">
        <Option
          id="option_one"
          value={question.option_one}
          handleClick={handleClick}
          className={getStyle("option_one", "bg-red-500")}
        />
        <Option
          id="option_two"
          value={question.option_two}
          handleClick={handleClick}
          className={getStyle("option_two", "bg-blue-500")}
        />
        <Option
          id="option_three"
          value={question.option_three}
          handleClick={handleClick}
          className={getStyle("option_three", "bg-yellow-500")}
        />
        <Option
          id="option_four"
          value={question.option_four}
          handleClick={handleClick}
          className={getStyle("option_four", "bg-green-500")}
        />
      </div>
      <button
        className={`
          self-start py-4 px-16 bg-purple-500 
          text-white text-2xl font-bold rounded-md 
          hover:bg-purple-700 disabled:bg-gray-400
        `}
        disabled={answered}
        onClick={() => handleAnswer()}
      >
        Answer
      </button>
    </div>
  );
};

export default AnswerQuestion;
