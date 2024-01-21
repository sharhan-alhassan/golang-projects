package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	err := writeToTempFile("Some Text")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Write to file successful")
}

func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}

	n, err := file.WriteString("Some text")
	if err != nil {
		return err
	}

	fmt.Printf("Number of bytes written: %d", n)
	file.Close()
	return nil
}	