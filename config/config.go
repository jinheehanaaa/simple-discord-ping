// Reads config.json

package config

import (
	"encoding/json" // to use unmarshalling function for converting json into struct
	"fmt"
	"io/ioutil" // ReadFile
)

// Create variables
var (
	Token     string
	BotPrefix string

	config *configStruct
)

// Create your own data type [Variable, Type, json]
type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// Read Config.json file
func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error)
		return err
	}

	fmt.Println(string(file))

	// bring the values into config of type configStruct
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Values in each variable is no longer json value but it is type that can be processed by golang
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
