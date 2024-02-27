package config

// Would be better to use github.com/joho/godotenv but as was required at the doc
// I'm trying to recreate that behavior in a simple way.

import "os"

// GetEnv retrieves the value of the environment variable with the given key.
// If the variable is not set, it returns the provided default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
