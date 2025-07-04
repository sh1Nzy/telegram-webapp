CREATE TABLE "categories" (
                              "id" SERIAL PRIMARY KEY,
                              "name" text UNIQUE NOT NULL,
                              "image" text
);