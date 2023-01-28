<script setup lang="ts">
defineProps<{
  isMobile: boolean;
}>();

const emit = defineEmits(['handle-click']);
</script>

<template>
  <button
    type="button"
    :class="`${!isMobile ? 'mobile' : 'desktop'}`"
    @click="emit('handle-click')"
  >
    Options
  </button>
</template>

<style scoped>
button {
  position: fixed;
}
.desktop {
  bottom: var(--spacer);
  right: var(--spacer);
}
.mobile {
  border-radius: 50%;
  height: calc(var(--spacer) * 3);
  width: calc(var(--spacer) * 3);
  bottom: var(--spacer);
  left: calc(50% - (var(--spacer) + var(--spacer-half)));
}
</style>
