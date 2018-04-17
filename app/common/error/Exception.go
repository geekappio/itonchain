package error

import "fmt"

type Throwable interface {
	getMessage() string
	getCause() *Exception
	toString() string
}

type Exception struct {
	Throwable
	message string
	cause *Exception
}

// Return message of exception
func (exception *Exception) getMessage() string {
	return exception.message
}

// Return cause of exception
func (exception *Exception) getCause() *Exception {
	return exception.cause
}

func (exception *Exception) toString() string {
	if exception.getCause() == nil {
		return fmt.Sprintf("%v", exception.message)
	} else {
		return fmt.Sprintf("%v\n%v", exception.message, exception.getCause().toString())
	}
}