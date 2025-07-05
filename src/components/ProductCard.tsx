import React from "react";
import "../style/home.css";

interface ProductCardProps {
    title: string;
    price: number;
    image: string;
}
// рейтинг + наличие + кнопки
const ProductCard: React.FC<ProductCardProps> = ({ title, price, image }) => {
    return (
        <div className="product_card">
            <img src={image} alt={title} className="product-img" />
            <div className="product-info">
                <h3 className="product-type">{title}</h3>
                <div className="product-meta">
                    {/* тут рейтинг + наличие */}
                </div>
                <p className="product-price">Цена: {price.toLocaleString()} руб.</p>
                <div className="product-act">
                    <button className="add-to-fav">like</button>  
                    <button className="add-to-card">В корзину</button>
                </div>
            </div>
        </div>
    )
}

export default ProductCard;