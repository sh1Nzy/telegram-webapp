CREATE TABLE "order_products" (
                                  "id" SERIAL PRIMARY KEY,
                                  "order_id" integer,
                                  "product_id" integer,
                                  "quantity" integer NOT NULL,
                                  "price" numeric(10,2) NOT NULL
);