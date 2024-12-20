package graphql

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalInt(123)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, 123, mustUnmarshalInt(t, 123))
		assert.Equal(t, 123, mustUnmarshalInt(t, int64(123)))
		assert.Equal(t, 123, mustUnmarshalInt(t, json.Number("123")))
		assert.Equal(t, 123, mustUnmarshalInt(t, "123"))
		assert.Equal(t, 0, mustUnmarshalInt(t, nil))
	})
}

func mustUnmarshalInt(t *testing.T, v any) int {
	res, err := UnmarshalInt(v)
	require.NoError(t, err)
	return res
}

func TestInt32(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalInt32(123)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, int32(123), mustUnmarshalInt32(t, 123))
		assert.Equal(t, int32(123), mustUnmarshalInt32(t, int64(123)))
		assert.Equal(t, int32(123), mustUnmarshalInt32(t, json.Number("123")))
		assert.Equal(t, int32(123), mustUnmarshalInt32(t, "123"))
		assert.Equal(t, int32(0), mustUnmarshalInt32(t, nil))
	})

	t.Run("overflow", func(t *testing.T) {
		cases := []struct {
			name string
			v    any
			err  string
		}{
			{"positive int overflow", math.MaxInt32 + 1, "2147483648 overflows 32-bit integer"},
			{"negative int overflow", math.MinInt32 - 1, "-2147483649 overflows 32-bit integer"},
			{"positive int overflow", int64(math.MaxInt32 + 1), "2147483648 overflows 32-bit integer"},
			{"negative int overflow", int64(math.MinInt32 - 1), "-2147483649 overflows 32-bit integer"},
			{"positive json.Number overflow", json.Number("2147483648"), "2147483648 overflows 32-bit integer"},
			{"negative json.Number overflow", json.Number("-2147483649"), "-2147483649 overflows 32-bit integer"},
			{"positive string overflow", "2147483648", "2147483648 overflows 32-bit integer"},
			{"negative string overflow", "-2147483649", "-2147483649 overflows 32-bit integer"},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				res, err := UnmarshalInt32(tc.v)
				assert.EqualError(t, err, tc.err) //nolint:testifylint // An error assertion makes more sense.
				assert.Equal(t, int32(0), res)
			})
		}
	})
}

func mustUnmarshalInt32(t *testing.T, v any) int32 {
	res, err := UnmarshalInt32(v)
	require.NoError(t, err)
	return res
}

func TestInt64(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalInt32(123)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, int64(123), mustUnmarshalInt64(t, 123))
		assert.Equal(t, int64(123), mustUnmarshalInt64(t, int64(123)))
		assert.Equal(t, int64(123), mustUnmarshalInt64(t, json.Number("123")))
		assert.Equal(t, int64(123), mustUnmarshalInt64(t, "123"))
		assert.Equal(t, int64(0), mustUnmarshalInt64(t, nil))
	})
}

func mustUnmarshalInt64(t *testing.T, v any) int64 {
	res, err := UnmarshalInt64(v)
	require.NoError(t, err)
	return res
}
