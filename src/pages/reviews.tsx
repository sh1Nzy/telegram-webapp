import React from "react";
import "../style/reviews.css";

const reviews = [
  {
    user: "Пользователь",
    rating: 5,
    text: "Рекомендую, все отлично"
  },
  {
    user: "Пользователь",
    rating: 5,
    text: ""
  },
  {
    user: "Пользователь",
    rating: 4,
    text: ""
  }
];

const Reviews: React.FC = () => {
  return (
    <div className="reviews-container">
      <h2 className="reviews-title">Отзывы</h2>
      <div className="reviews-list">
        {reviews.map((review, i) => (
          <div className="review-card" key={i}>
            <div className="review-header">
              <span className="review-user">{review.user}</span>
              <span className="review-stars">
                {Array.from({ length: 5 }).map((_, idx) =>
                  idx < review.rating ? (
                    <span key={idx} className="star filled">★</span>
                  ) : (
                    <span key={idx} className="star">☆</span>
                  )
                )}
              </span>
            </div>
            {review.text && <div className="review-text">{review.text}</div>}
          </div>
        ))}
      </div>
    </div>
  );
};

export default Reviews;
