-- Modify "buyers" table
ALTER TABLE "public"."buyers" DROP COLUMN "contractName", DROP COLUMN "contractPrice", DROP COLUMN "contractPriority", DROP COLUMN "contractCap", DROP COLUMN "contractCapFrequency", DROP COLUMN "contractStatus", DROP COLUMN "httpAthPassword";
-- Modify "customers" table
ALTER TABLE "public"."customers" ALTER COLUMN "created_at" DROP NOT NULL;
-- Modify "rules_Info" table
ALTER TABLE "public"."rules_Info" DROP COLUMN "flowName", DROP COLUMN "name", ADD COLUMN "flowName1" character varying(255) NULL;
-- Modify "userdetails" table
ALTER TABLE "public"."userdetails" DROP COLUMN "fid", DROP COLUMN "age", DROP COLUMN "phone", DROP COLUMN "address", DROP COLUMN "carBuyingInterest", DROP COLUMN "homeBuyInterest";
