-- Create "api_queue_details" table
CREATE TABLE "api_queue_details" ("id" bigserial NOT NULL, "thenscope" text NULL, "customer_id" bigint NULL, "cid" character varying(20) NULL, "status" character(1) NULL, "flowEnded" boolean NULL, "ReddisKey" character varying(20) NULL, "storeResponse" text NULL, "TimePicked" timestamptz NULL, "TimeFired" timestamptz NULL, "timeEnded" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "atlas_schema_revisions" table
CREATE TABLE "atlas_schema_revisions" ("version" character varying NOT NULL, "description" character varying NOT NULL, "type" bigint NOT NULL DEFAULT 2, "applied" bigint NOT NULL DEFAULT 0, "total" bigint NOT NULL DEFAULT 0, "executed_at" timestamptz NOT NULL, "execution_time" bigint NOT NULL, "error" text NULL, "error_stmt" text NULL, "hash" character varying NOT NULL, "partial_hashes" jsonb NULL, "operator_version" character varying NOT NULL, PRIMARY KEY ("version"));
-- Create "buyer_stats" table
CREATE TABLE "buyer_stats" ("id" bigserial NOT NULL, "buyer_id" bigint NULL, "per_hour_count" bigint NULL, "per_hour_date" timestamptz NULL, "per_day_count" bigint NULL, "per_day_date" date NULL, "per_week_count" bigint NULL, "per_week_date" date NULL, "per_month_count" bigint NULL, "per_month_date" date NULL, "per_year_count" bigint NULL, "per_year_date" date NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "buyergroups" table
CREATE TABLE "buyergroups" ("id" bigserial NOT NULL, "groupName" character varying(100) NOT NULL, "buyerIds" character varying(100) NULL, "priority" character varying(100) NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "idx_buyergroups_group_name" to table: "buyergroups"
CREATE UNIQUE INDEX "idx_buyergroups_group_name" ON "buyergroups" ("groupName");
-- Create "buyers" table
CREATE TABLE "buyers" ("id" bigserial NOT NULL, "name" character varying(255) NOT NULL, "price" numeric NULL, "webhookUrl" text NULL, "webhookMethod" character varying(20) NULL, "webhookTimeout" character varying(30) NULL, "headers" text NULL, "requestBody" text NULL, "httpAuthUsername" character varying(20) NULL, "httpAuthPassword" character varying(20) NULL, "responseModel" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "contractName" character varying(100) NULL, "contractPrice" numeric NULL, "contractPriority" bigint NULL, "contractCap" bigint NULL, "contractCapFrequency" text NULL, "contractStatus" boolean NULL, "httpAthPassword" character varying(20) NULL, PRIMARY KEY ("id"));
-- Create index "idx_buyers_name" to table: "buyers"
CREATE UNIQUE INDEX "idx_buyers_name" ON "buyers" ("name");
-- Create "customer_profiles" table
CREATE TABLE "customer_profiles" ("customer_id" bigserial NOT NULL, "user_id" bigint NULL, "edu_level" character varying(255) NULL, "edu_interest" character varying(255) NULL, "edu_updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "test_id" bigint NULL, "health_insurance" character varying(255) NULL, "health_multiple_prescription" character varying(255) NULL, "health_diabetic" character varying(255) NULL, "health_cpap" character varying(255) NULL, "health_low_income" character varying(255) NULL, "health_back_pain" character varying(255) NULL, "health_updated_at" timestamptz NULL, "finance_checking_account" character varying(255) NULL, "finance_debt_amount" character varying(255) NULL, "finance_debt" boolean NULL, "finance_updated_at" timestamptz NULL, "car_owner" character varying(255) NULL, "car_updated_at" timestamptz NULL, "car_insurance" character varying(255) NULL, "car_license" text NULL, "work_status" character varying(255) NULL, "work_updated_at" timestamptz NULL, "edu_jornaya_token" character varying(255) NULL, "captcha_score" numeric NULL, "captcha_updated_at" timestamptz NULL, "navy_program" character varying(255) NULL, "us_citizen" boolean NULL, "gpa_average" numeric NULL, "fq_score" bigint NULL, "fq_updated_at" timestamptz NULL, "ranch" character varying(255) NULL, "brta_filt" character varying(255) NULL, "glad1" character varying(255) NULL, "glad2" character varying(255) NULL, "home_owner" character varying(255) NULL, "home_updated_at" timestamptz NULL, "debt_type" character varying(255) NULL, "edu_grant" character varying(255) NULL, "finex" character varying(255) NULL, "demply" character varying(255) NULL, "clej" character varying(255) NULL, "auto_accident" character varying(255) NULL, "accident" character varying(50) NULL, "medicaid" character varying(50) NULL, "medicare" character varying(50) NULL, "auto_year" character varying(50) NULL, "auto_make" character varying(50) NULL, "auto_model" character varying(50) NULL, "additional_fields_db" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("customer_id"));
-- Create "customers" table
CREATE TABLE "customers" ("id" bigserial NOT NULL, "first_name" character varying(255) NULL, "last_name" character varying(255) NULL, "age" bigint NULL, "email" character varying(50) NOT NULL, "gender" character varying(10) NULL, "address1" character varying(255) NULL, "address2" character varying(255) NULL, "city" character varying(255) NULL, "state" character varying(100) NULL, "country" character varying(100) NULL, "zipcode" character varying(50) NULL, "subscribed" boolean NULL, "subscribed_at" timestamptz NULL, "s1" bigint NULL, "s2" bigint NULL, "s3" character varying(50) NULL, "user_agent" character varying(100) NULL, "optin_ip" character varying(50) NULL, "optin_url" character varying(100) NULL, "optin_time" timestamptz NULL, "utm_source" character varying(100) NULL, "utm_medium" character varying(100) NULL, "utm_campaign" character varying(100) NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NULL, "phone" character varying(20) NULL, "q" character varying(255) NULL, "l" character varying(255) NULL, "ongage_contact_id" character varying(255) NULL, "ongage_status" bigint NULL, "oversight_contact_id" character varying(255) NULL, "oversight_status" bigint NULL, "validation" character varying(255) NULL, "validation_date" timestamptz NULL, "processed" boolean NULL, "processed_date" timestamptz NULL, "tz" text NULL, "opt_in_phone" boolean NULL, "dob" text NULL, "click_ip" text NULL, "pushnami_id" character varying(255) NULL, "optin_phone" bigint NULL, "last_visit" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "idx_created_at" to table: "customers"
CREATE INDEX "idx_created_at" ON "customers" ("created_at");
-- Create index "idx_customers_email" to table: "customers"
CREATE UNIQUE INDEX "idx_customers_email" ON "customers" ("email");
-- Create index "idx_phone" to table: "customers"
CREATE INDEX "idx_phone" ON "customers" ("phone");
-- Create "lead_post_history" table
CREATE TABLE "lead_post_history" ("id" bigserial NOT NULL, "buyer_id" bigint NULL, "trace_id" text NULL, "buyer_group_id" bigint NULL, "customer_id" bigint NULL, "knowledge_base_name" character varying(100) NULL, "rule_manifest_id" character varying(255) NULL, "rulemanifest_status" character varying(100) NULL, "response_code" bigint NULL, "status" character varying(255) NULL, "response" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "idx_created_at_buyer_id_status" to table: "lead_post_history"
CREATE INDEX "idx_created_at_buyer_id_status" ON "lead_post_history" ("buyer_id", "status", "created_at");
-- Create index "idx_customer_id" to table: "lead_post_history"
CREATE INDEX "idx_customer_id" ON "lead_post_history" ("customer_id");
-- Create index "idx_rule_manifest_id_customer_id" to table: "lead_post_history"
CREATE INDEX "idx_rule_manifest_id_customer_id" ON "lead_post_history" ("customer_id", "rule_manifest_id");
-- Create index "idx_trace_id" to table: "lead_post_history"
CREATE INDEX "idx_trace_id" ON "lead_post_history" ("trace_id");
-- Create "national_audience" table
CREATE TABLE "national_audience" ("zeta_id" text NULL);
-- Create "post_history" table
CREATE TABLE "post_history" ("id" bigserial NOT NULL, "buyer_id" bigint NULL, "buyer_group_id" bigint NULL, "customer_id" bigint NULL, "knowledgebase_name" character varying(100) NULL, "rulemanifest_id" character varying(255) NULL, "rulemanifest_status" character varying(100) NULL, "responsecode" bigint NULL, "resp_msg" character varying(255) NULL, "resp_result" text NULL, "created_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "rule_queue_details" table
CREATE TABLE "rule_queue_details" ("id" bigserial NOT NULL, "actionName" character varying(100) NULL, "customer_id" bigint NULL, "cid" character varying(20) NULL, "ReddisKey" character varying(20) NULL, "status" text NULL, "endpoint" character varying(100) NULL, "method" character varying(10) NULL, "queryParams" text NULL, "requestBody" text NULL, "goRoutine" character varying(50) NULL, "TimePicked" timestamptz NULL, "TimeFired" timestamptz NULL, "timeEnded" timestamptz NULL, "trace_id" text NULL, PRIMARY KEY ("id"));
-- Create "rules" table
CREATE TABLE "rules" ("id" bigserial NOT NULL, "rule" text NULL, "ruleName" character varying(100) NULL, "ruleGroup" character varying(100) NULL, "priority" bigint NULL, PRIMARY KEY ("id"));
-- Create "rules_Info" table
CREATE TABLE "rules_Info" ("id" bigserial NOT NULL, "flowName" character varying(255) NOT NULL, "manifest" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "name" character varying(255) NULL, PRIMARY KEY ("id"));
-- Create index "idx_rules_Info_flow_name" to table: "rules_Info"
CREATE UNIQUE INDEX "idx_rules_Info_flow_name" ON "rules_Info" ("flowName");
-- Create "start_point_config" table
CREATE TABLE "start_point_config" ("id" bigserial NOT NULL, "name" character varying(255) NOT NULL, "start_point" bigint NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "idx_start_point_config_name" to table: "start_point_config"
CREATE UNIQUE INDEX "idx_start_point_config_name" ON "start_point_config" ("name");
-- Create "userdetails" table
CREATE TABLE "userdetails" ("id" bigserial NOT NULL, "userName" character varying(100) NULL, "email" character varying(50) NULL, "fid" character varying(100) NULL, "age" bigint NULL, "phone" character varying(13) NULL, "address" character varying(500) NULL, "carBuyingInterest" boolean NULL, "homeBuyInterest" boolean NULL, "password" character varying(100) NULL, PRIMARY KEY ("id"));
-- Create "work_flow_json_details" table
CREATE TABLE "work_flow_json_details" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "json" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "idx_name_created_at" to table: "work_flow_json_details"
CREATE INDEX "idx_name_created_at" ON "work_flow_json_details" ("name", "created_at");
-- Create index "idx_work_flow_json_details_name" to table: "work_flow_json_details"
CREATE UNIQUE INDEX "idx_work_flow_json_details_name" ON "work_flow_json_details" ("name");
-- Create "workflow_execution" table
CREATE TABLE "workflow_execution" ("trace_id" uuid NULL, "workflow_id" uuid NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "payload" text NULL);
