package params

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type EnvSuite struct {
	suite.Suite
	testEnvName string
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvSuite))
}

func (suite *EnvSuite) SetupTest() {
	suite.testEnvName = "TEST_ENV_VAR_12345"
	// Make sure the environment variable doesn't exist at the start of each test
	_ = os.Unsetenv(suite.testEnvName)
}

func (suite *EnvSuite) TearDownTest() {
	// Clean up after each test
	_ = os.Unsetenv(suite.testEnvName)
}

func (suite *EnvSuite) TestGetEnvAsString() {
	tests := []struct {
		name        string
		envValue    *string // nil means environment variable doesn't exist
		defaultVal  string
		want        string
		wantDefault bool
	}{
		{
			name:        "non-existent env var returns default",
			envValue:    nil,
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "empty env var returns default",
			envValue:    ptr(""),
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "whitespace env var returns default",
			envValue:    ptr("   "),
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "non-empty env var returns trimmed value",
			envValue:    ptr("  value  "),
			defaultVal:  "default",
			want:        "value",
			wantDefault: false,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				// Setup environment variable for this test case
				if scenario.envValue == nil {
					_ = os.Unsetenv(suite.testEnvName)
				} else {
					_ = os.Setenv(suite.testEnvName, *scenario.envValue)
				}

				// Call the function under test
				got, gotDefault := GetEnvAsString(suite.testEnvName, scenario.defaultVal)

				// Assert results
				suite.Equal(scenario.want, got, "GetEnvAsString() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetEnvAsString() defaultUsed")
			},
		)
	}
}

func (suite *EnvSuite) TestGetEnvAsInt() {
	tests := []struct {
		name        string
		envValue    *string // nil means environment variable doesn't exist
		defaultVal  int
		want        int
		wantDefault bool
	}{
		{
			name:        "non-existent env var returns default",
			envValue:    nil,
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "empty env var returns default",
			envValue:    ptr(""),
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "whitespace env var returns default",
			envValue:    ptr("   "),
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "valid integer env var returns parsed value",
			envValue:    ptr("123"),
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "valid integer with whitespace env var returns parsed value",
			envValue:    ptr("  123  "),
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "invalid integer env var returns default",
			envValue:    ptr("abc"),
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
	}

	for _, tt := range tests {
		suite.Run(
			tt.name, func() {
				// Setup environment variable for this test case
				if tt.envValue == nil {
					_ = os.Unsetenv(suite.testEnvName)
				} else {
					_ = os.Setenv(suite.testEnvName, *tt.envValue)
				}

				// Call the function under test
				got, gotDefault := GetEnvAsInt(suite.testEnvName, tt.defaultVal)

				// Assert results
				suite.Equal(tt.want, got, "GetEnvAsInt() value")
				suite.Equal(tt.wantDefault, gotDefault, "GetEnvAsInt() defaultUsed")
			},
		)
	}
}

func (suite *EnvSuite) TestGetEnvAsBool() {
	tests := []struct {
		name        string
		envValue    *string // nil means environment variable doesn't exist
		defaultVal  bool
		want        bool
		wantDefault bool
	}{
		{
			name:        "non-existent env var returns default",
			envValue:    nil,
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "empty env var returns default",
			envValue:    ptr(""),
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "whitespace env var returns default",
			envValue:    ptr("   "),
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "true returns true",
			envValue:    ptr("true"),
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "TRUE returns true",
			envValue:    ptr("TRUE"),
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "1 returns true",
			envValue:    ptr("1"),
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "false returns false",
			envValue:    ptr("false"),
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "FALSE returns false",
			envValue:    ptr("FALSE"),
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "0 returns false",
			envValue:    ptr("0"),
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "invalid boolean env var returns default",
			envValue:    ptr("not-a-bool"),
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				// Setup environment variable for this test case
				if scenario.envValue == nil {
					_ = os.Unsetenv(suite.testEnvName)
				} else {
					_ = os.Setenv(suite.testEnvName, *scenario.envValue)
				}

				// Call the function under test
				got, gotDefault := GetEnvAsBool(suite.testEnvName, scenario.defaultVal)

				// Assert results
				suite.Equal(scenario.want, got, "GetEnvAsBool() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetEnvAsBool() defaultUsed")
			},
		)
	}
}

func (suite *EnvSuite) TestGetEnvAsFloat() {
	tests := []struct {
		name        string
		envValue    *string // nil means environment variable doesn't exist
		defaultVal  float64
		want        float64
		wantDefault bool
	}{
		{
			name:        "non-existent env var returns default",
			envValue:    nil,
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "empty env var returns default",
			envValue:    ptr(""),
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "whitespace env var returns default",
			envValue:    ptr("   "),
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "valid float env var returns parsed value",
			envValue:    ptr("123.45"),
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "valid integer as float env var returns parsed value",
			envValue:    ptr("123"),
			defaultVal:  42.5,
			want:        123.0,
			wantDefault: false,
		},
		{
			name:        "valid float with whitespace env var returns parsed value",
			envValue:    ptr("  123.45  "),
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "invalid float env var returns default",
			envValue:    ptr("abc"),
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				// Setup environment variable for this test case
				if scenario.envValue == nil {
					_ = os.Unsetenv(suite.testEnvName)
				} else {
					_ = os.Setenv(suite.testEnvName, *scenario.envValue)
				}

				// Call the function under test
				got, gotDefault := GetEnvAsFloat(suite.testEnvName, scenario.defaultVal)

				// Assert results
				suite.Equal(scenario.want, got, "GetEnvAsFloat() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetEnvAsFloat() defaultUsed")
			},
		)
	}
}

func (suite *EnvSuite) TestGetEnvAsDuration() {
	tests := []struct {
		name        string
		envValue    *string // nil means environment variable doesn't exist
		defaultVal  time.Duration
		want        time.Duration
		wantDefault bool
	}{
		{
			name:        "non-existent env var returns default",
			envValue:    nil,
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "empty env var returns default",
			envValue:    ptr(""),
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "whitespace env var returns default",
			envValue:    ptr("   "),
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "valid duration env var returns parsed value",
			envValue:    ptr("10s"),
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "valid duration with whitespace env var returns parsed value",
			envValue:    ptr("  10s  "),
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "invalid duration env var returns default",
			envValue:    ptr("abc"),
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				// Setup environment variable for this test case
				if scenario.envValue == nil {
					_ = os.Unsetenv(suite.testEnvName)
				} else {
					_ = os.Setenv(suite.testEnvName, *scenario.envValue)
				}

				// Call the function under test
				got, gotDefault := GetEnvAsDuration(suite.testEnvName, scenario.defaultVal)

				// Assert results
				suite.Equal(scenario.want, got, "GetEnvAsDuration() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetEnvAsDuration() defaultUsed")
			},
		)
	}
}

// Helper function to create a pointer to a string
func ptr(s string) *string {
	return &s
}
