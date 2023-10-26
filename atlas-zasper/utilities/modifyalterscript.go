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
	*alter_statements += add_stmts + rename_stmts + ";"
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
	//fmt.Println("==========================Previous List==================:", PreviousList)
	file_content := UpdateMigrations(scanner, dbType) //we can take this from environment variables
	if file_content != "" {
		CreateFile(absolutePath, file_content)
	}
	PreviousList = New_list
}

func AssignListofStmts() {
	New_list = models.GetNewColumnList()
	if File_count == 1 {
		old_list = New_list
	} else {
		//old_list = PreviousList
		old_list = GetColumnList()
	}
}
func GetColumnList() map[string][]string {
	old_data := `CREATE TABLE "customers" ("id" bigserial,"first_name" varchar(255),"last_name" varchar(255),"age" bigint,"email" varchar(50) NOT NULL,"gender" varchar(10),"address1" varchar(255),"address2" varchar(255),"city" varchar(255),"state" varchar(100),"country" varchar(100),"zipcode" varchar(50),"phone" varchar(20),"subscribed" boolean,"subscribed_at" timestamptz,"s1" bigint,"s2" bigint,"s3" varchar(50),"q" varchar(255),"l" varchar(255),"ongage_contact_id" varchar(255),"ongage_status" bigint,"oversight_contact_id" varchar(255),"oversight_status" bigint,"validation" varchar(255),"validation_date" timestamptz,"processed" boolean,"processed_date" timestamptz,"tz" text,"opt_in_phone" boolean,"dob" text,"click_ip" text,"pushnami_id" varchar(255),"optin_phone" bigint,"last_visit" timestamptz,"user_agent" varchar(100),"optin_ip" varchar(50),"optin_url" varchar(100),"optin_time" timestamptz,"utm_source" varchar(100),"utm_medium" varchar(100),"utm_campaign" varchar(100),"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE INDEX IF NOT EXISTS "idx_created_at" ON "customers" ("created_at");
	CREATE INDEX IF NOT EXISTS "idx_phone" ON "customers" ("phone");
	CREATE UNIQUE INDEX IF NOT EXISTS "idx_customers_email" ON "customers" ("email");
	CREATE TABLE "work_flow_json_details" ("id" uuid DEFAULT gen_random_uuid(),"name" text NOT NULL,"json" text,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE UNIQUE INDEX IF NOT EXISTS "idx_work_flow_json_details_name" ON "work_flow_json_details" ("name");
	CREATE INDEX IF NOT EXISTS "idx_name_created_at" ON "work_flow_json_details" ("name","created_at");
	CREATE TABLE "buyergroups" ("id" bigserial,"groupName" varchar(100) NOT NULL,"buyerIds" varchar(100),"priority" varchar(100),"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE INDEX IF NOT EXISTS "idx_created_at" ON "buyergroups" ("created_at");
	CREATE UNIQUE INDEX IF NOT EXISTS "idx_buyergroups_group_name" ON "buyergroups" ("groupName");
	CREATE TABLE "buyers" ("id" bigserial,"name" varchar(255) NOT NULL,"price" decimal,"webhookUrl" text,"webhookMethod" varchar(20),"webhookTimeout" varchar(30),"headers" text,"requestBody" text,"httpAuthUsername" varchar(20),"httpAuthPassword" varchar(20),"responseModel" text,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE INDEX IF NOT EXISTS "idx_created_at" ON "buyers" ("created_at");
	CREATE UNIQUE INDEX IF NOT EXISTS "idx_buyers_name" ON "buyers" ("name");
	CREATE TABLE "lead_post_history" ("id" bigserial,"buyer_id" bigint,"trace_id" text,"buyer_group_id" bigint,"customer_id" bigint,"knowledge_base_name" 
	varchar(100),"rule_manifest_id" varchar(255),"rulemanifest_status" varchar(100),"response_code" bigint,"status" varchar(255),"response" TEXT,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE INDEX IF NOT EXISTS "idx_rule_manifest_id_customer_id" ON "lead_post_history" ("customer_id","rule_manifest_id");
	CREATE INDEX IF NOT EXISTS "idx_customer_id" ON "lead_post_history" ("customer_id");
	CREATE INDEX IF NOT EXISTS "idx_trace_id" ON "lead_post_history" ("trace_id");
	CREATE INDEX IF NOT EXISTS "idx_created_at_buyer_id_status" ON "lead_post_history" ("buyer_id","status","created_at");
	CREATE TABLE "rules_Info" ("id" bigserial,"flowName" varchar(255),"name" TEXT,"manifest" TEXT,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
	CREATE INDEX IF NOT EXISTS "idx_name_created_at" ON "rules_Info" ("created_at");
	CREATE TABLE "userdetails" ("id" bigserial,"userName" varchar(100),"password" varchar(100),"email" varchar(50),PRIMARY KEY ("id"));
	CREATE TABLE "customer_profiles" ("customer_id" bigserial,"user_id" bigint,"edu_level" varchar(255),"edu_interest" varchar(255),"edu_updated_at" timestamptz,"deleted_at" timestamptz,"test_id" bigint,"health_insurance" varchar(255),"health_multiple_prescription" varchar(255),"health_diabetic" varchar(255),"health_cpap" varchar(255),"health_low_income" varchar(255),"health_back_pain" varchar(255),"health_updated_at" timestamptz,"finance_checking_account" varchar(255),"finance_debt_amount" varchar(255),"finance_debt" boolean,"finance_updated_at" timestamptz,"car_owner" varchar(255),"car_updated_at" timestamptz,"car_insurance" varchar(255),"car_license" text,"work_status" varchar(255),"work_updated_at" timestamptz,"edu_jornaya_token" varchar(255),"captcha_score" decimal,"captcha_updated_at" timestamptz,"navy_program" varchar(255),"us_citizen" boolean,"gpa_average" decimal,"fq_score" bigint,"fq_updated_at" timestamptz,"ranch" varchar(255),"brta_filt" varchar(255),"glad1" varchar(255),"glad2" varchar(255),"home_owner" varchar(255),"home_updated_at" timestamptz,"debt_type" varchar(255),"edu_grant" varchar(255),"finex" varchar(255),"demply" varchar(255),"clej" varchar(255),"auto_accident" 
	varchar(255),"accident" varchar(50),"medicaid" varchar(50),"medicare" varchar(50),"auto_year" varchar(50),"auto_make" varchar(50),"auto_model" varchar(50),"additional_fields_db" text,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("customer_id"));
	CREATE TABLE "workflow_execution" ("trace_id" uuid,"workflow_id" uuid,"payload" TEXT,"created_at" timestamptz,"updated_at" timestamptz);
	CREATE INDEX IF NOT EXISTS "idx_name_created_at" ON "workflow_execution" ("created_at");
	`

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
