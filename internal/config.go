package internal

type Config struct {
	Port     string
	DbConfig DbConfig
}

type DbConfig struct {
	Address            string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}

func CreateConfig() Config {
	return Config{
		Port: ":8080",
		DbConfig: DbConfig{
			Address:            "postgres://admin:admin@localhost:54321/goad?sslmode=disable",
			MaxOpenConnections: 30,
			MaxIdleConnections: 30,
			MaxIdleTime:        "15m",
		},
	}

}
