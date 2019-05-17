package main

import (
	"syscall/js"
	"testing"
)

var tests = []struct {
	name   string
	input  string
	output string
	err    bool
}{
	{
		name: "success1",
		input: "" + // 見易さのため
			`package main` + "\n" +
			"\n" +
			`import "fmt"` + "\n" +
			"\n" +
			`func main() {` + "\n" +
			`	fmt.Println("Hello, 世界")` + "\n" +
			`}` + "\n",
		output: "" +
			`package main` + "\n" +
			"\n" +
			`import "fmt"` + "\n" +
			"\n" +
			`func main() {` + "\n" +
			`	fmt.Println("Hello, 世界")` + "\n" +
			`}` + "\n",
		err: false,
	},
	{
		name: "success2",
		input: "" +
			`package main` + "\n" +
			`import "fmt"` + "\n" +
			`func main() {` + "\n" +
			`	fmt.Println("Hello, 世界")` + "\n" +
			`}` + "\n",
		output: "" +
			`package main` + "\n" +
			"\n" +
			`import "fmt"` + "\n" +
			"\n" +
			`func main() {` + "\n" +
			`	fmt.Println("Hello, 世界")` + "\n" +
			`}` + "\n",
		err: false,
	},
	{
		name: "failure1",
		input: "" +
			`package main` + "\n" +
			"\n" +
			`import "fmt"` + "\n" +
			"\n" +
			`func main() {` + "\n" +
			`	fmt.Println("Hello, 世界")` + "\n",
		output: "6:31: expected '}', found 'EOF'",
		err:    true,
	},
	{
		name: "failure2",
		input: "" +
			`package main` + "\n" +
			"\n" +
			`import "fmt"` + "\n" +
			"\n" +
			`func main() {` + "\n",
		output: "5:15: expected '}', found 'EOF'",
		err:    true,
	},
}

func TestRegisterCallback(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("got err = %v, want err = syscall/js: Value.Call: property codeFormat is not a function, got undefined", err)
			}
		}()

		js.Global().Call("codeFormat", "hoge")
	})

	t.Run("Success", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("got err = %v, want err = <nil>", err)
			}
		}()

		registerCallback()
		js.Global().Call("codeFormat", "hoge")
	})
}

func TestCodeFormat(t *testing.T) {
	this := js.ValueOf(nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := codeFormat(this, []js.Value{js.ValueOf(test.input)}).(map[string]interface{})
			if got["err"] != test.err {
				t.Errorf("got err = %v, want err = %v", got["err"], test.err)
			} else if got["output"] != test.output {
				t.Errorf("got output = %v, want output = %v", got["output"], test.output)
			}
		})
	}
}

func TestJsCallCodeFormat(t *testing.T) {
	registerCallback()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := js.Global().Call("codeFormat", test.input)
			if got.Get("err").Bool() != test.err {
				t.Errorf("got err = %v, want err = %v", got.Get("err").Bool(), test.err)
			} else if got.Get("output").String() != test.output {
				t.Errorf("got output = %v, want output = %v", got.Get("output").String(), test.output)
			}
		})
	}
}
