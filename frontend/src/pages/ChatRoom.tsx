import { useParams } from "react-router-dom";
import { useUserStore } from "../store/useUserStore";
import { useSocket } from "../hooks/useSocket";
import ChatBox from "../components/ChatBox";
import MessageInput from "../components/MessageInput";

export default function ChatRoom() {
  const { id } = useParams<{ id: string }>();
  const { username } = useUserStore();
  const { messages, sendMessage } = useSocket(id!, username);

  return (
    <div className="flex flex-col h-screen">
      <div className="flex-1 overflow-y-auto p-4">
        <ChatBox messages={messages} />
      </div>
      <div className="p-4 border-t">
        <MessageInput onSend={sendMessage} />
      </div>
    </div>
  );
}
