CREATE UNIQUE INDEX ON "basket" ("user_id", "product_id");
CREATE UNIQUE INDEX ON "favorites" ("user_id", "product_id");
CREATE UNIQUE INDEX ON "prices" ("product_id", "count");
CREATE UNIQUE INDEX ON "prices" ("count", "product_id");
