-- Modify "users" table
ALTER TABLE "public"."users" ADD CONSTRAINT "chk_users_name" CHECK (name <> 'jinzhu'::text);