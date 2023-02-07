<script setup lang="ts">
const props = defineProps({
  size: {
    default: 24,
    required: false,
    type: Number,
  },
});
</script>

<template>
  <svg
    fill="none"
    viewBox="0 0 362 198"
    xmlns="http://www.w3.org/2000/svg"
    :height="props.size"
    :width="props.size"
  >
    <path
      d="M164 16.9706L180.971 1.85966e-05L361.99 181.019L345.019 197.99L164 16.9706Z"
      fill="white"
    />
    <path
      d="M181.019 0L197.99 16.9706L16.9706 197.99L0 181.019L181.019 0Z"
      fill="white"
    />
  </svg>
</template>

<style scoped>

</style>
