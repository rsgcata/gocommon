package env

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type EnvSuite struct {
	suite.Suite
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvSuite))
}

func (suite *EnvSuite) TestItCanGetEnvAsString() {
	scenarios := map[string]struct {
		envVal      string
		defaultVal  string
		expectedVal string
	}{
		"empty env val":       {"", "test 123", "test 123"},
		"empty space env val": {"    ", "test 123", "test 123"},
		"not empty env val":   {"not empty 123", "test 123", "not empty 123"},
	}

	envName := "TEST_1234567"
	val := GetAsString(envName, "default 123")
	suite.Assert().Equal("default 123", val)

	for name, scenario := range scenarios {
		suite.T().Setenv(envName, scenario.envVal)
		val := GetAsString(envName, scenario.defaultVal)
		suite.Assert().Equal(scenario.expectedVal, val, "Failed scenario %s", name)
	}
}

func (suite *EnvSuite) TestItCanGetEnvAsInt() {
	scenarios := map[string]struct {
		envVal      string
		defaultVal  int
		expectedVal int
	}{
		"empty env val":       {"", 123, 123},
		"empty space env val": {"    ", 1234, 1234},
		"not empty env val":   {"12345", 123, 12345},
	}

	envName := "TEST_1234567"
	val := GetAsInt(envName, 123)
	suite.Assert().Equal(123, val)

	for name, scenario := range scenarios {
		suite.T().Setenv(envName, scenario.envVal)
		val := GetAsInt(envName, scenario.defaultVal)
		suite.Assert().Equal(scenario.expectedVal, val, "Failed scenario %s", name)
	}
}

func (suite *EnvSuite) TestItCanGetEnvAsBool() {
	scenarios := map[string]struct {
		envVal      string
		defaultVal  bool
		expectedVal bool
	}{
		"empty env val":           {"", true, true},
		"empty space env val":     {"    ", false, false},
		"not empty env val 1":     {"1", false, true},
		"not empty env val 0":     {"0", true, false},
		"not empty env val true":  {"true", false, true},
		"not empty env val false": {"false", true, false},
	}

	envName := "TEST_1234567"
	val := GetAsBool(envName, true)
	suite.Assert().Equal(true, val)

	for name, scenario := range scenarios {
		suite.T().Setenv(envName, scenario.envVal)
		val := GetAsBool(envName, scenario.defaultVal)
		suite.Assert().Equal(scenario.expectedVal, val, "Failed scenario %s", name)
	}
}
