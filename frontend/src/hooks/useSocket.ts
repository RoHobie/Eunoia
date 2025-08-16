import { useEffect, useRef, useState } from "react";
import type { Message } from "../types";

export function useSocket(roomID: string, username: string) {
  const wsRef = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    const wsUrl = `${import.meta.env.VITE_BACKEND_WS}/ws/${roomID}?username=${username}`;
    wsRef.current = new WebSocket(wsUrl);

    wsRef.current.onmessage = (event) => {
      const msg: Message = JSON.parse(event.data);
      setMessages((prev) => [...prev, msg].slice(-50)); // keep last 50
    };

    return () => {
      wsRef.current?.close();
    };
  }, [roomID, username]);

  const sendMessage = (content: string) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(content);
    }
  };

  return { messages, sendMessage };
}
