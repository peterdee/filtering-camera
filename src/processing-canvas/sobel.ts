// reference: https://github.com/julyskies/brille/blob/release/processing/solarize.go

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
  const newImageData = new ImageData(imageData.width, imageData.height);
  for (let i = 0; i < data.length; i += 4) {
    // newImageData.data[i] = data[i] <= adjustedThreshold
    //   ? 255 - data[i]
    //   : data[i];
    // newImageData.data[i + 1] = data[i + 1] <= adjustedThreshold
    //   ? 255 - data[i + 1]
    //   : data[i + 1];
    // newImageData.data[i + 2] = data[i + 2] <= adjustedThreshold
    //   ? 255 - data[i + 2]
    //   : data[i + 2];
    // newImageData.data[i + 3] = data[i + 3];
  }
  return newImageData;
}
