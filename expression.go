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

import (
	"fmt"
	"regexp"
	"time"
)

var (
	RegularOfWebDomain    = regexp.MustCompile(`^(http://|https://)?(\d{1,3}(.\d{1,3}){3}|(www.)?[^www.][a-z]+\.[a-z]+)(:\d+)?$`)
	RegularOfIPv4         = regexp.MustCompile(`^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`)
	RegularOfInternalIPv4 = regexp.MustCompile(`(^127\.)|(^10\.)|(^172\.(1[6-9]|2[0-9]|3[0-1])\.)|(^192\.168\.)`)
	RegularOfCellphone    = regexp.MustCompile(`^(\+86(-|\s)?)?1[3-9][0-9]{9}$`)
	RegularOfTelephone    = regexp.MustCompile(`^(\d{3,4}-)?[1-9][0-9]{4,8}$`)
	RegularOfEmail        = regexp.MustCompile(`^[a-zA-Z0-9]{1,20}(([._\-])?[a-zA-Z0-9]{1,20})*@[a-z0-9]+(\.([a-z]+))+$`)
)

// IsIPv4 regular expression is used to check whether the IP address is IPV4.
func IsIPv4(str string) bool {
	return RegularOfIPv4.MatchString(str)
}

// IsInternalIPv4 Use a regular expression to check whether the ip address is ipv4.
func IsInternalIPv4(str string) bool {
	return RegularOfInternalIPv4.MatchString(str)
}

// IsCellphoneOfChina regular expression is used to check whether the telephone number is china.
func IsCellphoneOfChina(str string) bool {
	return RegularOfCellphone.MatchString(str)
}

// IsTelephoneOfChina regular expression is used to check whether the telephone number is china.
func IsTelephoneOfChina(str string) bool {
	return RegularOfTelephone.MatchString(str)
}

// IsEmail Use a regular expression to check whether the email address is valid.
func IsEmail(str string) bool {
	return RegularOfEmail.MatchString(str)
}

// IsWebDomain Use a regular expression to check whether the web domain is valid.
func IsWebDomain(str string) bool {
	return RegularOfWebDomain.MatchString(str)
}

// IsTimeHourPoint matched time str is time hour point.
// str similar 00:01/05:28/19:25/23:59
func IsTimeHourPoint(str string) bool {
	var daytime = func(value string) *time.Time {
		parsed, err := time.Parse("2006-01-02 15:04:05", value)
		if err != nil {
			return nil
		}

		return &parsed
	}

	today := time.Now().Format("2006-01-02")
	parsed := daytime(fmt.Sprintf("%s %s:00", today, str))
	if parsed == nil {
		return false
	}

	return true
}
