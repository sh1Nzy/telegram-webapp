import React from "react";
import { useFavorites } from "../context/FavoritesContext";
import "../style/favorites.css";
import ProductCard from "../components/ProductCard";
import { Link, useNavigate } from "react-router-dom";

const Favorites: React.FC = () => {
  const { favorites, removeFavorite, clearFavorites } = useFavorites();

  const isEmpty = favorites.length === 0;

  const navigate = useNavigate();

  return (
    <div className="favorites-container">
      <h1 className="h-home">Избранные товары</h1>
      {!isEmpty && (
        <div className="favorites-header">
          <span></span>
          <button className="favorites-clear" onClick={clearFavorites}>Очистить</button>
        </div>
      )}
      <div className="favorites-list">
        {favorites.map(item => (
          <ProductCard
            key={item.id}
            id={item.id}
            title={item.title}
            price={item.price}
            image={item.image}
            rating={item.rating}
            inStock={item.inStock}
            favoriteMode
            onRemoveFavorite={() => removeFavorite(item.id)}
          />
        ))}
      </div>
      {!isEmpty && (
        <div style={{ width: "90%", maxWidth: 400, margin: "24px auto 0 auto", textAlign: "center" }}>
          <Link to="/catalog" className="continue-shopping">
            Продолжить покупки
          </Link>
        </div>
      )}
      {isEmpty && (
        <div className="favorites-empty">
          <p className="favorites-empty-text">Тут пока ещё ничего нет...</p>
        </div>
      )}
    </div>
  );
};

export default Favorites;
