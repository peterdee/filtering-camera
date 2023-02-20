export default function getPixel(
  x: number,
  y: number,
  width: number,
): number {
  return ((y * width) + x) * 4;
}
