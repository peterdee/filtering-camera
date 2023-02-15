<script setup lang="ts">
import type {
  FilterType,
  GrayscaleType,
  ProcessingType,
} from '../types';
import { FILTER_TYPES } from '../constants';

defineProps<{
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
  <div class="f d-col j-start fade-in modal-background">
    <button
      class="backdrop"
      type="button"
      @click="emit('toggle-modal')"
    />
    <div class="f d-col p-2 content">
      <div class="f d-col mh-auto controls">
        <span class="ns">
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
      </div>
    </div>
  </div>
</template>

<style scoped>
.backdrop {
  background-color: transparent;
  cursor: default;
  height: 50%;
  width: 100%;
}
.content {
  background-color: rgba(255, 255, 255, .7);
  box-shadow: 0 0 2px 2px var(--text);
  height: 50%;
}
.controls {
  max-width: calc(var(--spacer) * 25);
  width: 50%;
}
.modal-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, .5);
  backdrop-filter: blur(var(--spacer-quarter));
}
</style>
