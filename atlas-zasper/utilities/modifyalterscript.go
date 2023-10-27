package utilities

import (
	"atlastask/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var File_count = 0
var PreviousList map[string][]string

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

	absolutePath := GetAbsolutePath(relativePath)
	scanner, file, err := GetFileScanner(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var filename string
	for scanner.Scan() {
		File_count += 1
		filename = scanner.Text()
	}
	return strings.Split(filename, " ")[0]

}

func GetFileScanner(absolutePath string) (*bufio.Scanner, *os.File, error) {
	f, err := os.Open(absolutePath)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return scanner, f, nil
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
// MySQL hasn't tested... need to do some modifications to alterations as it may create issues while Renaming..
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

// This was done.
func UpdateAltersPostgres(scanner *bufio.Scanner, alter_statements *string, table_name *string, flag *int, content string) {

	*table_name = strings.Split(strings.Split(content, " ")[2], ".")[1] //it excludes public

	add_stmts, _, rename_stmts := AlterScripts(*table_name)
	if add_stmts != "" || rename_stmts != "" {
		*alter_statements += add_stmts + rename_stmts + ";"
	}
}

// This need to be updated at line 125 and 126
func UpdateQueries(scanner *bufio.Scanner, alter_statements *string, file_content *string, table_name *string, flag *int, alters_map map[string]string, dbtype string) {
	AssignListofStmts()
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
				*alter_statements = ""
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
	var file_content, alter_statements, flag = "", "", 0
	var table_name string
	var alters_map = make(map[string]string)

	UpdateQueries(scanner, &alter_statements, &file_content, &table_name, &flag, alters_map, dbtype)
	if alter_statements != "" && file_content != "" && table_name != "" {
		if dbtype == "mysql" {
			for key, value := range alters_map {
				GenerateShFiles(value, key, dbtype)
			}
		} else {
			GenerateShFiles(alter_statements, `"postgres"`, dbtype) //-- this should generate the .sh file
		}
	} else {
		log.Fatal("No New Migrations Found")
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

	relativePath := "alter-scripts/alter_" + table_name[1:len(table_name)-1] + ".sh"
	absolutePath := GetAbsolutePath(relativePath)
	CreateFile(absolutePath, script_text)

}

// The function we will call to alter the migrations
func AlterMigrationScripts() {

	filename := GetFileName()
	relativePath := "migrations\\" + filename
	absolutePath := GetAbsolutePath(relativePath)
	scanner, file, err := GetFileScanner(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	LoadEnv()
	dbType := os.Getenv("DB_TYPE")
	file_content := UpdateMigrations(scanner, dbType)
	if file_content != "" {
		CreateFile(absolutePath, file_content)
	}
	UpdateOldList()

}

func UpdateOldList() {
	err := os.WriteFile("data.txt", []byte(models.Statements), 0644)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
}

func AssignListofStmts() {
	New_list = models.GetNewColumnList()
	fileData, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	// Convert the byte slice back to a string
	retrieved_data := string(fileData)
	old_list = GetColumnList(retrieved_data)
}

func GetColumnList(schema string) map[string][]string {
	old_data := schema
	column_names_map := make(map[string][]string)
	statements_arr := strings.Split(old_data, ";")
	for _, stats := range statements_arr {
		if strings.Contains(stats, "CREATE TABLE") {
			table_name := strings.Split(stats, " ")[2] //it excludes public
			stats := strings.TrimSpace(stats[strings.Index(stats, "\" (")+3 : len(stats)-2])
			column_names_map[table_name] = strings.Split(stats, ",")
		}
	}
	return column_names_map
}
