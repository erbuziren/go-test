package main

import "fmt"

func main() {
	str:= "my name is joshua"
	result := reverse(str)
	fmt.Println(result)
}

func reverse(text string) string{
	if len(text) == 1{
		return text
	}
	return text[len(text)-1:len(text)]+reverse(text[:len(text)-1])
}