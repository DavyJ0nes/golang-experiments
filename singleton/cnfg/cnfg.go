// Package cnfg provides simple config file parsing into map
package cnfg

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
)

type config struct {
	ConfigMap map[string]string
}

type configPair struct {
	Key   string
	Value string
}

// A reference to the singleton
var cnfg config

// MustLoad is used to instanciate the Config variables. It panics if there is an error rather than returning an error.
// This is because it MustLoad ;)
func MustLoad(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	configs, err := readConfigFile(file)
	if err != nil {
		panic(err.Error())
	}

	cnfg = config{
		ConfigMap: make(map[string]string),
	}

	for _, c := range configs {
		cnfg.ConfigMap[c.Key] = c.Value
	}
}

// Get is used to get the value of a Config Value
// In same vain as os.Env will return empty string if not found rather than erroring.
func Get(key string) string {
	return cnfg.ConfigMap[key]
}

// readConfigFile parses the contents of a file as yaml.
func readConfigFile(file *os.File) (output []configPair, err error) {
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return output, errors.Wrap(err, "Problem parsing file")
	}

	cm := make(map[string]string)

	err = yaml.Unmarshal(body, &cm)
	if err != nil {
		return output, errors.Wrap(err, "Problem unmarshalling file")
	}

	for k, v := range cm {
		output = append(output, configPair{k, v})
	}

	return output, nil
}
