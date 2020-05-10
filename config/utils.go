package config

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(fmt.Errorf("Error: unable get os env [%s]", key))
}

func IsDev() bool {
	return GetEnv("GO_ENV") == "development"
}
