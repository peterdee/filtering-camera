<script setup lang="ts">
import ArrowIcon from './ArrowIcon.vue';

defineProps<{
  isMobile: boolean;
}>();

const emit = defineEmits(['handle-click']);
</script>

<template>
  <button
    type="button"
    :class="`options ${isMobile ? 'mobile' : 'desktop'}`"
    @click="emit('handle-click')"
  >
    <ArrowIcon :size="24" />
  </button>
</template>

<style scoped>
button {
  border-radius: 50%;
  height: calc(var(--spacer) * 3);
  position: fixed;
  width: calc(var(--spacer) * 3);
}
.desktop {
  bottom: calc(var(--spacer) * 2);
  right: calc(var(--spacer) * 2);
}
.mobile {
  bottom: var(--spacer);
  left: calc(50% - (var(--spacer) + var(--spacer-half)));
}
.options {
  padding: 0;
}
</style>
