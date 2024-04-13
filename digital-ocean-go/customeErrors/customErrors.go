package customeerrors

type error interface {
	Error() string
}

type MyError struct {}

func (m *MyError) Error() string {
	return "boom"
}

func sayHello() (string, error) {
	return "", &MyError{}
}