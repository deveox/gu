package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadDotEnv loads a .env file.
func LoadDotEnv(filePath string) error {
	return godotenv.Load(filePath)
}

// Get returns the value of an environment variable as a string.
func Get(key, byDefault string) string {
	r := os.Getenv(key)
	if r == "" {
		return byDefault
	}
	return r
}

// GetInt returns the value of an environment variable as an int.
//
// If the value is not an int, it will return the default value.
func GetInt(v string, byDefault int) int {
	r := os.Getenv(v)
	if r == "" {
		return byDefault
	}
	n, _ := strconv.Atoi(r)
	if n == 0 {
		return byDefault
	}
	return n

}

// GetUint8 returns the value of an environment variable as an uint8.
//
// If the value is not an uint8, it will return the default value
func GetUint8(v string, byDefault uint8) uint8 {
	r := os.Getenv(v)
	if r == "" {
		return byDefault
	}
	n, _ := strconv.ParseUint(r, 10, 32)
	if n == 0 {
		return byDefault
	}
	return uint8(n)
}

// GetUint64 returns the value of an environment variable as an uint64.
//
// If the value is not an uint64, it will return the default value
func GetUint64(v string, byDefault uint64) uint64 {
	r := os.Getenv(v)
	if r == "" {
		return byDefault
	}
	n, _ := strconv.ParseUint(r, 10, 64)
	if n == 0 {
		return byDefault
	}
	return n
}

// GetInt64 returns the value of an environment variable as an int64.
//
// If the value is not an int64, it will return the default value
func GetInt64(v string, byDefault int64) int64 {
	r := os.Getenv(v)
	if r == "" {
		return byDefault
	}
	n, _ := strconv.ParseInt(r, 10, 64)
	if n == 0 {
		return byDefault
	}
	return n
}

// GetBool returns the value of an environment variable as a bool.
//
// If the value is not a bool, it will return the default value
func GetBool(v string, byDefault bool) bool {
	r := os.Getenv(v)
	switch r {
	case "true":
		return true
	case "false":
		return false
	default:
		return byDefault
	}
}
