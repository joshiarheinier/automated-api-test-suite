package testlib

import (
	"encoding/json"
)

type TestResponse struct {
	StatusCode	int
	Body	[]byte
}


func (tr TestResponse) ParseBody() (map[string]interface{}, error) {

	parsedBody := make(map[string]interface{})
	err := json.Unmarshal(tr.Body, &parsedBody)
	if err != nil {
		return nil, err
	}
	return parsedBody, nil
}

func (tr TestResponse) RetrieveValueFromBody(key string, body interface{}) interface{} {
	switch body.(type) {
	case map[string]interface{}:
		if body.(map[string]interface{})[key] != nil {
			return body.(map[string]interface{})[key]
		} else {
			for _, v := range body.(map[string]interface{}) {
				return tr.RetrieveValueFromBody(key, v)
			}
		}
	case []interface{}:
		for _, v := range body.([]interface{}) {
			return tr.RetrieveValueFromBody(key, v)
		}
	}
	return nil
}