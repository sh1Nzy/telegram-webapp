import React from "react";
import "../style/basket.css";
import { useCart } from "../context/CartContext";
import { Link } from "react-router-dom";

const Cart: React.FC = () => {
  const { cartItems } = useCart();
  const isEmpty = cartItems.length === 0;
  return (
    <div className="cart-container">
      <h1 className="h-home">Корзина</h1>
      {isEmpty ? (
        <div className="cart-empty">
          <div className="cart-empty-icon">
            <img src="/images/cart.svg" alt="cart" />
            <svg width="120" height="120" fill="none" viewBox="0 0 24 24"></svg>
          </div>
          <div className="cart-empty-title">Ваша корзина пуста</div>
          <div className="cart-empty-desc">
            <span><Link to="/" className="cart-empty-link">Нажмите здесь</Link>, чтобы продолжить покупки</span>
          </div>
        </div>
      ) : (
        <div className="cart-full">
          {cartItems.map(item => (
            <div className="cart-item" key={item.id}>
              <img src={item.image} alt={item.title} className="cart-item-img" />
              <div className="cart-item-info">
                <div className="cart-item-title">{item.title}</div>
                <div className="cart-item-controls">
                  <button>-</button>
                  <span>{item.count}</span>
                  <button>+</button>
                </div>
              </div>
              <div className="cart-item-price">{item.price.toLocaleString()} руб.</div>
            </div>
          ))}
          <div className="cart-summary">
            <div className="cart-summary-row">
              <span>Итого</span>
              <span>
                {cartItems.reduce((sum, i) => sum + i.price * i.count, 0).toLocaleString()} руб.
              </span>
            </div>
            <input className="cart-promo" placeholder="Есть промокод?" />
            <Link to="/checkout" className="cart-checkout">Перейти к оформлению</Link>
          </div>
        </div>
      )}
    </div>
  );
};

export default Cart;
