import { Question } from "@/types/question";
import { IoClose } from "react-icons/io5";

type SlideProps = {
  questions: Question[];
  selectQuestion: (id: string) => void;
  addNewQuestion: () => void;
  removeQuestion: (id: string) => void;
};

type QuestionSlideProps = {
  position: number;
  question: Question;
  selectQuestion: (id: string) => void;
  removeQuestion: (id: string) => void;
};

const QuestionSlide = ({
  position,
  question,
  selectQuestion,
  removeQuestion,
}: QuestionSlideProps) => {
  return (
    <div
      className={`
        relative flex items-center justify-center 
        text-4xl text-white bg-gray-300 
        h-40 w-full rounded-sm cursor-pointer
        hover:bg-gray-400 transition-all duration-300 ease-out
      `}
      onClick={(e) => {
        e.stopPropagation();

        if (e.target !== e.currentTarget) return;

        selectQuestion(question.id);
      }}
    >
      <button
        className="group absolute top-2 right-2"
        onClick={() => removeQuestion(question.id)}
      >
        <IoClose className="w-5 h-5 transition-colors duration-300 ease-out group-hover:text-red-500 group-focus:text-red-500" />
      </button>
      <div>{position}</div>
    </div>
  );
};

const Slide = ({
  questions,
  selectQuestion,
  addNewQuestion,
  removeQuestion,
}: SlideProps) => {
  return (
    <div className="grid col-span-2 content-start bg-white px-4 py-4 gap-4 overflow-y-auto">
      {questions.map((question, index) => {
        return (
          <QuestionSlide
            key={question.id}
            position={index + 1}
            question={question}
            selectQuestion={selectQuestion}
            removeQuestion={removeQuestion}
          />
        );
      })}
      <div className="flex justify-center py-4">
        <button
          className="bg-blue-500 text-white font-bold py-4 px-8 rounded-md hover:bg-blue-700"
          onClick={() => addNewQuestion()}
        >
          Add Question
        </button>
      </div>
    </div>
  );
};

export default Slide;
