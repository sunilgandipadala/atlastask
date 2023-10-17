-- Modify "pets" table
ALTER TABLE "public"."pets" DROP COLUMN "name", ADD COLUMN "pet_name" text NULL, ADD COLUMN "deleted_time" timestamptz NULL;
