import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useUserStore } from "../store/useUserStore";

export default function Home() {
  const [roomID, setRoomID] = useState("");
  const [creating, setCreating] = useState(false);
  const navigate = useNavigate();
  const { username, setUsername } = useUserStore();

  const handleCreate = async () => {
    setCreating(true);
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/create`, {
      method: "POST",
    });
    const data = await res.json();
    navigate(`/room/${data.roomID}`);
  };

  const handleJoin = () => {
    if (roomID) {
      navigate(`/room/${roomID}`);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen gap-4">
      <input
        type="text"
        placeholder="Enter username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className="border p-2 rounded"
      />
      <div className="flex gap-2">
        <button
          onClick={handleCreate}
          disabled={!username || creating}
          className="bg-blue-500 text-white px-4 py-2 rounded"
        >
          {creating ? "Creating..." : "Create Room"}
        </button>
        <input
          type="text"
          placeholder="Room ID"
          value={roomID}
          onChange={(e) => setRoomID(e.target.value)}
          className="border p-2 rounded"
        />
        <button
          onClick={handleJoin}
          disabled={!username || !roomID}
          className="bg-green-500 text-white px-4 py-2 rounded"
        >
          Join Room
        </button>
      </div>
    </div>
  );
}
