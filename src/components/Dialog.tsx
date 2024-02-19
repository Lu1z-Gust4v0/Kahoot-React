import { IoClose } from "react-icons/io5";

type DialogProps = {
  close: () => void;
};

const Dialog = ({ close }: DialogProps) => {
  return (
    <div
      className="absolute flex justify-center inset-0 bg-black/50 py-32 z-10"
      onClick={(e) => {
        e.stopPropagation();
        console.log(e.target, e.currentTarget);
        if (e.target !== e.currentTarget) return;

        close();
      }}
    >
      <div className="flex flex-col h-fit min-w-80 w-1/4 bg-white py-4 px-4 gap-4 rounded-md z-20">
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
        />
        <button className="py-4 px-16 rounded-md bg-blue-500 text-white text-2xl font-bold hover:bg-blue-700">
          Enter
        </button>
      </div>
    </div>
  );
};

export default Dialog;
