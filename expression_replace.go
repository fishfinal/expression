// Copyright 2022 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import "regexp"

// RegularReplaceMatch Use a regular expression to replace original string to novel string.
// The underlying call is regexp.Compile(patten).ReplaceAllString() method.
func RegularReplaceMatch(patten string, original string, replace string) (string, error) {
	regular, err := regexp.Compile(patten)
	if err != nil {
		return original, err
	}
	return regular.ReplaceAllString(original, replace), nil
}
