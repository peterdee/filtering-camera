export default function isMobile(): boolean {
  return /Mobi/i.test(window.navigator.userAgent);
}
