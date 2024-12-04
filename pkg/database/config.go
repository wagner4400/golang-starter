package database

type PGConfig struct {
	DBName          string `env:"PG_DBNAME"`
	User            string `env:"PG_USER"`
	Password        string `env:"PG_PASSWORD"`
	Host            string `env:"PG_HOST"`
	Port            int    `env:"PG_PORT,default=5432"`
	SSLMode         string `env:"PG_SSL_MODE"`
	FallbackAppName string `env:"PG_FALLBACK_APP_NAME"`
	ConnectTimeout  int    `env:"PG_CONNECT_TIMEOUT"`
	SSLClientCert   string `env:"PG_SSL_CERT"`
	SSLClientKey    string `env:"PG_SSL_KEY"`
	SSLRootCert     string `env:"PG_SSL_ROOT_CERT"`
}
