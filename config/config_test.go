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
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"runtime"
	"testing"
)

func Test_getName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"WithFullPath",
			"/a/b/c",
			"c",
		}, {
			"WithExtPath",
			"/a/b/c.yaml",
			"c",
		}, {
			"WithSinglePath",
			"a",
			"a",
		}, {
			"WithPathInWin",
			"D:\\a\\b",
			"b",
		}, {
			"WithExtPathInWin",
			"D:\\a\\b\\c.yaml",
			"c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(tt.in); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateConfigPath(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{
			"PathNotExist",
			"/a/b/c",
			true,
		},
		{
			"PathIsDir",
			".",
			true,
		},
		{
			"PathExistAndIsFile",
			"./testdata/conf.yaml",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfigPath(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfigPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getExt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			"WithNoExt",
			"/a/b/c",
			"",
		},
		{
			"WithYamlExt",
			"/a/b/c.yaml",
			".yaml",
		},
		{
			"WithJsonExt",
			"/a/b/c.json",
			".json",
		},
		{
			"WithJpegExt",
			"/a/b/c.jpeg",
			".jpeg",
		},
		{
			"WithMultiExt",
			"/a/b/c.metric.yaml",
			".yaml",
		},
		{
			"WithOnlyExt",
			".bashrc",
			".bashrc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExt(tt.in); got != tt.want {
				t.Errorf("getExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPath(t *testing.T) {
	testsForWin := []struct {
		name string
		in   string
		want string
	}{
		{"WithFullPath",
			"/a/b/c",
			"\\a\\b",
		}, {
			"WithExtPath",
			"/a/b/c.yaml",
			"\\a\\b",
		}, {
			"WithSinglePath",
			"a",
			".",
		}, {
			"WithPathInWin",
			"D:\\a\\b",
			"D:\\a",
		}, {
			"WithExtPathInWin",
			"D:\\a\\b\\c.yaml",
			"D:\\a\\b",
		},
	}

	testsForUnix := []struct {
		name string
		in   string
		want string
	}{
		{"WithFullPath",
			"/a/b/c",
			"/a/b",
		}, {
			"WithExtPath",
			"/a/b/c.yaml",
			"/a/b",
		}, {
			"WithSinglePath",
			"a",
			".",
		},
	}

	if runtime.GOOS == "windows" {
		for _, tt := range testsForWin {
			t.Run(tt.name, func(t *testing.T) {
				if got := getPath(tt.in); got != tt.want {
					t.Errorf("getPath() = %v, want %v", got, tt.want)
				}
			})
		}
	} else {
		for _, tt := range testsForUnix {
			t.Run(tt.name, func(t *testing.T) {
				if got := getPath(tt.in); got != tt.want {
					t.Errorf("getPath() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{
			"ConfigFileNotFound",
			"/a/b/c/conf.yaml",
			true,
		},
		{
			"ConfigFileWithYamlFound",
			"./testdata/conf.yaml",
			false,
		},
		{
			"ConfigFileWithJsonFound",
			"./testdata/conf.ok.json",
			false,
		},
		{
			"ConfigFileReadFailedTest",
			"./testdata/conf.bad.json",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initConfig(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("initConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_config(t *testing.T) {
	Faker()
	assert.Equal(t, viper.GetString("svc_name"), "fake")
	assert.Equal(t, viper.GetString("nothing"), "")
	assert.Equal(t, viper.GetString("db.user"), "fake_user")
	assert.Equal(t, viper.GetString("db.pass"), "fake_pass")
	assert.Equal(t, viper.GetString("db.dbname"), "fake_db")
	assert.Equal(t, viper.GetInt("port"), 30030)
	assert.Equal(t, viper.GetString("log.level"), "debug")
	assert.Equal(t, viper.GetBool("log.multi_staging"), true)
	assert.Equal(t, viper.GetString("log.rotate.all_log_path"), "logs/all.log")
	assert.Equal(t, viper.GetString("log.rotate.info_log_path"), "logs/info.log")
	assert.Equal(t, viper.GetString("log.rotate.warn_log_path"), "logs/warn.log")
	assert.Equal(t, viper.GetInt("log.rotate.max_size"), 1024)
	assert.Equal(t, viper.GetInt("log.rotate.max_backups"), 5)
	assert.Equal(t, viper.GetInt("log.rotate.max_age"), 30)
	assert.Equal(t, viper.GetBool("log.rotate.compress"), false)
}
