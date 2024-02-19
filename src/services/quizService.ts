import apiProvider from "@/providers/api";

type QuestionData = {
  title: string;
  body: string;
  option_one: string;
  option_two: string;
  option_three: string;
  option_four: string;
  correct: string;
};

export type CreateQuizRequest = {
  title: string;
  code: string;
  max_players: number;
  questions: QuestionData[];
};

type Game = {
  id: string;
  title: string;
  code: string;
  max_players: number;
};

type CreateQuizResponse = {
  message: string;
  game: Game;
  questions: QuestionData & { id: string };
};


async function CreateNewQuiz(data: CreateQuizRequest) {
  try {
    const response = await apiProvider.post<
      CreateQuizResponse,
      CreateQuizRequest
    >("/api/game", data);

    return response;
  } catch (error) {
    throw error;
  }
}

const quizService = {
  CreateNewQuiz,
};

export default quizService;
