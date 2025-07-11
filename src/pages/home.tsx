import React from "react";
import ProductCard from "../components/ProductCard";
import "../style/home.css";

const Home: React.FC = () => {
    const products = [
        {
            id: "xbox-series-x",
            title: "Игровая приставка Xbox Series X 1Tb",
            price: 59000,
            image: "xbox-x.png",
            rating: 4.7,
            inStock: true
        },
        {
            id: "xbox-series-s",
            title: "Игровая приставка Xbox Series S 512gb",
            price: 39000,
            image: "xbox-s.png",
            rating: 4.4,
            inStock: false
        }
    ];

    return (
        <div className="home-container">
            <h1 className="h-home">Главная</h1>
            <input
                type="text"
                placeholder="Поиск товаров"
                className="search-input"
            />
            <div className="product-list">
                <p className="section-title">Выбор покупателей</p>
                {products.map((p, i) => (
                    <ProductCard
                        key={i}
                        id={p.id}
                        title={p.title}
                        price={p.price}
                        image={`/images/${p.image}`}
                        rating={p.rating ?? 5.0}
                        inStock={p.inStock ?? true}
                    />
                ))}
            </div>
        </div>
    );
};

export default Home;