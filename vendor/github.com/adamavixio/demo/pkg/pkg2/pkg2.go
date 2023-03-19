package pkg2

import (
	"fmt"

	"github.com/adamavixio/demo/pkg/pkg1"
)

func Hello(arg string) {
	fmt.Printf("Calling pkg1 from pkg2 with arg %v\n", arg)
	pkg1.Hello(arg)
}
