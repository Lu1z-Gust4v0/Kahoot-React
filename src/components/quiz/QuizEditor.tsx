"use client"
import { Quiz } from "@/types/quiz";
import { Question } from "@/types/question";
import { useState } from "react";
import Slide from "@/components/quiz/Slide";
import QuestionEditor from "@/components/quiz/QuestionEditor";
import QuizControlPanel from "@/components/quiz/QuizControlPanel";

export const MIN_PLAYERS = 4;
export const MAX_PLAYERS = 12;

const DEFAULT_QUESTION: Question = {
  id: Math.random().toString(),
  title: "",
  body: "",
  option_one: "",
  option_two: "",
  option_three: "",
  option_four: "",
  correct: "",
};

const DEFAULT_QUIZ: Quiz = {
  title: "",
  code: "",
  players: MIN_PLAYERS,
  questions: [DEFAULT_QUESTION],
};

const QuizEditor = () => {
  const [quiz, setQuiz] = useState<Quiz>(DEFAULT_QUIZ);
  const [selectedQuestion, setSelectedQuestion] =
    useState<Question>(DEFAULT_QUESTION);

  const selectQuestion = (id: string) => {
    quiz.questions.forEach((question) => {
      if (question.id === id) setSelectedQuestion(question);
    });
  };

  const editQuestion = (questionToUpdate: Question) => {
    setQuiz((previous) => ({
      ...previous,
      questions: previous.questions.map((question) => {
        if (question.id === questionToUpdate.id) {
          return questionToUpdate;
        }
        return question;
      }),
    }));
  };

  const addNewQuestion = () => {
    setQuiz((previous) => ({
      ...previous,
      questions: [
        ...previous.questions,
        {
          id: Math.random().toString(),
          title: DEFAULT_QUESTION.title,
          body: DEFAULT_QUESTION.body,
          option_one: DEFAULT_QUESTION.option_one,
          option_two: DEFAULT_QUESTION.option_two,
          option_three: DEFAULT_QUESTION.option_three,
          option_four: DEFAULT_QUESTION.option_four,
          correct: DEFAULT_QUESTION.correct,
        },
      ],
    }));
  };

  const editQuiz = (quiz: Quiz) => setQuiz(quiz);

  return (
    <>
      <Slide
        questions={quiz.questions}
        selectQuestion={selectQuestion}
        addNewQuestion={addNewQuestion}
      />
      <QuestionEditor
        selectedQuestion={selectedQuestion}
        setSelectedQuestion={setSelectedQuestion}
        editQuestion={editQuestion}
      />
      <QuizControlPanel quiz={quiz} editQuiz={editQuiz} />
    </>
  );
};

export default QuizEditor;
