pg-online-schema-change perform --alter-statement 'ALTER TABLE "customers" ADD COLUMN "newone" varchar(255);ALTER TABLE "customers" RENAME COLUMN "first_name" TO "full_name";' --dbname "TestDB" --host "localhost" --username "postgres"