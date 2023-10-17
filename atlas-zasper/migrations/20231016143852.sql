-- Modify "buyers" table
ALTER TABLE "public"."buyers" DROP COLUMN "contractName", DROP COLUMN "contractPrice", DROP COLUMN "contractPriority", DROP COLUMN "contractCap", DROP COLUMN "contractCapFrequency", DROP COLUMN "contractStatus", DROP COLUMN "httpAthPassword";
-- Modify "customers" table
ALTER TABLE "public"."customers" ALTER COLUMN "created_at" DROP NOT NULL;
-- Modify "rules_Info" table
ALTER TABLE "public"."rules_Info" DROP COLUMN "flowName", DROP COLUMN "name", ADD COLUMN "flowName1" character varying(255) NULL;
-- Modify "userdetails" table
ALTER TABLE "public"."userdetails" DROP COLUMN "fid", DROP COLUMN "age", DROP COLUMN "phone", DROP COLUMN "address", DROP COLUMN "carBuyingInterest", DROP COLUMN "homeBuyInterest";
-- Drop "api_queue_details" table
DROP TABLE "public"."api_queue_details";
-- Drop "buyer_stats" table
DROP TABLE "public"."buyer_stats";
-- Drop "national_audience" table
DROP TABLE "public"."national_audience";
-- Drop "post_history" table
DROP TABLE "public"."post_history";
-- Drop "rule_queue_details" table
DROP TABLE "public"."rule_queue_details";
-- Drop "rules" table
DROP TABLE "public"."rules";
-- Drop "start_point_config" table
DROP TABLE "public"."start_point_config";
