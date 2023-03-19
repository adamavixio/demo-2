package pkg1

import "fmt"

// Notice that only Hello can be exported. You can call
// it in other packages. The lowercase hello method is
// not exported. That is how Go determines exports.

// This method is exported because it
// is capitalized
func Hello(arg string) {
	hello(arg)
}

// This method is NOT exported because it
// is NOT capitalized
func hello(arg string) {
	fmt.Printf("called from pkg1 with args %v\n", arg)
}
