/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 *
 * For any questions, please contact developer@tenta.io
 *
 * misc.go: Miscellaneous utilities
 */

package utils

import "fmt"

// Debug mode permits logging
const Debug = false

// Log logs a message from a running extractor
func Log(msg string, args ...interface{}) {
	if Debug {
		fmt.Printf(msg+"\n", args...)
	}
}

// MakeSet creates a map from the specified keys;
// the values are irrelevant, this just wants to simulate a set
func MakeSet(values []string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, value := range values {
		res[value] = struct{}{}
	}
	return res
}
