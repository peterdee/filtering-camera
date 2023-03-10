import getCoordinates from '../utilities/get-coordinates';
import getGradientPoint from '../utilities/get-gradient-point';
import getPixel from '../utilities/get-pixel';

const HORIZONTAL: number[][] = [
  [-1, 0, 1],
  [-2, 0, 2],
  [-1, 0, 1],
];

const VERTICAL: number[][] = [
  [1, 2, 1],
  [0, 0, 0],
  [-1, -2, -1],
];

export default function sobel(imageData: ImageData): ImageData {
  const { data, height, width } = imageData;
  for (let i = 0; i < data.length; i += 4) {
    const { x, y } = getCoordinates(i / 4, width);
    let gradientX = 0;
    let gradientY = 0;
    for (let m = 0; m < 3; m += 1) {
      for (let n = 0; n < 3; n += 1) {
        const k = getGradientPoint(x, m, width);
        const l = getGradientPoint(y, n, height);
        const pixel = getPixel(x + k, y + l, width);
        const average = Math.round((data[pixel] + data[pixel + 1] + data[pixel + 2]) / 3);
        gradientX += average * HORIZONTAL[m][n];
        gradientY += average * VERTICAL[m][n];
      }
    }
    const colorCode = 255 - Math.sqrt(gradientX ** 2 + gradientY ** 2);
    data[i] = colorCode;
    data[i + 1] = colorCode;
    data[i + 2] = colorCode;
  }
  return imageData;
}
