<script setup lang="ts">
const emit = defineEmits(['handle-change']);

const props = defineProps({
  ariaLabel: {
    required: true,
    type: String,
  },
  disabled: {
    default: false,
    required: false,
    type: Boolean,
  },
  globalClasses: {
    default: '',
    required: false,
    type: String,
  },
  value: {
    required: true,
    type: String,
  },
});
</script>

<template>
  <select
    :aria-label="props.ariaLabel"
    :class="`${props.globalClasses}`"
    :disabled="props.disabled"
    :value="props.value"
    @change="(event: Event): void => emit('handle-change', event)"
  >
    <slot />
  </select>
</template>

<style scoped>
select {
  appearance: none;
  background-color: transparent;
  background-image: url('/select-arrow.svg');
  background-position-x: calc(100% - var(--spacer) + var(--spacer-quarter));
  background-position-y: calc(var(--spacer-quarter) / 2);
  background-repeat: no-repeat;
  border: calc(var(--spacer-quarter) / 4) solid var(--accent);
  border-radius: var(--spacer-quarter);
  cursor: pointer;
  font-size: calc(var(--spacer) - var(--spacer-quarter));
  height: calc(var(--spacer) * 2 - var(--spacer-quarter));
  outline: none;
  padding: 0 var(--spacer);
  transition: border var(--transition) ease-out;
  -webkit-appearance: none;
}
select:focus {
  border: calc(var(--spacer-quarter) / 4) solid var(--accent-light);
  transition: border var(--transition) ease-in;
}
select:hover {
  border: calc(var(--spacer-quarter) / 4) solid var(--accent-light);
  transition: border var(--transition) ease-in;
}
</style>
