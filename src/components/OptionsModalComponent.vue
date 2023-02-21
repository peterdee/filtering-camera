<script setup lang="ts">
import { reactive } from 'vue';

import type {
  FilterType,
  GrayscaleType,
  ProcessingType,
} from '../types';
import { FILTER_TYPES } from '../constants';
import ModalWrap from './ModalWrapComponent.vue';
import StyledSelect from './StyledSelect.vue';
import StyledRange from './StyledRange.vue';

defineProps<{
  flipImage: boolean;
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
  'toggle-flip',
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
  <ModalWrap :is-closing="state.isClosing">
    <div
      :class="`f d-col mh-auto j-space-between controls ${isMobile
        ? 'controls-mobile'
        : ''}`"
    >
      <div class="f d-col">
        <h2 :class="`ns t-center title ${isMobile ? 'title-mobile' : ''}`">
          Camera: image processing in real time
        </h2>
        <div class="f ai-center j-space-between ns mt-2">
          <label
            class="switch-label"
            for="switch"
          >
            Flip image
          </label>
          <label class="switch">
            <input
              id="switch"
              type="checkbox"
              :checked="flipImage"
              @input="emit('toggle-flip')"
            >
            <span class="slider round" />
          </label>
        </div>
        <span class="mt-1 ns">
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
          <option value="luminance">
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
  </ModalWrap>
</template>

<style scoped>
.checkbox {
  background-color: transparent;
}
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
.title {
  font-size: calc(var(--spacer) + var(--spacer-half));
  font-weight: 300;
}
.title-mobile {
  font-size: calc(var(--spacer) + var(--spacer-quarter));
}
.switch {
  display: inline-block;
  height: calc(var(--spacer) + var(--spacer-half));
  position: relative;
  width: calc(var(--spacer) * 3);
}
.switch input {
  height: 0;
  opacity: 0;
  width: 0;
}
.switch-label {
  cursor: pointer;
}
.slider {
  background-color: var(--muted);
  bottom: 0;
  cursor: pointer;
  left: 0;
  position: absolute;
  right: 0;
  top: 0;
  transition: var(--transition);
  -webkit-transition: var(--transition);
}
.slider:before {
  background-color: var(--background);
  bottom: var(--spacer-quarter);
  content: "";
  height: var(--spacer);
  left: var(--spacer-quarter);
  position: absolute;
  transition: var(--transition);
  -webkit-transition: var(--transition);
  width: var(--spacer);
}
input:checked + .slider {
  background-color: var(--accent);
}
input:focus + .slider {
  box-shadow: 0 0 1px var(--accent);
}
input:checked + .slider:before {
  -ms-transform: translateX(calc(var(--spacer) + var(--spacer-half)));
  -webkit-transform: translateX(calc(var(--spacer) + var(--spacer-half)));
  transform: translateX(calc(var(--spacer) + var(--spacer-half)));
}
.slider.round {
  border-radius: calc(var(--spacer));
}
.slider.round:before {
  border-radius: 50%;
}
</style>
