CREATE TABLE "prices" (
                          "id" SERIAL PRIMARY KEY,
                          "product_id" int,
                          "count" int,
                          "price" numeric(10,2)
);