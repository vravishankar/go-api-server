package config

// Config - Definition
type Config struct {
	DB *DBConfig
}

// DBConfig - Definition
type DBConfig struct {
	Cluster  string
	Keyspace string
}

// GetConfig - Method Definition
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Cluster:  "127.0.0.1",
			Keyspace: "payments",
		},
	}
}
