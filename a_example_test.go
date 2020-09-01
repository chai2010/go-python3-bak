package python3_test

import (
	"fmt"

	"github.com/chai2010/go-python3"
)

func Example() {
	i, err := python3.Py_Main([]string{"go-python3", "./tests/test.py"})
	fmt.Println(i, err)

	// Output:
	// 0 <nil>
}
