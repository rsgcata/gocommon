package params

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type StrconvSuite struct {
	suite.Suite
}

func TestStrconvTestSuite(t *testing.T) {
	suite.Run(t, new(StrconvSuite))
}

func (suite *StrconvSuite) TestGetAsStr() {
	tests := []struct {
		name        string
		input       string
		defaultVal  string
		want        string
		wantDefault bool
	}{
		{
			name:        "empty string returns default",
			input:       "",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "whitespace returns default",
			input:       "   ",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "non-empty string returns trimmed value",
			input:       "  value  ",
			defaultVal:  "default",
			want:        "value",
			wantDefault: false,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := GetAsString(scenario.input, scenario.defaultVal)
				suite.Equal(scenario.want, got, "GetAsString() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetAsString() defaultUsed")
			},
		)
	}
}

func (suite *StrconvSuite) TestGetAsInt() {
	tests := []struct {
		name        string
		input       string
		defaultVal  int
		want        int
		wantDefault bool
	}{
		{
			name:        "empty string returns default",
			input:       "",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "whitespace returns default",
			input:       "   ",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "valid integer returns parsed value",
			input:       "123",
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "valid integer with whitespace returns parsed value",
			input:       "  123  ",
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "invalid integer returns default",
			input:       "abc",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "float string returns default",
			input:       "123.45",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := GetAsInt(scenario.input, scenario.defaultVal)
				suite.Equal(scenario.want, got, "GetAsInt() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetAsInt() defaultUsed")
			},
		)
	}
}

func (suite *StrconvSuite) TestGetAsBool() {
	tests := []struct {
		name        string
		input       string
		defaultVal  bool
		want        bool
		wantDefault bool
	}{
		{
			name:        "empty string returns default",
			input:       "",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "whitespace returns default",
			input:       "   ",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "true returns true",
			input:       "true",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "TRUE returns true",
			input:       "TRUE",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "True returns true",
			input:       "True",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "1 returns true",
			input:       "1",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "false returns false",
			input:       "false",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "FALSE returns false",
			input:       "FALSE",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "False returns false",
			input:       "False",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "0 returns false",
			input:       "0",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "invalid boolean returns default",
			input:       "not-a-bool",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := GetAsBool(scenario.input, scenario.defaultVal)
				suite.Equal(scenario.want, got, "GetAsBool() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetAsBool() defaultUsed")
			},
		)
	}
}

func (suite *StrconvSuite) TestGetAsFloat() {
	tests := []struct {
		name        string
		input       string
		defaultVal  float64
		want        float64
		wantDefault bool
	}{
		{
			name:        "empty string returns default",
			input:       "",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "whitespace returns default",
			input:       "   ",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "valid float returns parsed value",
			input:       "123.45",
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "valid integer as float returns parsed value",
			input:       "123",
			defaultVal:  42.5,
			want:        123.0,
			wantDefault: false,
		},
		{
			name:        "valid float with whitespace returns parsed value",
			input:       "  123.45  ",
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "invalid float returns default",
			input:       "abc",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := GetAsFloat(scenario.input, scenario.defaultVal)
				suite.Equal(scenario.want, got, "GetAsFloat() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetAsFloat() defaultUsed")
			},
		)
	}
}

func (suite *StrconvSuite) TestGetAsDuration() {
	tests := []struct {
		name        string
		input       string
		defaultVal  time.Duration
		want        time.Duration
		wantDefault bool
	}{
		{
			name:        "empty string returns default",
			input:       "",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "whitespace returns default",
			input:       "   ",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "valid duration returns parsed value",
			input:       "10s",
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "valid duration with whitespace returns parsed value",
			input:       "  10s  ",
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "invalid duration returns default",
			input:       "abc",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := GetAsDuration(scenario.input, scenario.defaultVal)
				suite.Equal(scenario.want, got, "GetAsDuration() value")
				suite.Equal(scenario.wantDefault, gotDefault, "GetAsDuration() defaultUsed")
			},
		)
	}
}
