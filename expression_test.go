// Copyright 2021 fishfinal <contact.fishfinal@gmail.com>;
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsIPv4(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "192.168.0.1",
			expected: true,
		},
		{
			ip:       "0.0.0.0",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsIPv4(grid.ip)
		assert.Equal(t, grid.expected, actual)
	}
}

func Test_IsCellphoneOfChina(t *testing.T) {
	grids := []struct {
		cellphone string
		expected  bool
	}{
		{
			cellphone: "12323434334",
			expected:  false,
		},
		{
			cellphone: "13717995000",
			expected:  true,
		},
		{
			cellphone: "+8613717995000",
			expected:  true,
		},
		{
			cellphone: "+86 13717995000",
			expected:  true,
		},
		{
			cellphone: "+86-13717995000",
			expected:  true,
		},
		{
			cellphone: "+86 -13717995000",
			expected:  false,
		},
		{
			cellphone: "+86- 13717995000",
			expected:  false,
		},
		{
			cellphone: "+86 - 13717995000",
			expected:  false,
		},
	}

	for _, grid := range grids {
		actual := IsCellphoneOfChina(grid.cellphone)
		assert.Equal(t, grid.expected, actual)
	}
}

func Test_IsTelephoneOfChina(t *testing.T) {
	grids := []struct {
		telephone string
		expected  bool
	}{
		{
			telephone: "010-10086",
			expected:  true,
		},
		{
			telephone: "10086",
			expected:  true,
		},
		{
			telephone: "96069",
			expected:  true,
		},
		{
			telephone: "0350-346321d",
			expected:  false,
		},
		{
			telephone: "0350-346321",
			expected:  true,
		},
		{
			telephone: "0350-3463241",
			expected:  true,
		},
		{
			telephone: "010-92832839",
			expected:  true,
		},
		{
			telephone: "010-02832839",
			expected:  false,
		},
	}

	for _, grid := range grids {
		actual := IsTelephoneOfChina(grid.telephone)
		assert.Equal(t, grid.expected, actual)
	}
}

func Test_IsEmail(t *testing.T) {
	grids := []struct {
		email    string
		expected bool
	}{
		{
			email:    "wu.shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "a@tom.com",
			expected: true,
		},
		{
			email:    "a.b@tom.com",
			expected: true,
		},
		{
			email:    "a.helloshaohua@tom.com.cn",
			expected: true,
		},
		{
			email:    "a@foxmail.com",
			expected: true,
		},
		{
			email:    "wu.shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "wu-shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "wu_shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "123456@163.com",
			expected: true,
		},
		{
			email:    "123456@qq.com",
			expected: true,
		},
		{
			email:    "123456@com",
			expected: false,
		},
		{
			email:    "hello.@foxmail.com",
			expected: false,
		},
		{
			email:    "hello.world....@foxmail.com",
			expected: false,
		},
		{
			email:    "helle..@foxmail.com",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsEmail(grid.email)
		assert.Equal(t, grid.expected, actual,
			"want: %t, but got: %t, email: %+v\n", grid.expected, actual, grid.email)
	}
}

func Test_IsInternalIPv4(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "127.0.0.1",
			expected: true,
		},
		{
			ip:       "10.0.0.1",
			expected: true,
		},
		{
			ip:       "172.16.0.1",
			expected: true,
		},
		{
			ip:       "172.26.0.1",
			expected: true,
		},
		{
			ip:       "172.30.0.1",
			expected: true,
		},
		{
			ip:       "172.30.0.1",
			expected: true,
		},
		{
			ip:       "172.10.0.1",
			expected: false,
		},
		{
			ip:       "172.1.0.1",
			expected: false,
		},
		{
			ip:       "172.0.0.1",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsInternalIPv4(grid.ip)
		assert.Equal(t, grid.expected, actual)
	}
}

func Test_IsWebDomain(t *testing.T) {
	grids := []struct {
		domain   string
		expected bool
	}{
		{
			domain:   "https://baidu.com",
			expected: true,
		},
		{
			domain:   "https://www.baidu.com",
			expected: true,
		},
		{
			domain:   `http://baidu.com:80`,
			expected: true,
		},
		{
			domain:   "https://www.www.baidu.com",
			expected: false,
		},
		{
			domain:   "www.www.com",
			expected: false,
		},
		{
			domain:   "ww.baidu.com",
			expected: false,
		},
		{
			domain:   "192.168.0.1",
			expected: true,
		},
		{
			domain:   "http:192.168.0.1",
			expected: false,
		},
		{
			domain:   "http://192.168.0.1",
			expected: true,
		},
		{
			domain:   "https://192.168.0.1:8092",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsWebDomain(grid.domain)
		assert.Equal(t, grid.expected, actual, "domain: %s", grid.domain)
	}
}

func Test_IsTimeHourPoint(t *testing.T) {
	grids := []struct {
		timeStr  string
		expected bool
	}{
		{
			timeStr:  "09:59",
			expected: true,
		},

		{
			timeStr:  "10:59",
			expected: true,
		},
		{
			timeStr:  "00:01",
			expected: true,
		},
		{
			timeStr:  "23:59",
			expected: true,
		},
		{
			timeStr:  "00:00",
			expected: true,
		},
		{
			timeStr:  "24:01",
			expected: false,
		},
		{
			timeStr:  "10:60",
			expected: false,
		},
		{
			timeStr:  "02:60",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsTimeHourPoint(grid.timeStr)
		assert.Equal(t, grid.expected, actual, "timeStr: %s", grid.timeStr)
	}
}
