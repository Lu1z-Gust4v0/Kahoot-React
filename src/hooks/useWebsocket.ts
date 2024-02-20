"use client";
import { useEffect, useState } from "react";

const HOST = "localhost:8000";

const useWebsocket = (
  path: string,
  handleMessage: (event: MessageEvent<any>) => void,
) => {
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const connection = new WebSocket(`ws://${HOST}${path}`);

    connection.onopen = () => {
      console.log("Connection opened successfully");
      setWs(connection);
    };

    connection.onclose = () => {
      console.log("Connection closed");
    };

    connection.onmessage = handleMessage;
  }, [path]);

  function sendMessage(message: string) {
    if (!ws) return;
    ws.send(message);
  }

  return { ws, sendMessage };
};

export default useWebsocket;
