package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

// To get absolutepath
func GetAbsolutePath(relativePath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	// Construct the absolute path by joining the current working directory and the relative path
	absolutePath := filepath.Join(cwd, relativePath)
	return absolutePath
}

// To getFileName
func GetFileName() string {
	relativePath := "migrations/atlas.sum"

	// Get the current working directory
	absolutePath := GetAbsolutePath(relativePath)
	f, err := os.Open(absolutePath) //we  have to modify the path
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

func GetFileScanner(absolutePath string) *bufio.Scanner {
	f, err := os.Open(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return scanner
}

// This will Create a file and writes
func CreateFile(filename string, file_body string) {
	err := os.WriteFile(filename, []byte(file_body), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//#---------------#-------------------#

// Make sure the migrations generated are the same type you are calling..
func UpdateAltersMysql(scanner *bufio.Scanner, alter_statements *string, table_name *string, flag *int, content string) {

	*table_name = strings.Split(content, " ")[2]
	content = strings.TrimSpace(content[strings.Index(content, *table_name) : len(content)-1])
	content_queries := strings.Split(content, ",")

	for i := 0; i < len(content_queries) && len(content_queries) > 1; i++ {
		*flag += 1
		if strings.Contains(content_queries[i], "DROP COLUMN") && strings.Contains(content_queries[i+1], "ADD COLUMN") {

			old_column_name := strings.TrimSpace(content_queries[i][strings.Index(content_queries[i], "DROP COLUMN")+11:])
			new_column_name := strings.TrimSpace(content_queries[i+1][strings.Index(content_queries[i+1], "ADD COLUMN")+10:])
			*alter_statements += "RENAME COLUMN " + old_column_name + " TO " + new_column_name + ","
			i += 1

		} else {
			*alter_statements += content_queries[i] + ","
		}
	}
	if *flag == 0 {
		*alter_statements += content_queries[0] + ","
	}

}

func UpdateAltersPostgres(scanner *bufio.Scanner, alter_statements *string, table_name *string, flag *int, content string) {

	//Here there causing error when ALTER TABLE public.tablename isn't followed..
	*table_name = strings.Split(strings.Split(content, " ")[2], ".")[1] //it excludes public
	content = strings.TrimSpace(content[strings.Index(content, *table_name)+len(*table_name) : len(content)-1])
	content_queries := strings.Split(content, ",")

	var i int
	for i = 0; i+1 < len(content_queries) && len(content_queries) > 1; i++ {
		*flag += 1

		if strings.Count(content, "DROP COUMN") == 1 && strings.Count(content, "ADD COLUMN") == 1 {
			old_column_name := strings.TrimSpace(content_queries[i][strings.Index(content_queries[i], "DROP COLUMN")+11:])
			new_column_name := strings.TrimSpace(content_queries[i+1][strings.Index(content_queries[i+1], "ADD COLUMN")+10:])
			*alter_statements += "ALTER TABLE " + *table_name + " RENAME COLUMN " + old_column_name + " TO " + new_column_name + ";"
			i += 1
		} else if strings.Contains(content_queries[i], "ADD COLUMN") {
			*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[i] + ";"
		}

	}
	if i == len(content_queries)-1 && strings.Contains(content_queries[i], "ADD COLUMN") {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[i] + ";"
	}
	if *flag == 0 {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[0] + ";"
	}

}

// This need to be updated at line 125 and 126
func UpdateQueries(scanner *bufio.Scanner, alter_statements *string, file_content *string, table_name *string, flag *int, alters_map map[string]string, dbtype string) {
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "DROP TABLE") {
			if strings.Contains(content, "ADD TABLE") {
				old_table_name := strings.TrimSpace(content[strings.Index(content, "DROP TABLE")+10:])
				new_table_name := strings.TrimSpace(content[strings.Index(content, "ADD TABLE")+9:])
				*file_content += "ALTER TABLE IF EXISTS " + old_table_name + " RENAME TO " + new_table_name
			} else {
				*file_content += ""
			}
		} else if strings.Contains(content, "ALTER TABLE") {

			if dbtype == "mysql" {
				//*alter_statements = ""
				UpdateAltersMysql(scanner, alter_statements, table_name, flag, content)
			} else if strings.Contains(content, "ALTER TABLE \"public\"") && (dbtype == "postgres") {
				UpdateAltersPostgres(scanner, alter_statements, table_name, flag, content)
			}
		} else {
			*file_content += content + "\n"
		}
		if dbtype == "mysql" && *table_name != "" {
			alters_map[*table_name] += *alter_statements
		}
	}

}

func UpdateMigrations(scanner *bufio.Scanner, dbtype string) (filec string) {
	//Variable Declaration & Initialization
	var file_content, alter_statements = "", ""
	var table_name string
	var flag = 0
	var alters_map = make(map[string]string)

	//UpdateQueries... and Update Alters...
	UpdateQueries(scanner, &alter_statements, &file_content, &table_name, &flag, alters_map, dbtype)

	if dbtype == "mysql" {
		for key, value := range alters_map {
			GenerateShFiles(value, key, dbtype)
		}
	} else {
		GenerateShFiles(alter_statements, `"postgres"`, dbtype) //-- this should generate the .sh file
	}
	return strings.Replace(file_content, "\n\n", "\n", 1)
}

// To Generate Sh file..
func GenerateShFiles(alter_statements string, table_name string, dbtype string) {
	if len(alter_statements) > 1 {
		alter_statements = strings.TrimSpace(alter_statements[:len(alter_statements)-1])
	}
	//Here we can change the database value from .env file
	script_text := `pt-online-schema-change D=` + os.Getenv("DB_NAME") + `,t=` + table_name + ` --alter '` + alter_statements + `' ` + `--alter-foreign-keys-method "rebuild_constraints" --ask-pass --check-alter --check-unique-key-change --null-to-not-null --swap-tables --no-drop-old-table`
	if dbtype == "postgres" {
		script_text = "pg-online-schema-change perform --alter-statement '" + alter_statements + `' --dbname "` + os.Getenv("DB_NAME") + `" --host "` + os.Getenv("DB_HOST") + `" --username "` + os.Getenv("DB_USER") + `"`
	}
	var relativePath string
	relativePath = "alter-scripts/alter_" + table_name[1:len(table_name)-1] + ".sh"
	absolutePath := GetAbsolutePath(relativePath)
	CreateFile(absolutePath, script_text)

}

// The function we will call to alter the migrations
func AlterMigrationScripts() {

	filename := GetFileName()
	relativePath := "migrations\\" + filename
	absolutePath := GetAbsolutePath(relativePath)
	//scanner := GetFileScanner(absolutePath) //scanner is not generating..
	//From here have to check once...
	f, err := os.Open(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//===== the method isn't working.. so will see later..
	LoadEnv()
	dbType := os.Getenv("DB_TYPE")
	file_content := UpdateMigrations(scanner, dbType) //we can take this from environment variables
	CreateFile(absolutePath, file_content)

}
