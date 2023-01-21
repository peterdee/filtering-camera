// reference: https://github.com/julyskies/brille/blob/release/processing/binary.go

import clamp from '../utilities/clamp';

export default function binary(
  imageData: ImageData,
  threshold: number,
): ImageData {
  const { data } = imageData;
  const newImageData = new ImageData(imageData.width, imageData.height);
  const adjustedThreshold = clamp(threshold, 0, 255);
  for (let i = 0; i < data.length; i += 4) {
    const average = Math.round((data[i] + data[i + 1] + data[i + 2]) / 3);
    let partial = 0;
    if (average > adjustedThreshold) {
      partial = 255;
    }
    newImageData.data[i] = partial;
    newImageData.data[i + 1] = partial;
    newImageData.data[i + 2] = partial;
    newImageData.data[i + 3] = data[i + 3];
  }
  return newImageData;
}
