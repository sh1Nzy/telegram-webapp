import React from "react";
import { useCart } from "../context/CartContext";
import "../style/home.css";
import "../style/productCard.css";

interface ProductCardProps {
    id: string;
    title: string;
    price: number;
    image: string;
    rating: number;
    inStock: boolean;
}

const ProductCard: React.FC<ProductCardProps> = ({ id, title, price, image, rating, inStock }) => {
    const { addToCart } = useCart();
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
                    <button className="add-to-fav" title="В избранное">♡</button>
                    <button
                        className="add-to-cart"
                        disabled={!inStock}
                        onClick={() => {
                            addToCart({ id, title, price, image });
                            console.log("addToCart called", id);
                        }}
                    >
                        В корзину
                    </button>
                </div>
            </div>
        </div>
    );
};

export default ProductCard;