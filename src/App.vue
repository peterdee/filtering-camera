<script setup lang="ts">
import {
  onMounted,
  reactive,
  ref,
} from 'vue';

import bridge from './bridge';
import canvasFilters from './processing-canvas';
import ErrorModal from './components/ErrorModalComponent.vue';
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
  flipImage: boolean;
  fpsCount: number;
  frameTime: number[];
  isMobile: boolean;
  processingType: ProcessingType;
  selectedFilter: FilterType;
  selectedGrayscaleType: GrayscaleType;
  selectedThreshold: number;
  showErrorModal: boolean;
  showOptionsModal: boolean;
  wasmLoaded: boolean;
}

const state = reactive<ComponentState>({
  ctx: null,
  flipImage: false,
  fpsCount: 0,
  frameTime: [],
  isMobile: true,
  processingType: 'canvas',
  selectedFilter: FILTER_TYPES[0],
  selectedGrayscaleType: 'average',
  selectedThreshold: FILTER_TYPES[0].defaultThreshold || 0,
  showErrorModal: false,
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

  let canvasHeight = window.innerHeight;
  let canvasWidth = window.innerWidth;
  if (canvasRef.value) {
    canvasHeight = canvasRef.value.height;
    canvasWidth = canvasRef.value.width;
  }

  const imageData = ((): ImageData => {
    const frame = ctx.getImageData(0, 0, canvasWidth, canvasHeight);
    if (state.processingType === 'canvas') {
      if (state.selectedFilter.value === 'binary') {
        return canvasFilters.binary(frame, state.selectedThreshold);
      }
      if (state.selectedFilter.value === 'colorInversion') {
        return canvasFilters.colorInversion(frame);
      }
      if (state.selectedFilter.value === 'eightColors') {
        return canvasFilters.eightColors(frame);
      }
      if (state.selectedFilter.value === 'grayscale') {
        return canvasFilters.grayscale(frame, state.selectedGrayscaleType);
      }
      if (state.selectedFilter.value === 'laplacian') {
        return canvasFilters.laplacian(frame);
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

const handleError = (): void => {
  state.showErrorModal = true;
};

const handleSuccess = (stream: MediaStream): null | void => {
  // check if video track is available
  const [videoTrack = null] = stream.getVideoTracks();
  if (!videoTrack) {
    state.showErrorModal = true;
    return null;
  }

  const windowHeight = window.innerHeight;
  const windowWidth = window.innerWidth;
  let canvasHeight = windowHeight;
  let canvasWidth = windowWidth;

  // get video resolution
  const capabilities = videoTrack.getCapabilities();
  if (capabilities.height && capabilities.height.max) {
    canvasHeight = capabilities.height.max > windowHeight
      ? windowHeight
      : capabilities.height.max;
  }
  if (capabilities.width && capabilities.width.max) {
    canvasWidth = capabilities.width.max > windowWidth
      ? windowWidth
      : capabilities.width.max;
  }

  // set canvas size
  if (canvasRef.value) {
    canvasRef.value.height = canvasHeight;
    canvasRef.value.width = canvasWidth;
  }

  // adjust video track resolution for mobile devices
  if (state.isMobile) {
    videoTrack.applyConstraints({
      height: canvasWidth,
      width: canvasHeight,
    });
  }

  const video = document.createElement('video');
  video.onplay = (): null | NodeJS.Timeout | void => draw(video);
  video.muted = true;
  video.playsInline = true;
  video.srcObject = stream;
  video.play();
  return null;
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

const handleThresholdInput = (value: string): void => {
  state.selectedThreshold = Number(value);
};

const toggleFlipImage = (): void => {
  state.flipImage = !state.flipImage;
};

const toggleOptionsModal = (): void => {
  state.showOptionsModal = !state.showOptionsModal;
};

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

  if (!mobile) {
    state.flipImage = true;
  }

  if (canvasRef.value) {
    const ctx = canvasRef.value.getContext('2d', { willReadFrequently: true }) || null;
    if (ctx) {
      state.ctx = ctx;
    }
  }

  const windowHeight = window.innerHeight;
  const windowWidth = window.innerWidth;

  const constraints: MediaStreamConstraints = {
    audio: false,
    video: {
      height: {
        ideal: windowHeight,
        max: windowHeight,
        min: 1,
      },
      width: {
        ideal: windowWidth,
        max: windowWidth,
        min: 1,
      },
    },
  };
  if (mobile) {
    (constraints.video as MediaTrackConstraints).facingMode = { exact: 'environment' };
  }

  navigator.mediaDevices.getUserMedia(constraints)
    .then(handleSuccess)
    .catch(handleError);
});
</script>

<template>
  <div
    :class="`f j-center ai-center wrap ${state.isMobile
      ? 'wrap-mobile'
      : ''}`"
  >
    <canvas
      ref="canvasRef"
      :class="`${state.flipImage ? 'flip' : ''}`"
    />
    <ErrorModal
      v-if="state.showErrorModal"
      :is-mobile="state.isMobile"
    />
    <FPSCounter
      v-if="!state.showErrorModal && !state.showOptionsModal"
      :count="state.fpsCount"
      :is-mobile="state.isMobile"
    />
    <OptionsButton
      v-if="!state.showErrorModal && !state.showOptionsModal"
      :is-mobile="state.isMobile"
      @handle-click="toggleOptionsModal"
    />
    <OptionsModal
      v-if="state.showOptionsModal"
      :flip-image="state.flipImage"
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
      @toggle-flip="toggleFlipImage"
      @toggle-modal="toggleOptionsModal"
    />
  </div>
</template>

<style scoped>
.flip {
  -moz-transform: scale(-1, 1);
  -o-transform: scale(-1, 1);
  -webkit-transform: scale(-1, 1);
  transform: scale(-1, 1);
}
.wrap {
  height: 100vh;
}
.wrap-mobile {
  height: 100%;
}
</style>
