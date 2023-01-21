// reference: https://github.com/julyskies/brille/blob/release/processing/solarize.go

import clamp from '../utilities/clamp';

export default function solarize(
  imageData: ImageData,
  threshold: number,
): ImageData {
  const { data } = imageData;
  const newImageData = new ImageData(imageData.width, imageData.height);
  const adjustedThreshold = clamp(threshold, 0, 255);
  for (let i = 0; i < data.length; i += 4) {
    newImageData.data[i] = data[i] <= adjustedThreshold
      ? 255 - data[i]
      : data[i];
    newImageData.data[i + 1] = data[i + 1] <= adjustedThreshold
      ? 255 - data[i + 1]
      : data[i + 1];
    newImageData.data[i + 2] = data[i + 2] <= adjustedThreshold
      ? 255 - data[i + 2]
      : data[i + 2];
    newImageData.data[i + 3] = data[i + 3];
  }
  return newImageData;
}
