"use client";
import { useState } from "react";
import Slide from "@/components/Slide";
import QuestionEditor from "@/components/QuestionEditor";
import { Question } from "@/types/question";
import QuizControlPanel from "@/components/QuizControlPanel";

const defaultQuestion: Question = {
  id: Math.random().toString(),
  title: "",
  body: "",
  option_one: "",
  option_two: "",
  option_three: "",
  option_four: "",
  correct: [],
};

export default function EditQuiz() {
  const [questions, setQuestions] = useState<Question[]>([defaultQuestion]);
  const [selectedQuestion, setSelectedQuestion] =
    useState<Question>(defaultQuestion);

  const selectQuestion = (id: string) => {
    questions.forEach((question) => {
      if (question.id === id) setSelectedQuestion(question);
    });
  };

  const addNewQuestion = () => {
    setQuestions((previous) => [
      ...previous,
      {
        id: Math.random().toString(),
        title: "",
        body: "",
        option_one: "",
        option_two: "",
        option_three: "",
        option_four: "",
        correct: [],
      },
    ]);
  };

  const editQuestion = (questionToUpdate: Question) => {
    setQuestions(
      questions.map((question) => {
        if (question.id === questionToUpdate.id) {
          return questionToUpdate;
        }
        return question;
      }),
    );
  };

  return (
    <main className="grid grid-cols-12 h-screen w-full bg-gray-200">
      <Slide
        questions={questions}
        selectQuestion={selectQuestion}
        addNewQuestion={addNewQuestion}
      />
      <QuestionEditor
        selectedQuestion={selectedQuestion}
        setSelectedQuestion={setSelectedQuestion}
        editQuestion={editQuestion}
      />
      <QuizControlPanel />
    </main>
  );
}
