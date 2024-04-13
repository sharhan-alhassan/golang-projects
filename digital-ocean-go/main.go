package main

import (
	"errors"
	"fmt"
	"os"

	handleerrors "github.com/sharhan-alhassan/golang-projects/digital-ocean-go/handleErrors"
)

type RequestError struct {
	StatusCode int
	Err error			
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err: errors.New("unavailable"),
	}
}

func main() {
	// Run custom error from "errors" package
	handleerrors.CustomError()

	// Run using Errorf package
	handleerrors.CustomErrorf()

	// Run Boom() error function
	// err := handleerrors.Boom()
	// if err != nil {
	// 	fmt.Println("An error occured:", err)
	// }
	// if err := handleerrors.Boom(); err != nil {
	// 	fmt.Println("An error occured:", err)
	// 	// exit the process
	// 	return
	// }
	// fmt.Println("Yeeih...no errors")

	// run 
	name, size, err := handleerrors.ReturnValueandError("sharhan")
	if err != nil {
		fmt.Println("An error occured: ", err)
		return
	}
	fmt.Printf("Capitalized name: %v, length: %d", name, size)

	err = doRequest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("success")

}
