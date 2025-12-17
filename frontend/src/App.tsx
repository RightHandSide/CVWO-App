import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Pages/home.tsx";
import Register from "./Pages/register.tsx"
import RegisterComplete from "./Pages/register_completed.tsx"
import Users from "./Pages/users.tsx"
import Test from "./Pages/test.tsx"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/register" element={<Register/>} />
        <Route path="/register/completed" element={<RegisterComplete/>} />
        <Route path="/users" element={<Users/>} />
        <Route path="/test" element={<Test/>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;