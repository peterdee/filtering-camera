let wakeLock: WakeLockSentinel;

export default async function requestWakeLock() {
  if ('wakeLock' in navigator && 'request' in navigator.wakeLock) {
    wakeLock = await navigator.wakeLock.request('screen');
  }
}

const handleVisibilityChange = async () => {
  if (wakeLock !== null && document.visibilityState === 'visible') {
    await requestWakeLock();
  }
};

document.addEventListener('visibilitychange', handleVisibilityChange);
