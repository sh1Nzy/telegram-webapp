CREATE TABLE "orders" (
    "id" SERIAL PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "status" varchar(50) NOT NULL DEFAULT 'pending',
    "created_at" timestamp DEFAULT (current_timestamp)
);
