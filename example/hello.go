package main

import (
	"fmt"
	"github.com/jakha/talib/talib"
	"math"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
	fmt.Println(talib.Sin([]float64{1, math.Pi / 2}))
}
