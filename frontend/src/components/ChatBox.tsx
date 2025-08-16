import type { Message } from "../types";

export default function ChatBox({ messages }: { messages: Message[] }) {
  return (
    <div className="flex flex-col gap-2">
      {messages.map((msg, i) => (
        <div key={i} className="p-2 rounded bg-gray-100">
          <span className="font-semibold">{msg.username}:</span>{" "}
          <span>{msg.content}</span>
        </div>
      ))}
    </div>
  );
}
