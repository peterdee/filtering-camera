import type { GrayscaleType } from '../types/processing';

export interface FilterType {
  defaultGrayscaleType?: GrayscaleType;
  defaultThreshold?: number;
  isGrayscale: boolean;
  maxThreshold?: number;
  minThreshold?: number;
  name: string;
  step?: number;
  value: string;
  withThreshold: boolean;
}

export const FILTER_TYPES: FilterType[] = [
  {
    defaultThreshold: 122,
    isGrayscale: false,
    maxThreshold: 255,
    minThreshold: 0,
    name: 'Binary',
    step: 1,
    value: 'binary',
    withThreshold: true,
  },
  {
    isGrayscale: false,
    name: 'Eight colors',
    value: 'eightColors',
    withThreshold: false,
  },
  {
    defaultGrayscaleType: 'luminosity',
    isGrayscale: true,
    name: 'Grayscale',
    value: 'grayscale',
    withThreshold: false,
  },
  {
    isGrayscale: false,
    name: 'Sobel filter',
    value: 'sobel',
    withThreshold: false,
  },
  {
    defaultThreshold: 105,
    isGrayscale: false,
    maxThreshold: 255,
    minThreshold: 0,
    name: 'Solarize',
    step: 1,
    value: 'solarize',
    withThreshold: true,
  },
];
