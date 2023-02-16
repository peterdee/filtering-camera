<script setup lang="ts">
import {
  onBeforeUnmount,
  onMounted,
  reactive,
  ref,
} from 'vue';

import bridge from './bridge';
import canvasFilters from './processing-canvas';
import type {
  FilterType,
  GrayscaleType,
  ProcessingType,
} from './types';
import { FILTER_TYPES } from './constants';
import FPSCounter from './components/FPSCounterComponent.vue';
import isMobile from './utilities/is-mobile';
import OptionsModal from './components/OptionsModalComponent.vue';
import OptionsButton from './components/OptionsButtonComponent.vue';

interface ComponentState {
  ctx: CanvasRenderingContext2D | null;
  fpsCount: number;
  frameTime: number[];
  isMobile: boolean;
  processingType: ProcessingType;
  selectedFilter: FilterType;
  selectedGrayscaleType: GrayscaleType;
  selectedThreshold: number;
  showOptionsModal: boolean;
  wasmLoaded: boolean;
}

const state = reactive<ComponentState>({
  ctx: null,
  fpsCount: 0,
  frameTime: [],
  isMobile: true,
  processingType: 'canvas',
  selectedFilter: FILTER_TYPES[0],
  selectedGrayscaleType: 'luminosity',
  selectedThreshold: FILTER_TYPES[0].defaultThreshold || 0,
  showOptionsModal: false,
  wasmLoaded: false,
});

const canvasRef = ref<HTMLCanvasElement>();

const draw = (video: HTMLVideoElement): null | NodeJS.Timeout | void => {
  const { ctx } = state;
  if (!ctx) {
    return null;
  }

  ctx.drawImage(video, 0, 0);
  state.frameTime.push(Date.now());
  if (state.frameTime.length === 10) {
    const diff = state.frameTime[9] - state.frameTime[0];
    state.fpsCount = Math.round(10000 / diff);
    state.frameTime = [];
  }

  const imageData = ((): ImageData => {
    const frame = ctx.getImageData(0, 0, window.innerWidth, window.innerHeight);
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

  return setTimeout(draw, 10, video);
};

const handleError = (error: MediaStreamError): void => {
  console.log('media device error', error);
};

const handleSuccess = (stream: MediaStream): null | void => {
  console.log(stream.getVideoTracks()[0].getSettings(), stream);
  const video = document.createElement('video');
  video.onplay = (): null | NodeJS.Timeout | void => draw(video);
  video.muted = true;
  video.playsInline = true;
  video.srcObject = stream;
  video.play();
};

const handleFilterSelection = (filter: FilterType): void => {
  state.selectedFilter = filter;
};

const handleGrayscaleTypeSelection = (type: GrayscaleType): void => {
  state.selectedGrayscaleType = type;
};

const handleProcessingTypeSelection = (type: ProcessingType): void => {
  state.processingType = type;
};

const handleResize = (): void => {
  if (canvasRef.value) {
    canvasRef.value.height = window.innerHeight;
    canvasRef.value.width = window.innerWidth;
  }
};

const handleThresholdInput = (value: string): void => {
  state.selectedThreshold = Number(value);
};

const toggleOptionsModal = (): void => {
  state.showOptionsModal = !state.showOptionsModal;
};

onBeforeUnmount((): void => {
  window.removeEventListener('resize', handleResize);
});

onMounted(async (): Promise<void> => {
  // set correct favicon
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    const faviconLink = document.querySelector<HTMLLinkElement>(`link[rel~='${'icon'}']`);
    if (faviconLink) {
      faviconLink.href = 'favicon-light.png';
    }
  }

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

  // check if this is a mobile device
  const mobile = isMobile();
  state.isMobile = mobile;

  window.addEventListener('resize', handleResize);

  const height = window.innerHeight;
  const width = window.innerWidth;

  if (canvasRef.value) {
    canvasRef.value.height = height;
    canvasRef.value.width = width;
    const ctx = canvasRef.value.getContext('2d', { willReadFrequently: true }) || null;
    if (ctx) {
      state.ctx = ctx;

      // flip image horizontally
      // ctx.translate(width, 0);
      // ctx.scale(-1, 1);
    }
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
  if (mobile) {
    (constraints.video as MediaTrackConstraints).facingMode = { exact: 'environment' };
    (constraints.video as MediaTrackConstraints).height = {
      // ideal: height,
      max: height,
      // min: height,
    };
    (constraints.video as MediaTrackConstraints).width = {
      // ideal: width,
      max: width,
      // min: width,
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
    <FPSCounter
      v-if="!state.showOptionsModal"
      :count="state.fpsCount"
      :is-mobile="state.isMobile"
    />
    <OptionsButton
      v-if="!state.showOptionsModal"
      :is-mobile="state.isMobile"
      @handle-click="toggleOptionsModal"
    />
    <OptionsModal
      v-if="state.showOptionsModal"
      :is-mobile="state.isMobile"
      :processing-type="state.processingType"
      :selected-filter="state.selectedFilter"
      :selected-grayscale-type="state.selectedGrayscaleType"
      :selected-threshold="state.selectedThreshold"
      :wasm-loaded="state.wasmLoaded"
      @handle-filter="handleFilterSelection"
      @handle-grayscale-type="handleGrayscaleTypeSelection"
      @handle-processing-type="handleProcessingTypeSelection"
      @handle-threshold="handleThresholdInput"
      @toggle-modal="toggleOptionsModal"
    />
  </div>
</template>

<style scoped>
canvas {
  height: 100%;
  width: 100%;
}
.wrap {
  height: 100%;
}
</style>
