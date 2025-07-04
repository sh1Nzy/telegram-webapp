CREATE TABLE "basket" (
                          "id" SERIAL PRIMARY KEY,
                          "user_id" integer,
                          "product_id" integer,
                          "quantity" integer NOT NULL,
                          "added_at" timestamp DEFAULT (current_timestamp)
);