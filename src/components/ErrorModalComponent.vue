<script setup lang="ts">
import ModalWrap from './ModalWrapComponent.vue';

defineProps<{
  isMobile: boolean;
}>();
</script>

<template>
  <ModalWrap :is-closing="false">
    <div
      :class="`f d-col mh-auto content ${isMobile
        ? 'content-mobile'
        : ''}`"
    >
      <div class="ns t-center title">
        Could not access the camera!
      </div>
      <div class="mt-1 ns t-center text">
        This application requires an access to the camera, and
        it seems that camera is not available!
      </div>
    </div>
  </ModalWrap>
</template>

<style scoped>
.content {
  max-width: calc(var(--spacer) * 25);
  min-width: calc(var(--spacer) * 15);
  width: 50%;
}
.content-mobile {
  max-width: 90%;
  width: 90%;
}
.text {
  font-size: calc(var(--spacer) - var(--spacer-quarter) / 2);
  font-weight: 200;
}
.title {
  color: var(--error);
  font-size: calc(var(--spacer) + var(--spacer-half));
  font-weight: 300;
}
.title-mobile {
  font-size: calc(var(--spacer) + var(--spacer-quarter));
}
</style>
