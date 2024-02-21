"use client"

import { useState, ChangeEvent, FormEvent } from "react";
import { WebsocketMessage } from "@/types/quiz";

type QuizSubscribeFormProps = {
  sendMessage: (message: string) => void;
  setSubscribed: () => void;
};

const QuizSubscribeForm = ({ sendMessage, setSubscribed }: QuizSubscribeFormProps) => {
  const [name, setName] = useState("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    sendMessage(
      JSON.stringify({
        type: WebsocketMessage.REGISTER,
        name: name,
      }),
    );

    setSubscribed();
  };

  return (
    <form
      className="flex flex-col h-fit col-start-5 col-span-4 bg-white p-8 gap-8 rounded-md shadow-md"
      onSubmit={handleSubmit}
    >
      <h2 className="text-5xl text-gradient font-bold mb-4">Select Name</h2>
      <input
        id="name"
        className="py-6 rounded-md px-4 shadow-md"
        type="text"
        placeholder="Name"
        onChange={handleChange}
        value={name}
        autoComplete="off"
        required
      />
      <button className="py-4 px-16 rounded-md bg-purple-500 text-white text-2xl font-bold hover:bg-purple-700">
        Enter
      </button>
    </form>
  );
};

export default QuizSubscribeForm;
