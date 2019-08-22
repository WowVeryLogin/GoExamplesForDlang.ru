//Package sums is generated. Do not edit it.
package main

func GeneratedSumInts(a interface{}) interface{} {
	switch v := a.(type) {
	case []int:
		var res int
		for i := range v {
			res += v[i]
		}
		return res
	case []int16:
		var res int16
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
	case []int64:
		var res int64
		for i := range v {
			res += v[i]
		}
		return res
	default:
		return nil
	}
}
