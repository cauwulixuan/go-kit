/*
 * Copyright 2022 The Inspur AIStation Group Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Note: the example only works with the code within the same release/branch.

package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// Retry strategies:
//- No BackOff
//- Linear BackOff
//- LinearRandom BackOff
//- Exponential Backoff
//- ExponentialRandom BackOff

var (
	Client = resty.New()
)

func init() {
	Init()
}

func Init() {
	SetRetry()
	SetTimeout()
}

func SetRetry() {
	Client.
		SetRetryCount(viper.GetInt("http.retries.max_num_of_attempts")).
		SetRetryMaxWaitTime(time.Duration(viper.GetInt("http.retries.max_backoff_delay")) * time.Second).
		AddRetryCondition(
			func(response *resty.Response, err error) bool {
				return err != nil || response.StatusCode() != http.StatusOK
			},
		)
}

func SetTimeout() {
	Client.SetTimeout(time.Duration(viper.GetInt("http.timeout")) * time.Second)
}
