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
	"github.com/cauwulixuan/go-kit/log"
	"net/http"
)

//GetUrl Get data from a given url.
func GetUrl(url string) string {
	resp, err := Client.R().Get(url)
	if err != nil {
		log.SErrorf("Error happend while get url %s, error message: %v", url, err.Error())
		return ""
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.String()
	} else {
		log.SInfof("Status code is %d not 200 while getting url %s", resp.StatusCode(), url)
		return ""
	}
}
