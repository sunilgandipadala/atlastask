-- Drop index "idx_buyers_name" from table: "buyers"
DROP INDEX "public"."idx_buyers_name";
-- Modify "buyers" table
ALTER TABLE "public"."buyers" DROP COLUMN "name", DROP COLUMN "price", DROP COLUMN "httpAuthPassword", ADD COLUMN "name_new" character varying(255) NOT NULL, ADD COLUMN "pricees" numeric NULL;
-- Create index "idx_buyers_name" to table: "buyers"
CREATE UNIQUE INDEX "idx_buyers_name" ON "public"."buyers" ("name_new");
