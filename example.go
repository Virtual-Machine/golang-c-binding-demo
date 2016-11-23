package main

/*
#include "example_lib.h"
#cgo LDFLAGS: -lexample
*/
import "C"

import "fmt"

func main() {

	fmt.Println(C.GoString(C.helloWorld()))
	fmt.Println(C.addSomeNumbers(2, 45))
}
