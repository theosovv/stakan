// Package config provides 12-factor env-based configuration helpers.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Get returns env value or def.
func Get(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

// MustGet returns env value or panics. Use only in main during bootstrap.
func MustGet(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("config: required env %s is not set", key))
	}
	return v
}

// Duration parses env as time.Duration or returns def.
func Duration(key string, def time.Duration) time.Duration {
	if v, ok := os.LookupEnv(key); ok {
		d, err := time.ParseDuration(v)
		if err == nil {
			return d
		}
	}
	return def
}

// Int parses env as int or returns def.
func Int(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		n, err := strconv.Atoi(v)
		if err == nil {
			return n
		}
	}
	return def
}
