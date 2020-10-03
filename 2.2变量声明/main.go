package main

import (
	"fmt"
	"math"
)

func main() {
	var z float32
	//分别对应值 0 -0 +Inf -Inf NaN
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	//math.IsNaN用来测试结果是否是NaN，math.IsInf用来测试结果是否是正无穷大或者负无穷大
	nan := math.NaN()
	inf := math.Inf(0)
	fmt.Println(math.IsNaN(nan))
	fmt.Println(math.IsInf(inf, 0))
}
