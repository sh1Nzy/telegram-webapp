import React from "react";
import { useCart } from "../context/CartContext";
import { useFavorites } from "../context/FavoritesContext";
import "../style/home.css";
import "../style/productCard.css";

interface ProductCardProps {
    id: string;
    title: string;
    price: number;
    image: string;
    rating: number;
    inStock: boolean;
    favoriteMode?: boolean; 
    onRemoveFavorite?: () => void; 
}

const ProductCard: React.FC<ProductCardProps> = ({ id, title, price, image, rating, inStock, favoriteMode, onRemoveFavorite }) => {
    const { addToCart } = useCart();
    const { favorites, addFavorite, removeFavorite } = useFavorites();
    const isFavorite = favorites.some(fav => fav.id === id);
    console.log("ProductCard rendered", id);

    return (
        <div className="product-card">
            <img src={image} alt={title} className="product-image" />
            <div className="product-info">
                <div className="product-title">{title}</div>
                <div className="product-meta">
                    <span className="product-rating">
                        <span className="star">⭐</span>
                        <span className="product-rating-value">{' '}{rating.toFixed(1)}</span>
                    </span>
                    <span className={inStock ? "in-stock" : "out-of-stock"}>
                        {inStock ? "Есть в наличии" : "Нет в наличии"}
                    </span>
                </div>
                <div className="product-price">Цена: {price.toLocaleString()} руб.</div>
                <div className="product-actions">
                    {favoriteMode ? (
                        <button
                            className="remove-favorite"
                            title="Удалить из избранного"
                            onClick={onRemoveFavorite}
                        >
                            🗑
                        </button>
                    ) : (
                        <button
                            className="add-to-fav"
                            title={isFavorite ? "Убрать из избранного" : "В избранное"}
                            onClick={() => {
                                isFavorite
                                    ? removeFavorite(id)
                                    : addFavorite({ id, title, price, image, rating, inStock });
                            }}
                            style={{ color: isFavorite ? "#a349a4" : "grey" }}
                        >
                            {isFavorite ? "♥" : "♡"}
                        </button>
                    )}
                    <button
                        className="add-to-cart"
                        disabled={!inStock}
                        onClick={() => addToCart({ id, title, price, image })}
                    >
                        В корзину
                    </button>
                </div>
            </div>
        </div>
    );
};

export default ProductCard;