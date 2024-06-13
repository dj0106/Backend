package Configuration

type ServerConfiguration struct {
	PRODUCTION  int
	SERVER_PORT string
	CLIENT_NAME string

	MYSQL_DB_NAME  string
	MYSQL_HOST     string
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_PORT     int
}
