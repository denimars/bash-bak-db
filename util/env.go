package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadEnv() (Config, bool) {
	location := location()
	fmt.Println(location)
	var config Config
	data, err := os.ReadFile(fmt.Sprintf("%v/%v", location, "config.yaml"))
	if err != nil {
		fmt.Println(err)
		return config, false
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return config, false
	}
	return config, true

}
