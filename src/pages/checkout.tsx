import React, { useState } from "react";
import "../style/checkout.css";

const deliveryOptions = [
  { id: "mkad", label: "Доставка курьером в пределах МКАД", price: "0 руб." },
  { id: "out_mkad", label: "Доставка курьером за МКАД", price: "от 1000 руб.", note: "Точная сумма будет рассчитана менеджером" },
  { id: "pickup", label: "Самовывоз", price: "0 руб." },
  { id: "yandex", label: "ЯндексДоставка (ПВЗ)", price: "X руб." },
  { id: "cdek", label: "СДЭК (ПВЗ)", price: "X руб." }
];

const cities = ["Минск", "Гомель", "Брест", "Витебск", "Могилёв", "Гродно"];

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
  const [city, setCity] = useState("Минск");
  const [showCitySelect, setShowCitySelect] = useState(false);

  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const getDeliveryCost = () => {
    const opt = deliveryOptions.find(opt => opt.id === delivery);
    if (!opt || opt.price === "0 руб.") return 0;
    if (opt.id === "out_mkad") return 1000;
    return 500; // Примерная цена для X руб.
  };

  const total = 66990 + getDeliveryCost();

  const needsAddress = ["mkad", "out_mkad", "yandex", "cdek"].includes(delivery);
  const needsZip = delivery === "mkad";

  return (
    <div className="checkout-container">
      <h1 className="checkout-title">Оформление заказа</h1>

      <section className="checkout-section">
        <h2 className="checkout-subtitle">Покупатель</h2>
        <label>Ф.И.О.*
          <input name="name" value={form.name} onChange={handleInput} required />
        </label>
        <label>E-mail
          <input name="email" value={form.email} onChange={handleInput} />
        </label>
        <label>Мобильный телефон*
          <input name="phone" value={form.phone} onChange={handleInput} required />
        </label>
        <button className="checkout-save">Сохранить</button>
      </section>

      <section className="checkout-section">
        <h2 className="checkout-subtitle">Способ доставки</h2>

        <div
          className="checkout-city-select"
          onClick={() => setShowCitySelect(!showCitySelect)}
        >
        {city || "Город"}
        </div>
        {showCitySelect && (
        <div style={{ marginBottom: '12px' }}>
        {cities.map((c) => (
        <button
          key={c}
          onClick={() => {
            setCity(c);
            setShowCitySelect(false);
          }}
          style={{
            marginRight: '8px',
            marginBottom: '6px',
            background: '#f0f0f0',
            border: 'none',
            padding: '6px 12px',
            borderRadius: '6px',
            cursor: 'pointer'
          }}
        >
          {c}
        </button>
      ))}
    </div>
  )}

  <div className="delivery-options">
    {deliveryOptions.map((opt) => (
      <label
        key={opt.id}
        className={`delivery-option${delivery === opt.id ? " selected" : ""}`}
      >
        <div className="delivery-text">
          <div>{opt.label}</div>
          <div className="delivery-price">{opt.price}</div>
          {opt.note && <div className="delivery-note">{opt.note}</div>}
        </div>
        <input
          type="radio"
          name="delivery"
          value={opt.id}
          checked={delivery === opt.id}
          onChange={() => setDelivery(opt.id)}
        />
      </label>
    ))}
  </div>
</section>

      <section className="checkout-section">
        {needsAddress && (
          <label>Адрес доставки*
            <input name="address" value={form.address || ""} onChange={handleInput} required />
          </label>
        )}
        {needsZip && (
          <label>Индекс*
            <input name="zip" value={form.zip || ""} onChange={handleInput} required />
          </label>
        )}
        <label>Комментарий к заказу:
          <input name="comment" value={form.comment || ""} onChange={handleInput} />
        </label>
        <button className="checkout-save">Сохранить</button>
      </section>

      <section className="checkout-summary-section">
        <div className="checkout-summary-block">
          <div className="checkout-summary-row">
            <span>Итого:</span>
            <span className="checkout-summary-total">{total.toLocaleString()} руб.</span>
          </div>
          <div className="checkout-summary-row">
            <span>Общий вес:</span>
            <span>2.499 кг</span>
          </div>
          <div className="checkout-summary-row">
            <span>Доставка:</span>
            <span>
              {deliveryOptions.find(opt => opt.id === delivery)?.price || "-"}
            </span>
          </div>
          {delivery === "out_mkad" && (
            <div className="checkout-summary-row">
              <span></span>
              <span className="delivery-note">Точная сумма будет рассчитана менеджером</span>
            </div>
          )}
          <div className="checkout-summary-row">
            <span>Способ:</span>
            <span>{deliveryOptions.find(opt => opt.id === delivery)?.label}</span>
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
