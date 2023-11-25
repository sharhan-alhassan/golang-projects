package main

import "fmt"

func pointerExample(x *string) string {
	return *x
}

func main()  {
	name := "Sharhan"
	ptr := &name

	n := pointerExample(ptr)
	fmt.Printf("Pointer value: %v", n)
}