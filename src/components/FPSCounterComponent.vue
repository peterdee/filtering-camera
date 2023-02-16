<script setup lang="ts">
import { SPACER } from '../constants';

defineProps<{
  count: number;
  isMobile: boolean;
}>();
</script>

<template>
  <div
    class="f fade-in ns fps"
    :style="{
      right: isMobile ? `${SPACER / 2}px` : `${SPACER * 2}px`,
      top: isMobile ? `${SPACER / 2}px` : `${SPACER * 2}px`,
    }"
  >
    <div>
      FPS:
    </div>
    <div class="t-right value">
      {{ count }}
    </div>
  </div>
</template>

<style scoped>
.fps {
  background-color: rgba(255, 255, 255, .7);
  backdrop-filter: blur(var(--spacer-quarter));
  border-radius: var(--spacer-half);
  color: var(--accent-dark);
  font-size: calc(var(--spacer) + var(--spacer-quarter));
  font-weight: 500;
  padding: var(--spacer-half);
  position: absolute;
}
.value {
  width: calc(var(--spacer) * 2);
}
</style>
