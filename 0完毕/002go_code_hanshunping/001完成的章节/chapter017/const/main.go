package main

import(
	"fmt"
)

func main()  {
	const(
		a = iota
		b
		c,d = iota,iota
		e = iota
		f
	)
	fmt.Println(a,b,c,d,e,f)
}