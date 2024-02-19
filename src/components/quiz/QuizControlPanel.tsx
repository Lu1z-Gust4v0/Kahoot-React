import { useState } from "react";
import { Quiz } from "@/types/quiz";
import { ChangeEvent, CSSProperties, FormEvent } from "react";
import { MAX_PLAYERS, MIN_PLAYERS } from "./QuizEditor";
import quizService from "@/services/quizService";
import { quizToCreateQuizRequest } from "@/utils/validate";
import { isError } from "@/types/error";
import Toast from "../Toast";

type QuizControlPanelProps = {
  quiz: Quiz;
  editQuiz: (quiz: Quiz) => void;
};

const QuizControlPanel = ({ quiz, editQuiz }: QuizControlPanelProps) => {
  const [toast, setToast] = useState({
    open: false,
    type: "error",
    message: "",
  });

  const toggle = () => {
    setToast((previous) => ({ ...previous, open: !toast.open }));
  };

  const newToast = (type: string, message: string) =>
    setToast({ open: true, type: type, message: message });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const target = e.target;
    
    editQuiz({
      ...quiz,
      [target.id]: target.id === "players" ? parseInt(target.value) : target.value,
    });
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    try {
      console.log(quiz)
      const request = quizToCreateQuizRequest(quiz);

      const response = await quizService.CreateNewQuiz(request);

      newToast("success", "question created successfully")
      console.log(response)
    } catch (error) {
      if (isError(error)) {
        console.log(error.message);
        newToast("error", error.message)
        return;
      }
      console.log(String(error));
      newToast("error", String(error))
    }
  };

  return (
    <>
      {toast.open && (
        <Toast close={toggle} type={toast.type} message={toast.message} />
      )}
      <form
        className="flex flex-col gap-4 col-span-3 bg-white my-8 mx-4 rounded-md p-4"
        onSubmit={handleSubmit}
      >
        <h2 className="text-5xl text-purple-500 font-bold mb-8">
          Quiz Options
        </h2>
        <label
          htmlFor="title"
          className="text-purple-500 text-xl font-semibold"
        >
          Quiz title
        </label>
        <input
          id="title"
          className="py-6 rounded-md px-4 shadow-md"
          type="text"
          placeholder="Quiz Title"
          value={quiz.title}
          onChange={handleChange}
        />
        <label htmlFor="code" className="text-purple-500 text-xl font-semibold">
          Quiz code
        </label>
        <input
          id="code"
          className="py-6 rounded-md px-4 shadow-md"
          type="text"
          placeholder="Quiz Entry Code"
          value={quiz.code}
          onChange={handleChange}
        />
        <label
          htmlFor="players"
          className="text-purple-500 text-xl font-semibold"
        >
          Nº of players
        </label>
        <div className="flex gap-4">
          <div className="flex gap-6 w-full shadow-md py-4 px-4 rounded-md">
            <input
              style={
                {
                  "--percentage": `${((quiz.players - MIN_PLAYERS) /
                      (MAX_PLAYERS - MIN_PLAYERS)) *
                    100
                    }%`,
                } as CSSProperties
              }
              className="w-4/5"
              id="players"
              type="range"
              min={MIN_PLAYERS}
              max={MAX_PLAYERS}
              value={quiz.players}
              onChange={handleChange}
            />
            <span className="text-2xl text-purple-500">
              {quiz.players < 10 ? "0" + quiz.players : quiz.players}
            </span>
          </div>
        </div>
        <button className="py-4 px-12 rounded-md bg-purple-500 text-white text-xl font-bold hover:bg-purple-700">
          Create Quiz
        </button>
      </form>
    </>
  );
};

export default QuizControlPanel;
