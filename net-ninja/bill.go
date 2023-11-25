
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func createBill() bill {
	// Read from the std input i.e the terminal
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Create a new bill name: ")
	// Read everthing after when they type enter, thus the newline "\n" and store it in "name" variable
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}


func main()  {
	// myBill := newBill("Sharhan")
	myBill := createBill()

	myBill.addItem("coffee", 4.99)
	myBill.updateTip(5.99)

	fmt.Println(myBill.format())
}