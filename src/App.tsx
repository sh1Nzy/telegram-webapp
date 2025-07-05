import './App.css';
import { Routes, Route } from "react-router-dom";
import { useEffect, useState } from "react";
import Home from "./pages/home";

function App() {
  
  const [user, setUser] = useState(null);
  
  useEffect(() => {
  
    const tg = window.Telegram.WebApp as any;
    tg.ready();
    tg.expand();
    const initDataUnsafe = tg.initDataUnsafe;
  
    if (initDataUnsafe?.user) {
      setUser(initDataUnsafe.user);
    }
  }, []);

  return (
    <Routes>
      <Route path="/" element={<Home />}></Route>
      
    </Routes>
    // <div style={{ padding: 20 }}>
    //   <h1>Hello, {user?.first_name || "Unknown"}!</h1>
    //   <p>Welcome to Telegram Web App!</p>
    // </div>
  );
}

export default App;
