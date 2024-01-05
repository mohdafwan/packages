package pub

import (
	"fmt"
)

func Barks() string {
	return "Barks"
}

func Bark() string {
	return "Another Barks"
}

func From11() {
	fmt.Println("I Am from version 1.1.0")
}

func From12(s string) interface{} {
	return "This is the version 1.2.0" + s
}
