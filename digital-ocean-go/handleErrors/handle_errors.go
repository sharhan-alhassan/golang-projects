package handleerrors

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CustomError()  {
	err := errors.New("my new error")
	fmt.Println("Error is : ", err)
}

func CustomErrorf()  {
	err := fmt.Errorf("error occured at: %v", time.Now())
	fmt.Println("An error happened: ", err)
}

func Boom() error {
	return errors.New("boom")
}

func ReturnValueandError(name string) (string, int, error) {
	handle := func (err error) (string, int, error) {
		return "", 0, err
	}

	if name == "" {
		// return "", errors.New("no name provided ")
		return handle(errors.New("no name provided"))
	}

	return strings.ToTitle(name), len(name), nil
}