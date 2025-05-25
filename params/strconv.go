package params

import (
	"strconv"
	"strings"
	"time"
)

// GetAsString converts the input string to a trimmed string value.
// It returns the trimmed input string and false if the input is not empty after trimming.
// If the input is empty or contains only whitespace, it returns the defaultVal and true.
//
// Parameters:
//   - val: The input string to be converted
//   - defaultVal: The default value to return if the input is empty or whitespace
//
// Returns:
//   - parsedVal: The parsed string value or defaultVal if input is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetAsString(val string, defaultVal string) (parsedVal string, defaultUsed bool) {
	parsedVal = strings.TrimSpace(val)

	if parsedVal == "" {
		return defaultVal, true
	}

	return parsedVal, false
}

// GetAsInt converts the input string to an integer value.
// It returns the parsed integer and false if the input is a valid integer.
// If the input is empty, contains only whitespace, or is not a valid integer,
// it returns the defaultVal and true.
//
// Parameters:
//   - val: The input string to be converted to an integer
//   - defaultVal: The default value to return if the input is invalid
//
// Returns:
//   - parsedVal: The parsed integer value or defaultVal if input is invalid
//   - defaultUsed: True if defaultVal was used due to parsing error, false otherwise
func GetAsInt(val string, defaultVal int) (parsedVal int, defaultUsed bool) {
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal, true
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal, true
	}

	return valInt, false
}

// GetAsBool converts the input string to a boolean value.
// It returns the parsed boolean and false if the input is a valid boolean representation.
// Valid boolean values include: "true", "TRUE", "True", "1", "false", "FALSE", "False", "0".
// If the input is empty, contains only whitespace, or is not a valid boolean,
// it returns the defaultVal and true.
//
// Parameters:
//   - val: The input string to be converted to a boolean
//   - defaultVal: The default value to return if the input is invalid
//
// Returns:
//   - parsedVal: The parsed boolean value or defaultVal if input is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetAsBool(val string, defaultVal bool) (parsedVal bool, defaultUsed bool) {
	val = strings.TrimSpace(val)

	if val == "" {
		return defaultVal, true
	}

	valBool, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal, true
	}

	return valBool, false
}

// GetAsFloat converts the input string to a float64 value.
// It returns the parsed float64 and false if the input is a valid floating-point number.
// If the input is empty, contains only whitespace, or is not a valid floating-point number,
// it returns the defaultVal and true.
//
// Parameters:
//   - val: The input string to be converted to a float64
//   - defaultVal: The default value to return if the input is invalid
//
// Returns:
//   - parsedVal: The parsed float64 value or defaultVal if input is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetAsFloat(val string, defaultVal float64) (parsedVal float64, defaultUsed bool) {
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal, true
	}

	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultVal, true
	}

	return valFloat, false
}

// GetAsDuration converts the input string to a time.Duration value.
// It returns the parsed duration and false if the input is a valid duration string.
// If the input is empty, contains only whitespace, or is not a valid duration format,
// it returns the defaultVal and true.
//
// Valid duration strings are those accepted by time.ParseDuration, such as "300ms", "1.5h", "2h45m".
//
// Parameters:
//   - val: The input string to be converted to a time.Duration
//   - defaultVal: The default value to return if the input is invalid
//
// Returns:
//   - parsedVal: The parsed time.Duration value or defaultVal if input is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func GetAsDuration(val string, defaultVal time.Duration) (
	parsedVal time.Duration,
	defaultUsed bool,
) {
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal, true
	}

	valDuration, err := time.ParseDuration(val)
	if err != nil {
		return defaultVal, true
	}

	return valDuration, false
}

type RawVal string

func (rawVal RawVal) GetAsString(defaultVal string) (string, bool) {
	return GetAsString(string(rawVal), defaultVal)
}

func (rawVal RawVal) GetAsInt(defaultVal int) (int, bool) {
	return GetAsInt(string(rawVal), defaultVal)
}

func (rawVal RawVal) GetAsBool(defaultVal bool) (bool, bool) {
	return GetAsBool(string(rawVal), defaultVal)
}

func (rawVal RawVal) GetAsFloat(defaultVal float64) (float64, bool) {
	return GetAsFloat(string(rawVal), defaultVal)
}

func (rawVal RawVal) GetAsDuration(defaultVal time.Duration) (time.Duration, bool) {
	return GetAsDuration(string(rawVal), defaultVal)
}
