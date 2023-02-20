import clamp from '../utilities/clamp';
import getCoordinates from '../utilities/get-coordinates';
import getGradientPoint from '../utilities/get-gradient-point';
import getPixel from '../utilities/get-pixel';

const LAPLACIAN: number[][] = [
  [-1, -1, -1],
  [-1, 8, -1],
  [-1, -1, -1],
];

export default function laplacian(imageData: ImageData): ImageData {
  const { data, height, width } = imageData;
  for (let i = 0; i < data.length; i += 4) {
    let averageSum = 0;
    const { x, y } = getCoordinates(i / 4, width);
    for (let m = 0; m < 3; m += 1) {
      for (let n = 0; n < 3; n += 1) {
        const k = getGradientPoint(x, m, width);
        const l = getGradientPoint(y, n, height);
        const pixel = getPixel(x + k, y + l, width);
        const average = Math.round((data[pixel] + data[pixel + 1] + data[pixel + 2]) / 3);
        averageSum += average * LAPLACIAN[m][n];
      }
    }
    const channel = 255 - clamp(averageSum, 0, 255);
    data[i] = channel;
    data[i + 1] = channel;
    data[i + 2] = channel;
  }
  return imageData;
}
