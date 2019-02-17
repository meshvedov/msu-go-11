package main

import (
	"bytes"
	"fmt"
	"strings"
)

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	fibonacci = memoize(func(i int, i2 ...int) interface{} {
		if i <= 0 || i == 1 {
			return i
		}
		return fibonacci(i-2).(int) + fibonacci(i-1).(int)
	})

	desimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanForDecimal = memoize(func(i int, i2 ...int) interface{} {
		var buffer bytes.Buffer
		for i, desimal := range desimals {
			remainder := i / desimal
			i %= desimal
			if remainder > 0 {
				buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}
		return buffer.String()
	})
}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
