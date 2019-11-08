// +build slimming_analyse

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
 * slimming.go: Symbol usage analysis methods for code slimming
 */

package runtime

import "fmt"

// slimmingAnalyse toggles run-time symbol usage analysis
const slimmingAnalyse = true

var usedSymbols = make(map[string]bool)

// UsedSymbol marks the symbol having the specified name, as used.
func UsedSymbol(name string, v Object) Object {
	usedSymbols[name] = true
	return v
}

func usedSymbolDictAttr(dict Dict, attr string) {
	tnamei := dict.GetItem(NewStr("__symbol__"))
	if tnamei == nil {
		// No symbol will be defined for runtime created types,
		// for example using namedtuple.
		return
	}
	tname := tnamei.(Str).Value()
	name := fmt.Sprintf("%s.%s", tname, attr)
	usedSymbols[name] = true
}

func usedSymbolAttr(t Type, attr Str) {
	if t.Builtin() {
		return
	}
	usedSymbolDictAttr(t.Dict(), attr.Value())
}

// GetUsedSymbols returns the set of used symbols.
func GetUsedSymbols() []string {
	names := []string{}
	for name := range usedSymbols {
		names = append(names, name)
	}
	return names
}
