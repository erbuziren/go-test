package main

import "fmt"

func main() {
	f := fib();
	for i:=0;i<20;i++{
		fmt.Println(f())
	}

}

func fib() func() int{
	a,b:=0,1
	f:=func() int{
		a,b = b,a+b
		return a
	}
	return f
}