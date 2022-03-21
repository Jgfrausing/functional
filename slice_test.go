package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	val int
}

type TestCase[T any] struct {
	name     string
	val      T
	lst      []T
	err      bool
	contains bool
}

func Test_Contains_String(t *testing.T) {
	testData := []TestCase[string]{
		{
			name:     "string value in list",
			val:      "a",
			lst:      []string{"a", "b"},
			err:      false,
			contains: true,
		},
		{
			name:     "string value not in list",
			val:      "c",
			lst:      []string{"a", "b"},
			err:      false,
			contains: false,
		},
	}

	for _, value := range testData {
		t.Run(value.name, func(t *testing.T) {
			asserter := assert.New(t)

			got := Contains(value.val, value.lst)
			asserter.Equal(value.contains, got)
		})
	}
}

func Test_Contains_Int(t *testing.T) {
	testData := []TestCase[int]{
		{
			name:     "int value in list",
			val:      1,
			lst:      []int{1, 2},
			err:      false,
			contains: true,
		},
		{
			name:     "int value not in list",
			val:      3,
			lst:      []int{1, 2},
			err:      false,
			contains: false,
		},
		{
			name:     "list is nil",
			val:      4,
			lst:      nil,
			err:      true,
			contains: false,
		},
	}

	for _, value := range testData {
		t.Run(value.name, func(t *testing.T) {
			asserter := assert.New(t)

			got := Contains(value.val, value.lst)
			asserter.Equal(value.contains, got)
		})
	}
}

func Test_Contains_Struct(t *testing.T) {
	testData := []TestCase[testStruct]{
		{
			name:     "struct value first in list",
			val:      testStruct{1},
			lst:      []testStruct{{1}, {2}},
			err:      false,
			contains: true,
		},
		{
			name:     "struct value last in list",
			val:      testStruct{2},
			lst:      []testStruct{{1}, {2}},
			err:      false,
			contains: true,
		},
		{
			name:     "struct value not in list",
			val:      testStruct{3},
			lst:      []testStruct{{1}, {2}},
			err:      false,
			contains: false,
		},
	}

	for _, value := range testData {
		t.Run(value.name, func(t *testing.T) {
			asserter := assert.New(t)

			got := Contains(value.val, value.lst)
			asserter.Equal(value.contains, got)
		})
	}
}

func Test_Filter(t *testing.T) {
	a := assert.New(t)
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	filtered := Filter(lst, func(i int) bool {
		return 5 < i
	})
	a.Equal(4, len(filtered))

}
func Test_Map(t *testing.T) {
	a := assert.New(t)
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapped := Map(lst, func(t int) string {
		return fmt.Sprintf("'%d'", t)
	})

	a.Equal(len(mapped), len(lst))

}
func Test_Reduce(t *testing.T) {
	a := assert.New(t)
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := "123456789"
	actual := Reduce(lst, "", func(r string, v int) string {
		return fmt.Sprintf("%s%d", r, v)
	})
	a.Equal(expected, actual)
}
