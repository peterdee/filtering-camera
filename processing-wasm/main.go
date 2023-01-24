package main

import "syscall/js"

func sum() js.Func {
	sumFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 2 {
			return "MISSING_ARGUMENTS"
		}
		var accumulator int
		for i := range arguments {
			accumulator += arguments[i].Int()
		}
		return accumulator
	})
	return sumFunction
}

func main() {
	js.Global().Set("getSum", sum())
	select {}
}
