import React from "react";
import ProductCard from "../components/ProductCard";
// import "../style/home";

const Home: React.FC = () => {
    const products = [
    {
        title: "Игровая приставка Xbox Series X 1Tb",
        price: 59000,
        image: "/images/xbox-x.png"
    },
    {
        title: "Игровая приставка Xbox Series S 512gb",
        price: 39000,
        image: "/images/xbox-s.png"
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
            <p className="section-title">Выбор покупателей</p>
            {products.map((p, i) => (
                <ProductCard key={i} {...p} />
            ))}
        </div>
    );
};

export default Home;