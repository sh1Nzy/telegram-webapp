CREATE TABLE "comments" (
                            "id" SERIAL PRIMARY KEY,
                            "comment" text,
                            "user_id" integer,
                            "product_id" integer,
                            "created_at" timestamp DEFAULT (current_timestamp)
);