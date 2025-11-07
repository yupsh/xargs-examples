package xargs_test

import (
	"fmt"
	"os"

	"github.com/yupsh/echo"
	yup "github.com/gloo-foo/framework"
	. "github.com/gloo-foo/pipe"
	"github.com/yupsh/xargs"
)

func ExampleXargs() {
	// echo "one\ntwo\nthree" | xargs echo
	if err := yup.Run(Pipeline(
		echo.Echo("one\ntwo\nthree"),
		xargs.Xargs(echo.Echo()),
	)); err != nil {
		fmt.Fprintf(os.Stderr, "xargs: %v\n", err)
	}
	// Output:
	// one two three
}
