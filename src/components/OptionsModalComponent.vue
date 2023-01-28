<script setup lang="ts">
const emit = defineEmits(['toggle-modal']);
</script>

<template>
  <div class="f d-col j-start fade-in modal-background">
    <button
      class="backdrop"
      type="button"
      @click="emit('toggle-modal')"
    />
    <div class="f d-col p-2 content">
      Content
    </div>
  </div>
</template>

<style scoped>
.backdrop {
  background-color: transparent;
  cursor: default;
  width: 100%;
}
.content {
  background-color: var(--background);
  box-shadow: 0 0 2px 2px var(--text);
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
