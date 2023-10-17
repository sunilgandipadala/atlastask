-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "email" text NULL;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
