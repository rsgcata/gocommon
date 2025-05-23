package params

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
	"time"
)

type URLSuite struct {
	suite.Suite
}

func TestURLTestSuite(t *testing.T) {
	suite.Run(t, new(URLSuite))
}

func (suite *URLSuite) TestNewQueryParamsFromUrl() {
	// Create a URL with query parameters
	u, err := url.Parse("https://example.com?key=value&int=123")
	suite.NoError(err)

	// Create QueryParams from URL
	qp := NewQueryParamsFromUrl(*u)

	suite.Equal("value", qp.params.Get("key"))
	suite.Equal("123", qp.params.Get("int"))
}

func (suite *URLSuite) TestQueryParamsGetAsString() {
	tests := []struct {
		name        string
		params      *QueryParams
		key         string
		defaultVal  string
		want        string
		wantDefault bool
	}{
		{
			name:        "nil QueryParams returns default",
			params:      nil,
			key:         "key",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "key not present returns default",
			params:      &QueryParams{params: url.Values{"otherkey": []string{"value"}}},
			key:         "key",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "empty value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{""}}},
			key:         "key",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "whitespace value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"   "}}},
			key:         "key",
			defaultVal:  "default",
			want:        "default",
			wantDefault: true,
		},
		{
			name:        "valid value returns trimmed value",
			params:      &QueryParams{params: url.Values{"key": []string{"  value  "}}},
			key:         "key",
			defaultVal:  "default",
			want:        "value",
			wantDefault: false,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := scenario.params.GetAsString(scenario.key, scenario.defaultVal)
				suite.Equal(scenario.want, got, "QueryParams.GetAsString() value")
				suite.Equal(
					scenario.wantDefault,
					gotDefault,
					"QueryParams.GetAsString() defaultUsed",
				)
			},
		)
	}
}

func (suite *URLSuite) TestQueryParamsGetAsInt() {
	tests := []struct {
		name        string
		params      *QueryParams
		key         string
		defaultVal  int
		want        int
		wantDefault bool
	}{
		{
			name:        "nil QueryParams returns default",
			params:      nil,
			key:         "key",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "key not present returns default",
			params:      &QueryParams{params: url.Values{"otherkey": []string{"123"}}},
			key:         "key",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "empty value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{""}}},
			key:         "key",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "whitespace value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"   "}}},
			key:         "key",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
		{
			name:        "valid integer returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"123"}}},
			key:         "key",
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "valid integer with whitespace returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"  123  "}}},
			key:         "key",
			defaultVal:  42,
			want:        123,
			wantDefault: false,
		},
		{
			name:        "invalid integer returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"abc"}}},
			key:         "key",
			defaultVal:  42,
			want:        42,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := scenario.params.GetAsInt(scenario.key, scenario.defaultVal)
				suite.Equal(scenario.want, got, "QueryParams.GetAsInt() value")
				suite.Equal(scenario.wantDefault, gotDefault, "QueryParams.GetAsInt() defaultUsed")
			},
		)
	}
}

func (suite *URLSuite) TestQueryParamsGetAsBool() {
	tests := []struct {
		name        string
		params      *QueryParams
		key         string
		defaultVal  bool
		want        bool
		wantDefault bool
	}{
		{
			name:        "nil QueryParams returns default",
			params:      nil,
			key:         "key",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "key not present returns default",
			params:      &QueryParams{params: url.Values{"otherkey": []string{"true"}}},
			key:         "key",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "empty value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{""}}},
			key:         "key",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "whitespace value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"   "}}},
			key:         "key",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
		{
			name:        "true returns true",
			params:      &QueryParams{params: url.Values{"key": []string{"true"}}},
			key:         "key",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "TRUE returns true",
			params:      &QueryParams{params: url.Values{"key": []string{"TRUE"}}},
			key:         "key",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "1 returns true",
			params:      &QueryParams{params: url.Values{"key": []string{"1"}}},
			key:         "key",
			defaultVal:  false,
			want:        true,
			wantDefault: false,
		},
		{
			name:        "false returns false",
			params:      &QueryParams{params: url.Values{"key": []string{"false"}}},
			key:         "key",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "FALSE returns false",
			params:      &QueryParams{params: url.Values{"key": []string{"FALSE"}}},
			key:         "key",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "0 returns false",
			params:      &QueryParams{params: url.Values{"key": []string{"0"}}},
			key:         "key",
			defaultVal:  true,
			want:        false,
			wantDefault: false,
		},
		{
			name:        "invalid boolean returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"not-a-bool"}}},
			key:         "key",
			defaultVal:  true,
			want:        true,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := scenario.params.GetAsBool(scenario.key, scenario.defaultVal)
				suite.Equal(scenario.want, got, "QueryParams.GetAsBool() value")
				suite.Equal(scenario.wantDefault, gotDefault, "QueryParams.GetAsBool() defaultUsed")
			},
		)
	}
}

func (suite *URLSuite) TestQueryParamsGetAsFloat() {
	tests := []struct {
		name        string
		params      *QueryParams
		key         string
		defaultVal  float64
		want        float64
		wantDefault bool
	}{
		{
			name:        "nil QueryParams returns default",
			params:      nil,
			key:         "key",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "key not present returns default",
			params:      &QueryParams{params: url.Values{"otherkey": []string{"123.45"}}},
			key:         "key",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "empty value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{""}}},
			key:         "key",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "whitespace value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"   "}}},
			key:         "key",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
		{
			name:        "valid float returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"123.45"}}},
			key:         "key",
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "valid integer as float returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"123"}}},
			key:         "key",
			defaultVal:  42.5,
			want:        123.0,
			wantDefault: false,
		},
		{
			name:        "valid float with whitespace returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"  123.45  "}}},
			key:         "key",
			defaultVal:  42.5,
			want:        123.45,
			wantDefault: false,
		},
		{
			name:        "invalid float returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"abc"}}},
			key:         "key",
			defaultVal:  42.5,
			want:        42.5,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := scenario.params.GetAsFloat(scenario.key, scenario.defaultVal)
				suite.Equal(scenario.want, got, "QueryParams.GetAsFloat() value")
				suite.Equal(
					scenario.wantDefault,
					gotDefault,
					"QueryParams.GetAsFloat() defaultUsed",
				)
			},
		)
	}
}

func (suite *URLSuite) TestQueryParamsGetAsDuration() {
	tests := []struct {
		name        string
		params      *QueryParams
		key         string
		defaultVal  time.Duration
		want        time.Duration
		wantDefault bool
	}{
		{
			name:        "nil QueryParams returns default",
			params:      nil,
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "key not present returns default",
			params:      &QueryParams{params: url.Values{"otherkey": []string{"10s"}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "empty value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{""}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "whitespace value returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"   "}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
		{
			name:        "valid duration returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"10s"}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "valid duration with whitespace returns parsed value",
			params:      &QueryParams{params: url.Values{"key": []string{"  10s  "}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        10 * time.Second,
			wantDefault: false,
		},
		{
			name:        "invalid duration returns default",
			params:      &QueryParams{params: url.Values{"key": []string{"abc"}}},
			key:         "key",
			defaultVal:  5 * time.Second,
			want:        5 * time.Second,
			wantDefault: true,
		},
	}

	for _, scenario := range tests {
		suite.Run(
			scenario.name, func() {
				got, gotDefault := scenario.params.GetAsDuration(scenario.key, scenario.defaultVal)
				suite.Equal(scenario.want, got, "QueryParams.GetAsDuration() value")
				suite.Equal(
					scenario.wantDefault,
					gotDefault,
					"QueryParams.GetAsDuration() defaultUsed",
				)
			},
		)
	}
}
