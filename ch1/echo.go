//echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func divider(x int) {
	fmt.Printf("\n------- echo %d -------", x)
}

func echo1() {
	divider(1)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	divider(2)
	s, sep := "", "-"
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	divider(3)
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo1()
	echo2()
	echo3()
}
