import React from "react";
import { useParams } from "react-router-dom";
import { products, catalogCategories, subcategories } from "../data/products";
import ProductCard from "../components/ProductCard";
import "../style/catalog.css";

export const CategoryPage: React.FC = () => {
  const { categoryId } = useParams();

  const category = catalogCategories.find(cat => cat.id === categoryId);
  const items = products[categoryId as keyof typeof products] || [];
  const subs = subcategories[categoryId as keyof typeof subcategories] || [];

  return (
    <div className="category-container">
      <h1 className="h-home">{category?.name || "Категория"}</h1>
      <div className="subcategory-scroll">
        {subs.map((sub: { id: string; name: string; image: string }) => (
          <button className="subcategory-chip" key={sub.id}>
            <img src={`/images/${sub.image}`} alt={sub.name} />
            <span>{sub.name}</span>
          </button>
        ))}
      </div>
      <div className="category-filters">
        <button className="filter-btn">
          <span className="filter-icon">⚲</span> Фильтр
        </button>
        <select className="sort-select">
          <option>По возрастанию цены</option>
          <option>По убыванию цены</option>
          <option>По рейтингу</option>
        </select>
      </div>
      <div className="product-list">
        {items.map((item, i) => (
          <ProductCard
            key={i}
            id={item.id}
            title={item.name}
            price={item.price}
            image={`/images/${item.image}`}
            rating={item.rating ?? 5.0}
            inStock={item.inStock ?? true}
          />
        ))}
      </div>
    </div>
  );
};
