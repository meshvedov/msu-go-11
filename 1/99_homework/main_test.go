package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestReturnInt(t *testing.T) {
	if ReturnInt() != 1 {
		t.Error("expected 1")
	}
}

func ReturnInt() int {
	return 1
}

func TestReturnFloat(t *testing.T) {
	if ReturnFloat() != float32(1.1) {
		t.Error("expected 1.1")
	}
}

func ReturnFloat() float32 {
	return 1.1
}
func TestReturnIntArray(t *testing.T) {
	if ReturnIntArray() != [3]int{1, 3, 4} {
		t.Error("expected '[3]int{1, 3, 4}'")
	}
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}
func TestReturnIntSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	result := ReturnIntSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}
func TestIntSliceToString(t *testing.T) {
	expected := "1723100500"
	result := IntSliceToString([]int{17, 23, 100500})
	if expected != result {
		t.Error("expected", expected, "have", result)
	}
}

func IntSliceToString(s []int) (str string) {
	for _, v := range s {
		str += fmt.Sprint(v)
	}
	return str
}
func TestMergeSlices(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func MergeSlices(sflt []float32, sint []int32) (res []int) {
	for _, v := range sflt {
		res = append(res, int(v))
	}
	for _, v := range sint {
		res = append(res, int(v))
	}
	return res
}
func TestGetMapValuesSortedByKey(t *testing.T) {

	var cases = []struct {
		expected []string
		input    map[int]string
	}{
		{
			expected: []string{
				"Январь",
				"Февраль",
				"Март",
				"Апрель",
				"Май",
				"Июнь",
				"Июль",
				"Август",
				"Сентябрь",
				"Октябрь",
				"Ноябрь",
				"Декарь",
			},
			input: map[int]string{
				9:  "Сентябрь",
				1:  "Январь",
				2:  "Февраль",
				10: "Октябрь",
				5:  "Май",
				7:  "Июль",
				8:  "Август",
				12: "Декарь",
				3:  "Март",
				6:  "Июнь",
				4:  "Апрель",
				11: "Ноябрь",
			},
		},

		{
			expected: []string{
				"Зима",
				"Весна",
				"Лето",
				"Осень",
			},
			input: map[int]string{
				10: "Зима",
				30: "Лето",
				20: "Весна",
				40: "Осень",
			},
		},
	}

	for _, item := range cases {
		result := GetMapValuesSortedByKey(item.input)
		if !reflect.DeepEqual(result, item.expected) {
			t.Error("expected", item.expected, "have", result)
		}
	}
}

func GetMapValuesSortedByKey(m map[int]string) (res []string) {
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, v := range keys {
		res = append(res, m[v])
	}
	return res
}
