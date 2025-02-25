/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cast_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-spring/spring-base/assert"
	"github.com/go-spring/spring-base/cast"
)

func ptr(i interface{}) interface{} {
	switch v := i.(type) {
	case bool:
		return &v
	case int:
		return &v
	case int8:
		return &v
	case int16:
		return &v
	case int32:
		return &v
	case int64:
		return &v
	case uint:
		return &v
	case uint8:
		return &v
	case uint16:
		return &v
	case uint32:
		return &v
	case uint64:
		return &v
	case float32:
		return &v
	case float64:
		return &v
	case string:
		return &v
	default:
		return nil
	}
}

func TestToInt(t *testing.T) {

	testcases := []struct {
		param  interface{}
		expect int64
	}{
		{int64(10), int64(10)},
		{ptr(int64(10)), int64(10)},
		{10.0, int64(10)},
		{ptr(10.0), int64(10)},
		{"10", int64(10)},
		{ptr("10"), int64(10)},
		{true, int64(1)},
		{ptr(true), int64(1)},
	}

	for i, testcase := range testcases {
		v := cast.ToInt64(testcase.param)
		assert.Equal(t, v, testcase.expect, fmt.Sprintf("index %d", i))
	}
}

func TestToTime(t *testing.T) {

	t.Run("unit", func(t *testing.T) {

		testcases := []struct {
			value  int64
			unit   time.Duration
			expect time.Time
		}{
			{1, time.Nanosecond, time.Unix(0, 1)},
			{1, time.Millisecond, time.Unix(0, 1*1e6)},
			{1, time.Second, time.Unix(1, 0)},
			{1, time.Hour, time.Unix(0, 0).Add(time.Hour)},
		}

		for i, testcase := range testcases {
			s := cast.ToTime(testcase.value, cast.TimeArg{Unit: testcase.unit})
			assert.Equal(t, s, testcase.expect, fmt.Sprintf("index %d", i))
		}
	})

	t.Run("format", func(t *testing.T) {

		testcases := []struct {
			value  string
			format string
			expect time.Time
		}{
			{
				"1970-01-01 08:00:00.000000001 +0800",
				"2006-01-02 15:04:05.000000000 -0700",
				time.Unix(0, 1),
			},
		}

		for i, testcase := range testcases {
			s := cast.ToTime(testcase.value, cast.TimeArg{Format: testcase.format})
			assert.Equal(t, s, testcase.expect, fmt.Sprintf("index %d", i))
		}
	})
}
