import clamp from '../utilities/clamp';

export default function solarize(
  imageData: ImageData,
  threshold: number,
): ImageData {
  const { data } = imageData;
  const adjustedThreshold = clamp(threshold, 0, 255);
  for (let i = 0; i < data.length; i += 4) {
    data[i] = data[i] <= adjustedThreshold
      ? 255 - data[i]
      : data[i];
    data[i + 1] = data[i + 1] <= adjustedThreshold
      ? 255 - data[i + 1]
      : data[i + 1];
    data[i + 2] = data[i + 2] <= adjustedThreshold
      ? 255 - data[i + 2]
      : data[i + 2];
  }
  return imageData;
}
