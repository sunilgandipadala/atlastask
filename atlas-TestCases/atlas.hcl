variable "destructive" {
  type    = bool
  default = false
}


data "external_schema" "gorm" {
  
  program = [
    "go",
    "run",
    "-mod=mod",
    "./",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  //url = "mysql://root:Sunil@513@localhost:3306/"

  url = "postgres://postgres:postgres@904@0.0.0.0:5432/atlasdb?search_path=public&sslmode=disable"
  //to use versioned migration
  //dev = "docker://mysql/8/schema"
  dev = "docker://postgres/15/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  diff {
    skip {
      drop_schema = !var.destructive
      drop_table  = !var.destructive
    }
  }
}