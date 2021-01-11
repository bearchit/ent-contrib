// Copyright 2004-present Facebook
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

// Code generated by entc, DO NOT EDIT.

package ent_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/facebookincubator/ent-contrib/entgql/internal/todo/ent"
	"github.com/stretchr/testify/assert"
)

func TestCursorEncoding(t *testing.T) {
	t.Run("EncodeDecode", func(t *testing.T) {
		const (
			id    = 42
			value = "foobar"
			quote = uint8('"')
		)
		var buf bytes.Buffer
		c := ent.Cursor{ID: id, Value: value}
		c.MarshalGQL(&buf)
		s := buf.String()
		assert.Equal(t, quote, s[0])
		n := len(s) - 1
		assert.Equal(t, quote, s[n])
		c = ent.Cursor{}
		err := c.UnmarshalGQL(s[1:n])
		assert.NoError(t, err)
		assert.Equal(t, id, c.ID)
		assert.Equal(t, value, c.Value)
	})
	t.Run("EncodeNoValue", func(t *testing.T) {
		const id = 55
		var buf bytes.Buffer
		c := ent.Cursor{ID: id}
		c.MarshalGQL(&buf)
		s, err := strconv.Unquote(buf.String())
		assert.NoError(t, err)
		c = ent.Cursor{}
		err = c.UnmarshalGQL(s)
		assert.NoError(t, err)
		assert.Equal(t, id, c.ID)
		assert.Nil(t, c.Value)
	})
	t.Run("DecodeBadInput", func(t *testing.T) {
		inputs := []interface{}{
			0xbadbeef,
			"cursor@bad123",
			"Y3Vyc29yQGJhZDEyMw==",
		}
		for _, input := range inputs {
			var c ent.Cursor
			err := c.UnmarshalGQL(input)
			assert.Error(t, err)
		}
	})
}
