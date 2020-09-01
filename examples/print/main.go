package main

import (
	. "github.com/chai2010/go-python3"
)

func main() {
	Py_Initialize()
	defer Py_Finalize()

	builtins := PyEval_GetBuiltins()
	builtins_print := PyDict_GetItemString(builtins, "print")

	arg := PyUnicode_FromString("hi print")
	defer arg.DecRef()

	code := builtins_print.CallFunctionObjArgs(arg)
	defer code.DecRef()
}
