package main

import (
	"syscall/js"
	"testing"
)

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
		t.Fatal("failed test")
	}
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
		t.Fatal("failed test")
	}
}