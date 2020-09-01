// +build ignore

package main

import (
	"fmt"

	. "github.com/chai2010/go-python3"
)

func main() {
	Py_Initialize()
	defer Py_Finalize()

	pystr := PyUnicode_FromString("go-python3")
	str := PyUnicode_AsUTF8(pystr)

	fmt.Println("hello", str)
}
