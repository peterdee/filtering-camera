// reference: https://github.com/julyskies/brille/blob/release/processing/grayscale.go

import clamp from '../utilities/clamp';
import type { GrayscaleType } from '../types';

export default function grayscale(
  imageData: ImageData,
  grayscaleType: GrayscaleType = 'average',
): ImageData {
  const { data } = imageData;
  const processed = new Uint8ClampedArray(data.length);
  for (let i = 0; i < data.length; i += 4) {
    if (grayscaleType === 'average') {
      const average = Math.round((data[i] + data[i + 1] + data[i + 2]) / 3);
      processed[i] = average;
      processed[i + 1] = average;
      processed[i + 2] = average;
    } else {
      const adjustedR = Math.round(data[i] * 0.21);
      const adjustedG = Math.round(data[i + 1] * 0.72);
      const adjustedB = Math.round(data[i + 2] * 0.07);
      const average = clamp(adjustedR + adjustedG + adjustedB, 0, 255);
      processed[i] = average;
      processed[i + 1] = average;
      processed[i + 2] = average;
    }
    processed[i + 3] = data[i + 3];
  }
  imageData.data.set(processed);
  return imageData;
}
