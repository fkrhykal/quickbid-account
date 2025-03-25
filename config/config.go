package config

import (
	"log"
	"os"
	"strconv"
)

func Configure[T any](defaultConfig T, configFuncs ...func(prevConfig T) T) T {
	for _, configFunc := range configFuncs {
		configFunc(defaultConfig)
	}
	return defaultConfig
}

func MustEnvString(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env error: %s didn't exist", key)
	}
	return v
}

func MustEnvInt(key string) int {
	v := MustEnvString(key)
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Fatalf("env error: failed to parse %s because %v", key, err)
	}
	return int(result)
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
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(result)
}
