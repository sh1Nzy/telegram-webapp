-- Заполняем таблицу users
INSERT INTO "users" ("id", "telegram_id", "username", "created_at") VALUES
(1, 123456789, 'user1', NOW()),
(2, 987654321, 'user2', NOW()),
(3, 456789123, 'user3', NOW());

-- Заполняем таблицу categories
INSERT INTO "categories" ("id", "name") VALUES
(1, 'Электроника'),
(2, 'Одежда'),
(3, 'Книги');

-- Заполняем таблицу firms
INSERT INTO "firms" ("id", "name") VALUES
(1, 'Apple'),
(2, 'Samsung'),
(3, 'Nike');

-- Заполняем таблицу products
INSERT INTO "products" ("id", "name", "description", "category_id", "firm_id", "attributes", "sell_count", "stock", "image") VALUES
(1, 'iPhone 13', 'Смартфон Apple', 1, 1, '{"color": "black", "memory": "128GB"}', 0, 10, ARRAY['iphone13.jpg']),
(2, 'Galaxy S21', 'Смартфон Samsung', 1, 2, '{"color": "white", "memory": "256GB"}', 0, 15, ARRAY['galaxys21.jpg']),
(3, 'Кроссовки Air Max', 'Спортивная обувь', 2, 3, '{"color": "red", "size": "42"}', 0, 20, ARRAY['airmax.jpg']);

