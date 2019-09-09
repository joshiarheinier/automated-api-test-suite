package testlib

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
)

type TestExecutor struct {
	R	*gin.Engine
	W	*httptest.ResponseRecorder
}


func Initiate() TestExecutor {
	executor := TestExecutor{
		R: gin.Default(),
		W: httptest.NewRecorder(),
	}
	return executor
}

func (te *TestExecutor) Flush()  {
	te.W = httptest.NewRecorder()
}

func (te TestExecutor) CheckHeader(header map[string]string, storedData map[string]interface{}) {
	for k, v := range header {
		values := strings.Split(v, "%")
		if len(values) > 1 {
			for i, e := range values {
				if i % 2 != 0 {
					values[i] = storedData[e].(string)
				}
			}
			newValue := ""
			for _, e := range values {
				newValue += e
			}
			header[k] = newValue
		}
	}
}

func (te TestExecutor) CheckBody(body []byte, storedData map[string]interface{}) []byte {
	strBody := string(body)
	values := strings.Split(strBody, "%")
	if len(values) > 1 {
		for i, e := range values {
			if i % 2 != 0 {
				values[i] = storedData[e].(string)
			}
		}
		newValue := ""
		for _, e := range values {
			newValue += e
		}
		return []byte(newValue)
	}
	return body
}

func (te TestExecutor) SetHeader(req *http.Request, header map[string]string) {
	for k, v := range header {
		req.Header.Add(k, v)
	}
}

func (te TestExecutor) SetParams(url string, params map[string]string) string {
	url += "?"
	i := 0
	for k, v := range params {
		url = url + k + "=" + v
		if i != len(params) - 1 {
			url += "&"
		}
		i += 1
	}
	return url
}

func (te TestExecutor) MakeRequest(method string, url string, body []byte) (*http.Request, error) {
	return http.NewRequest(method, url, bytes.NewReader(body))
}

func (te TestExecutor) Execute(req *http.Request) TestValidator {
	te.R.ServeHTTP(te.W, req)
	res := TestResponse{
		StatusCode: te.W.Code,
		Body:       te.W.Body.Bytes(),
	}
	return TestValidator{Response:res}
}