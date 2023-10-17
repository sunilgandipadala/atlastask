package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetFileName() string {
	f, err := os.Open("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/atlas.sum") //we  have to modify the path
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var filename string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		filename = scanner.Text()
	}
	ofilename := strings.Split(filename, " ")
	return ofilename[0]
}

func UpdateMigrations(scanner *bufio.Scanner) (filec string) {
	//Variable Declaration & Initialization
	var file_content, alter_statements = "", ""
	var table_name string
	var flag = 0

	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "ALTER TABLE") {
			var flag = 1
			//Here have to check length conditions ---- Used for postgres -- need to check for mysql
			table_name = strings.Split(content, " ")[2]
			content = strings.TrimSpace(content[strings.Index(content, table_name) : len(content)-1]) //it will return a new string with trialing and ending removal of white spaces
			//no need of TrimSpace even
			content_queries := strings.Split(content, ",")

			for i := 0; i < len(content_queries) && len(content_queries) > 1; i++ {
				flag = 0
				if strings.Contains(content_queries[i], "DROP COLUMN") && strings.Contains(content_queries[i+1], "ADD COLUMN") {

					old_column_name := strings.TrimSpace(content_queries[i][strings.Index(content_queries[i], "DROP COLUMN")+11:])
					new_column_name := strings.TrimSpace(content_queries[i+1][strings.Index(content_queries[i+1], "ADD COLUMN")+10:])
					alter_statements += "RENAME COLUMN " + old_column_name + " TO " + new_column_name + ","
					i += 1

				} else {
					alter_statements += content_queries[i] + ","
				}
			}
			if flag == 1 {
				alter_statements += content_queries[0] + ","
			}
		} else {
			file_content += content + "\n"
		}
	}
	if flag == 1 {
		GenerateShFiles(alter_statements, table_name) //-- this should generate the .sh file
	}
	return strings.Replace(file_content, "\n\n", "\n", 1)
}

// for this we have to pass database and table name too
func GenerateShFiles(alter_statements string, table_name string) {
	alter_statements = strings.TrimSpace(alter_statements[:len(alter_statements)-1])
	//Here we can change the database value from .env file
	script_text := `pt-online-schema-change D=employees,t=` + table_name + ` --alter '` + alter_statements + `' ` + `--alter-foreign-keys-method "rebuild_constraints" --ask-pass --check-alter --check-unique-key-change --null-to-not-null --swap-tables --no-drop-old-table`
	err := os.WriteFile("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/models/first.sh", []byte(script_text), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(script_text)
}

func AlterMigrationScripts() {

	filename := GetFileName()
	//we got the file name name .. now need to open that file name
	f, err := os.Open("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	file_content := UpdateMigrations(scanner)

	file := fmt.Sprintf("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/%s", filename)
	err = os.WriteFile(file, []byte(file_content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
