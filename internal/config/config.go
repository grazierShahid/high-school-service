package config

type Config struct {
	BaseConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		BaseConfig: LoadBaseConfig(),
	}

	return cfg, nil
}

func (cfg *Config) IsDevelopment() bool {
	return cfg.App.AppEnv == "development"
}
