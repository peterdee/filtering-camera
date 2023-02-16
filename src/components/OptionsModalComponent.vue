<script setup lang="ts">
import { reactive } from 'vue';

import type {
  FilterType,
  GrayscaleType,
  ProcessingType,
} from '../types';
import { FILTER_TYPES } from '../constants';
import StyledSelect from './StyledSelect.vue';
import StyledRange from './StyledRange.vue';

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

interface ComponentState {
  isClosing: boolean;
}

const state = reactive<ComponentState>({
  isClosing: false,
});

const handleClose = (): void => {
  state.isClosing = true;
  setTimeout((): void => emit('toggle-modal'), 150);
};

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
  <div
    :class="`f d-col j-space-between fade-in modal-background ${state.isClosing
      ? 'fade-out'
      : ''}`"
  >
    <div class="empty" />
    <div
      :class="`f d-col mh-auto j-space-between controls ${isMobile
        ? 'controls-mobile'
        : ''}`"
    >
      <div class="f d-col">
        <h2 :class="`ns t-center title ${isMobile ? 'title-mobile' : ''}`">
          Camera: image processing in real time
        </h2>
        <span class="mt-2 ns">
          Processing type
        </span>
        <StyledSelect
          ariaLabel="Processing type"
          globalClasses="mt-half"
          :disabled="!wasmLoaded"
          :value="(processingType as string)"
          @handle-change="handleProcessingTypeSelection"
        >
          <option value="canvas">
            Canvas
          </option>
          <option value="wasm">
            WASM
          </option>
        </StyledSelect>
        <span class="mt-1 ns">
          Available filters
        </span>
        <StyledSelect
          ariaLabel="Available filters"
          globalClasses="mt-half"
          :value="selectedFilter.value"
          @handle-change="handleFilterSelection"
        >
          <option
            v-for="filter in FILTER_TYPES"
            :key="filter.value"
            :value="filter.value"
          >
            {{ filter.name }}
          </option>
        </StyledSelect>
        <span
          v-if="selectedFilter.isGrayscale"
          class="mt-1 ns"
        >
          Grayscale type
        </span>
        <StyledSelect
          v-if="selectedFilter.isGrayscale"
          ariaLabel="Grayscale type"
          globalClasses="mt-half"
          :value="selectedGrayscaleType"
          @handle-change="handleGrayscaleTypeSelection"
        >
          <option value="average">
            Average
          </option>
          <option value="luminosity">
            Luminance
          </option>
        </StyledSelect>
        <StyledRange
          v-if="selectedFilter.withThreshold"
          global-classes="mt-1"
          :selected-filter="selectedFilter"
          :threshold-value="selectedThreshold"
          @handle-input="handleThresholdInput"
        />
      </div>
      <button
        class="mt-1"
        type="button"
        @click="handleClose"
      >
        Close
      </button>
    </div>
    <div class="f j-center mt-2 footer">
      <span class="ns mr-1 footer-text">
        {{ `© ${new Date().getFullYear()}` }}
      </span>
      <span class="ns mr-1 footer-text">
        <a
          href="https://github.com/peterdee/filtering-camera"
          rel="noopener noreferrer"
          target="_blank"
        >
          Camera image processing
        </a>
      </span>
      <span class="ns footer-text">
        <a
          href="https://github.com/peterdee"
          rel="noopener noreferrer"
          target="_blank"
        >
          Peter Dyumin
        </a>
      </span>
    </div>
  </div>
</template>

<style scoped>
.controls {
  max-width: calc(var(--spacer) * 25);
  min-height: calc(var(--spacer) * 22);
  min-width: calc(var(--spacer) * 15);
  width: 50%;
}
.controls-mobile {
  max-width: 90%;
  width: 90%;
}
.empty {
  background-color: transparent;
  height: 1px;
}
.footer {
  height: calc(var(--spacer) * 2);
}
.footer-text {
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
