table "employees" {
  schema = schema.practice_db
  column "id" {
    null = false
    type = varchar(191)
  }
  column "name" {
    null = false
    type = longtext
  }
  column "gender" {
    null = true
    type = longtext
  }
  column "role" {
    null = true
    type = longtext
  }
  column "district" {
    null = true
    type = longtext
  }
  column "pincode" {
    null = true
    type = bigint
  }
  column "age" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  index "id" {
    unique  = true
    columns = [column.id]
  }
}
table "sunil_practicetable" {
  schema = schema.practice_db
  column "tint" {
    null = true
    type = tinyint
  }
  column "sint" {
    null = true
    type = smallint
  }
  column "intval" {
    null = true
    type = int
  }
  column "bint" {
    null = true
    type = bigint
  }
}
table "tasks" {
  schema = schema.practice_db
  column "name" {
    null = true
    type = longtext
  }
  column "webhook_url" {
    null = true
    type = varchar(191)
  }
  column "webhook_method" {
    null = true
    type = longtext
  }
  column "webhook_timeout" {
    null = true
    type = longtext
  }
  column "response" {
    null = true
    type = longtext
  }
  index "webhook_url" {
    unique  = true
    columns = [column.webhook_url]
  }
}
table "users" {
  schema = schema.practice_db
  column "name" {
    null = true
    type = text
  }
  column "email" {
    null = true
    type = varchar(191)
  }
  column "phone" {
    null = true
    type = longtext
  }
  column "password" {
    null = true
    type = longtext
  }
  column "confirm_password" {
    null = true
    type = longtext
  }
  column "security_id" {
    null = true
    type = longtext
  }
  column "security_key" {
    null = true
    type = longtext
  }
  column "premium_plan" {
    null = true
    type = longtext
  }
  column "age" {
    null = true
    type = bigint
  }
  index "email" {
    unique  = true
    columns = [column.email]
  }
}


schema "practice_db" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
