/*
Copyright 2022 The Inspur AIStation Group Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// Init initial config and watch config on change
func Init(cfgPath string) {
	if err := initConfig(cfgPath); err != nil {
		panic(fmt.Errorf("Init config failed, error: %v", err))
	}
	watchConfig()
}

func getExt(path string) string {
	return filepath.Ext(path)
}

func getPath(path string) string {
	return filepath.Dir(path)
}

func getName(path string) string {
	ext := getExt(path)
	return strings.Trim(filepath.Base(path), ext)
}

// initConfig initial config with adding config path,
// setting config name and setting config type.
// if configType not in one of "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"
// ConfigFileNotFoundError will be returned.
func initConfig(cfgPath string) error {
	viper.AddConfigPath("/etc/config/") // path to look for the config file in
	viper.AddConfigPath(".")
	viper.AddConfigPath(getPath(cfgPath))
	viper.SetConfigName(getName(cfgPath))

	cfgType := strings.Trim(getExt(cfgPath), ".")
	viper.SetConfigType(cfgType)

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("Config file not found.")
			return err.(viper.ConfigFileNotFoundError)
		} else {
			// Config file was found but another error was produced
			log.Printf("Error occurred while reading config file. error: %v\n.", err.Error())
			return err
		}
	}
	return nil
}

// watchConfig watch config file on change by using fsnotify module
func watchConfig() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: ", e.Name)
	})
	viper.WatchConfig()
}
