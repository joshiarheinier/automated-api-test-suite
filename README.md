# Automated API Test Suite
This Go-written test suite is an automated test where a list of test configurations in JSON format will be scanned automatically through single script, using Gin as the framework.
How to use:
  1. Run the server: `go run main.go` (The port can be changed on the main.go file)
  2. Run the test: `go test .` (It will execute the test script located on the same level as main.go, in this case is main_test.go)
