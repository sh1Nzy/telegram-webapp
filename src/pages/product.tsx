import React from "react";
import { useParams, useNavigate, Link } from "react-router-dom";
import { products } from "../data/products";
import { useFavorites } from "../context/FavoritesContext";
import "../style/product.css";

const ProductPage: React.FC = () => {
  const { productId } = useParams();
  const navigate = useNavigate();
  const { favorites, addFavorite, removeFavorite } = useFavorites();

  // Найти товар по id (поиск по всем категориям)
  let product: any;
  for (const cat of Object.values(products)) {
    product = (cat as any[]).find((p) => p.id === productId);
    if (product) break;
  }
  if (!product) return <div>Товар не найден</div>;

  const isFavorite = favorites.some(fav => fav.id === product.id);

  return (
    <div className="product-container">
      <div className="product-header">
        <button className="back-btn" onClick={() => navigate(-1)}>
          <span className="back-arrow">&lt;</span>
        </button>
        <span className="product-title-top">Магазин</span>
        <button className="close-btn" onClick={() => navigate("/")}>×</button>
      </div>
      <div className="product-image-slider-block">
        <img src={`/images/${product.image}`} alt={product.title} className="product-image-large" />
        <div className="product-slider-dots">
          <span className="dot active"></span>
          <span className="dot"></span>
          <span className="dot"></span>
          <span className="dot"></span>
        </div>
      </div>
      <div className="product-title">{product.title}</div>
      <div className="product-rating-row">
        <span className="star">★</span>
        <span className="product-rating-value">{product.rating.toFixed(1)}</span>
        <Link to={`/product/${product.id}/reviews`} className="product-reviews-link">
          Отзывы
        </Link>
        <button
          className="product-fav-btn"
          title={isFavorite ? "Убрать из избранного" : "В избранное"}
          onClick={() => isFavorite ? removeFavorite(product.id) : addFavorite(product)}
          style={{ color: isFavorite ? "#a349a4" : "#bbb", marginLeft: "12px" }}
        >
          {isFavorite ? "♥" : "♡"}
        </button>
      </div>
      <div className="product-specs">
        <b>Характеристики</b>
        <ul>
          <li>Тип — {product.type || "..."}</li>
          <li>Модель — {product.model || "..."}</li>
          <li>Размеры — {product.size || "..."}</li>
          <li>Разрешение — {product.resolution || "..."}</li>
          <li>Процессор — {product.cpu || "..."}</li>
        </ul>
      </div>
      <div className="product-description">
        {product.description || "Описание товара..."}
      </div>
      <div className="product-buy-block">
        <div className="product-price-block">
          <span className="product-price-label">Цена:</span>
          <span className="product-price-value">{product.price.toLocaleString()} руб.</span>
        </div>
        <button
          className="product-add-to-cart"
          disabled={!product.inStock}
          onClick={() => {/* addToCart */}}
        >
          Добавить в корзину
        </button>
      </div>
      <div className="product-stock-row">
        {product.inStock ? (
          <>
            <span className="stock-icon in-stock">✔</span>
            <span className="stock-text in-stock">В наличии</span>
          </>
        ) : (
          <>
            <span className="stock-icon out-of-stock">✖</span>
            <span className="stock-text out-of-stock">Нет в наличии</span>
          </>
        )}
      </div>
      <div className="product-tg-row">
        <span className="tg-icon">✈️</span>
        <a
          href="https://t.me/your_channel"
          target="_blank"
          rel="noopener noreferrer"
          className="tg-link"
        >
          Суперцены на всю технику в нашем Tg-канале!
        </a>
      </div>
      <div className="product-tg-note">
        <span className="tg-star">*</span>
        <span className="tg-note-text">Предложение действует в случае подписки на канал</span>
      </div>
    </div>
  );
};

export default ProductPage;
