package config

import (
	"fmt"
	"os"
	"strconv"
)

func MustEnvInt(key string) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("env %s doesn't exist", key))
	}
	r, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(r)
}

func MustEnvString(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("env %s doesn't exist", key))
	}
	return v
}

func EnvString(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}

func EnvInt(key string, defaultValue int) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	r, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(r)
}
