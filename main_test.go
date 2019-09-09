package main

import (
	"bytes"
	"encoding/json"
	"github.com/joshia/automated-api-test-suite/controllers"
	"github.com/joshia/automated-api-test-suite/docs"
	"github.com/joshia/automated-api-test-suite/testlib"
	"testing"
)


func TestAPIFlow(t *testing.T)  {
	ex := testlib.Initiate()
	controllers.SetUpRouter(ex.R)
	testConfig := docs.New()
	for _, schema := range testConfig.Schema {
		storedData := make(map[string]interface{})
		for _, job := range schema.Jobs {
			tmp, _ := job.Body.MarshalJSON()
			stringBody := new(bytes.Buffer)
			json.Compact(stringBody, tmp)
			byteBody := stringBody.Bytes()
			if len(storedData) != 0 {
				ex.CheckHeader(job.Header, storedData)
				ex.CheckHeader(job.Params, storedData)
				byteBody = ex.CheckBody(byteBody, storedData)
			}
			job.Url = ex.SetParams(job.Url, job.Params)
			var validator testlib.TestValidator
			for i := 0; i < job.Try; i++ {
				req, _ := ex.MakeRequest(job.Method, job.Url, byteBody)
				ex.SetHeader(req, job.Header)
				validator = ex.Execute(req)
				ex.Flush()
			}
			validator.ExpectResponseStatus(t, job.Expected.Status)
			validator.ExpectBody(t, job.Expected.Body)
			if len(job.SaveKeys) != 0 {
				for i := range job.SaveKeys {
					storedData[job.SaveKeys[i]] = validator.ExpectNotNilAndSave(t, job.SaveKeys[i])
				}
			}
		}
	}
}
