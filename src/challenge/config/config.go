/*======================Travel Audience Number Server===========================

	Author(s):
		Sugath Fernando

==============================================================================*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// configuration file name
const (
	ConfigFile string = "settings.json"
)

var (
	configLoaded bool = false
	App          AppConfig
)

// AppConfig struct that defines the port
type AppConfig struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

// init function that load configuration settings, and returns error if presents.
func init() {

	// check the config file is already loaded.
	if configLoaded {
		return
	}

	// read the configuration file.
	appConfigContents, err := ioutil.ReadFile(ConfigFile)

	// if error present, panic the error.
	if err != nil {
		log.Panic("Application configuration file unreadable: ", err)
	}
	appConfig := AppConfig{}
	err = json.Unmarshal(appConfigContents, &appConfig)
	if err != nil {
		log.Panic("Application configuration file unreadable: ", err)
	}
	App = appConfig

	// set confiLoaded variable to true.
	configLoaded = true
}
