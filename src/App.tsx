import './App.css';
import { Routes, Route, Link } from "react-router-dom";
import { useEffect, useState } from "react";
import Home from "./pages/home";
import BottomNav from './components/BottomNav';
import Catalog from './pages/catalog';
import { CategoryPage } from './pages/category';
import Cart from './pages/basket';
import Favorites from "./pages/favorites";
import ProductPage from "./pages/product";
import Reviews from "./pages/reviews";
import Checkout from "./pages/checkout";

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
    <>
    <Routes>
      <Route path="/" element={<Home />}></Route>
      <Route path="/catalog" element={<Catalog />} />
      <Route path="/category/:categoryId" element={<CategoryPage />} />
      <Route path="/basket" element={<Cart />} />
      <Route path="/favorites" element={<Favorites />} />
      <Route path="/product/:productId" element={<ProductPage />} />
      <Route path="/product/:productId/reviews" element={<Reviews />} />
      <Route path="/checkout" element={<Checkout />} />
    </Routes>
    <BottomNav/>
    </>
  );
}

export default App;
