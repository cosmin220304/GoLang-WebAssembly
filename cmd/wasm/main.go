package main

import (
	"fmt"
	"syscall/js"
)

func getFullName(this js.Value, p []js.Value) interface{} {
	if len(p) < 2 {
		return "Invalid number of arguments passed"
	}
	return p[0].String() + " " + p[1].String()
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("getFullName", js.FuncOf(getFullName))
}
