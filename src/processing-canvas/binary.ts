// reference: https://github.com/julyskies/brille/blob/release/processing/binary.go

import clamp from '../utilities/clamp';

export default function binary(
  imageData: ImageData,
  threshold: number,
): ImageData {
  const { data } = imageData;
  const processed = new Uint8ClampedArray(data.length);
  const adjustedThreshold = clamp(threshold, 0, 255);
  for (let i = 0; i < data.length; i += 4) {
    const average = Math.round((data[i] + data[i + 1] + data[i + 2]) / 3);
    let partial = 0;
    if (average > adjustedThreshold) {
      partial = 255;
    }
    processed[i] = partial;
    processed[i + 1] = partial;
    processed[i + 2] = partial;
    processed[i + 3] = data[i + 3];
  }
  imageData.data.set(processed);
  return imageData;
}
