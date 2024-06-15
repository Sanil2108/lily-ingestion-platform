package config

// ServerConfig is server specific config
type ServerConfig struct {
	Host string `json:"address"`
	Port int    `json:"port"`
}

type TemporalConfig struct {
	Host string `json:"address"`
	Port int    `json:"port"`
}

// Config is the definition for global config
type Config struct {
	Env           string       `json:"env"`
	Server        ServerConfig `json:"server"`
	Temporal 			TemporalConfig `json:"temporal"`
}

// NewConfig is construtor for the config
func NewConfig() *Config {
	return &Config{
		Env:           "dev",
		Server:        ServerConfig{Host: "127.0.0.1", Port: 4000},
		Temporal: 			TemporalConfig{Host: "127.0.0.1", Port: 7233},
	}
}
