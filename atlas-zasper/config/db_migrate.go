package config

/*
import (
	"fmt"

	model "zasper-be/zasper/models"
	"zasper-be/zasper/shared"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgresDatabase() {

	db, err := GetConnectionByType(Config.PostgresUsername, Config.PostgresPassword, Config.PostgresHost, Config.PostgresDb, shared.ConvertToString(Config.PostgresPort), "postgres")
	if err != nil {
		log.Error().Msg("error while connecting postgress database")
		panic(err.Error())
	}
	log.Info().Msg("Successfully connected to PostGres database")
	DB = db
	if Config.LogLevel == "DEBUG" {
		DB = db.Debug()
	}

}

func ConnectDatabase() {
	// root:@tcp(localhost:3306)/hij_local?charset=utf8&parseTime=True&loc=UTC&timeout=10s
	// db, err := gorm.Open(mysql.Open(`smartuser:smart12#$@tcp(10.218.18.157)/test_db?charset=utf8&parseTime=True&loc=UTC&timeout=10s`), &gorm.Config{})

	db, err := GetConnectionByType(Config.MySqlUsername, "", Config.MySqlHost, Config.MySqlDb, shared.ConvertToString(Config.MySqlPort), "mysql")
	if err != nil {
		log.Error().Msg("error")
		panic(err.Error())
	}
	log.Info().Msg("success")
	Db := db
	if Config.LogLevel == "DEBUG" {
		DB = Db.Debug()
	}
}

// postgres://your_postgres_username:your_postgres_password@localhost:5432/your_database_name?sslmode=disable

func MigrateDB() {
	MigrateDatabaseTables(model.DATABASE_APPLICATION_TABLES)
}

func MigrateDatabaseTable(table model.DatabaseTableInterface) {
	// _ = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table).Error
	//DB.AutoMigrate(table)
	// .Error
	// if err != nil {
	// 	fmt.Printf("Unable to migrate %s: err=%v\n", table.TableName(), err)
	// 	log.Fatalf("Unable to migrate %s: err=%v\n", table.TableName(), err)
	// }
}

func MigrateDatabaseTables(tables []model.DatabaseTableInterface) {
	for _, table := range tables {
		MigrateDatabaseTable(table)
	}
}

func GetConnectionByType(userName, passwd, host, dbName, port, dbType string) (*gorm.DB, error) {
	var connectionString string
	var db *gorm.DB
	var err error

	if dbType == "postgres" {
		connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			userName,
			passwd,
			dbName,
			port,
		)
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	} else if dbType == "mysql" {
		connectionString = fmt.Sprintf(`%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC&timeout=10s`,
			userName,
			host,
			port,
			dbName)
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	}

	return db, err
}
*/
