import { Question } from "./question"

export type Quiz = {
  title: string 
  code: string
  players: number
  questions: Question[]
}
