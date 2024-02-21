import { Question } from "@/types/question";
import { ChangeEvent, Dispatch, SetStateAction } from "react";

type QuestionEditorProps = {
  selectedQuestion: Question;
  setSelectedQuestion: Dispatch<SetStateAction<Question>>;
  editQuestion: (questionToUpdate: Question) => void;
};

type QuestionOptionProps = {
  id: string;
  name: string;
  value: string;
  correct: string;
  handleChange: (
    e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLTextAreaElement>,
  ) => void;
  addCorrectOption: (option: string) => void;
  className?: string;
};

const QuestionOption = ({
  id,
  name,
  value,
  correct,
  handleChange,
  addCorrectOption,
  className,
}: QuestionOptionProps) => {
  const isChecked = correct === id

  const toggleCheckbox = () => {
    addCorrectOption(id)
  };

  return (
    <div className="relative flex items-center">
      <input
        id={id}
        className={`w-full py-8 px-4 pr-12 rounded-md text-white text-2xl placeholder:text-white ${className}`}
        type="text"
        placeholder={name}
        onChange={handleChange}
        value={value}
      />
      <label
        className="group absolute flex items-center right-4 cursor-pointer"
        htmlFor={`checkbox-${id}`}
      >
        <input
          className="checkbox"
          type="checkbox"
          name={id}
          id={`checkbox-${id}`}
          onChange={() => toggleCheckbox()}
          checked={isChecked}
        />
        <span className="checkmark"></span>
      </label>
    </div>
  );
};

const QuestionEditor = ({
  selectedQuestion,
  setSelectedQuestion,
  editQuestion,
}: QuestionEditorProps) => {
  const handleChange = (
    e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLTextAreaElement>,
  ) => {
    const target = e.target;

    setSelectedQuestion((previous) => ({
      ...previous,
      [target.id]: target.value,
    }));
  };

  const selectCorrectOption = (option: string) => {
    setSelectedQuestion((previous) => ({
      ...previous,
      correct: option,
    }));
  };

  return (
    <div className="flex flex-col col-span-7 py-8 px-8 gap-8">
      <input
        id="title"
        className="py-6 rounded-md px-4"
        type="text"
        placeholder="Choose a title"
        value={selectedQuestion.title}
        onChange={handleChange}
      />
      <textarea
        id="body"
        className="h-1/3 p-4 resize-none rounded-md"
        placeholder="Question body (markdown)"
        value={selectedQuestion.body}
        onChange={handleChange}
      ></textarea>
      <div className="grid grid-cols-2 gap-4">
        <QuestionOption
          id="option_one"
          name="option 01"
          className="bg-red-500"
          value={selectedQuestion.option_one}
          correct={selectedQuestion.correct}
          handleChange={handleChange}
          addCorrectOption={selectCorrectOption}
        />
        <QuestionOption
          id="option_two"
          name="option 02"
          className="bg-blue-500"
          value={selectedQuestion.option_two}
          correct={selectedQuestion.correct}
          handleChange={handleChange}
          addCorrectOption={selectCorrectOption}
        />
        <QuestionOption
          id="option_three"
          name="option 03"
          className="bg-yellow-500"
          value={selectedQuestion.option_three}
          correct={selectedQuestion.correct}
          handleChange={handleChange}
          addCorrectOption={selectCorrectOption}
        />
        <QuestionOption
          id="option_four"
          name="option 04"
          className="bg-green-500"
          value={selectedQuestion.option_four}
          correct={selectedQuestion.correct}
          handleChange={handleChange}
          addCorrectOption={selectCorrectOption}
        />
      </div>
      <button
        className="self-start py-4 px-16 rounded-md bg-blue-500 text-white text-2xl font-bold hover:bg-blue-700"
        onClick={() => editQuestion(selectedQuestion)}
      >
        Save Question
      </button>
    </div>
  );
};

export default QuestionEditor;
