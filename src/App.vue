<script setup lang="ts">
import {
  onMounted,
  reactive,
  ref,
} from 'vue';

import { FILTER_TYPES, type FilterType } from './constants';
import bridge from './bridge';
import canvasFilters from './processing-canvas';
import type { GrayscaleType, ProcessingType } from './types/processing';
import isMobile from './utilities/is-mobile';

interface ComponentState {
  ctx: CanvasRenderingContext2D | null;
  processingType: ProcessingType;
  selectedFilter: FilterType;
  selectedGrayscaleType: GrayscaleType;
  selectedThreshold: number;
  wasmLoaded: boolean;
}

const state = reactive<ComponentState>({
  ctx: null,
  processingType: 'canvas',
  selectedFilter: FILTER_TYPES[0],
  selectedGrayscaleType: 'luminosity',
  selectedThreshold: FILTER_TYPES[0].defaultThreshold || 0,
  wasmLoaded: false,
})

const canvasRef = ref<HTMLCanvasElement>();

const draw = (video: HTMLVideoElement): null | void => {
  const { ctx } = state;
  if (!ctx) {
    return null;
  }

  ctx.drawImage(video, 0, 0);

  const imageData = ((): ImageData => {
    const frame = ctx.getImageData(0, 0,  window.innerWidth, window.innerHeight);
    if (state.processingType === 'canvas') {
      if (state.selectedFilter.value === 'binary') {
        return canvasFilters.binary(frame, state.selectedThreshold);
      }
      if (state.selectedFilter.value === 'eightColors') {
        return canvasFilters.eightColors(frame);
      }
      if (state.selectedFilter.value === 'grayscale') {
        return canvasFilters.grayscale(frame, state.selectedGrayscaleType);
      }
      if (state.selectedFilter.value === 'sobel') {
        return canvasFilters.sobel(frame);
      }
      return canvasFilters.solarize(frame, state.selectedThreshold);
    }
    return bridge(
      frame,
      state.selectedFilter.value,
      state.selectedThreshold,
      state.selectedGrayscaleType,
    );
  })();

  ctx.putImageData(imageData, 0, 0);

  setTimeout(draw, 10, video);
}

const handleError = (error: MediaStreamError): void => {
  console.log('media device error', error);
}

const handleSuccess = (stream: MediaStream): null | void => {
  const video = document.createElement('video');
  video.onplay = (): null | void => draw(video);
  video.muted = true;
  video.playsInline = true;
  video.srcObject = stream;
  video.play();
}

const handleFilterSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  const [filter] = FILTER_TYPES.filter(
    (entry: FilterType): boolean => entry.value === value,
  );
  state.selectedFilter = filter;
}

const handleGrayscaleTypeSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  state.selectedGrayscaleType = value as GrayscaleType;
}

const handleProcessingTypeSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  state.processingType = value as ProcessingType;
}

const handleThresholdInput = (event: Event): void => {
  const { value } = event.target as HTMLInputElement;
  state.selectedThreshold = Number(value);
}

onMounted(async (): Promise<void> => {
  // load WASM
  const go = new (window as any).Go();
  try {
    const waInstance = await WebAssembly.instantiateStreaming(
      fetch('bin.wasm'),
      go.importObject,
    );
    go.run(waInstance.instance);
    state.wasmLoaded = true;
  } catch {
    state.wasmLoaded = false;
  }

  const height = window.innerHeight;
  const width = window.innerWidth;

  if (canvasRef.value) {
    canvasRef.value.height = height;
    canvasRef.value.width = width;
    state.ctx = canvasRef.value.getContext('2d', { willReadFrequently: true });
  }

  // TODO: better image ratios for mobile devices

  const constraints: MediaStreamConstraints = {
    audio: false,
    video: {
      height: {
        ideal: 720,
        max: 720,
        min: 600,
      },
      width: {
        ideal: 1280,
        max: 1280,
        min: 800,
      },
    },
  };
  if (isMobile()) {
    (constraints.video as MediaTrackConstraints).facingMode = { exact: 'environment' };
    (constraints.video as MediaTrackConstraints).height = {
      ideal: height,
      max: height,
      min: height,
    };
    (constraints.video as MediaTrackConstraints).width = {
      ideal: width,
      max: width,
      min: width,
    };
  }

  navigator.mediaDevices.getUserMedia(constraints)
    .then(handleSuccess)
    .catch(handleError);
});
</script>

<template>
  <div class="f j-center ai-center wrap">
    <canvas ref="canvasRef" />
    <div class="f d-col controls">
      <select
        :value="state.processingType"
        @change="handleProcessingTypeSelection"
      >
        <option value="canvas">
          Canvas
        </option>
        <option value="wasm">
          WASM
        </option>
      </select>
      <select
        class="mt-1"
        :value="state.selectedFilter.value"
        @change="handleFilterSelection"
      >
        <option
          v-for="filter in FILTER_TYPES"
          :value="filter.value"
        >
          {{ filter.name }}
        </option>
      </select>
      <select
        v-if="state.selectedFilter.isGrayscale"
        class="mt-1"
        :value="state.selectedGrayscaleType"
        @change="handleGrayscaleTypeSelection"
      >
        <option value="average">
          Average
        </option>
        <option value="luminosity">
          Luminosity
        </option>
      </select>
      <div
        v-if="state.selectedFilter.withThreshold"
        class="f d-col j-center mt-1"
      >
        <div class="f j-center ai-center">
          <span>
            {{ state.selectedFilter.minThreshold || 0 }}
          </span>
          <input
            class="mh-1"
            type="range"
            :max="state.selectedFilter.maxThreshold || 255"
            :min="state.selectedFilter.minThreshold || 0"
            :step="state.selectedFilter.step || 1"
            :value="state.selectedThreshold"
            @input="handleThresholdInput"
          />
          <span>
            {{ state.selectedFilter.maxThreshold || 255 }}
          </span>
        </div>
        <span class="t-center">
          {{ state.selectedThreshold }}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
canvas {
  max-height: 100vh;
  max-width: 100vw;
}
.controls {
  backdrop-filter: blur(var(--spacer-half));
  border-radius: var(--spacer-half);
  left: var(--spacer);
  padding: var(--spacer);
  position: absolute;
  top: var(--spacer);
  width: calc(var(--spacer) * 15);
}
.wrap {
  height: 100vh;
}
</style>
