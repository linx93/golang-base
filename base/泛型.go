package main

// Number 数字类型，是我们自定义的类型限制，含所有的数字类型
type Number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

//简单版：支持int64和float64
func sum[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, item := range m {
		sum += item
	}
	return sum
}

//超强版：支持int | int8 | int16 | int32 | int64 | float64 | float32    就是Number
func sumNumber[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, item := range m {
		sum += item
	}
	return sum
}

//func main() {
//	//简单版测试
//	m := map[string]int64{"a": 1, "b": 2}
//	sumIntResult := sum(m)
//	fmt.Println(sumIntResult)
//
//	//超强版本测试
//	mm := map[string]float32{"a": 1.1, "b": 2.2}
//	fmt.Println(sumNumber(mm))
//}
