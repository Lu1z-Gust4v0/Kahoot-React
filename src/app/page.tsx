"use client";
import Link, { LinkProps } from "next/link";
import { ReactNode, useState } from "react";
import Dialog from "@/components/Dialog";

type LinkButtonProps = {
  children: ReactNode;
  className?: string;
} & LinkProps;

const LinkButton = ({ children, className, ...props }: LinkButtonProps) => {
  return (
    <Link
      className={`self-start py-4 px-16 rounded-md bg-blue-500 text-white text-2xl font-bold hover:bg-blue-700 ${className}`}
      {...props}
    >
      {children}
    </Link>
  );
};

export default function Home() {
  const [open, setOpen] = useState(false);

  const toggle = () => {
    setOpen(!open);
  };

  return (
    <div className="grid grid-cols-12 h-screen w-full bg-gray-200">
      <nav className="flex items-center bg-white h-20 col-span-full px-8 shadow-md">
        <h1 className="text-3xl font-bold text-gradient">Kahoot React</h1>
      </nav>
      <main className="min-h-[90vh] col-span-full md:col-start-2 md:col-span-10 lg:col-start-3 lg:col-span-8 flex flex-col py-8 px-2 overflow-y-auto">
        {open && <Dialog close={toggle} />}
        <div className="flex flex-col items-center gap-8 py-4 px-8 bg-white shadow-md rounded-md">
          <h2 className="text-4xl text-gradient font-bold">
            Select an option
          </h2>
          <div className="flex flex-col md:flex-row justify-center gap-8">
            <LinkButton
              href="/quiz/edit"
              className="bg-purple-500 hover:bg-purple-700"
            >
              New Quiz
            </LinkButton>
            <LinkButton href="#" onClick={() => toggle()}>
              Enter a Quiz
            </LinkButton>
          </div>
        </div>
      </main>
    </div>
  );
}
