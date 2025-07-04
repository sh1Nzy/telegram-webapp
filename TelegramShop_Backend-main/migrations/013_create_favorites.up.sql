CREATE TABLE "favorites" (
                             "id" SERIAL PRIMARY KEY,
                             "user_id" integer,
                             "product_id" integer,
                             "added_at" timestamp DEFAULT (current_timestamp)
);