package config

import (
	"os"

	"github.com/cecep31/fiberapi/exception"
	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filnames ...string) Config {
	err := godotenv.Load(filnames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
