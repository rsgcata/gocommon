package params

import (
	"os"
	"time"
)

// GetEnvAsString retrieves an environment variable and converts it to a string.
// It returns the string value of the environment variable and false if the variable exists and is
// not empty after trimming. If the environment variable doesn't exist or is empty/whitespace,
// it returns the defaultVal and true.
//
// Parameters:
//   - envName: The name of the environment variable to retrieve
//   - defaultVal: The default value to return if the environment variable doesn't exist or is invalid
//
// Returns:
//   - parsedVal: The parsed string value or defaultVal if environment variable doesn't exist or is
//     invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetEnvAsString(envName string, defaultVal string) (parsedVal string, defaultUsed bool) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, true
	}
	return GetAsString(val, defaultVal)
}

// GetEnvAsInt retrieves an environment variable and converts it to an integer.
// It returns the integer value of the environment variable and false if the variable exists and
// contains a valid integer. If the environment variable doesn't exist or cannot be parsed as an
// integer, it returns the defaultVal and true.
//
// Parameters:
//   - envName: The name of the environment variable to retrieve
//   - defaultVal: The default value to return if the environment variable doesn't exist or is invalid
//
// Returns:
//   - parsedVal: The parsed integer value or defaultVal if environment variable doesn't exist or
//     is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetEnvAsInt(envName string, defaultVal int) (parsedVal int, defaultUsed bool) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, true
	}
	return GetAsInt(val, defaultVal)
}

// GetEnvAsBool retrieves an environment variable and converts it to a boolean.
// It returns the boolean value of the environment variable and false if the variable exists and
// contains a valid boolean representation. Valid boolean values include:
// "true", "TRUE", "True", "1", "false", "FALSE", "False", "0". If the environment variable doesn't
// exist or cannot be parsed as a boolean, it returns the defaultVal and true.
//
// Parameters:
//   - envName: The name of the environment variable to retrieve
//   - defaultVal: The default value to return if the environment variable doesn't exist or is
//     invalid
//
// Returns:
//   - parsedVal: The parsed boolean value or defaultVal if environment variable doesn't exist or
//     is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetEnvAsBool(envName string, defaultVal bool) (parsedVal bool, defaultUsed bool) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, true
	}
	return GetAsBool(val, defaultVal)
}

// GetEnvAsFloat retrieves an environment variable and converts it to a float64.
// It returns the float64 value of the environment variable and false if the variable exists and
// contains a valid floating-point number. If the environment variable doesn't exist or cannot be
// parsed as a float64, it returns the defaultVal and true.
//
// Parameters:
//   - envName: The name of the environment variable to retrieve
//   - defaultVal: The default value to return if the environment variable doesn't exist or is
//     invalid
//
// Returns:
//   - parsedVal: The parsed float64 value or defaultVal if environment variable doesn't exist or
//     is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetEnvAsFloat(envName string, defaultVal float64) (parsedVal float64, defaultUsed bool) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, true
	}
	return GetAsFloat(val, defaultVal)
}

// GetEnvAsDuration retrieves an environment variable and converts it to a time.Duration.
// It returns the duration value of the environment variable and false if the variable exists and
// contains a valid duration string. Valid duration strings are those accepted by
// time.ParseDuration, such as "300ms", "1.5h", "2h45m". If the environment variable doesn't
// exist or cannot be parsed as a duration, it returns the defaultVal and true.
//
// Parameters:
//   - envName: The name of the environment variable to retrieve
//   - defaultVal: The default value to return if the environment variable doesn't exist or is
//     invalid
//
// Returns:
//   - parsedVal: The parsed duration value or defaultVal if environment variable doesn't exist or
//     is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetEnvAsDuration(envName string, defaultVal time.Duration) (
	parsedVal time.Duration,
	defaultUsed bool,
) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, true
	}
	return GetAsDuration(val, defaultVal)
}
