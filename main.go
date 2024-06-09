package main

import "fmt"

const (
	google = 1
	apple  = 2
)

func StubFactory(partner int) Stub {
	switch partner {
	case apple:
		return Apple{}
	case google:
		return Google{}
	default:
		return nil
	}
}
func main() {
	stubGoogle := StubFactory(google)
	if stubGoogle != nil {
		fmt.Println(stubGoogle.configWallet()) // Output: true
	} else {
		fmt.Println("Failed to create Google instance")
	}

	stubApple := StubFactory(apple)
	if stubApple != nil {
		fmt.Println(stubApple.configWallet()) // Output: false
	} else {
		fmt.Println("Failed to create Apple instance")
	}

	stubInvalid := StubFactory(3)
	if stubInvalid == nil {
		fmt.Println("Invalid type provided to factory")
	}
}
