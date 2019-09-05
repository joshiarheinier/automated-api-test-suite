package testlib

import (
	"encoding/json"
	"log"
	"strings"
)

type TestResponse struct {
	StatusCode	int
	Body	[]byte
}

//For GetValueFromBody, only string values are retrievable for now
func (tr TestResponse) GetValueFromBody(key string) string {
	res := ""
	tmp := string(tr.Body)
	idx := strings.Index(tmp, key) + len(key) + 3
	for string(tmp[idx]) != `"` {
		res += string(tmp[idx])
		idx += 1
	}
	return res
}

//For ParseBody, only 1 level JSON body with any key:value pairs with no nested is parsable
func (tr TestResponse) ParseBody() map[string]string {
	parsedBody := make(map[string]string)
	err := json.Unmarshal(tr.Body, &parsedBody)
	if err != nil {
		log.Fatalf("Failed to encode response body: %v\n", err)
	}
	return parsedBody
}