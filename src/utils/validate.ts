import { MAX_PLAYERS, MIN_PLAYERS } from "@/components/quiz/QuizEditor"
import { Quiz } from "@/types/quiz";
import { CreateQuizRequest } from "@/services/quizService";

const validateString = (field: string, input: string) => {
  if (input === "") {
    throw Error(`Missing ${field}`);
  }
  return input;
};

const validatePlayerCount = (players: number) => {
  if (players < MIN_PLAYERS || players > MAX_PLAYERS) {
    throw Error("Invalid quiz player count");
  }
  return players;
};

const validateCorrectOption = (field: string, option: string) => {
  const string = validateString(field, option);

  const validOptions = [
    "option_one",
    "option_two",
    "option_three",
    "option_four",
  ];

  if (!validOptions.includes(string)) {
    throw Error(`Invalid ${field} correct option`);
  }

  return option;
};

export const quizToCreateQuizRequest = (quiz: Quiz): CreateQuizRequest => {
  try {
    return {
      title: validateString("quiz title", quiz.title),
      code: validateString("quiz code", quiz.code),
      max_players: validatePlayerCount(quiz.players),
      questions: quiz.questions.map((question, index) => {
        return {
          title: validateString(`question ${index + 1} title`, question.title),
          body: validateString(`question ${index + 1} body`, question.body),
          option_one: validateString(
            `question ${index + 1} option one`,
            question.option_one,
          ),
          option_two: validateString(
            `question ${index + 1} option two`,
            question.option_two,
          ),
          option_three: validateString(
            `question ${index + 1} option three`,
            question.option_three,
          ),
          option_four: validateString(
            `question ${index + 1} option four`,
            question.option_four,
          ),
          correct: validateCorrectOption(
            `question ${index + 1}`,
            question.correct,
          ),
        };
      }),
    };
  } catch (error) {
    throw error;
  }
};
