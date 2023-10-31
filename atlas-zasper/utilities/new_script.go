package utilities

import (
	"atlastask/models"
	"fmt"
	"os"
	"strings"
)

var New_list, old_list map[string][]string

func PresentInList(samplecol string, sampleli []string) bool {
	for _, col := range sampleli {
		if col == samplecol {
			return true
		}
	}
	return false
}

func AssignListofStmts() {
	New_list = models.GetNewColumnList()
	fileData, err := os.ReadFile("data.txt") //in data.txt initially we don't need to place any creational schema...
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	// Convert the byte slice back to a string
	retrieved_data := string(fileData)
	old_list = GetColumnList(retrieved_data)
}

func GetColumnNames(column_list []string) []string {
	var column_names []string
	for _, col := range column_list {
		column_names = append(column_names, strings.Split(col, " ")[0])
	}
	return column_names
}

func CategorizeStmts(table_name string, rename_stmts *string, add_queries *string, old_column_names []string, new_column_names []string, new_column_list []string, i, j *int) {
	//better doing with column names only without properties
	old_col := old_column_names[*i]
	new_col := new_column_names[*j]
	if ok := PresentInList(old_col, new_column_names); ok {
		if old_col == new_col {
			*i += 1
			*j += 1
			return
		} else if !strings.Contains(new_column_list[*j], "PRIMARY KEY") {
			*add_queries += "ALTER TABLE " + table_name + " " + "ADD COLUMN " + new_column_list[*j] + ";"
			*j += 1
			return
		}
	} else if not_ok := PresentInList(new_col, old_column_names); !not_ok {
		*rename_stmts += "ALTER TABLE " + table_name + " " + "RENAME COLUMN " + old_col + " TO " + new_col + ";"
		*i += 1
		*j += 1
	} else {
		*i += 1
	}
}

func AlterScripts(table_name string) (add_querie string, rename_stmt string) {

	i, j := 0, 0
	var rename_stmts, add_queries = "", ""
	old_column_list := old_list[table_name] //we have to check that old_columns will be updated with new_columns once done with migration updates
	new_column_list := New_list[table_name]
	old_column_names := GetColumnNames(old_column_list)
	new_column_names := GetColumnNames(new_column_list)

	for i < len(old_column_list) && j < len(new_column_list) {
		CategorizeStmts(table_name, &rename_stmts, &add_queries, old_column_names, new_column_names, new_column_list, &i, &j)
	}
	for j < len(new_column_names) {
		if !strings.Contains(new_column_list[j], "PRIMARY KEY") {
			add_queries += "ALTER TABLE " + table_name + " " + "ADD COLUMN " + new_column_list[j] + ";"
			j += 1
		}
	}

	return add_queries, rename_stmts
}
