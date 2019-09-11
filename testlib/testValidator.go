package testlib

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

type TestValidator struct {
	Response	TestResponse
}

func (tv TestValidator) ExpectResponseStatus(t *testing.T, status string) {
	s := strconv.Itoa(tv.Response.StatusCode)
	if string(status[0]) == "!" {
		tv.ExpectValueNotEqual(t, status[1:], s)
	} else {
		tv.ExpectValueEqual(t, status, s)
	}
}

func (tv TestValidator) ExpectBody(t *testing.T, body map[string]interface{})  {
	for k, v := range body {
		val := tv.Response.RetrieveValueFromBody(k, body)
		if val != nil {
			assert.IsType(t, v, val)
			tv.ExpectValueEqual(t, v, val)
		}
	}
}

func (tv TestValidator) ExpectValueEqual(t *testing.T, expected interface{}, actual interface{})  {
	assert.Equal(t, expected, actual)
}

func (tv TestValidator) ExpectValueNotEqual(t *testing.T, expected interface{}, actual interface{})  {
	assert.NotEqual(t, expected, actual)
}

func (tv TestValidator) ExpectNotNilAndSave(t *testing.T, key string) interface{} {
	body, err := tv.Response.ParseBody()
	if err != nil {
		log.Fatalf("Failed to encode response body: %v\n", err)
	}
	value := tv.Response.RetrieveValueFromBody(key, body)
	assert.NotEqual(t, "", value)
	return value
}
