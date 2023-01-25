import clamp from '../utilities/clamp';
import type { GrayscaleType } from '../types/processing';

export default function bridge(
  imageData: ImageData,
  filter: string,
  threshold: number = 0,
  grayscaleType: GrayscaleType = 'luminosity',
): ImageData {
  const { data } = imageData;
  const processed = new Uint8ClampedArray(data.length);
  if (filter === 'binary') {
    const adjustedThreshold = clamp(threshold || 0, 0, 255);
    (window as any).binary(data, adjustedThreshold, processed);
  }
  if (filter === 'eightColors') {
    (window as any).eightColors(data, processed);
  }
  if (filter === 'grayscale') {
    (window as any).grayscale(data, grayscaleType, processed);
  }
  if (filter === 'sobel') {
    (window as any).sobel(
      data,
      imageData.width,
      imageData.height,
      processed,
    );
  }
  if (filter === 'solarize') {
    const adjustedThreshold = clamp(threshold || 0, 0, 255);
    (window as any).solarize(data, adjustedThreshold, processed);
  }
  imageData.data.set(processed);
  return imageData;
}
