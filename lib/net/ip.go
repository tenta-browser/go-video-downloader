/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * ip.go: Methods for parsing and generating IP addresses
 */

package http

import (
	"encoding/binary"
	"math/rand"
	"net"

	// imported to seed the random number generator
	_ "github.com/tenta-browser/go-video-downloader/lib/random"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// RandomIPv4 generates a random IPv4 address in the range denoted by cidr
var RandomIPv4 = rnt.NewSimpleFunction("RandomIP",
	[]string{"cidr"},
	func(args []rnt.Object) rnt.Object {
		cidr := args[0].(rnt.Str).Value()
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(rnt.RaiseType(rnt.ValueErrorType, err.Error()))
		}
		ip := make(net.IP, net.IPv4len)
		binary.LittleEndian.PutUint32(ip, rand.Uint32())
		for i := range ip {
			ip[i] = ipnet.IP[i] | (ip[i] & ^ipnet.Mask[i])
		}
		return rnt.NewStr(ip.String())
	})
