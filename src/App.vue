<script setup lang="ts">
import {
  onMounted,
  reactive,
  ref,
} from 'vue';

import binary from './processing-canvas/binary';
import eightColors from './processing-canvas/eight-colors';
import grayscale from './processing-canvas/grayscale';
import sobel from './processing-canvas/sobel';
import solarize from './processing-canvas/solarize';

interface ComponentState {
  ctx: CanvasRenderingContext2D | null;
}

const state = reactive<ComponentState>({
  ctx: null,
})

const canvasRef = ref<HTMLCanvasElement>();

const draw = (video: HTMLVideoElement): null | void => {
  const { ctx } = state;
  if (!ctx) {
    return null;
  }

  ctx.drawImage(video, 0, 0);

  // const imageData = grayscale(
  //   ctx.getImageData(0, 0, 400, 400),
  //   'luminosity',
  // );
  // const imageData = binary(
  //   ctx.getImageData(0, 0, 400, 400),
  //   105,
  // );
  // const imageData = eightColors(ctx.getImageData(0, 0, 400, 400));
  // const imageData = solarize(
  //   ctx.getImageData(0, 0, 400, 400),
  //   105,
  // );
  const imageData = sobel(ctx.getImageData(0, 0, 400, 400));

  ctx.putImageData(imageData, 0, 0);

  setTimeout(draw, 10, video);
}

const handleError = (error: MediaStreamError): void => {
  console.log(error);
}

const handleSuccess = (stream: MediaStream): null | void => {
  const video = document.createElement('video');
  video.onplay = (): null | void => draw(video)
  video.srcObject = stream;
  video.play();
}

onMounted((): void => {
  if (canvasRef.value) {
    canvasRef.value.height = window.innerHeight;
    canvasRef.value.width = window.innerWidth;
    state.ctx = canvasRef.value.getContext('2d');
  }

  navigator.getUserMedia(
    {
      audio: false,
      video: true,
    },
    handleSuccess,
    handleError,
  );
});
</script>

<template>
  <div class="f j-center ai-center wrap">
    <canvas ref="canvasRef" />
  </div>
</template>

<style scoped>
canvas {
  max-height: 100vh;
  max-width: 100vw;
}
.wrap {
  height: 100vh;
  width: 100%;
}
</style>
