# Automated API Test Suite
This Go-written API level test suite is an automated test where a list of test configurations in JSON format will be scanned automatically through single script, using [Gin as the framework](https://github.com/gin-gonic/gin).
## How to use
  1. Run the server: `go run main.go` (The port can be changed on the main.go file)
  2. Run the test: `go test .` (It will execute the test script located on the same level as main.go, in this case is main_test.go)
## Main Components
### Executor, Validator, and Response Handler
TestExecutor is a struct where request is created. It manages the URL, headers, body, and other request's contents. TestResponse is a container of the response retrieved, storing status code and response body (by default, any contents can be added if necessary). TestValidator is a module that check the response values whether those values match the expected ones.
### JSON Configuration
JSON file named `testConfig.json` is the script where all test cases are placed, stored in `/docs/json/test/`. All test cases must have the same format with each other, and each test case can have 1 or more jobs, in case there is a job that needs a values that must be retrieved by another job.
```
{
  "config_name": "",
  "hostname": "",
  "port": "",
  "schema": [
    {
      "test_title": "",
      "jobs": [
        {
          "test_name": "",
          "try": 0,
          "url": "",
          "method": "",
          "params": {},
          "header": {},
          "body": {},
          "save_keys": [],
          "expected": {
            "response_status": "",
            "body": {}
          }
        }
      ]
    }
  ]
}
```
The template above can also be found in `/docs/json/test/template.json`.
