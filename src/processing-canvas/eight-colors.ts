// reference: https://github.com/julyskies/brille/blob/release/processing/eight-colors.go

interface Color {
  b: number;
  g: number;
  r: number;
}

const COLORS: Color[] = [
  { b: 255, r: 0, g: 0 },
  { b: 0, r: 255, g: 0 },
  { b: 0, r: 0, g: 255 },
  { b: 255, r: 255, g: 0 },
  { b: 255, r: 0, g: 255 },
  { b: 0, r: 255, g: 255 },
  { b: 255, r: 255, g: 255 },
  { b: 0, r: 0, g: 0 },
];

export default function eightColors(imageData: ImageData): ImageData {
  const { data } = imageData;
  for (let i = 0; i < data.length; i += 4) {
    let minDelta = 195076;
    let selectedColor: Color = COLORS[0];
    for (let j = 0; j < COLORS.length; j += 1) {
      const indexColor = COLORS[j];
      const rDifference = (data[i] - indexColor.r) ** 2;
      const gDifference = (data[i + 1] - indexColor.g) ** 2;
      const bDifference = (data[i + 2] - indexColor.b) ** 2;
      const delta = rDifference + gDifference + bDifference;
      if (delta < minDelta) {
        minDelta = delta;
        selectedColor = indexColor;
      }
    }
    data[i] = selectedColor.r;
    data[i + 1] = selectedColor.g;
    data[i + 2] = selectedColor.b;
  }
  return imageData;
}
