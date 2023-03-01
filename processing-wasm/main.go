package main

import (
	"math"
	"runtime"
	"sync"
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

var LAPLACIAN = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
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

func clampMax[T float64 | int | uint](value T, max T) T {
	if value > max {
		return max
	}
	return value
}

func getCoordinates(pixel, width int) (int, int) {
	return pixel % width, int(math.Floor(float64(pixel) / float64(width)))
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

func getPixPerThread(pixLen, threads int) int {
	pixPerThreadRaw := float64(pixLen) / float64(threads)
	module := math.Mod(pixPerThreadRaw, 4.0)
	if module == 0 {
		return int(pixPerThreadRaw)
	}
	return int(pixPerThreadRaw + (float64(threads) - math.Mod(pixPerThreadRaw, 4.0)))
}

func gray(r, g, b uint8) uint8 {
	return uint8((int(r) + int(g) + int(b)) / 3)
}

func binary() js.Func {
	binaryFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 2 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		threshold := arguments[1].Int()
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			average := gray(buffer[i], buffer[i+1], buffer[i+2])
			var channel uint8 = 255
			if uint8(average) < uint8(threshold) {
				channel = 0
			}
			buffer[i], buffer[i+1], buffer[i+2] = channel, channel, channel
		}
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return binaryFunction
}

func colorInversion() js.Func {
	colorInversionFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 1 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			buffer[i] = 255 - buffer[i]
			buffer[i+1] = 255 - buffer[i+1]
			buffer[i+2] = 255 - buffer[i+2]
		}
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return colorInversionFunction
}

func eightColors() js.Func {
	eightColorsFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 1 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
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
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return eightColorsFunction
}

func grayscale() js.Func {
	grayscaleFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 2 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		grayscaleType := arguments[1].String()
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			var average uint8
			if grayscaleType == "average" {
				average = gray(buffer[i], buffer[i+1], buffer[i+2])
			} else {
				average = uint8(
					(float64(buffer[i])*0.21 + float64(buffer[i+1])*0.72 + float64(buffer[i+2])*0.07),
				)
			}
			buffer[i], buffer[i+1], buffer[i+2] = average, average, average
		}
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return grayscaleFunction
}

func laplacian() js.Func {
	laplacianFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 3 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		width := arguments[1].Int()
		height := arguments[2].Int()
		buffer := make([]uint8, pixelData.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, pixelData)
		for i := 0; i < len(buffer); i += 4 {
			averageSum := 0
			x, y := getCoordinates(i/4, width)
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					k := getGradientPoint(x, m, width)
					l := getGradientPoint(y, n, height)
					pixel := getPixel(x+k, y+l, width)
					average := gray(buffer[pixel], buffer[pixel+1], buffer[pixel+2])
					averageSum += int(average) * LAPLACIAN[m][n]
				}
			}
			channel := 255 - uint8(clamp(averageSum, 255, 0))
			buffer[i], buffer[i+1], buffer[i+2] = channel, channel, channel
		}
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return laplacianFunction
}

func sobel() js.Func {
	sobelFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 3 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		width := arguments[1].Int()
		height := arguments[2].Int()

		pixLen := pixelData.Get("byteLength").Int()
		buffer := make([]uint8, pixLen)
		result := make([]uint8, pixLen)
		js.CopyBytesToGo(buffer, pixelData)

		threads := runtime.NumCPU()
		pixPerThread := getPixPerThread(pixLen, threads)

		var wg sync.WaitGroup

		processing := func(thread int) {
			defer wg.Done()
			startIndex := pixPerThread * thread
			endIndex := clampMax(startIndex+pixPerThread, pixLen)
			for i := startIndex; i < endIndex; i += 4 {
				x, y := getCoordinates(i/4, width)
				gradientX, gradientY := 0, 0
				for m := 0; m < 3; m += 1 {
					for n := 0; n < 3; n += 1 {
						px := getPixel(
							clamp(x-(3/2-m), width-1, 0),
							clamp(y-(3/2-n), height-1, 0),
							width,
						)
						average := gray(buffer[px], buffer[px+1], buffer[px+2])
						gradientX += int(average) * SOBEL_HORIZONTAL[m][n]
						gradientY += int(average) * SOBEL_VERTICAL[m][n]
					}
				}
				channel := uint8(
					255 - clamp(
						math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
						255,
						0,
					),
				)
				result[i], result[i+1], result[i+2], result[i+3] = channel, channel, channel, buffer[i+3]
			}
		}

		for t := 0; t < threads; t += 1 {
			wg.Add(1)
			go processing(t)
		}

		wg.Wait()

		return js.CopyBytesToJS(pixelData, result)
	})
	return sobelFunction
}

func solarize() js.Func {
	solarizeFunction := js.FuncOf(func(this js.Value, arguments []js.Value) any {
		if len(arguments) < 2 {
			return "MISSING_ARGUMENTS"
		}
		pixelData := arguments[0]
		threshold := arguments[1].Int()
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
			buffer[i], buffer[i+1], buffer[i+2] = r, g, b
		}
		return js.CopyBytesToJS(pixelData, buffer)
	})
	return solarizeFunction
}

func main() {
	js.Global().Set("binary", binary())
	js.Global().Set("colorInversion", colorInversion())
	js.Global().Set("eightColors", eightColors())
	js.Global().Set("grayscale", grayscale())
	js.Global().Set("laplacian", laplacian())
	js.Global().Set("sobel", sobel())
	js.Global().Set("solarize", solarize())
	select {}
}
