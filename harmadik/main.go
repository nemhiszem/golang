package main

import (
	"fmt"
)

func main() {
	//	bool tipus
	/*var n bool = false
	fmt.Printf("%v, %T\n", n, n)*/

	/*n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)*/

	/*var n bool a zero érték false
	fmt.Printf("%v, %T\n", n, n)*/

	//	numerikus tipus
	/*var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)*/

	/*a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b) menet közben nem változtatható az változó típusa*/

	/*var a int = 10
	var b int8 = 3
	fmt.Println(a + b) eltérő típusúakkal nem lehet műveletet végezni

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b)) de így ha convertálva van akkor már lehet*/

	/*a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100  ezek a bit operation-ok */

	/*a := 8 // 2^3
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 bitshifting */

	/*var n float64 = 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)*/

	/*a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) float-al nem lehet bit shiftet vagy bit operationokat használni ahoz int tipus kell*/

	//	complex number
	/*var n complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", n, n)*/

	/*a := 1 + 2i
	b := 2 + 5.2i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)*/

	/*var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))*/

	// text tipus
	/*s := "this is a string"
	fmt.Printf("%v, %T\n", s, s) // fmt.Printf("%v, %T\n", string(s[2]), s[2]) */

	/*s := "this is a string"
	b  := []byte(s)
	fmt.Printf("%v, %T\n", b, b)*/

	/*r := 'a'
	fmt.Printf("%v, %T\n", r, r)*/

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // a videóban amit néztem elmagyarázta mi a run szóval azt nézzétek: https://youtu.be/YS4e4q9oBaU
}
