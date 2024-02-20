import { QuizStatus, WebsocketMessage } from "./quiz";

export type WebsocketErrorMessage = {
  type: WebsocketMessage.ERROR;
  message: string;
};

type Player = {
  id: string;
  game_id: string;
  name: string;
  score: number;
};

export type WebsocketGameStateMessage = {
  type: WebsocketMessage.GAME_STATE;
  title: string;
  code: string;
  current_question: number;
  question_count: number;
  max_players: number;
  active_players: number;
  status: QuizStatus;
  players: Player[];
};

export type WebsocketQuestionMessage = {
  type: WebsocketMessage.QUESTION;
  gameid: string;
  title: string;
  body: string;
  option_one: string;
  option_two: string;
  option_three: string;
  option_four: string;
  correct: string;
};

export type WebsocketScoresMessage = {
  type: WebsocketMessage.SCORES;
  game_id: string;
  players: Player[];
};
