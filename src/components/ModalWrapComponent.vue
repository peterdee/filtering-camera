<script setup lang="ts">
defineProps<{
  isClosing: boolean;
}>();
</script>

<template>
  <div
    :class="`f d-col j-space-between fade-in modal-background ${isClosing
      ? 'fade-out'
      : ''}`"
  >
    <div class="empty" />
    <slot />
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
</style>
