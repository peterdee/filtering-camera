// reference: https://github.com/julyskies/brille
package main

import (
	"math"
	"syscall/js"
)

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

var SOBEL_HORIZONTAL = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var SOBEL_VERTICAL = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func clamp[T float64 | int | uint](value T, max T, min T) T {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func getCoordinates(pixel, width int) (int, int) {
	x := pixel % width
	y := math.Floor(float64(pixel) / float64(width))
	return x, int(y)
}

func getGradientPoint(axisValue, shift, axisLength int) int {
	if axisValue+shift >= axisLength {
		return axisLength - axisValue - 1
	}
	return shift
}

func getPixel(x, y, width int) int {
	return ((y * width) + x) * 4
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

func grayscale() js.Func {
	grayscaleFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 3 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		grayscaleType := arguments[1].String()
		result := arguments[2]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			if grayscaleType == "average" {
				average := gray(buffer[i], buffer[i+1], buffer[i+2])
				buffer[i] = average
				buffer[i+1] = average
				buffer[i+2] = average
			} else {
				average := uint8(
					(float64(buffer[i])*0.21 + float64(buffer[i+1])*0.72 + float64(buffer[i+2])*0.07),
				)
				buffer[i] = average
				buffer[i+1] = average
				buffer[i+2] = average
			}
		}
		return js.CopyBytesToJS(result, buffer)
	})
	return grayscaleFunction
}

func sobel() js.Func {
	sobelFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 4 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		width := arguments[1].Int()
		height := arguments[2].Int()
		result := arguments[3]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			x, y := getCoordinates(i/4, width)
			gradientX := 0
			gradientY := 0
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					k := getGradientPoint(x, m, width)
					l := getGradientPoint(y, n, height)
					pixel := getPixel(x+k, y+l, width)
					average := gray(buffer[pixel], buffer[pixel+1], buffer[pixel+2])
					gradientX += int(average) * SOBEL_HORIZONTAL[m][n]
					gradientY += int(average) * SOBEL_VERTICAL[m][n]
				}
			}
			colorCode := uint8(255 - clamp(
				math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
				255,
				0,
			))
			buffer[i] = colorCode
			buffer[i+1] = colorCode
			buffer[i+2] = colorCode
		}
		return js.CopyBytesToJS(result, buffer)
	})
	return sobelFunction
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
	js.Global().Set("grayscale", grayscale())
	js.Global().Set("sobel", sobel())
	js.Global().Set("solarize", solarize())
	select {}
}
