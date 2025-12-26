import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Pages/home.tsx";
import Register from "./Pages/register.tsx";
import RegisterComplete from "./Pages/register_completed.tsx";
import Login from "./Pages/login.tsx";
import Threads from "./Pages/threads.tsx";
import Posts from "./Pages/posts.tsx";
import Comments from "./Pages/comments.tsx";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/register" element={<Register/>} />
        <Route path="/register/completed" element={<RegisterComplete/>} />
        <Route path="/login" element={<Login/>} />
        <Route path="/threads" element={<Threads/>} />
        <Route path="/thread/:id" element={<Posts/>} />
        <Route path="/post/:id" element={<Comments/>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;