<script setup lang="ts">
import type {
  FilterType,
  GrayscaleType,
  ProcessingType,
} from '../types';
import { FILTER_TYPES } from '../constants';

defineProps<{
  isMobile: boolean;
  processingType: ProcessingType;
  selectedFilter: FilterType;
  selectedGrayscaleType: GrayscaleType;
  selectedThreshold: number;
  wasmLoaded: boolean;
}>();

const emit = defineEmits([
  'handle-filter',
  'handle-grayscale-type',
  'handle-processing-type',
  'handle-threshold',
  'toggle-modal',
]);

const handleFilterSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  const [filter] = FILTER_TYPES.filter(
    (entry: FilterType): boolean => entry.value === value,
  );
  return emit('handle-filter', filter);
};

const handleGrayscaleTypeSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  return emit('handle-grayscale-type', value);
};

const handleProcessingTypeSelection = (event: Event): void => {
  const { value } = event.target as HTMLSelectElement;
  return emit('handle-processing-type', value);
};

const handleThresholdInput = (event: Event): void => {
  const { value } = event.target as HTMLInputElement;
  return emit('handle-threshold', value);
};
</script>

<template>
  <div class="f d-col j-center fade-in modal-background">
    <div
      :class="`f d-col mh-auto controls ${isMobile ? 'controls-mobile' : ''}`"
    >
      <h2 :class="`ns t-center title ${isMobile ? 'title-mobile' : ''}`">
        Camera: image processing in real time
      </h2>
      <span class="mt-2 ns">
        Processing type
      </span>
      <select
        aria-label="Processing type"
        class="mt-half"
        :disabled="!wasmLoaded"
        :value="processingType"
        @change="handleProcessingTypeSelection"
      >
        <option value="canvas">
          Canvas
        </option>
        <option value="wasm">
          WASM
        </option>
      </select>
      <span class="mt-1 ns">
        Available filters
      </span>
      <select
        aria-label="Available filters"
        class="mt-half"
        :value="selectedFilter.value"
        @change="handleFilterSelection"
      >
        <option
          v-for="filter in FILTER_TYPES"
          :key="filter.value"
          :value="filter.value"
        >
          {{ filter.name }}
        </option>
      </select>
      <span
        v-if="selectedFilter.isGrayscale"
        class="mt-1 ns"
      >
        Grayscale type
      </span>
      <select
        v-if="selectedFilter.isGrayscale"
        aria-label="Grayscale type"
        class="mt-half"
        :value="selectedGrayscaleType"
        @change="handleGrayscaleTypeSelection"
      >
        <option value="average">
          Average
        </option>
        <option value="luminosity">
          Luminance
        </option>
      </select>
      <div
        v-if="selectedFilter.withThreshold"
        class="f d-col j-center mt-1 ns"
      >
        <span>
          Filter threshold
        </span>
        <div class="f j-center ai-center mt-half">
          <span>
            {{ selectedFilter.minThreshold || 0 }}
          </span>
          <input
            aria-label="Threshold value"
            class="mh-1"
            type="range"
            :max="selectedFilter.maxThreshold || 255"
            :min="selectedFilter.minThreshold || 0"
            :step="selectedFilter.step || 1"
            :value="selectedThreshold"
            @input="handleThresholdInput"
          />
          <span>
            {{ selectedFilter.maxThreshold || 255 }}
          </span>
        </div>
        <span class="t-center">
          {{ selectedThreshold }}
        </span>
      </div>
      <button
        class="mt-1"
        type="button"
        @click="emit('toggle-modal')"
      >
        Close
      </button>
      <div class="f j-center mt-2">
        <span class="footer ns mr-1">
          {{ `© ${new Date().getFullYear()}` }}
        </span>
        <span class="footer ns">
          <a hre="https://github.com/peterdee">Peter Dyumin</a>
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.controls {
  max-width: calc(var(--spacer) * 25);
  min-width: calc(var(--spacer) * 15);
  width: 50%;
}
.controls-mobile {
  max-width: 90%;
  width: 90%;
}
.footer {
  font-size: calc(var(--spacer) - var(--spacer-quarter));
}
.modal-background {
  background-color: rgba(255, 255, 255, .9);
  height: 100%;
  left: 0;
  position: fixed;
  top: 0;
  width: 100%;
}
.title {
  font-size: calc(var(--spacer) + var(--spacer-half));
  font-weight: 300;
}
.title-mobile {
  font-size: calc(var(--spacer) + var(--spacer-quarter));
}
</style>
