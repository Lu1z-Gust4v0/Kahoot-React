import { Question } from "@/types/question";

type SlideProps = {
  questions: Question[];
  selectQuestion: (id: string) => void;
  addNewQuestion: () => void;
};

type QuestionSlideProps = {
  position: number;
  question: Question;
  selectQuestion: (id: string) => void;
};

const QuestionSlide = ({
  position,
  question,
  selectQuestion,
}: QuestionSlideProps) => {
  return (
    <div
      className={`
        flex items-center justify-center 
        text-4xl text-white bg-gray-300 
        h-40 w-full rounded-sm cursor-pointer
        hover:bg-gray-400 transition-all duration-300 ease-out
      `}
      onClick={() => selectQuestion(question.id)}
    >
      {position}
    </div>
  );
};

const Slide = ({ questions, selectQuestion, addNewQuestion }: SlideProps) => {
  return (
    <div className="grid col-span-2 content-start bg-white px-4 py-4 gap-4 overflow-y-auto">
      {questions.map((question, index) => {
        return (
          <QuestionSlide
            key={question.id}
            position={index + 1}
            question={question}
            selectQuestion={selectQuestion}
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
