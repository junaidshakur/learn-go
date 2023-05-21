package greetings

import (
	"fmt"
)

func Greet(name string) string {
	return fmt.Sprintf("Hi %v, Welcome to GO.", name)
}
