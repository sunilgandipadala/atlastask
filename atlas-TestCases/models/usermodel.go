package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"check:name <> 'jinzhu'"`
	Email string `gorm:"unique;primaryKey"`
	Pets  []Pet
}
type Pet struct {
	gorm.Model
	PetNamee    string
	User        User
	UserID      uint
	DeletedTime *time.Time `gorm:"index:idx_users_deleted_at,sort:ASC,nulls:LAST"`
}

// We have to work on this struct from now with MYSQL DB
type Employee struct {
	Emp_No     int    `gorm:"primarykey"`
	Birth_Date string `gorm:"type:date"`
	First_Name string `gorm:"type:Varchar(16)"`
	Last_Name  string `gorm:"type:Varchar(16)"`
	Gender     string
	hire_date  string `gorm:"type:date"`
}

/*func ConnectingDB() {
	//mysql_conn := "root:Sunil@513@/practice_db"
	//db, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	//dsn := "host=localhost port=5432 dbname=atlasdb user=postgres password=Postgres@904 sslmode=prefer connect_timeout=10"
	//db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(&UserData{}, &Pet{})
	//user := User{Name: "SunilK", Gender: "Male"}
	//db.Create(&user)
	//pet := Pet{Name: "Kitty", UserID: user.ID}
	//	db.Create(&pet)

	// Create a Pet associated with the User

	//Now we have to perform the remaining task

}

func AlterMigrationScripts() {
	f, err := os.Open("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/atlas.sum")
	if err != nil {
		log.Fatal(err)
	}

	var filename string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		filename = scanner.Text()
	}
	ofilename := strings.Split(filename, " ")

	//we got the file name name .. now need to open that file name
	f, err = os.Open("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/" + ofilename[0])
	if err != nil {
		log.Fatal(err)
	}

	//Variable Declaration & Initialization
	scanner = bufio.NewScanner(f)
	var file_content, alter_statements = "", ""
	var table_name string

	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "ALTER TABLE") {
			var flag = 1
			//Here have to check length conditions ---- Used for postgres -- need to check for mysql
			table_name = strings.Split(content, " ")[3]
			content = strings.TrimSpace(content[strings.Index(content, table_name) : len(content)-1])
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

	GenerateShFiles(alter_statements, table_name) //-- this should generate the .sh file
	defer f.Close()
	file := fmt.Sprintf("C:/Users/sunil.gandipadala/Desktop/Atlas/atlas-automigrations/migrations/%s", ofilename[0])
	err = os.WriteFile(file, []byte(file_content), 0644)

	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
*/
