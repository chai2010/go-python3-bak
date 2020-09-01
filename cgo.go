package python3

/*
#cgo linux pkg-config: python3

#cgo darwin pkg-config: python-3.7
#cgo darwin LDFLAGS: -lintl

#cgo windows CFLAGS: -I./internal/python3.7/include/python3.7m
#cgo windows LDFLAGS: -L${SRCDIR} -lpython37
*/
import "C"
