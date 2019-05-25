package main

import (
	"io/ioutil"
	"path/filepath"
	"syscall/js"
	"testing"
)

var tests = []struct {
	name string
	err  bool
}{
	{
		name: "success1",
		err:  false,
	},
	{
		name: "success2",
		err:  false,
	},
	{
		name: "failure1",
		err:  true,
	},
	{
		name: "failure2",
		err:  true,
	},
}

func TestRegisterCallback(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("got err: %v, want err: syscall/js: Value.Call: property formatter is not a function, got undefined", err)
			}
		}()

		js.Global().Call("formatter", "hoge")
	})

	t.Run("Success", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("got err: %v, want err: <nil>", err)
			}
		}()

		registerCallback()
		js.Global().Call("formatter", "hoge")
	})
}

func TestFormatter(t *testing.T) {
	this := js.ValueOf(nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := filepath.Join("testdata", "formatter", "input", test.name+".golden")
			output := filepath.Join("testdata", "formatter", "output", test.name+".golden")
			inbuf, _ := ioutil.ReadFile(input)
			outbuf, _ := ioutil.ReadFile(output)
			got := formatter(this, []js.Value{js.ValueOf(string(inbuf))}).(map[string]interface{})

			if got["err"] != test.err {
				t.Errorf("got err = %v, want err = %v", got["err"], test.err)
			} else if got["output"] != string(outbuf) {
				t.Errorf("got output: %v, want output: %v", got["output"], string(outbuf))
			}
		})
	}
}

func TestJsCallFormatter(t *testing.T) {
	registerCallback()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := filepath.Join("testdata", "formatter", "input", test.name+".golden")
			output := filepath.Join("testdata", "formatter", "output", test.name+".golden")
			inbuf, _ := ioutil.ReadFile(input)
			outbuf, _ := ioutil.ReadFile(output)
			got := js.Global().Call("formatter", string(inbuf))

			if got.Get("err").Bool() != test.err {
				t.Errorf("got err = %v, want err = %v", got.Get("err").Bool(), test.err)
			} else if got.Get("output").String() != string(outbuf) {
				t.Errorf("got output: %v, want output: %v", got.Get("output").String(), string(outbuf))
			}
		})
	}
}
