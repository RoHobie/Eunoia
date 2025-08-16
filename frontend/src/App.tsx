import { Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import ChatRoom from "./pages/ChatRoom";

function App() {
  return (
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/room/:id" element={<ChatRoom />} />
      </Routes>
  );
}

export default App;
