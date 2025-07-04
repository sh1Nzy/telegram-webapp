CREATE TABLE "marks" (
                         "product_id" integer,
                         "user_id" integer,
                         "mark" numeric(10,2),
                         "created_at" timestamp DEFAULT (current_timestamp)
);