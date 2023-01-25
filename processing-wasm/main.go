// reference: https://github.com/julyskies/brille
package main

import "syscall/js"

type Color struct {
	R, G, B int
}

var COLORS = [8]Color{
	{255, 0, 0},
	{0, 255, 0},
	{0, 0, 255},
	{255, 255, 0},
	{255, 0, 255},
	{0, 255, 255},
	{255, 255, 255},
	{0, 0, 0},
}

func gray(r, g, b uint8) uint8 {
	average := (int(r) + int(g) + int(b)) / 3
	return uint8(average)
}

func binary() js.Func {
	binaryFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 3 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		threshold := arguments[1].Int()
		result := arguments[2]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			average := gray(buffer[i], buffer[i+1], buffer[i+2])
			var colorCode uint8 = 255
			if uint8(average) < uint8(threshold) {
				colorCode = 0
			}
			buffer[i] = colorCode
			buffer[i+1] = colorCode
			buffer[i+2] = colorCode
		}
		return js.CopyBytesToJS(result, buffer)
	})
	return binaryFunction
}

func eightColors() js.Func {
	eightColorsFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 2 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		result := arguments[1]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			r, g, b := buffer[i], buffer[i+1], buffer[i+2]
			minDelta := 195076
			var selectedColor Color
			for i := range COLORS {
				indexColor := COLORS[i]
				rDifference := int(r) - indexColor.R
				gDifference := int(g) - indexColor.G
				bDifference := int(b) - indexColor.B
				delta := rDifference*rDifference + gDifference*gDifference + bDifference*bDifference
				if delta < minDelta {
					minDelta = delta
					selectedColor = indexColor
				}
			}
			buffer[i] = uint8(selectedColor.R)
			buffer[i+1] = uint8(selectedColor.G)
			buffer[i+2] = uint8(selectedColor.B)
		}
		return js.CopyBytesToJS(result, buffer)
	})
	return eightColorsFunction
}

func solarize() js.Func {
	solarizeFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 3 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		threshold := arguments[1].Int()
		result := arguments[2]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			r := buffer[i]
			if r <= uint8(threshold) {
				r = 255 - r
			}
			g := buffer[i+1]
			if g <= uint8(threshold) {
				g = 255 - g
			}
			b := buffer[i+2]
			if b <= uint8(threshold) {
				b = 255 - b
			}
			buffer[i] = r
			buffer[i+1] = g
			buffer[i+2] = b
		}
		return js.CopyBytesToJS(result, buffer)
	})
	return solarizeFunction
}

func main() {
	js.Global().Set("binary", binary())
	js.Global().Set("eightColors", eightColors())
	js.Global().Set("solarize", solarize())
	select {}
}
