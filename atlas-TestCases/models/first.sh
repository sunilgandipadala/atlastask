pt-online-schema-change D=employees,t="employees" --execute --alter \
 'RENAME COLUMN "name" TO "pet_name" text NULL, ADD COLUMN "sunil" timestamptz NULL' \
 --alter-foreign-keys-method "rebuild_constraints" \
 --ask-pass --check-alter --check-unique-key-change --null-to-not-null --swap-tables --no-drop-old-table

