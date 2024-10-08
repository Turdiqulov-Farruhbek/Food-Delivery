package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	// "github.com/spf13/cast"
)

type Config struct {
  HTTPPort string

  PostgresHost     string
  PostgresPort     int
  PostgresUser     string
  PostgresPassword string
  PostgresDatabase string
  LogPath          string
  KafkaHost        string
  KafkaPort        string
  NotificationHost string
  NotificationPort string


  DefaultOffset string
  DefaultLimit  string
}

func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":"))

  config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localst"))
  config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 542))
  config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "pogres"))
  config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
  config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "proycts"))
  config.LogPath = cast.ToString(getOrReturnDefaultValue("LOG_PATH", "logs/info.log"))
  config.KafkaHost = cast.ToString(getOrReturnDefaultValue("KAFKA_HOST", "q"))
  config.KafkaPort = cast.ToString(getOrReturnDefaultValue("KAFKA_PORT", "q"))
  config.NotificationHost = cast.ToString(getOrReturnDefaultValue("NOTIFICATION_HOST", "q"))
  config.NotificationPort = cast.ToString(getOrReturnDefaultValue("NOTIFICATION_PORT", "q"))

  config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))


  return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}