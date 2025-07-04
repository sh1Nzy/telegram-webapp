CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "telegram_id" bigint UNIQUE NOT NULL,
                         "username" text UNIQUE NOT NULL,
                         "created_at" timestamp DEFAULT current_timestamp
);
