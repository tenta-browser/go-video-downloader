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
 * log.go: Basic logging utils
 */

package utils

import "fmt"

const (
	// NoLogLevel turns off logging
	NoLogLevel = iota
	// StdLogLevel permits normal logging
	StdLogLevel
	// VerboseLogLevel turns on verbose logging
	VerboseLogLevel
)

// LogLevel controls the amount of logging
var LogLevel = NoLogLevel

// Log logs a message from a running extractor
func Log(msg string, args ...interface{}) {
	if LogLevel >= StdLogLevel {
		fmt.Printf(msg+"\n", args...)
	}
}

// LogV logs a verbose message from a running extractor
func LogV(msg string, args ...interface{}) {
	if LogLevel >= VerboseLogLevel {
		Log(msg, args...)
	}
}
