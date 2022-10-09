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

package utils

import (
	"fmt"
	"github.com/cauwulixuan/go-kit/http"
	"github.com/cauwulixuan/go-kit/log"
	"github.com/spf13/viper"
)

func TestLog() {
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.DPanic("DPanic")
	log.Panic("Panic")
	log.Fatal("Fatal")

	log.Slogger.Debug("Slogger debug")
	log.Slogger.Info("Slogger info")
	log.Slogger.Warn("Slogger warn")
	log.Slogger.Error("Slogger error")
	log.Slogger.DPanic("Slogger dpanic")
	log.Slogger.Panic("Slogger panic")
	log.Slogger.Fatal("Slogger fatal")
}

func TestConfig() {
	var name = viper.GetString("db.name")
	fmt.Println(name)
}

func TestHttpClient() {
	http.GetUrl("http://hello")
}
