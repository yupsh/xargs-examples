package xargs_test

import (
	"fmt"
	"os"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/echo"
	"github.com/yupsh/xargs"
)

func ExampleXargs() {
	// echo "one\ntwo\nthree" | xargs echo
	if err := gloo.Run(Pipeline(
		echo.Echo("one\ntwo\nthree"),
		xargs.Xargs(echo.Echo()),
	)); err != nil {
		fmt.Fprintf(os.Stderr, "xargs: %v\n", err)
	}
	// Output:
	// one two three
}
