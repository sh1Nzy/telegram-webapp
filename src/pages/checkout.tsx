import React, { useState } from "react";
import "../style/checkout.css";

const deliveryOptions = [
  { id: "mkad", label: "Доставка курьером в пределах МКАД", price: "0 руб." },
  { id: "out_mkad", label: "Доставка курьером за МКАД", price: "от 1000 руб.", note: "Точная сумма будет рассчитана менеджером" },
  { id: "pickup", label: "Самовывоз", price: "0 руб." },
  { id: "yandex", label: "ЯндексДоставка (ПВЗ)", price: "X руб." },
  { id: "cdek", label: "СДЭК (ПВЗ)", price: "0 руб." }
];

const Checkout: React.FC = () => {
  const [form, setForm] = useState({
    name: "",
    email: "",
    phone: "",
    address: "",
    zip: "",
    comment: ""
  });
  const [delivery, setDelivery] = useState("mkad");
  const [checkoutVariation, setCheckoutVariation] = useState<1 | 2 | 3>(3); // По умолчанию показываем третью вариацию

  const handleInput = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  // Вариация 3 - упрощенная версия с только выбором доставки (как на скриншоте)
  if (checkoutVariation === 3) {
    return (
      <div className="checkout-container checkout-v3">
        <div className="checkout-header">
          <button className="back-button" onClick={() => window.history.back()}>
            Назад
          </button>
          <h1 className="checkout-title">Оформление заказа</h1>
          <button className="close-button" onClick={() => window.location.href = "/"}>
            ✕
          </button>
        </div>

        {/* Временные кнопки для переключения вариаций */}
        <div style={{padding: '10px', background: '#fff', margin: '8px 0', textAlign: 'center'}}>
          <button onClick={() => setCheckoutVariation(1)} style={{margin: '0 5px', padding: '5px 10px'}}>
            Вариация 1
          </button>
          <button onClick={() => setCheckoutVariation(2)} style={{margin: '0 5px', padding: '5px 10px'}}>
            Вариация 2
          </button>
          <button onClick={() => setCheckoutVariation(3)} style={{margin: '0 5px', padding: '5px 10px', background: '#8B5CF6', color: 'white'}}>
            Вариация 3 (текущая)
          </button>
        </div>

        <section className="checkout-section">
          <h2 className="checkout-subtitle">Способ доставки</h2>
          <div className="checkout-label">Город</div>
          
          <div className="delivery-options-v3">
            {deliveryOptions.map(opt => (
              <div 
                key={opt.id} 
                className={`delivery-option-v3 ${delivery === opt.id ? "selected" : ""}`}
                onClick={() => setDelivery(opt.id)}
              >
                <div className="delivery-option-content">
                  <div className="delivery-option-label">{opt.label}</div>
                  <div className="delivery-option-price">{opt.price}</div>
                </div>
                <div className={`radio-button ${delivery === opt.id ? "checked" : ""}`}>
                  {delivery === opt.id && <div className="radio-inner"></div>}
                </div>
              </div>
            ))}
          </div>
        </section>

        <section className="checkout-section">
          <label className="comment-label">
            Комментарий к заказу:
            <textarea 
              name="comment" 
              value={form.comment || ""} 
              onChange={handleInput}
              className="comment-textarea"
              placeholder="Введите комментарий..."
            />
          </label>
        </section>

        <section className="checkout-summary-section">
          <div className="checkout-summary-block">
            <div className="checkout-summary-row">
              <span>Итого:</span>
              <span className="checkout-summary-total">66 990 руб.</span>
            </div>
            <div className="checkout-summary-row">
              <span>Общий вес:</span>
              <span>2.499 кг</span>
            </div>
            <div className="checkout-summary-row">
              <span>Доставка:</span>
              <span>
                {delivery === "mkad" && "Бесплатно"}
                {delivery === "out_mkad" && "от 1000 руб."}
                {delivery === "pickup" && "0 руб."}
                {delivery === "yandex" && "X руб."}
                {delivery === "cdek" && "0 руб."}
              </span>
            </div>
            <div className="checkout-summary-row">
              <span>Доставка:</span>
              <span>
                {deliveryOptions.find(opt => opt.id === delivery)?.label}
              </span>
            </div>
            <div className="checkout-summary-row">
              <span>Оплата:</span>
              <span>Наличными</span>
            </div>
            <button className="checkout-order-btn">Оформить заказ</button>
          </div>
        </section>
      </div>
    );
  }

  // Вариация 1 и 2 - полная версия (существующий код)
  return (
    <div className="checkout-container">
      <h1 className="checkout-title">Оформление заказа</h1>
      
      {/* Временные кнопки для переключения вариаций */}
      <div style={{padding: '10px', background: '#fff', margin: '8px 0', textAlign: 'center', width: '92%', maxWidth: '400px', borderRadius: '16px'}}>
        <button onClick={() => setCheckoutVariation(1)} style={{margin: '0 5px', padding: '5px 10px', background: checkoutVariation === 1 ? '#a349a4' : '#ddd', color: checkoutVariation === 1 ? 'white' : 'black'}}>
          Вариация 1 {checkoutVariation === 1 && '(текущая)'}
        </button>
        <button onClick={() => setCheckoutVariation(2)} style={{margin: '0 5px', padding: '5px 10px', background: checkoutVariation === 2 ? '#a349a4' : '#ddd', color: checkoutVariation === 2 ? 'white' : 'black'}}>
          Вариация 2 {checkoutVariation === 2 && '(текущая)'}
        </button>
        <button onClick={() => setCheckoutVariation(3)} style={{margin: '0 5px', padding: '5px 10px'}}>
          Вариация 3
        </button>
      </div>
      <section className="checkout-section">
        <h2 className="checkout-subtitle">Покупатель</h2>
        <label>
          Ф.И.О.<span className="required">*</span>
          <input name="name" value={form.name} onChange={handleInput} required />
        </label>
        <label>
          E-mail
          <input name="email" value={form.email} onChange={handleInput} />
        </label>
        <label>
          Мобильный телефон<span className="required">*</span>
          <input name="phone" value={form.phone} onChange={handleInput} required />
        </label>
        <button className="checkout-save">Сохранить</button>
      </section>

      <section className="checkout-section">
        <h2 className="checkout-subtitle">Способ доставки</h2>
        <div className="checkout-label">Город</div>
        <div className="delivery-options">
          {deliveryOptions.map(opt => (
            <label key={opt.id} className={`delivery-option${delivery === opt.id ? " selected" : ""}`}>
              <input
                type="radio"
                name="delivery"
                value={opt.id}
                checked={delivery === opt.id}
                onChange={() => setDelivery(opt.id)}
              />
              <div>
                <div>{opt.label}</div>
                <div className="delivery-price">{opt.price}</div>
                {opt.note && <div className="delivery-note">{opt.note}</div>}
              </div>
            </label>
          ))}
        </div>
      </section>

      <section className="checkout-section">
        <label>
          Адрес доставки<span className="required">*</span>
          <input name="address" value={form.address || ""} onChange={handleInput} required />
        </label>
        <label>
          Индекс<span className="required">*</span>
          <input name="zip" value={form.zip || ""} onChange={handleInput} required />
        </label>
        <label>
          Комментарий к заказу:
          <input name="comment" value={form.comment || ""} onChange={handleInput} />
        </label>
        <button className="checkout-save">Сохранить</button>
      </section>

      <section className="checkout-summary-section">
        <div className="checkout-summary-block">
          <div className="checkout-summary-row">
            <span>Итого:</span>
            <span className="checkout-summary-total">66 990 руб.</span>
          </div>
          <div className="checkout-summary-row">
            <span>Общий вес:</span>
            <span>2.499 кг</span>
          </div>
          <div className="checkout-summary-row">
            <span>Доставка:</span>
            <span>
              {delivery === "mkad" && "Бесплатно"}
              {delivery === "out_mkad" && "от 1000 руб."}
              {delivery === "pickup" && "0 руб."}
              {delivery === "yandex" && "X руб."}
              {delivery === "cdek" && "0 руб."}
            </span>
          </div>
          {delivery === "out_mkad" && (
            <div className="checkout-summary-row">
              <span></span>
              <span className="delivery-note">Точная сумма будет рассчитана менеджером</span>
            </div>
          )}
          <div className="checkout-summary-row">
            <span>Доставка:</span>
            <span>
              {deliveryOptions.find(opt => opt.id === delivery)?.label}
            </span>
          </div>
          <div className="checkout-summary-row">
            <span>Оплата:</span>
            <span>Наличными курьеру</span>
          </div>
          <button className="checkout-order-btn">Оформить заказ</button>
        </div>
        <div className="checkout-summary-note">
          Нажимая на кнопку, вы соглашаетесь на <a href="#" className="checkout-summary-link">обработку персональных данных</a> и с <a href="#" className="checkout-summary-link">публичной офертой</a>
        </div>
      </section>
    </div>
  );
};

export default Checkout;