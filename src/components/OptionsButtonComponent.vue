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
    :class="`options ${!isMobile ? 'mobile' : 'desktop'}`"
    @click="emit('handle-click')"
  >
    <ArrowIcon
      v-if="!isMobile"
      :size="24"
    />
    <span v-if="isMobile">
      Options
    </span>
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
.options {
  padding: 0;
}
</style>
