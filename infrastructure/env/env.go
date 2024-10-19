package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetAsString Returns the env value for the provided env variable name.
// It will use the defaultVal if the env variable is not set or empty.
// It also trims empty spaces from the begging and the end if there are any.
func GetAsString(envName string, defaultVal string) string {
	val, exists := os.LookupEnv(envName)

	if !exists {
		return defaultVal
	}

	val = strings.TrimSpace(val)

	if val == "" {
		return defaultVal
	}

	return val
}

// GetAsInt Returns the env value as int for the provided env variable name.
// It will use the defaultVal if the env variable is not set, empty or invalid integer.
func GetAsInt(envName string, defaultVal int) int {
	val, exists := os.LookupEnv(envName)

	if !exists {
		return defaultVal
	}

	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}

	return valInt
}

// GetAsBool Returns the env value as bool for the provided env variable name.
// It will use the defaultVal if the env variable is not set, empty or invalid bool.
// Accepted bool string formats are the same as the ones accepted by strconv.ParseBool.
func GetAsBool(envName string, defaultVal bool) bool {
	val, exists := os.LookupEnv(envName)

	if !exists {
		return defaultVal
	}

	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal
	}

	valBool, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return valBool
}

// GetAsFloat Returns the env value as float for the provided env variable name.
// It will use the defaultVal if the env variable is not set, empty or invalid float.
func GetAsFloat(envName string, defaultVal float64) float64 {
	val, exists := os.LookupEnv(envName)

	if !exists {
		return defaultVal
	}

	val = strings.TrimSpace(val)

	if val == "" {
		return defaultVal
	}

	valFloat, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return defaultVal
	}

	return valFloat
}

// GetAsDuration Returns the env value as bool for the provided env variable name.
// It will use the defaultVal if the env variable is not set, empty or invalid duration format.
// Accepted duration string formats are the same as the ones accepted by time.ParseDuration.
func GetAsDuration(envName string, defaultVal time.Duration) time.Duration {
	val, exists := os.LookupEnv(envName)

	if !exists {
		return defaultVal
	}

	val = strings.TrimSpace(val)

	if val == "" {
		return defaultVal
	}

	valDuration, err := time.ParseDuration(val)

	if err != nil {
		return defaultVal
	}

	return valDuration
}
