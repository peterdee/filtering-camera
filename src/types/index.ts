export type GrayscaleType = 'average' | 'luminance';

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

export type ProcessingType = 'canvas' | 'wasm';
