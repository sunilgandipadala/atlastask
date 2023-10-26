package utilities

import (
	"fmt"
	"strings"
)

//var old_list = []string{"id", "name", "price", "webhookurl", "another", "headers", "requestBody", "gender", "lastone"}
//var new_list = []string{"id", "pet_name", "what..", "name", "price", "hewllooo", "webhookurl", "headersnew", "requestBody", "lastonee"}
/*
func GetOldList() map[string][]string {

	//we will do this from the imported schema.. is more better....
	var table_column_map = make(map[string][]string)
	var columns = "id, first_name, last_name, age, email, gender, address1, address2, city, state, country, zipcode, subscribed, subscribed_at, s1, s2, s3, user_agent, optin_ip, optin_url, optin_time, utm_source, utm_medium, utm_campaign, created_at, updated_at, phone, q, l, ongage_contact_id, ongage_status, oversight_contact_id, oversight_status, validation, validation_date, processed, processed_date, tz, opt_in_phone, dob, click_ip, pushnami_id, optin_phone, last_visit"
	table_column_map[`"customers"`] = strings.Split(columns, ",")
	table_column_map[`"work_flow_json_details"`] = strings.Split("id, name, json, created_at, updated_at", ",")
	table_column_map[`"buyergroups"`] = strings.Split(`id, "groupName", "buyerIds", priority, created_at, updated_at`, ",")
	table_column_map[`"buyers"`] = strings.Split(`id, name, price, "webhookUrl", "webhookMethod", "webhookTimeout", headers, "requestBody", "httpAuthUsername", "httpAuthPassword", "responseModel", created_at, updated_at, "contractName", "contractPrice", "contractPriority", "contractCap", "contractCapFrequency", "contractStatus", "httpAthPassword"`, ",")
	table_column_map[`"lead_post_history"`] = strings.Split(`id, buyer_id, trace_id, buyer_group_id, customer_id, knowledge_base_name, rule_manifest_id, rulemanifest_status, response_code, status, response, created_at, updated_at`, ",")
	table_column_map[`"rules"`] = strings.Split(`id, rule, "ruleName", "ruleGroup", priority`, ",")
	table_column_map[`"userdetails"`] = strings.Split(`id, "userName", email, fid, age, phone, address, "carBuyingInterest", "homeBuyInterest", password`, ",")
	table_column_map[`"customer_profiles"`] = strings.Split(`customer_id, user_id, edu_level, edu_interest, edu_updated_at, deleted_at, test_id, health_insurance, health_multiple_prescription, health_diabetic, health_cpap, health_low_income, health_back_pain, health_updated_at, finance_checking_account, finance_debt_amount, finance_debt, finance_updated_at, car_owner, car_updated_at, car_insurance, car_license, work_status, work_updated_at, edu_jornaya_token, captcha_score, captcha_updated_at, navy_program, us_citizen, gpa_average, fq_score, fq_updated_at, ranch, brta_filt, glad1, glad2, home_owner, home_updated_at, debt_type, edu_grant, finex, demply, clej, auto_accident, accident, medicaid, medicare, auto_year, auto_make, auto_model, additional_fields_db, created_at, updated_at`, ",")
	table_column_map[`"workflow_execution"`] = strings.Split(`trace_id, workflow_id, created_at, updated_at, payload`, ",")

	return table_column_map
}
*/

var new_list = models.GetNewColumnList()

func PresentInList(samplecol string, sampleli []string) bool {
	for _, col := range sampleli {
		if col == samplecol {
			return true
		}
	}
	return false
}
func GetOnlyColumns(columns []string) []string {
	var new_columns []string
	for _, col := range columns {
		col = strings.Split(col, " ")[0]
		col = col[1 : len(col)-1]
		new_columns = append(new_columns, col)
	}
	fmt.Println(new_columns)
	return new_columns
}

// this method will recieve .. table_name as param
func AlterScripts() {
	table_name := `"customers"`
	i, j := 0, 0
	var AddedColumns, DroppedColumns = []string{}, []string{}
	var rename_stmts = ""
	old_column_list := new_list[table_name] //we have to check that old_columns will be updated with new_columns once done with migration updates
	new_column_list := GetOnlyColumns(new_list[table_name])
	for _, col := range new_column_list {
		fmt.Println("column:", col)
	}
	fmt.Println("OLD---COLUMNS';", len(old_column_list))
	fmt.Println("NEW - --- COlumns", len(new_column_list))
	for i < len(old_column_list) && j < len(new_column_list) {

		if ok := PresentInList(old_column_list[i], new_column_list); ok {
			if old_column_list[i] == new_column_list[j] {
				i += 1
				j += 1
				continue
			} else {
				AddedColumns = append(AddedColumns, new_column_list[j])
				j += 1
				continue
			}
		} else if not_ok := PresentInList(new_column_list[j], old_column_list); !not_ok && strings.Contains(new_column_list[j], old_column_list[i]) {
			rename_stmts += "RENAME " + old_column_list[i] + " COLUMN TO " + new_column_list[j] + "\n"
			i += 1
			j += 1
		} else {
			DroppedColumns = append(DroppedColumns, old_column_list[i])
			i += 1
		}
	}
	//here if i<len(old_column_list)-- add remaining columns to DROP LIST and j<len(new_column_list) -- add remaining columns to ADD list
	if i < len(old_column_list) {
		DroppedColumns = append(DroppedColumns, old_column_list[i:]...)
	}
	if j < len(new_column_list) {
		AddedColumns = append(AddedColumns, new_column_list[i:]...)
	}
	fmt.Println("Added Columns:", AddedColumns)
	fmt.Println("Dropped Columns:", DroppedColumns)
	fmt.Println("Renamed Columns: ", rename_stmts)
	GetOldList()
}

//To get column names from the generated script

/*CREATE TABLE "customers" ("id" bigserial,"first_name" varchar(255),"last_name" varchar(255),"age" bigint,"email" varchar(50) NOT NULL,"gender" varchar(10),"address1" varchar(255),"address2" varchar(255),"city" varchar(255),"state" varchar(100),"country" varchar(100),"zipcode" varchar(50),"phone" varchar(20),"subscribed" boolean,"subscribed_at" timestamptz,"s1" bigint,"s2" bigint,"s3" varchar(50),"q" varchar(255),"l" varchar(255),"ongage_contact_id" varchar(255),"ongage_status" bigint,"oversight_contact_id" varchar(255),"oversight_status" bigint,"validation" varchar(255),"validation_date" timestamptz,"processed" boolean,"processed_date" timestamptz,"tz" text,"opt_in_phone" boolean,"dob" text,"click_ip" text,"pushnami_id" varchar(255),"optin_phone" bigint,"last_visit" timestamptz,"user_agent" varchar(100),"optin_ip" varchar(50),"optin_url" varchar(100),"optin_time" timestamptz,"utm_source" varchar(100),"utm_medium" varchar(100),"utm_campaign" varchar(100),"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
CREATE INDEX IF NOT EXISTS "idx_created_at" ON "customers" ("created_at");
CREATE INDEX IF NOT EXISTS "idx_phone" ON "customers" ("phone");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_customers_email" ON "customers" ("email");
CREATE TABLE "work_flow_json_details" ("id" uuid DEFAULT gen_random_uuid(),"name" text NOT NULL,"json" text,"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
CREATE UNIQUE INDEX IF NOT EXISTS "idx_work_flow_json_details_name" ON "work_flow_json_details" ("name");
CREATE INDEX IF NOT EXISTS "idx_name_created_at" ON "work_flow_json_details" ("name","created_at");
CREATE TABLE "buyergroups" ("id" bigserial,"groupName" varchar(100) NOT NULL,"buyerIds" varchar(100),"priority" varchar(100),"created_at" timestamptz,"updated_at" timestamptz,PRIMARY KEY ("id"));
*/

/*store in the form of
{
	"tablename": []string{},
	"tablename": []string{},
}
 make(map[string] []string)
 ====== To extract column names=======
 in each sentence, if CREATE TABLE "--table--name--" (  then, after ( this character index to len(sentance)-2 are our columns names..
 now split the column names string with "," now each element is a particular column.. so, now we can access the column names from this list by using element.split(" ")[0]
 so, we can perform our operation...

 Everytime, we have to update this map and previous version into old_list map (a copy)

 Once - from the migrations if we found ALTER TABLE for that particular tablename, we will compare the columnd update that table statment only in the migrations.*/

/* -========== Get Columns names of each table into a Map===============-

 column_names_map := make(map[string] []string)
 for scanner.Scan(){
	content := scanner.Text()
	if 	strings.Contain(content,"CREATE TABLE"){
		content := strings.TrimSpace(content[strings.Index(content, "\" (")+3:len(content)-2])
		column_names_map["table_name"] = content.split(",")
	}
 }
*/
