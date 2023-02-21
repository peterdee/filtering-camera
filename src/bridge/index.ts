import clamp from '../utilities/clamp';
import type { GrayscaleType } from '../types';

export default function bridge(
  imageData: ImageData,
  filter: string,
  threshold: number = 0,
  grayscaleType: GrayscaleType = 'luminance',
): ImageData {
  const { data } = imageData;
  if (filter === 'binary') {
    const adjustedThreshold = clamp(threshold || 0, 0, 255);
    (window as any).binary(data, adjustedThreshold);
  }
  if (filter === 'colorInversion') {
    (window as any).colorInversion(data);
  }
  if (filter === 'eightColors') {
    (window as any).eightColors(data);
  }
  if (filter === 'grayscale') {
    (window as any).grayscale(data, grayscaleType);
  }
  if (filter === 'laplacian') {
    (window as any).laplacian(
      data,
      imageData.width,
      imageData.height,
    );
  }
  if (filter === 'sobel') {
    (window as any).sobel(
      data,
      imageData.width,
      imageData.height,
    );
  }
  if (filter === 'solarize') {
    const adjustedThreshold = clamp(threshold || 0, 0, 255);
    (window as any).solarize(data, adjustedThreshold);
  }
  return imageData;
}
