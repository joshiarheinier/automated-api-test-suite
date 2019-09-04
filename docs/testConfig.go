package docs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

const (
	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v\n"
)

type TestConfiguration struct {
	ConfigName	string			`json:"config_name"`
	Hostname	string			`json:"hostname"`
	Port		string			`json:"port"`
	Schema		[]TestSchema	`json:"schema"`
}

type TestSchema struct {
	TestTitle		string				`json:"test_title"`
	Jobs			[]SubTestConfig		`json:"jobs"`
}

type SubTestConfig struct {
	TestName		string				`json:"test_name"`
	Url				string				`json:"url"`
	Method			string				`json:"method"`
	Params			map[string]string	`json:"params"`
	Header			map[string]string	`json:"header"`
	Body			json.RawMessage		`json:"body"`
	SaveKeys		[]string			`json:"save_keys"`
	Expected		ExpectedResult	`json:"expected"`
}

type ExpectedResult struct {
	Status	string				`json:"response_status"`
	Body	map[string]string	`json:"body"`
}

var TestConfigData  = &TestConfiguration{}

func init()  {
	_, filename, _, _ := runtime.Caller(0)
	testConfigFilePath := path.Join(path.Dir(filename), "./json/test/testConfig.json")
	testConfigFile, err := ioutil.ReadFile(testConfigFilePath)
	if err != nil {
		log.Fatalf(ErrFailedToDecodeConfigurationFile, err)
	}
	err = json.Unmarshal(testConfigFile, &TestConfigData)
	if err != nil {
		log.Fatalf(ErrFailedToDecodeConfigurationFile, err)
	}
}

func New() *TestConfiguration {
	return TestConfigData
}