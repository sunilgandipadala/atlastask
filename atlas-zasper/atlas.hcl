variable "destructive" {
  type    = bool
  default = false
}


data "external_schema" "gorm" {
  
  program = [
    "go",
    "run",
    "main.go",
    "migrate",
    "-mod=mod",
    "./",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  //url = "mysql://root:Sunil@513@:3306/employees"
  url = "postgres://postgres:postgres@904@172.31.50.134:5432/TestDB?search_path=public&sslmode=disable"
  
  //to use versioned migration
  dev = "docker://postgres/15/dev"
  //dev = "docker://mysql/8/dev"
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