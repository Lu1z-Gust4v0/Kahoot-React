"use client"
import { ChangeEvent, FormEvent, useState } from "react";
import { IoClose } from "react-icons/io5";
import { useRouter } from "next/navigation";

type DialogProps = {
  close: () => void;
};

const Dialog = ({ close }: DialogProps) => {
  const [code, setCode] = useState("")
  
  const router = useRouter()

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => setCode(e.target.value)
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()

    router.push(`/quiz/play?code=${code}`)
  }

  return (
    <div
      className="absolute flex justify-center inset-0 bg-black/50 py-32 z-10"
      onClick={(e) => {
        e.stopPropagation();

        if (e.target !== e.currentTarget) return;

        close();
      }}
    >
      <form 
        className="flex flex-col h-fit min-w-80 w-1/4 bg-white py-4 px-4 gap-4 rounded-md z-20"
        onSubmit={handleSubmit}
      >
        <div className="flex mb-4 justify-between">
          <h3 className="text-4xl text-blue-500 font-bold">Quiz Code</h3>
          <button className="group self-start" onClick={() => close()}>
            <IoClose className="w-4 h-4 transition-colors duration-200 group-hover:text-blue-500 group-focus:text-blue-500" />
          </button>
        </div>
        <input
          id="title"
          className="py-6 rounded-md px-4 shadow-md"
          type="text"
          placeholder="Quiz Code"
          onChange={handleChange}
          value={code}
        />
        <button className="py-4 px-16 rounded-md bg-blue-500 text-white text-2xl font-bold hover:bg-blue-700">
          Enter
        </button>
      </form>
    </div>
  );
};

export default Dialog;
