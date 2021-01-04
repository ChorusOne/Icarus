package common

import "os"

func Getenv(name string, defaultValue string) string {
	value, exists := os.LookupEnv(name) 
	if !exists {
		value = defaultValue
	}
	return value
}
