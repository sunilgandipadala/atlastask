package models

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func MigrateDB() {
	MigrateDatabaseTables(DATABASE_APPLICATION_TABLES)
}

func MigrateDatabaseTables(tables []DatabaseTableInterface) {
	modelss := []interface{}{}
	for _, table := range tables {
		modelss = append(modelss, table)
	}
	AtlasSchema(modelss)
}
func AtlasSchema(modelss []interface{}) {
	stmts, err := gormschema.New("postgres").Load(modelss...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
