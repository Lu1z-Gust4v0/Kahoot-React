import { Question } from "./question"

export type Quiz = {
  title: string 
  code: string
  players: number
  questions: Question[]
}

export enum WebsocketMessage {
	SCORES,
	QUESTION,
	REGISTER,
	ANSWER,
	START_GAME,
	GAME_STATE,
	ERROR,
}

export enum QuizStatus {
  WAITING,
	STARTED,
	FINISHED,
}
