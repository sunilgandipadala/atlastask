package utilities

import (
	"fmt"
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

// this method will recieve .. table_name as param
func AlterScripts(table_name string) (add_querie string, drop_querie string, rename_stmt string) {
	//table_name := `"customers"`
	i, j := 0, 0
	var rename_stmts, add_queries, drop_queries = "", "", ""

	old_column_list := old_list[table_name] //we have to check that old_columns will be updated with new_columns once done with migration updates
	new_column_list := New_list[table_name]

	for i < len(old_column_list) && j < len(new_column_list) {
		CategorizeStmts(table_name, &rename_stmts, &add_queries, &drop_queries, old_column_list, new_column_list, &i, &j)
	}

	for i < len(old_column_list) {
		col := strings.Split(old_column_list[i], " ")[0]
		drop_queries += "DROP COLUMN " + col + ";"
		i += 1
	}
	for j < len(new_column_list) {
		add_queries += "ALTER TABLE " + table_name + " " + "ADD COLUMN " + new_column_list[j] + ";"
		j += 1
	}

	return add_queries, drop_queries, rename_stmts
}

func CategorizeStmts(table_name string, rename_stmts *string, add_queries *string, drop_queries *string, old_column_list []string, new_column_list []string, i, j *int) {
	if ok := PresentInList(old_column_list[*i], new_column_list); ok {
		if old_column_list[*i] == new_column_list[*j] {
			*i += 1
			*j += 1
			return
		} else {
			*add_queries += "ALTER TABLE " + table_name + " " + "ADD COLUMN " + new_column_list[*j] + ";"
			*j += 1
			return
		}
	} else if not_ok := PresentInList(new_column_list[*j], old_column_list); !not_ok { //there rasining some logical error will look into it..
		//may be here... there is a special case we need to check for --renaming...in the end,...
		col := strings.Split(old_column_list[*i], " ")[0]
		*rename_stmts += "ALTER TABLE " + table_name + " " + "RENAME " + col + " COLUMN TO " + new_column_list[*j] + ";"
		*i += 1
		*j += 1
	} else {
		col := strings.Split(old_column_list[*i], " ")[0]
		fmt.Println(old_column_list[*i], new_column_list[*j])
		*drop_queries += "DROP COLUMN " + col + ";"
		*i += 1
	}
}

/*

func CreateRenameStatements(alter_statements *string, table_name *string, content_queries []string, content string, i int) {
	drop_count := strings.Count(content, "DROP COLUMN")
	add_count := strings.Count(content, "ADD COLUMN")

	//Even there is an issue - as if some one has done one column removal and added one column to the struct it causes error
	if drop_count == 1 && add_count == 1 {
		old_column_name := strings.TrimSpace(content_queries[i][strings.Index(content_queries[i], "DROP COLUMN")+11:])
		new_column_name := strings.TrimSpace(content_queries[i+1][strings.Index(content_queries[i+1], "ADD COLUMN")+10:])
		*alter_statements += "ALTER TABLE " + *table_name + " RENAME COLUMN " + old_column_name + " TO " + new_column_name + ";"
		i += 1
	} else if strings.Contains(content_queries[i], "ADD COLUMN") {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[i] + ";"
	}
}


for i = 0; i+1 < len(content_queries) && len(content_queries) > 1; i++ {
		*flag += 1
		CreateRenameStatements(alter_statements, table_name, content_queries, content, i)
	}

	if i == len(content_queries)-1 && strings.Contains(content_queries[i], "ADD COLUMN") {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[i] + ";"
	}
	if *flag == 0 {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[0] + ";"
	}
*/

/* ======== The major task is to store the old_list and new_list updation =============

--- we will get new_list everytime that is fine....
--- In the same order how to get old_list

*/
