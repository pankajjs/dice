// Copyright (c) 2022-present, DiceDB contributors
// All rights reserved. Licensed under the BSD 3-Clause License. See LICENSE file in the project root for full license information.

package ironhawk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHGET(t *testing.T) {
	client := getLocalConnection()
	defer client.Close()
	defer client.FireString("DEL key_hGet key")

	testCases := []TestCase{
		{
			commands: []string{"HGET", "HGET KEY", "HGET KEY FIELD ANOTHER_FIELD"},
			expected: []interface{}{"ERR wrong number of arguments for 'hget' command",
				"ERR wrong number of arguments for 'hget' command",
				"ERR wrong number of arguments for 'hget' command"},
		},
		{
			commands: []string{"HSET key_hGet field value", "HSET key_hGet field newvalue"},
			expected: []interface{}{ONE, ZERO},
		},
		{
			commands: []string{"HGET wrong_key_hGet field"},
			expected: []interface{}{"(nil)"},
		},
		{
			commands: []string{"HGET key_hGet wrong_field"},
			expected: []interface{}{"(nil)"},
		},
		{
			commands: []string{"HGET key_hGet field"},
			expected: []interface{}{"newvalue"},
		},
		{
			commands: []string{"SET key value", "HGET key field"},
			expected: []interface{}{"OK", "WRONGTYPE Operation against a key holding the wrong kind of value"},
		},
	}

	for _, tc := range testCases {
		for i, cmd := range tc.commands {
			result := client.FireString(cmd)
			assert.Equal(t, tc.expected[i], result)
		}
	}
}
