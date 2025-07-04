import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from "react";

function App() {
  
  const [user, setUser] = useState(null);
  
  useEffect(() => {
  
    const tg = window.Telegram.WebApp;
    tg.ready();
    tg.expand();
    const initDataUnsafe = tg.initDataUnsafe;
  
    if (initDataUnsafe?.user) {
      setUser(initDataUnsafe.user);
    }
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h1>Hello, {user?.first_name || "Unknown"}!</h1>
      <p>Welcome to Telegram Web App!</p>
    </div>
  );
}

export default App;
