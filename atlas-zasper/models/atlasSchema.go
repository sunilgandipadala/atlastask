package models

import (
	"fmt"
	"os"
	"strings"

	"ariga.io/atlas-provider-gorm/gormschema"
)

var Statements string

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
	Statements = stmts
	//io.WriteString(os.Stdout, stmts)
}

func GetNewColumnList() map[string][]string {
	column_names_map := make(map[string][]string)
	MigrateDB()
	statements_arr := strings.Split(Statements, ";")
	for _, stats := range statements_arr {
		if strings.Contains(stats, "CREATE TABLE") {
			table_name := strings.Split(stats, " ")[2] //it excludes public
			stats := strings.TrimSpace(stats[strings.Index(stats, "\" (")+3 : len(stats)-2])
			column_names_map[table_name] = strings.Split(stats, ",")
		}
	}
	return column_names_map
}
