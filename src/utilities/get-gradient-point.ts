export default function getGradientPoint(
  axisValue: number,
  shift: number,
  axisLength: number,
): number {
  if (axisValue + shift >= axisLength) {
    return axisLength - axisValue - 1;
  }
  return shift;
}
