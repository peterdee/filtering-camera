// reference: https://github.com/julyskies/brille/blob/release/processing/grayscale.go

import clamp from '../utilities/clamp';
import type { GrayscaleType } from '../types/processing';

export default function grayscale(
  imageData: ImageData,
  grayscaleType: GrayscaleType = 'average',
): ImageData {
  const { data } = imageData;
  for (let i = 0; i < data.length; i += 4) {
    if (grayscaleType === 'average') {
      const average = Math.round((data[i] + data[i + 1] + data[i + 2]) / 3);
      data[i] = average;
      data[i + 1] = average;
      data[i + 2] = average;
    } else {
      const adjustedR = Math.round(data[i] * 0.21);
      const adjustedG = Math.round(data[i + 1] * 0.72);
      const adjustedB = Math.round(data[i + 2] * 0.07);
      const average = clamp(adjustedR + adjustedG + adjustedB, 0, 255);
      data[i] = average;
      data[i + 1] = average;
      data[i + 2] = average;
    }
  }
  imageData.data.set(data);
  return imageData;
}
