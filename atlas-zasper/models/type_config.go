package models

var Configdetails *ConfigManifest

// Define the Config struct to hold the parsed data
type ConfigManifest struct {
	AppName          string   `mapstructure:"APP_NAME"`
	AppVersion       string   `mapstructure:"APP_VERSION"`
	AppEnv           string   `mapstructure:"APP_ENV"`
	DebugMode        bool     `mapstructure:"DEBUG_MODE"`
	ServerHost       string   `mapstructure:"SERVER_HOST"`
	ServerPort       string   `mapstructure:"SERVER_PORT"`
	PostgresHost     string   `mapstructure:"DATABASE_HOST"`
	PostgresPort     int      `mapstructure:"DATABASE_PORT"`
	PostgresUsername string   `mapstructure:"DATABASE_USERNAME"`
	PostgresPassword string   `mapstructure:"DATABASE_PASSWORD"`
	PostgresDb       string   `mapstructure:"DATABASE_NAME"`
	RedisHost        string   `mapstructure:"REDIS_HOST"`
	RedisPort        string   `mapstructure:"REDIS_PORT"`
	KAFKA_HOST       string   `mapstructure:"KAFKA_HOST"`
	MySqlHost        string   `mapstructure:"MYSQL_HOST"`
	MySqlPort        int      `mapstructure:"MYSQL_PORT"`
	MySqlUsername    string   `mapstructure:"MYSQL_USERNAME"`
	MySqlPassword    string   `mapstructure:"MYSQL_PASSWORD"`
	MySqlDb          string   `mapstructure:"MYSQL_NAME"`
	KafkaBrokers     []string `mapstructure:"KAFKA_BROKERS"`
	TempoGrpcHost    string   `mapstructure:"TEMPO_GRPC_HOST"`
	LoggerFileName   string   `mapstructure:"LOGGER_FILE_NAME"`
	LoggerPath       string   `mapstructure:"LOGGER_PATH"`
	LogLevel         string   `mapstructure:"LOG_LEVEL"`
	RulesBucket      string   `mapstructure:"RULES_BUCKET"`
	RulesTopic       string   `mapstructure:"RULES_TOPIC"`
	LeadsTopic       string   `mapstructure:"LEADS_TOPIC"`
}
