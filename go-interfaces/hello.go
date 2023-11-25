package main
import "fmt"

func take_pointer(s *int)  {
	fmt.Println(s)
}

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

// func main()  {
// 	// Using a pointer
// 	s := "hello"
// 	var b *string = &s 

// 	var c = &s

// 	d := &s 

// 	// Instantiate the struct "Student"
// 	var std Student
// 	std.name = "sharhan"
// 	std.rollNo = 45
// 	std.marks = []int{23, 324, 23}

// 	st := new(Student)

// 	fmt.Println("Address of s is ", b)
// 	fmt.Println("Address of s is ", c)
// 	fmt.Println("Address of s is ", d)

// 	// Deferencing from a pointer
// 	fmt.Println("Original value: ", *d)

// 	// Print student struct std
// 	fmt.Printf("%+v :", std)
// 	fmt.Printf("%+v :", st)
// }

