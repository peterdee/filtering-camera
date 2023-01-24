package main

import (
	"fmt"
	"syscall/js"
)

func binary() js.Func {
	binaryFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 4 {
			return "MISSING_ARGUMENTS"
		}
		imageData := arguments[0].String()
		threshold := arguments[1].Int()
		width := arguments[2].Int()
		height := arguments[3].Int()

		return fmt.Sprintf("%d %d %d %s", threshold, width, height, imageData)
	})
	return binaryFunction
}

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
	js.Global().Set("binary", binary())
	js.Global().Set("getSum", sum())
	select {}
}
