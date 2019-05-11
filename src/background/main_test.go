package main

import (
	"syscall/js"
	"testing"
)

func TestRegisterCallbackFailed(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("got %v, want syscall/js: Value.Call: property codeFormat is not a function, got undefined", err)
		}
	}()

	js.Global().Call("codeFormat", "hoge")
}

func TestRegisterCallbackSuccess(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("got %v, want nil", err)
		}
	}()

	registerCallback()
	js.Global().Call("codeFormat", "hoge")
}

func TestCodeFormatFailed(t *testing.T) {
	this := js.ValueOf(nil)
	args := []js.Value{
		js.ValueOf(
			`
			package main

			import "fmt"
			
			func main() {
				fmt.Println("Hello, 世界")
			`,
		),
	}

	result := codeFormat(this, args).(map[string]interface{})

	if _, ok := result["newCode"]; ok {
		t.Errorf("got newCode, want err")
	}
}

func TestCodeFormatSuccess(t *testing.T) {
	this := js.ValueOf(nil)
	args := []js.Value{
		js.ValueOf(
			`
			package main

			import "fmt"
			
			func main() {
				fmt.Println("Hello, 世界")
			}
			`,
		),
	}

	result := codeFormat(this, args).(map[string]interface{})

	if _, ok := result["err"]; ok {
		t.Errorf("got err, want newCode")
	}
}
