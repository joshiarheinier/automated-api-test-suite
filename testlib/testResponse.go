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

func (tr TestResponse) ParseBody() map[string]string {

	parsedBody := make(map[string]string)
	err := json.Unmarshal(tr.Body, &parsedBody)
	if err != nil {
		log.Fatalf("Failed to encode response body: %v\n", err)
	}
	return parsedBody
}