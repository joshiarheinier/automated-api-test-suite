package testlib

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type TestValidator struct {
	Response	TestResponse
}

func (tv TestValidator) ExpectResponseStatus(t *testing.T, status string) {
	s, _ := strconv.Atoi(status)
	assert.Equal(t, s, tv.Response.StatusCode)
}

func (tv TestValidator) ExpectBody(t *testing.T, body map[string]string)  {
	for k, v := range body {
		tv.ExpectStringEqual(t, v, tv.Response.GetValueFromBody(k))
	}
}

func (tv TestValidator) ExpectStringEqual(t *testing.T, expected string, actual string)  {
	assert.Equal(t, expected, actual)
}

func (tv TestValidator) ExpectNotNilAndSave(t *testing.T, key string) string {
	body := tv.Response.ParseBody()
	value := body[key]
	assert.NotEqual(t, "", value)
	return value
}
