export interface Point {
  x: number;
  y: number;
}

export default function getCoordinates(
  pixel: number,
  width: number,
): Point {
  return {
    x: pixel % width,
    y: Math.floor(pixel / width),
  };
}
