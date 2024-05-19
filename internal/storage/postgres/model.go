package postgres

const (
	HOST        = "DB_HOST"
	PORT        = "DB_PORT"
	USER        = "DB_USER"
	PASSWORD    = "DB_PASSWORD"
	NAME        = "DB_NAME"
	SCHEME_NAME = "DB_SHEME"
	SSL_MODE    = "DB_SSL_MODE"
	LOG_LEVEL   = "DB_LOG_LEVEL"
)

type ConfigDB struct {
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	SchemeName string
	SSLMode    string
	LogLevel   string
}
