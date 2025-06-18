package config

type AppConfig struct {
	AppName string
	AppEnv  string
	AppURL  string
	AppPort string
}
type DbConfig struct {
	HOST     string
	PORT     string
	DATABASE string
	USERNAME string
	PASSWORD string
	NAME     string
}

type BaseConfig struct {
	App AppConfig
	DB  DbConfig
}

func LoadBaseConfig() BaseConfig {
	return BaseConfig{
		App: AppConfig{
			AppName: GetEnv("APP_NAME", "High School Management"),
			AppEnv:  GetEnv("APP_ENV", "development"),
			AppURL:  GetEnv("APP_URL", "localhost"),
			AppPort: GetEnv("APP_PORT", "8080"),
		},

		DB: DbConfig{
			HOST:     GetEnv("DB_HOST", "localhost"),
			PORT:     GetEnv("DB_PORT", "5432"),
			USERNAME: GetEnv("DB_USERNAME", "postgres"),
			PASSWORD: GetEnv("DB_PASSWORD", "postgres"),
			NAME:     GetEnv("DB_NAME", "postgres"),
		},
	}
}
