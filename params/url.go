package params

import (
	"net/url"
	"time"
)

type QueryParams struct {
	params url.Values
}

func NewQueryParamsFromUrl(url url.URL) *QueryParams {
	return &QueryParams{params: url.Query()}
}

// GetAsString is a helper function that retrieves first value from url.Values[key] and converts it
// to a string. It returns the string value and false if the key exists and the value is not
// empty after trimming. If the key doesn't exist or the value is empty/whitespace,
// it returns the defaultVal and true.
//
// Parameters:
//   - key: The key to look up in the url.Values
//   - defaultVal: The default value to return if the key doesn't exist or the value is invalid
//
// Returns:
//   - parsedVal: The parsed string value or defaultVal if key doesn't exist or value is invalid
//   - defaultUsed: True if defaultVal was used, false otherwise
func (v *QueryParams) GetAsString(key string, defaultVal string) (
	parsedVal string,
	defaultUsed bool,
) {
	if v == nil || !v.params.Has(key) {
		return defaultVal, true
	}

	return GetAsString(v.params.Get(key), defaultVal)
}

func (v *QueryParams) GetAsInt(key string, defaultVal int) (
	parsedVal int,
	defaultUsed bool,
) {
	if v == nil || !v.params.Has(key) {
		return defaultVal, true
	}

	return GetAsInt(v.params.Get(key), defaultVal)
}

func (v *QueryParams) GetAsBool(key string, defaultVal bool) (
	parsedVal bool,
	defaultUsed bool,
) {
	if v == nil || !v.params.Has(key) {
		return defaultVal, true
	}

	return GetAsBool(v.params.Get(key), defaultVal)
}

func (v *QueryParams) GetAsFloat(key string, defaultVal float64) (
	parsedVal float64,
	defaultUsed bool,
) {
	if v == nil || !v.params.Has(key) {
		return defaultVal, true
	}

	return GetAsFloat(v.params.Get(key), defaultVal)
}

func (v *QueryParams) GetAsDuration(key string, defaultVal time.Duration) (
	parsedVal time.Duration,
	defaultUsed bool,
) {
	if v == nil || !v.params.Has(key) {
		return defaultVal, true
	}

	return GetAsDuration(v.params.Get(key), defaultVal)
}
