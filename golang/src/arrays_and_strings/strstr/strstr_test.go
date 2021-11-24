package strstr

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := []string{
		"a",
		"hello",
		"aaaaa",
		"",
	}
	needle := []string{
		"a",
		"ll",
		"bba",
		"",
	}
	expected := []int{
		0,
		2,
		-1,
		0,
	}
	testFuncs := []func(haystack string, needle string) int{
		strStr,
	}

	for _, testFunc := range testFuncs {
		for index, hs := range haystack {
			if res := testFunc(hs, needle[index]); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
