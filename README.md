# go-python3: Python C-API 绑定

http://docs.python.org/c-api/index.html

安装环境(macOS):

1. `brew install pkg-config`
2. `brew install python@3.7`
3. `brew install intltool`

## Python命令行

查看Python版本:

```
$ go run ./examples/python3 -V
Python 3.7.8
```

执行 Python 文件:

```
$ go run ./examples/python3 ./examples/python3/hello.py
hello go-python3
```

打开Python交互环境:

```shell
$ go run ./examples/python3
Python 3.7.8 (default, Jul  8 2020, 14:18:28) 
[Clang 11.0.3 (clang-1103.0.32.62)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> print(123)
123
>>> quit()
```

执行 Python 脚本字符串:

```shell
$ go run ./examples/python3 -c "import sys;print(sys.path)"
$ PYTHONPATH="" go run ./examples/python3 -c "import sys;print(sys.path)"
```

## 例子: Go代码嵌入Python

```go
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
```

```
$ go run hello.go 
hello go-python3
```

## 例子: 调用print内置函数

```go
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
```

```
$ go run ./examples/print
hi print
```

## 补充说明

- Macos: 依赖 `/usr/local/opt/python@3.7/Frameworks/Python.framework/Versions/3.7/Python` 文件. 如果缺少, 可以用替代 `./libgo-python3.7.dylib`.
- Linux: 需要自己构建 `libpython3.7.so`, 然后改名为 `./libgo-python3.7.so`.

<!--
TODO: Go通过CGO为Py编写扩展
-->
