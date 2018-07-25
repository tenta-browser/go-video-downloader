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
 * scraps.go: Simple scraps of code (ex: casts, conversion) generated for the types used in extractors
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
	"reflect"
)

type τ_T00ω struct {
	Φ0 interface{}
	Φ1 interface{}
}

type τ_TOsOsω struct {
	Φ0 rnt.OptString
	Φ1 rnt.OptString
}

type τ_TOssω struct {
	Φ0 rnt.OptString
	Φ1 string
}

type τ_TTssωTssωω struct {
	Φ0 τ_Tssω
	Φ1 τ_Tssω
}

type τ_TiTssωω struct {
	Φ0 int
	Φ1 τ_Tssω
}

type τ_Tisω struct {
	Φ0 int
	Φ1 string
}

type τ_Tiαω struct {
	Φ0 int
	Φ1 interface{}
}

type τ_Tsssssω struct {
	Φ0 string
	Φ1 string
	Φ2 string
	Φ3 string
	Φ4 string
}

type τ_Tssssω struct {
	Φ0 string
	Φ1 string
	Φ2 string
	Φ3 string
}

type τ_Tsssω struct {
	Φ0 string
	Φ1 string
	Φ2 string
}

type τ_Tssω struct {
	Φ0 string
	Φ1 string
}

type τ_Tsαω struct {
	Φ0 string
	Φ1 interface{}
}

type τ_TsζResponseωω struct {
	Φ0 string
	Φ1 rnt.Response
}

type τ_Tsω struct {
	Φ0 string
}

type τ_Tα0ω struct {
	Φ0 interface{}
	Φ1 interface{}
}

type τ_TαOiω struct {
	Φ0 interface{}
	Φ1 rnt.OptInt
}

type τ_Tαsω struct {
	Φ0 interface{}
	Φ1 string
}

func τ_all_Lb(list []bool) bool {
	for _, el := range list {
		if el {
			return false
		}
	}
	return true
}

func τ_any_Lb(list []bool) bool {
	for _, el := range list {
		if el {
			return true
		}
	}
	return false
}

func τ_cast_α_to_Lα(any interface{}) []interface{} {
	anyv := reflect.ValueOf(any)
	len := anyv.Len()
	list := make([]interface{}, len)
	for i := 0; i < len; i++ {
		anyel := anyv.Index(i).Interface()
		list[i] = anyel
	}
	return list
}

func τ_cast_α_to_d(val interface{}) rnt.SDict {
	res, ok := val.(rnt.SDict)
	if !ok {
		panic(rnt.NewCastError(val, "rnt.SDict"))
	}
	return res
}

func τ_constr_SOs_from_LOs(iter []rnt.OptString) map[rnt.OptString]struct{} {
	set := map[rnt.OptString]struct{}{}
	for _, element := range iter {
		set[element] = struct{}{}
	}
	return set
}

func τ_constr_Ss_from_Ls(iter []string) map[string]struct{} {
	set := map[string]struct{}{}
	for _, element := range iter {
		set[element] = struct{}{}
	}
	return set
}

func τ_constr_Sα_from_Lα(iter []interface{}) map[interface{}]struct{} {
	set := map[interface{}]struct{}{}
	for _, element := range iter {
		set[element] = struct{}{}
	}
	return set
}

func τ_conv_LOs_to_Lα(src []rnt.OptString) []interface{} {
	dest := []interface{}{}
	for _, element := range src {
		dest = append(dest, element)
	}
	return dest
}

func τ_conv_Ld_to_Lα(src []rnt.SDict) []interface{} {
	dest := []interface{}{}
	for _, element := range src {
		dest = append(dest, element)
	}
	return dest
}

func τ_conv_Ls_to_Ls(src []string) []string {
	dest := []string{}
	for _, element := range src {
		dest = append(dest, element)
	}
	return dest
}

func τ_conv_Ls_to_Lα(src []string) []interface{} {
	dest := []interface{}{}
	for _, element := range src {
		dest = append(dest, element)
	}
	return dest
}

func τ_conv_Lα_to_Ld(src []interface{}) []rnt.SDict {
	dest := []rnt.SDict{}
	for _, element := range src {
		dest = append(dest, τ_cast_α_to_d(element))
	}
	return dest
}

func τ_conv_SOs_to_LOs(set map[rnt.OptString]struct{}) []rnt.OptString {
	list := []rnt.OptString{}
	for element := range set {
		list = append(list, element)
	}
	return list
}

func τ_conv_Ss_to_Ls(set map[string]struct{}) []string {
	list := []string{}
	for element := range set {
		list = append(list, element)
	}
	return list
}

func τ_conv_T00ω_to_TOsOsω(src τ_T00ω) τ_TOsOsω {
	return τ_TOsOsω{
		Φ0: τ_sink_Os(src.Φ0),
		Φ1: τ_sink_Os(src.Φ1),
	}
}

func τ_conv_TTssωTssωω_to_LTssω(tuple τ_TTssωTssωω) []τ_Tssω {
	return []τ_Tssω{tuple.Φ0, tuple.Φ1}
}

func τ_conv_Tsssssω_to_Ls(tuple τ_Tsssssω) []string {
	return []string{tuple.Φ0, tuple.Φ1, tuple.Φ2, tuple.Φ3, tuple.Φ4}
}

func τ_conv_Tssssω_to_Ls(tuple τ_Tssssω) []string {
	return []string{tuple.Φ0, tuple.Φ1, tuple.Φ2, tuple.Φ3}
}

func τ_conv_Tsssω_to_Ls(tuple τ_Tsssω) []string {
	return []string{tuple.Φ0, tuple.Φ1, tuple.Φ2}
}

func τ_conv_Tssω_to_Ls(tuple τ_Tssω) []string {
	return []string{tuple.Φ0, tuple.Φ1}
}

func τ_conv_Tsω_to_Ls(tuple τ_Tsω) []string {
	return []string{tuple.Φ0}
}

func τ_enumerate_LTssω(list []τ_Tssω, start int) []τ_TiTssωω {
	enum := []τ_TiTssωω{}
	for idx, el := range list {
		enum = append(enum, τ_TiTssωω{idx + start, el})
	}
	return enum
}

func τ_enumerate_Lα(list []interface{}, start int) []τ_Tiαω {
	enum := []τ_Tiαω{}
	for idx, el := range list {
		enum = append(enum, τ_Tiαω{idx + start, el})
	}
	return enum
}

func τ_index_Ls(list []string, idx int) string {
	if idx < 0 {
		idx += len(list)
	}
	return list[idx]
}

func τ_isTruthy_Oi(val rnt.OptInt) bool {
	return val.IsSet() && (val.Get()) != 0
}

func τ_isTruthy_Os(val rnt.OptString) bool {
	return val.IsSet() && (val.Get()) != ""
}

func τ_isinstance_d(value interface{}) bool {
	_, ok := value.(rnt.SDict)
	return ok
}

func τ_isinstance_s(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func τ_map_Fssω_over_Ls(fn func(string) string, list []string) []string {
	result := []string{}
	for _, el := range list {
		result = append(result, fn(el))
	}
	return result
}

func τ_map_Fsαω_over_Ls(fn func(string) interface{}, list []string) []interface{} {
	result := []interface{}{}
	for _, el := range list {
		result = append(result, fn(el))
	}
	return result
}

func τ_multiply_L0_by_i(list []interface{}, factor int) []interface{} {
	result := []interface{}{}
	for idx := 0; idx < factor; idx++ {
		result = append(result, list...)
	}
	return result
}

func τ_search_s_in_Ls(needle string, haystack []string) bool {
	for _, haystackEl := range haystack {
		if haystackEl == needle {
			return true
		}
	}
	return false
}

func τ_search_s_notin_Ls(needle string, haystack []string) bool {
	for _, haystackEl := range haystack {
		if haystackEl == needle {
			return false
		}
	}
	return true
}

func τ_search_s_notin_Sα(needle string, haystack map[interface{}]struct{}) bool {
	for haystackEl := range haystack {
		if haystackEl == needle {
			return false
		}
	}
	return true
}

func τ_search_α_in_Ls(needle interface{}, haystack []string) bool {
	for _, haystackEl := range haystack {
		if haystackEl == needle {
			return true
		}
	}
	return false
}

func τ_search_α_in_Sα(needle interface{}, haystack map[interface{}]struct{}) bool {
	for haystackEl := range haystack {
		if haystackEl == needle {
			return true
		}
	}
	return false
}

func τ_search_α_notin_Sα(needle interface{}, haystack map[interface{}]struct{}) bool {
	for haystackEl := range haystack {
		if haystackEl == needle {
			return false
		}
	}
	return true
}

func τ_sink_Fssω(nothing interface{}) func(string) string {
	return nil
}

func τ_sink_LOs(nothing interface{}) []rnt.OptString {
	return nil
}

func τ_sink_Ls(nothing interface{}) []string {
	return nil
}

func τ_sink_Lα(nothing interface{}) []interface{} {
	return nil
}

func τ_sink_Of(nothing interface{}) rnt.OptFloat {
	return rnt.OptFloat{}
}

func τ_sink_Oi(nothing interface{}) rnt.OptInt {
	return rnt.OptInt{}
}

func τ_sink_Os(nothing interface{}) rnt.OptString {
	return rnt.OptString{}
}

func τ_sink_d(nothing interface{}) rnt.SDict {
	return nil
}

func τ_sink_ζbytesω(nothing interface{}) []byte {
	return nil
}

func τ_zip_i_s(list1 []int, list2 []string) []τ_Tisω {
	result := []τ_Tisω{}
	len := rnt.Min(len(list1), len(list2))
	for i := 0; i < len; i++ {
		result = append(result, τ_Tisω{list1[i], list2[i]})
	}
	return result
}
