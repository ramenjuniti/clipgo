package main

import (
	"go/format"
	"syscall/js"
)

type error interface {
	Error() string
}

func codeFormat(this js.Value, args []js.Value) interface{} {
	code := args[0].String()
	newCode, err := format.Source([]byte(code))
	if err != nil {
		return map[string]interface{}{"output": err.Error(), "err": true}
	}
	return map[string]interface{}{"output": string(newCode), "err": false}
}

func registerCallback() {
	js.Global().Set("codeFormat", js.FuncOf(codeFormat))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallback()
	<-c
}
