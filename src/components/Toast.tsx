import { IoClose } from "react-icons/io5";

type ToastProps = {
  close: () => void;
  type: string;
  message: string;
  className?: string;
};

const Toast = ({ message, type, close, className }: ToastProps) => {
  return (
    <div
      className={
        "absolute justify-self-end flex flex-col h-fit min-h-48 min-w-80 w-1/4 m-8 bg-white py-4 px-4 rounded-md z-20 shadow-md " +
        className
      }
    >
      <div className="flex justify-end">
        <button className="group self-start" onClick={() => close()}>
          <IoClose className="w-4 h-4 transition-colors duration-200 group-hover:text-blue-500 group-focus:text-blue-500" />
        </button>
      </div>
      <p
        className={`text-xl ${type === "error" ? "text-red-500" : "text-green-700"
          }`}
      >
        {message}
      </p>
    </div>
  );
};

export default Toast;
