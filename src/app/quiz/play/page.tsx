"use client";

import { useCallback, useState } from "react";
import { useSearchParams } from "next/navigation";
import useWebsocket from "@/hooks/useWebsocket";
import Toast from "@/components/Toast";
import { WebsocketMessage } from "@/types/quiz";
import {
  WebsocketQuestionMessage,
  WebsocketScoresMessage,
} from "@/types/websocket";
import QuizSubscribeForm from "@/components/quiz/QuizSubscribeForm";
import AnswerQuestion from "@/components/quiz/AnswerQuestion";
import PlayersScore from "@/components/quiz/PlayersScore";

type PageState = {
  subscribed: boolean;
  question: WebsocketQuestionMessage | null;
  scores: WebsocketScoresMessage | null;
};

export default function PlayGamePage() {
  const [pageState, setPageState] = useState<PageState>({
    subscribed: false,
    question: null,
    scores: null,
  });
  const [toast, setToast] = useState({
    open: false,
    type: "error",
    message: "",
  });

  const newToast = (open: boolean, type: string, message: string) => {
    setToast({
      open: open,
      type: type,
      message: message,
    });
  };

  const searchParams = useSearchParams();
  const code = searchParams.get("code");
  
  const handleMessage = useCallback((event: MessageEvent<any>) => {
    const message = JSON.parse(event.data) as { type: WebsocketMessage };

    switch (message.type) {
      case WebsocketMessage.ERROR:
        const error = (message as { type: WebsocketMessage; error: string })
          .error;
        newToast(true, "error", error);

        break;

      case WebsocketMessage.QUESTION:
        const question = message as WebsocketQuestionMessage;

        setPageState((previous) => ({
          ...previous,
          question: question,
          scores: null,
        }));

        break;

      case WebsocketMessage.SCORES:
        const scores = message as WebsocketScoresMessage;

        setPageState((previous) => ({
          ...previous,
          question: null,
          scores: scores,
        }));

        break;

      default:
        console.log("Unknown websocket message: ", message);
    }
  }, [])

  const { sendMessage } = useWebsocket(`/api/ws/client/${code}`, handleMessage);

  const toggle = () => {
    setToast((previous) => ({ ...previous, open: !toast.open }));
  };

  const subscribe = () =>
    setPageState((previous) => ({ ...previous, subscribed: true }));

  return (
    <main className="grid grid-cols-12 h-screen w-full bg-gradient py-12">
      {toast.open && (
        <Toast
          className="justify-self-center"
          close={toggle}
          type={toast.type}
          message={toast.message}
        />
      )}

      {!pageState.subscribed && (
        <QuizSubscribeForm
          sendMessage={sendMessage}
          setSubscribed={subscribe}
        />
      )}

      {pageState.subscribed && !pageState.question && !pageState.scores && (
        <div className="flex h-fit col-start-4 col-span-6 bg-white p-8 gap-8 rounded-md shadow-md">
          <h2 className="text-5xl text-gradient font-bold">
            Waiting the game to start...
          </h2>
        </div>
      )}

      {pageState.subscribed && pageState.question && (
        <AnswerQuestion
          question={pageState.question}
          sendMessage={sendMessage}
        />
      )}

      {pageState.subscribed && pageState.scores && (
        <PlayersScore scores={pageState.scores} />
      )}
    </main>
  );
}
