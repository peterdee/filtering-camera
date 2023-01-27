<script setup lang="ts">
import { SPACER } from '../constants';

defineProps<{
  count: number;
  isMobile: boolean;
}>();
</script>

<template>
  <div
    class="f ns fps"
    :style="{
      right: isMobile ? `${SPACER / 2}px` : `${SPACER}px`,
      top: isMobile ? `${SPACER / 2}px` : `${SPACER}px`,
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
  background-color: rgba(255, 255, 255, .9);
  border-radius: var(--spacer-half);
  color: var(--accent);
  font-size: calc(var(--spacer) + var(--spacer-quarter));
  font-weight: 500;
  padding: var(--spacer-half);
  position: absolute;
}
.value {
  width: calc(var(--spacer) * 2);
}
</style>
