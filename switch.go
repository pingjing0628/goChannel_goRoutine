package main

var (
	i interface{}
)

// 依照 case 順序依序執行
func main()  {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))
}

func convert(i interface{})  {
	switch t := i.(type) {
	case int:
		println("i is integer", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	default:
		println("type not found.")
	}
}
