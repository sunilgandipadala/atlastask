package main

import (
	"atlastask/models"
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {

	//models.AlterMigrationScripts()

	stmts, err := gormschema.New("postgres").Load(models.User{}, models.Pet{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)

}
