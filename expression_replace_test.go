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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularReplaceMatch(t *testing.T) {
	grids := []struct {
		patten   string
		original string
		replace  string
		expected string
	}{
		{
			patten:   `\{\{\.NAME\}\}`,
			original: "The {{.NAME}} Application starting...",
			replace:  `CallMe`,
			expected: "The CallMe Application starting...",
		},
	}

	for _, grid := range grids {
		actual, err := RegularReplaceMatch(grid.patten, grid.original, grid.replace)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, actual)
	}
}
