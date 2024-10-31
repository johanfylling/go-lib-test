package say

import (
	"fmt"

	v0 "github.com/johanfylling/go-lib-test/say"
)

func Hello(n string) {
	fmt.Printf("Hello %s", n)
	v0.Hello(n)
}
