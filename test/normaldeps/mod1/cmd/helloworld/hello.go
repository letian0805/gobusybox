package main

import (
	"fmt"

	// A package whose import path does not match its $GOPATH.
	"github.com/letian0805/gobusybox/test/normaldeps/mod2/v2/pkg/hello"
)

func main() {
	fmt.Printf("test/normaldeps/mod2/hello: %s\n", hello.Hello())
}
