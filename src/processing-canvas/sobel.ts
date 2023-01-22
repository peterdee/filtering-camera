// reference: https://github.com/julyskies/brille/blob/release/processing/sobel-filter.go

interface Point {
  x: number;
  y: number;
}

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

const getCoordinates = (
  pixel: number,
  width: number,
): Point => ({
  x: pixel % width,
  y: Math.floor(pixel / width),
});

const getGradientPoint = (
  axisValue: number,
  shift: number,
  axisLength: number,
): number => {
  if (axisValue + shift >= axisLength) {
    return axisLength - axisValue - 1;
  }
  return shift;
};

const getPixel = (x: number, y: number, width: number) => (y * width) + x;

export default function sobel(imageData: ImageData): ImageData {
  const { data, height, width } = imageData;
  const newImageData = new ImageData(imageData.width, imageData.height);
  for (let i = 0; i < data.length; i += 4) {
    const { x, y } = getCoordinates(Math.round(i / 4), width);
    let gradientX = 0;
    let gradientY = 0;
    for (let m = 0; m < 3; m += 1) {
      for (let n = 0; n < 3; n += 1) {
        const k = getGradientPoint(x, m, width);
        const l = getGradientPoint(y, n, height);
        const pixel = getPixel(x + k, y + l, width);
        // console.log('x', x, 'y', y, 'm', m, 'n', n, 'k', k, 'l', l, 'pix', pixel);
        const average = Math.round((data[pixel] + data[pixel + 1] + data[pixel + 2]) / 3);
        gradientX += average * HORIZONTAL[m][n];
        gradientY += average * VERTICAL[m][n];
      }
    }

    const colorCode = 255 - Math.sqrt(gradientX ** 2 + gradientY ** 2);
    newImageData.data[i] = colorCode;
    newImageData.data[i + 1] = colorCode;
    newImageData.data[i + 2] = colorCode;
    newImageData.data[i + 3] = data[i + 3];
  }
  return newImageData;
}
