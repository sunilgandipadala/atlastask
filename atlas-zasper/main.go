package main

import (
	"atlastask/models"
	"atlastask/utilities"
	"fmt"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <argument>")
		return
	}
	argument := os.Args[1]
	if argument == "migrate" {
		models.MigrateDB()

	} else {
		//Rest of the Code
		//Server Initialization and switch casess
		fmt.Println("Alter migrate")
		//utilities.AlterMigrationScripts()
		utilities.AlterScripts()
	}
}
