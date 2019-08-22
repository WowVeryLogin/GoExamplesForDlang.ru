package main

import (
	"fmt"
	"reflect"
)

//go:generate sums

func SlowSumInts(a interface{}) int64 {
	if reflect.ValueOf(a).Kind() != reflect.Slice {
		return 0
	}

	var s int64
	for i := 0; i < reflect.ValueOf(a).Len(); i++ {
		s += reflect.ValueOf(a).Index(i).Int()
	}
	return s
}

func FastSumInts(a interface{}) interface{} {
	switch v := a.(type) {
	case []int:
		var res int
		for i := range v {
			res += v[i]
		}
		return res
	case []int32:
		var res int32
		for i := range v {
			res += v[i]
		}
		return res
	}
	return 0
}

func main() {
	//With slow but general reflection
	fmt.Println(SlowSumInts([]int32{1, 2, 3, 4, 5}))

	//Fast type assertion but with code duplication
	fmt.Println(FastSumInts([]int32{1, 2, 3, 4, 5}))

	//Same as type assertion but generated with go generate
	fmt.Println(GeneratedSumInts([]int32{1, 2, 3, 4, 5}))

	//Fast in-place for concrete type
	var res int32
	for _, v := range []int32{1, 2, 3, 4, 5} {
		res += v
	}
}
