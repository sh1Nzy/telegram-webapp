CREATE TABLE "products" (
                            "id" SERIAL PRIMARY KEY,
                            "name" text NOT NULL,
                            "firm_id" integer,
                            "description" text NOT NULL,
                            "category_id" integer,
                            "attributes" text,
                            "sell_count" integer DEFAULT 0,
                            "stock" integer DEFAULT 1,
                            "image" text[]
);

