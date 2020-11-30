package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func HelloWorld() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	f, t bool    = true, false
)

const Pi = 3.14
const (
	usename  = "test_user"
	password = "test_password"
)

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	HelloWorld()
	fmt.Println(i, f64, s, t, f)
	r1, r2 := add(1, 2)
	fmt.Println(r1, r2)
	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	c2 := circleArea(3)
	fmt.Println(c2(2))

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xf, xt)

	fmt.Println(string("Hello World!"[0]))
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "test"))

	var x int = 1
	xx := float64(x)
	var y float64 = 1.2
	xy := int(y)
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Println(xx, xy, i)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2:4])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{2, 3, 4},
		[]int{5, 6, 7},
	}
	fmt.Println(board)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m["apple"])
	m["new"] = 500

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
