package utilities

import (
	"bufio"
	"strings"
)

func UpdatedAltersPostgres(scanner *bufio.Scanner, alter_statements *string, table_name *string, flag *int, content string) {

	//Here there causing error when ALTER TABLE public.tablename isn't followed..
	*table_name = strings.Split(strings.Split(content, " ")[2], ".")[1] //it excludes public
	content = strings.TrimSpace(content[strings.Index(content, *table_name)+len(*table_name) : len(content)-1])
	content_queries := strings.Split(content, ",")

	for i := 0; i+1 < len(content_queries) && len(content_queries) > 1; i++ {
		*flag += 1
		if strings.Count(content, "DROP COUMN") == 1 && strings.Count(content, "ADD COLUMN") == 1 {
			if strings.Contains(content_queries[i], "DROP COLUMN") {
				if strings.Contains(content_queries[i+1], "ADD COLUMN") {
					old_column_name := strings.TrimSpace(content_queries[i][strings.Index(content_queries[i], "DROP COLUMN")+11:])
					new_column_name := strings.TrimSpace(content_queries[i+1][strings.Index(content_queries[i+1], "ADD COLUMN")+10:])
					*alter_statements += "ALTER TABLE " + *table_name + " RENAME COLUMN " + old_column_name + " TO " + new_column_name + ";"
					i += 1
					//Here this nested condition for the purpose of Drop Column Removal ...
				}
			} else {
				*alter_statements += content_queries[i] + ","
			}
		} else {
			*alter_statements += content_queries[i] + ","
		}
	}

	if *flag == 0 {
		*alter_statements += "ALTER TABLE " + *table_name + " " + content_queries[0] + ";"
	}

}

//The alternative solution is, that we have to follow some conditions while using this..
//In a single struct - whenever you are going to perform Renaming a column, then you are not supposed to Drop or Add column before running atlas migrate diff
//And in a single struct whenever you are going to perform more than one drop columns and one or more Add Columns Vice Versa - Don't Perform renamig a column before running atlas migrate diff.

//If we follow this, we can go with Count(DROP COLUMN) or Count(ADD COLUMN) --> to generate alter statements.
