import React, { useState } from "react";
import "../style/ordering.css";
import { useCart } from "../context/CartContext";
import { useNavigate } from "react-router-dom";

const Ordering: React.FC = () => {
  const { cartItems } = useCart();
  const navigate = useNavigate();
  const [selectedDelivery, setSelectedDelivery] = useState<string>("");
  const [comment, setComment] = useState<string>("");

  const subtotal = cartItems.reduce((sum, item) => sum + item.price * item.count, 0);
  
  const deliveryOptions = [
    { id: "mkad", label: "Доставка курьером в пределах МКАД", price: 0 },
    { id: "outside-mkad", label: "Доставка курьером за МКАД", price: 1000 },
    { id: "pickup", label: "Самовывоз", price: 0 },
    { id: "yandex", label: "Яндекс.Доставка (ПВЗ)", price: 0 }, // X руб.
    { id: "cdek", label: "СДЭК (ПВЗ)", price: 0 }, // X руб.
  ];

  const selectedOption = deliveryOptions.find(option => option.id === selectedDelivery);
  const deliveryPrice = selectedOption ? selectedOption.price : 0;
  const total = subtotal + deliveryPrice;

  const handleDeliveryChange = (deliveryId: string) => {
    setSelectedDelivery(deliveryId);
  };

  const handleOrderSubmit = () => {
    // Логика оформления заказа
    console.log("Order submitted:", {
      items: cartItems,
      delivery: selectedDelivery,
      comment,
      total
    });
  };

  return (
    <div className="ordering-container">
      <div className="ordering-header">
        <button className="back-button" onClick={() => navigate(-1)}>
          Назад
        </button>
        <h1 className="ordering-title">Оформление заказа</h1>
        <button className="close-button" onClick={() => navigate("/")}>
          ✕
        </button>
      </div>

      <div className="ordering-content">
        <div className="delivery-section">
          <h2 className="section-title">Способ доставки</h2>
          <div className="delivery-city">Город</div>
          
          <div className="delivery-options">
            {deliveryOptions.map((option) => (
              <div 
                key={option.id} 
                className={`delivery-option ${selectedDelivery === option.id ? 'selected' : ''}`}
                onClick={() => handleDeliveryChange(option.id)}
              >
                <div className="delivery-option-content">
                  <div className="delivery-option-label">{option.label}</div>
                  <div className="delivery-option-price">
                    {option.price === 0 ? "0 руб." : 
                     option.id === "outside-mkad" ? "от 1000 руб." : 
                     "X руб."}
                  </div>
                </div>
                <div className={`radio-button ${selectedDelivery === option.id ? 'checked' : ''}`}>
                  {selectedDelivery === option.id && <div className="radio-inner"></div>}
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="comment-section">
          <label htmlFor="comment" className="comment-label">Комментарий к заказу</label>
          <textarea
            id="comment"
            className="comment-input"
            value={comment}
            onChange={(e) => setComment(e.target.value)}
            placeholder="Введите комментарий..."
          />
        </div>

        <div className="order-summary">
          <div className="summary-row">
            <span>Итого:</span>
            <span className="summary-total">{total.toLocaleString()} руб.</span>
          </div>
          <div className="summary-details">
            <div className="summary-item">
              <span>Общий вес:</span>
              <span>{(cartItems.reduce((sum, item) => sum + item.count, 0) * 0.5).toFixed(1)} кг</span>
            </div>
            <div className="summary-item">
              <span>Доставка:</span>
              <span>{deliveryPrice === 0 ? "Бесплатно" : `${deliveryPrice.toLocaleString()} руб.`}</span>
            </div>
            <div className="summary-item">
              <span>Оплата:</span>
              <span>Наличными</span>
            </div>
          </div>

          <button 
            className="order-button"
            onClick={handleOrderSubmit}
            disabled={!selectedDelivery}
          >
            Оформить заказ
          </button>
        </div>
      </div>
    </div>
  );
};

export default Ordering;