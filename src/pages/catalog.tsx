import React from "react";
import { Link } from "react-router-dom";
import { catalogCategories } from "../data/products";
import "../style/catalog.css";

const Catalog: React.FC = () => {
    return (
        <div className="catalog-container">
            <h1 className="h-home">Каталог</h1>
            <div className="catalog-grid">
                {catalogCategories.map((cat) => (
                <Link to={`/category/${cat.id}`} key={cat.id}>
                <img src={`/images/${cat.image}`} alt={cat.name}/>
                <span>{cat.name}</span>
                </Link>
      ))}
    </div>
  </div>
    )
}
export default Catalog;